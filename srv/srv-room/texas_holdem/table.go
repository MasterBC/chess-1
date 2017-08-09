package texas_holdem

import (
	"chess/common/define"
	pb "chess/srv/srv-room/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"sync"
	"time"
	"github.com/satori/go.uuid"
	"chess/common/log"
	"chess/srv/srv-room/signal"
)

const (
	actionWait = 20
	MaxN       = 10

	// 四轮发牌 preflop (底牌), flop (翻牌), turn (转牌), river(河牌)
	DealPreflop = "preflop"
	DealFlop    = "flop"
	DealTurn    = "turn"
	DealRiver   = "river"
)



type Table struct {
	Id         string
	RoomId     int
	SmallBlind int
	BigBlind   int
	MinCarry int  // 最大携带
	MaxCarry int // 最小携带
	Cards      Cards
	Pot        []int32 // 奖池筹码数, 第一项为主池，其他项(若存在)为边池
	PotList PotDetails // 奖池 详细信息
	Timeout    int
	Button     int  // 庄家位置
	Players    Players // 坐下的玩家
	Bystanders Players // 站起的玩家
	Chips      []int32 // 玩家最终下注筹码，摊牌时为玩家最终获得筹码
	Bet        int  // 当前回合 上一玩家下注额
	N          int // 当前牌桌玩家数
	Max        int // 牌桌最大玩家数

	MaxChips int
	MinChips int
	remain   int
	allin    int
	EndChan  chan int  // 牌局结束通知
	exitChan chan interface{}  // 销毁牌桌
	lock     sync.Mutex
	dm       *DealMachine
}

func genTableId(rid int) string {
	u := uuid.NewV4()
	return fmt.Sprintf("%d-%s", rid, u.String())
}

func NewTable(rid, max, sb, bb, minC, maxC int) *Table {
	if max <= 0 || max > MaxN {
		max = 9 // default 9 players
	}

	table := &Table{
		Id:         genTableId(rid),
		RoomId: 	rid,
		Players:    make([]*Player, max, MaxN),
		Chips:      make([]int32, max, MaxN),
		SmallBlind: sb,
		BigBlind:   bb,
		MinCarry: minC,
		MaxCarry: maxC,
		Pot:        make([]int32, 1),
		Timeout:    actionWait,
		Max:        max,
		lock:       sync.Mutex{},
		dm:         NewDealMachine(),
		EndChan:    make(chan int),
		exitChan:   make(chan interface{}, 1),
	}

	// 初始化发牌器
	table.dm.Init()
	signal.TableWg.Add(1)
	go func() {
		defer signal.TableWg.Done()
		timer := time.NewTimer(time.Second * 6)
		for {
			select {
			case <-timer.C:
				table.start()
				timer.Reset(time.Second * 6)
			case <-table.exitChan:
				return
			case <-signal.TableDie: // server is shuting down...
				log.Debug("------Receive signal.TableDie")
				table.shutdown()
				return
			}
		}
	}()

	return table
}

// 当前游戏玩家数
func (t *Table) Cap() int {
	return len(t.Players)
}

func (t *Table) Player(id int) *Player {
	for _, p := range t.Players {
		if p != nil && p.Id == id {
			return p
		}
	}
	return nil
}

func (t *Table) AddBystander(p *Player) {
	t.lock.Lock()
	defer t.lock.Unlock()

	// table not exists
	if len(t.Id) == 0 {
		return
	}

	for _, v := range t.Bystanders {
		if v.Id == p.Id {
			log.Debugf("玩家%d已在旁观列表中！", p.Id)
			return
		}
	}
	t.Bystanders = append(t.Bystanders, p)
}

func (t *Table) DelBystander(p *Player) {
	if p == nil {
		return
	}

	t.lock.Lock()
	defer t.lock.Unlock()

	for k, v := range t.Bystanders {
		if v.Id == p.Id {
			t.Bystanders = append(t.Bystanders[:k], t.Bystanders[k+1:]...)
			return
		}
	}
}

// return player pos
func (t *Table) AddPlayer(p *Player) int {
	t.lock.Lock()
	defer t.lock.Unlock()

	// table not exists
	if len(t.Id) == 0 {
		return 0
	}

	for pos := range t.Players {
		if t.Players[pos] == nil {
			t.Players[pos] = p
			t.N++
			p.Table = t
			p.Pos = pos + 1
			p.Action = ActSitdown
			break
		}
	}

	return p.Pos
}

func (t *Table) DelPlayer(p *Player) {
	if p == nil || p.Pos == 0 {
		return
	}

	t.lock.Lock()
	defer t.lock.Unlock()

	t.Players[p.Pos-1] = nil
	t.N--
	if len(p.Cards) > 0 {
		t.remain--
	}

	if t.N == 0 {
		log.Debugf("(%s)牌桌上已无玩家，销毁之！", t.Id)
		DelTable(t.Id)
		select {
		case t.exitChan <- 0:
		default:
		}
	}

	if t.remain <= 1 {
		select {
		case t.EndChan <- 0:
		default:
		}
	}
}

// 广播牌桌上的玩家
func (t *Table) Broadcast(code int16, msg proto.Message) {
	for _, p := range t.Players {
		if p != nil {
			p.SendMessage(code, msg)
		}
	}
}

// 广播旁观的玩家
func (t *Table) BroadcastBystanders(code int16, msg proto.Message) {
	for _, p := range t.Bystanders {
		if p != nil {
			p.SendMessage(code, msg)
		}
	}
}

// 广播 牌桌上的玩家和旁观的玩家
func (t *Table) BroadcastAll(code int16, msg proto.Message) {
	for _, p := range t.Players {
		if p != nil {
			p.SendMessage(code, msg)
		}
	}

	for _, p := range t.Bystanders {
		if p != nil {
			p.SendMessage(code, msg)
		}
	}
}

// start starts from 0
func (t *Table) Each(start int, f func(p *Player) bool) {
	end := (t.Cap() + start - 1) % t.Cap()
	i := start
	for ; i != end; i = (i + 1) % t.Cap() {
		if t.Players[i] != nil && !f(t.Players[i]) {
			return
		}
	}

	// end
	if t.Players[i] != nil {
		f(t.Players[i])
	}
}

func (t *Table) start() {
	var dealer *Player

	t.Each(0, func(p *Player) bool {
		if p.Chips < t.BigBlind {
			p.Leave()
		}
		p.Bet = 0
		p.Cards = nil
		p.Action = ActReady
		p.Hand.Init()
		return true
	})

	// Select Dealer
	button := t.Button - 1
	t.Each((button+1)%t.Cap(), func(p *Player) bool {
		t.Button = p.Pos
		dealer = p
		return false
	})

	if dealer == nil {
		log.Debugf("(%s)找不到庄家", t.Id)
		return
	}

	t.lock.Lock()
	if t.N < 2 {
		t.lock.Unlock()
		log.Debugf("(%s)牌桌上玩家小于2人", t.Id)
		return
	}

	// 洗牌
	t.dm.Shuffle()

	// Small Blind
	sb := dealer.Next()
	if t.N == 2 { // one-to-one
		sb = dealer
	}
	if sb == nil {
		t.lock.Unlock()
		log.Debugf("(%s)找不到小盲注玩家", t.Id)
		return
	}

	// Big Blind
	bb := sb.Next()
	if bb == nil {
		t.lock.Unlock()
		log.Debugf("(%s)找不到大盲注玩家", t.Id)
		return
	}

	bbPos := bb.Pos

	t.Pot = nil
	t.PotList = nil
	t.Chips = make([]int32, t.Max)
	t.Bet = 0
	t.Cards = nil
	t.remain = 0
	t.allin = 0
	t.Each(0, func(p *Player) bool {
		if p.Action == ActReady {
			p.Bet = 0
			p.Cards = Cards{t.dm.Deal(), t.dm.Deal()}
			t.remain++
		}
		return true
	})
	t.lock.Unlock()

	// 2107, 通报本局庄家 (服务器广播此消息，代表游戏开始并确定本局庄家)
	t.BroadcastAll(define.Code["room_button_ack"], &pb.RoomButtonAck{
		BaseAck:   &pb.BaseAck{Ret: 1, Msg: "ok"},
		TableId:   t.Id,
		ButtonPos: int32(t.Button),
	})

	t.betting(sb.Pos, t.SmallBlind) // 小盲注
	t.betting(bb.Pos, t.BigBlind)   // 大盲注

	// Round 1 : preflop
	t.Each(sb.Pos-1, func(p *Player) bool {
		if p.Cards != nil {
			// 2108, 发牌
			p.SendMessage(define.Code["room_deal_ack"], &pb.RoomDealAck{
				BaseAck:   &pb.BaseAck{Ret: 1, Msg: "ok"},
				Action:    DealPreflop,
				Cards:     p.Cards.ToProtoMessage(),
				HandLevel: -1,
				HandFinalValue: -1,
			})
		}
		return true
	})
	// 旁观玩家通报
	t.BroadcastBystanders(define.Code["room_deal_ack"], &pb.RoomDealAck{
		BaseAck:   &pb.BaseAck{Ret: 1, Msg: "ok"},
		Action:    DealPreflop,
		HandLevel: -1,
		HandFinalValue: -1,
	})

	t.action(bbPos%t.Cap() + 1)
	if t.remain <= 1 {
		goto showdown
	}
	t.calc()

	// Round 2 : Flop
	t.ready()
	t.Cards = Cards{
		t.dm.Deal(),
		t.dm.Deal(),
		t.dm.Deal(),
	}
	t.Each(0, func(p *Player) bool {
		if len(p.Cards) > 0 {
			p.Hand.Init()
			p.Hand.SetCard(t.Cards[0])
			p.Hand.SetCard(t.Cards[1])
			p.Hand.SetCard(t.Cards[2])
			p.Hand.SetCard(p.Cards[0])
			p.Hand.SetCard(p.Cards[1])
			p.Hand.AnalyseHand()
		}
		// 2108,  翻牌
		p.SendMessage(define.Code["room_deal_ack"], &pb.RoomDealAck{
			BaseAck:   &pb.BaseAck{Ret: 1, Msg: "ok"},
			Action:    DealFlop,
			Cards:     t.Cards.ToProtoMessage(),
			HandLevel: int32(p.Hand.Level),
			HandFinalValue: int32(p.Hand.FinalValue),
		})

		return true
	})
	// 旁观玩家通报
	t.BroadcastBystanders(define.Code["room_deal_ack"], &pb.RoomDealAck{
		BaseAck:   &pb.BaseAck{Ret: 1, Msg: "ok"},
		Action:    DealFlop,
		Cards:     t.Cards.ToProtoMessage(),
	})

	t.action(0)

	if t.remain <= 1 {
		goto showdown
	}
	t.calc()

	// Round 3 : Turn
	t.ready()
	t.Cards = append(t.Cards, t.dm.Deal())
	t.Each(0, func(p *Player) bool {
		if len(p.Cards) > 0 {
			p.Hand.Init()
			p.Hand.SetCard(t.Cards[0])
			p.Hand.SetCard(t.Cards[1])
			p.Hand.SetCard(t.Cards[2])
			p.Hand.SetCard(t.Cards[3])
			p.Hand.SetCard(p.Cards[0])
			p.Hand.SetCard(p.Cards[1])
			p.Hand.AnalyseHand()
		}
		// 2108,  转牌
		p.SendMessage(define.Code["room_deal_ack"], &pb.RoomDealAck{
			BaseAck:   &pb.BaseAck{Ret: 1, Msg: "ok"},
			Action:    DealTurn,
			Cards:     t.Cards.ToProtoMessage(),
			HandLevel: int32(p.Hand.Level),
			HandFinalValue: int32(p.Hand.FinalValue),
		})

		return true
	})
	// 旁观玩家通报
	t.BroadcastBystanders(define.Code["room_deal_ack"], &pb.RoomDealAck{
		BaseAck:   &pb.BaseAck{Ret: 1, Msg: "ok"},
		Action:    DealTurn,
		Cards:     t.Cards.ToProtoMessage(),
	})

	t.action(0)
	if t.remain <= 1 {
		goto showdown
	}
	t.calc()

	// Round 4 : River   河牌
	t.ready()
	t.Cards = append(t.Cards, t.dm.Deal())
	t.Each(0, func(p *Player) bool {
		if len(p.Cards) > 0 {
			p.Hand.Init()
			p.Hand.SetCard(t.Cards[0])
			p.Hand.SetCard(t.Cards[1])
			p.Hand.SetCard(t.Cards[2])
			p.Hand.SetCard(t.Cards[3])
			p.Hand.SetCard(t.Cards[4])
			p.Hand.SetCard(p.Cards[0])
			p.Hand.SetCard(p.Cards[1])
			p.Hand.AnalyseHand()
		}
		// 2108,  河牌
		p.SendMessage(define.Code["room_deal_ack"], &pb.RoomDealAck{
			BaseAck:   &pb.BaseAck{Ret: 1, Msg: "ok"},
			Action:    DealRiver,
			Cards:     t.Cards.ToProtoMessage(),
			HandLevel: int32(p.Hand.Level),
			HandFinalValue: int32(p.Hand.FinalValue),
		})

		return true
	})
	// 旁观玩家通报
	t.BroadcastBystanders(define.Code["room_deal_ack"], &pb.RoomDealAck{
		BaseAck:   &pb.BaseAck{Ret: 1, Msg: "ok"},
		Action:    DealRiver,
		Cards:     t.Cards.ToProtoMessage(),
	})

	t.action(0)

showdown:
	t.showdown()
	// Final : Showdown   2111, 摊牌和比牌
	t.BroadcastAll(define.Code["room_showdown_ack"], &pb.RoomShowdownAck{
		BaseAck: &pb.BaseAck{Ret: 1, Msg: "ok"},
		Table:   t.ToProtoMessage(),
		PotList: t.PotList.ToProtoMessage(),
	})
}

func (t *Table) action(pos int) {
	if t.allin+1 >= t.remain {
		return
	}

	skip := 0
	if pos == 0 { // start from left hand of button
		pos = (t.Button)%t.Cap() + 1
	}

	for {
		raised := 0

		t.Each(pos-1, func(p *Player) bool {
			if t.remain <= 1 {
				return false
			}

			// 跳过玩家
			if p.Pos == skip || p.Chips == 0 || len(p.Cards) == 0 {
				return true
			}

			p.Action = ActBetting

			// 2110, 通报当前下注玩家
			t.BroadcastAll(define.Code["room_action_ack"], &pb.RoomActionAck{
				BaseAck: &pb.BaseAck{Ret: 1, Msg: "ok"},
				Pos:     int32(p.Pos),
				BaseBet: int32(t.Bet),
			})

			msg, _ := p.GetActionBet(time.Duration(t.Timeout) * time.Second)
			if t.remain <= 1 {
				return false
			}

			n := 0
			// timeout or leave
			if msg == nil {
				n = -1
			} else {
				n = int(msg.Bet)
			}

			if t.betting(p.Pos, n) { // 玩家加注
				raised = p.Pos
				return false
			}

			return true
		})

		if raised == 0 {
			break
		}

		pos = raised
		skip = pos
	}
}

// 计算奖池
func (t *Table) calc() (pots []handPot) {
	pots = calcPot(t.Chips)
	t.Pot = nil
	for _, pot := range pots {
		t.Pot = append(t.Pot, int32(pot.Pot))
	}

	// 2109, 通报奖池
	t.BroadcastAll(define.Code["room_pot_ack"], &pb.RoomPotAck{
		BaseAck: &pb.BaseAck{Ret: 1, Msg: "ok"},
		Pot:     t.Pot,
	})

	return
}

// 关闭桌子
func (t *Table) shutdown() {
	t.BroadcastAll(define.Code["room_shutdown_table_ack"], &pb.RoomShutdownTableAck{
		BaseAck: &pb.BaseAck{Ret: 1, Msg: "ok"},
	})
	for _, p := range t.Players {
		if p != nil {
			p.Leave()
		}
	}

	for _, p := range t.Bystanders {
		if p != nil {
			p.Leave()
		}
	}
}

// 摊牌
func (t *Table) showdown() {
	pots := t.calc()

	for i := range t.Chips {
		t.Chips[i] = 0
	}

	t.lock.Lock()
	defer t.lock.Unlock()

	t.PotList = nil

	for _, pot := range pots { // 遍历奖池

		maxHandLevel := -1
		maxHandFinalValue := -1

		// 计算该池子最大牌型和牌值
		for _, pos := range pot.OPos {
			p := t.Players[pos-1]
			if p != nil {
				if p.Hand.Level > maxHandLevel {
					maxHandLevel = p.Hand.Level
					maxHandFinalValue = p.Hand.FinalValue
				} else if p.Hand.Level == maxHandLevel && p.Hand.FinalValue > maxHandFinalValue {
					maxHandFinalValue = p.Hand.FinalValue
				}
			}
		}

		var winners []int

		for _, pos := range pot.OPos {
			p := t.Players[pos-1]
			if p != nil && len(p.Cards) > 0 {
				if p.Hand.Level == maxHandLevel && p.Hand.FinalValue == maxHandFinalValue {
					winners = append(winners, pos)
				}
			}
		}

		if len(winners) == 0 {
			fmt.Println("!!!no winners!!!")
			return
		}
		potDetail := &PotDetail{
			Pot: pot.Pot,
			Ps: make([]int32, len(t.Chips)),
		}
		for _, winner := range winners {
			potDetail.Ps[winner-1] += int32(pot.Pot / len(winners))
			t.Chips[winner-1] += int32(pot.Pot / len(winners))
		}
		potDetail.Ps[winners[0]-1] += int32(pot.Pot % len(winners)) // odd chips
		t.Chips[winners[0]-1] += int32(pot.Pot % len(winners)) // odd chips

		t.PotList = append(t.PotList, potDetail)
	}

	for i := range t.Chips {
		if t.Players[i] != nil {
			t.Players[i].Chips += int(t.Chips[i])
			if t.Players[i].Chips <= t.BigBlind { // 补筹码
				carry := t.MaxCarry/2 - t.Players[i].Chips
				if t.Players[i].CurrChips >= carry {
					t.Players[i].Chips += carry
					t.Players[i].CurrChips -= carry
				} else {
					t.Players[i].Chips += t.Players[i].CurrChips
					t.Players[i].CurrChips = 0
				}
			}
		}
	}
}

func (t *Table) ready() {
	t.Bet = 0
	t.lock.Lock()
	defer t.lock.Unlock()

	t.Each(0, func(p *Player) bool {
		p.Bet = 0
		if p.Action == ActAllin || p.Action == ActFold || p.Action == ActSitdown {
			return true
		}
		p.Action = ActReady
		return true
	})

}

func (t *Table) betting(pos, n int) (raised bool) {
	if pos <= 0 {
		return
	}

	t.lock.Lock()
	defer t.lock.Unlock()

	p := t.Players[pos-1]
	if p == nil {
		return
	}

	// 下注合法性判断
	if n > p.Chips || // 手上筹码不足
		(n == 0 && p.Bet != t.Bet) ||  // 让牌
		(n > 0 && n != p.Chips && ((n + p.Bet) < t.Bet)){
		log.Debugf("下注筹码不合法!!！ n:%d  p.Bet:%d  p.Chips:%d  t.Bet:%d",n,p.Bet,p.Chips,t.Bet)
		return
	}

	raised = p.Betting(n)
	if p.Action == ActFold {
		t.remain--
	}
	if p.Action == ActAllin {
		t.allin++
	}

	// 2106， 通报玩家下注结果
	t.BroadcastAll(define.Code["room_player_bet_ack"], &pb.RoomPlayerBetAck{
		BaseAck: &pb.BaseAck{Ret: 1, Msg: "ok"},
		TableId: t.Id,
		Action:  p.Action,
		Bet:     int32(p.Bet),
		Chips:   int32(p.Chips),
		Pos: int32(pos),
	})

	return
}

func (t *Table) ToProtoMessage() *pb.TableInfo {
	return &pb.TableInfo{
		Id:         t.Id,
		SmallBlind: int32(t.SmallBlind),
		BigBlind:   int32(t.BigBlind),
		Bet:        int32(t.Bet),
		Timeout:    int32(t.Timeout),
		Cards:      t.Cards.ToProtoMessage(),
		Pot:        t.Pot,
		Chips:      t.Chips,
		Button:     int32(t.Button),
		N:          int32(t.N),
		Max:        int32(t.Max),
		Players:    t.Players.ToProtoMessage(),
	}
}
