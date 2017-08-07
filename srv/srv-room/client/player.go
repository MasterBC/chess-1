package main

import (
	pb "chess/srv/srv-room/proto"
	"fmt"
	"github.com/golang/protobuf/proto"

	"os"
	"bufio"
	"strings"
	"strconv"
	"chess/common/log"
	"chess/srv/srv-room/misc/packet"
	//"time"
	"time"
)

var (
	Levels = map[int32]string{
		1: "高牌",
		2: "一对",
		3: "两对",
		4: "三条",
		5: "顺子",
		6: "同花",
		7: "葫芦",
		8: "四条",
		9: "同花顺",
		10: "皇家同花顺",
	}

	RANKNAME = []string{"2","3","4","5","6","7","8","9","10","J","Q","K","A"}
	SUITNAME = []string{"黑桃", "红桃", "梅花", "方块"}

	Actions = map[string]string{
		"call":"跟注",
		"check":"让牌",
		"raise":"加注",
		"fold":"弃牌",
		"allin":"全押",
	}
)


type Player struct {
	pb.PlayerInfo

	HandLevel int32
	HandFinalValue int32

	Table *pb.TableInfo

	Stream pb.RoomService_StreamClient // 后端游戏服数据流 - 双向流
	Die     chan struct{}
	autobet bool
}

func NewPlayer() *Player {
	p := &Player{}
	p.autobet = false;
	return p
}

func (p *Player) SendMessage(tos int16, msg proto.Message) error {
	return p.Stream.Send(&pb.Room_Frame{
		Type:    pb.Room_Message,
		Message: packet.Pack(tos, msg),
	})
}

func (p *Player)HandleMQ(tos int16, data []byte) {
	switch tos {
	case 40:
		ack := &pb.KickedOutAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}
		fmt.Print("您的账号在另一地点登录！！\n",)

	case 2006: // 加入游戏，获取牌桌信息
		ack := &pb.RoomGetTableAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}

		fmt.Printf("成功加入牌桌(%s)，当前牌桌玩家信息：\n", ack.Table.Id)

		p.Table = ack.Table
		for _, v := range ack.Table.Players {
			if v.Id == p.Id {
				p.Pos = v.Pos
				p.Nickname  = v.Nickname
				p.Avatar  = v.Avatar
				p.Level  = v.Level
				p.Chips   = v.Chips
				p.Bet     = v.Bet
				p.Action  = v.Action
			}
			if v.Id != 0 {
				fmt.Printf("玩家%d: 位置:%d 筹码:%d  \n", v.Id, v.Pos, v.Chips)
			}
		}
		//fmt.Print("poker>")
	case 2102: // 其他玩家加入
		ack := &pb.RoomPlayerJoinAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}

		if p.Table != nil {
			p.Table.Players[ack.Player.Pos-1] = ack.Player
			p.Table.N++
		}

		fmt.Printf("\n\t玩家%d加入牌桌\n", ack.Player.Id)
		//fmt.Print("poker>")

	case 2104: // 玩家离开牌桌
		ack := &pb.RoomPlayerGoneAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}
		if p.Table != nil {
			p.Table.Players[ack.Player.Pos-1] = nil
			p.Table.N--
		}

		fmt.Printf("\n\t玩家%d离开牌桌\n", ack.Player.Id)
		fmt.Print("poker>")

	case 2107: // 游戏开始
		fmt.Println("==============================================")
		fmt.Println("                 游戏开始")
		fmt.Println("==============================================")

		ack := &pb.RoomButtonAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}

		p.Table.Button = ack.ButtonPos
		p.Table.Id = ack.TableId
		p.Table.Bet = 0
		p.Table.Cards = nil
		p.Table.Pot = make([]int32, 1)
		for k,v := range p.Table.Players {
			if v != nil {
				p.Table.Players[k].Bet = 0
				p.Table.Players[k].Action = ""
				p.Table.Players[k].Cards = nil
				//p.Table.Players[k]. = 0
			}
		}

		dealer := p.Table.Players[ack.ButtonPos-1]
		fmt.Printf("\t庄家ID: %d.\n", dealer.Id)
		//fmt.Print("poker>")

	case 2108: // 发底牌 翻牌 转牌 河牌
		ack := &pb.RoomDealAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}

		switch ack.Action {
		case "preflop": // 发底牌
			p.Cards = ack.Cards
			if len(p.Cards) == 2 {
				fmt.Printf("[手牌]: %s %s.\n",
					SUITNAME[p.Cards[0].Suit]+RANKNAME[p.Cards[0].Val],
					SUITNAME[p.Cards[1].Suit]+RANKNAME[p.Cards[1].Val],
				)
			}

			//fmt.Print("poker>")
		case "flop": //  翻牌
			for k,v := range p.Table.Players {
				if v != nil {
					p.Table.Players[k].Bet = 0
					p.Table.Players[k].Action = ""
				}
			}
			p.Table.Cards = ack.Cards
			p.HandLevel = ack.HandLevel
			p.HandFinalValue = ack.HandFinalValue
			fmt.Printf("[翻牌]: %s %s %s.   ",
				SUITNAME[p.Table.Cards[0].Suit]+RANKNAME[p.Table.Cards[0].Val],
				SUITNAME[p.Table.Cards[1].Suit]+RANKNAME[p.Table.Cards[1].Val],
				SUITNAME[p.Table.Cards[2].Suit]+RANKNAME[p.Table.Cards[2].Val],
			)
			if len(p.Cards) == 2 {
				fmt.Printf("你的手牌: %s %s， 牌型:%s.\n",
					SUITNAME[p.Cards[0].Suit]+RANKNAME[p.Cards[0].Val],
					SUITNAME[p.Cards[1].Suit]+RANKNAME[p.Cards[1].Val],
					Levels[ack.HandLevel],
				)
			}

			//fmt.Print("poker>")
		case "turn": // 转牌
			for k,v := range p.Table.Players {
				if v != nil {
					p.Table.Players[k].Bet = 0
					p.Table.Players[k].Action = ""
				}
			}
			p.Table.Cards = ack.Cards
			p.HandLevel = ack.HandLevel
			p.HandFinalValue = ack.HandFinalValue
			fmt.Printf("[转牌]: %s %s %s %s.    ",
				SUITNAME[p.Table.Cards[0].Suit]+RANKNAME[p.Table.Cards[0].Val],
				SUITNAME[p.Table.Cards[1].Suit]+RANKNAME[p.Table.Cards[1].Val],
				SUITNAME[p.Table.Cards[2].Suit]+RANKNAME[p.Table.Cards[2].Val],
				SUITNAME[p.Table.Cards[3].Suit]+RANKNAME[p.Table.Cards[3].Val],
			)
			if len(p.Cards) == 2 {
				fmt.Printf("你的手牌: %s %s， 牌型:%s.\n",
					SUITNAME[p.Cards[0].Suit]+RANKNAME[p.Cards[0].Val],
					SUITNAME[p.Cards[1].Suit]+RANKNAME[p.Cards[1].Val],
					Levels[ack.HandLevel],
				)
			}
			//fmt.Print("poker>")
		case "river": // 河牌
			for k,v := range p.Table.Players {
				if v != nil {
					p.Table.Players[k].Bet = 0
					p.Table.Players[k].Action = ""
				}
			}
			p.Table.Cards = ack.Cards
			p.HandLevel = ack.HandLevel
			p.HandFinalValue = ack.HandFinalValue

			fmt.Printf("[河牌]: %s %s %s %s %s.    ",
				SUITNAME[p.Table.Cards[0].Suit]+RANKNAME[p.Table.Cards[0].Val],
				SUITNAME[p.Table.Cards[1].Suit]+RANKNAME[p.Table.Cards[1].Val],
				SUITNAME[p.Table.Cards[2].Suit]+RANKNAME[p.Table.Cards[2].Val],
				SUITNAME[p.Table.Cards[3].Suit]+RANKNAME[p.Table.Cards[3].Val],
				SUITNAME[p.Table.Cards[4].Suit]+RANKNAME[p.Table.Cards[4].Val],
			)
			if len(p.Cards) == 2 {
				fmt.Printf("你的手牌: %s %s， 牌型:%s.\n",
					SUITNAME[p.Cards[0].Suit] + RANKNAME[p.Cards[0].Val],
					SUITNAME[p.Cards[1].Suit] + RANKNAME[p.Cards[1].Val],
					Levels[ack.HandLevel],
				)
			}
			//fmt.Print("poker>")
		}

	case  2111: // 摊牌和比牌
		ack := &pb.RoomShowdownAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}
		fmt.Print("[比牌]:\n")
		for _, v := range ack.Table.Players {
			if v.Id != 0 {
				if len(ack.Table.Cards) < 3 {
					fmt.Printf("玩家%d: 赢得%d筹码\n", v.Id, ack.Table.Chips[v.Pos-1])
				} else {
					fmt.Printf("玩家%d: %s, 赢得%d筹码\n", v.Id, Levels[v.HandLevel], ack.Table.Chips[v.Pos-1])
				}
			}
		}

		//fmt.Print("poker>")
	case 2110: // 通报当前下注玩家
		ack := &pb.RoomActionAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}

		p.Table.Bet = ack.BaseBet
		if p.Table.Players[ack.Pos-1].Id == p.Id {
/*			fmt.Printf("\n你的手上剩余筹码：%d\n", p.Table.Players[ack.Pos - 1].Chips)
			fmt.Printf("你的本轮下注筹码：%d\n", p.Table.Players[ack.Pos - 1].Bet)
			fmt.Printf("上一玩家本轮下注筹码：%d\n", p.Table.Bet)*/
			fmt.Printf("\n[当前玩家广播]\n")
			fmt.Printf("\t你的剩余筹码:%d,最少跟注:%d\n",p.Table.Players[ack.Pos - 1].Chips,
				p.Table.Bet-p.Table.Players[ack.Pos - 1].Bet)
			if p.Table.Bet == 0 && p.Table.Players[ack.Pos - 1].Bet == 0 {
				fmt.Printf("\t该你下注了: (-1:弃牌 0:让牌 大于0:加注 %d:allin)\n", p.Table.Players[ack.Pos - 1].Chips)
			}
			if p.Table.Bet != 0 && p.Table.Bet == p.Table.Players[ack.Pos - 1].Bet {
				// 可让牌  可加注
				fmt.Printf("\t该你下注了: (-1:弃牌 0:让牌 大于0:加注 %d:allin)\n", p.Table.Players[ack.Pos - 1].Chips)
			}
			if p.Table.Bet != 0 && p.Table.Bet > p.Table.Players[ack.Pos - 1].Bet {
				// 不可让牌  可加注 可跟注
				fmt.Printf("\t该你下注了: (-1:弃牌 %d:跟注 大于%d:加注 %d:allin)\n",
					p.Table.Bet - p.Table.Players[ack.Pos - 1].Bet,
					p.Table.Bet - p.Table.Players[ack.Pos - 1].Bet,
					p.Table.Players[ack.Pos - 1].Chips,
				)
			}

			//一直跟的策略
			bet := p.Table.Bet - p.Table.Players[ack.Pos - 1].Bet
			if bet < 0 {
				bet = p.Table.Players[ack.Pos - 1].Bet
			}
			//下注
			if p.autobet  {
				st:=time.Now().Second()%10+3
				time.Sleep(time.Second*time.Duration(st))
				fmt.Printf("----------->%d秒后,自动下注: %d\n\n",st,bet)
				p.SendMessage(2105, &pb.RoomPlayerBetReq{
					BaseReq: &pb.BaseReq{AppFrom:"CMD"},
					TableId: p.Table.Id,
					Bet: int32(bet),
				})
			}else {
				fmt.Print("poker>")
			}



		} else {
			fmt.Printf("\n[当前玩家广播]\n")
			fmt.Printf("\t当前玩家,ID：%d,位置:%d\n",p.Table.Players[ack.Pos-1].Id,ack.Pos)
		}


	case 2109: // 通报奖池
		ack := &pb.RoomPotAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}

		p.Table.Pot = ack.Pot
		fmt.Printf("\n[当前奖池]: %v\n", ack.Pot)
		//fmt.Print("poker>")
	case 2106: // 下注结果
		ack := &pb.RoomPlayerBetAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}
		if p.Table.Players[ack.Pos - 1] != nil {
			p.Table.Players[ack.Pos - 1].Action = ack.Action
			p.Table.Players[ack.Pos - 1].Bet = ack.Bet
			p.Table.Players[ack.Pos - 1].Chips = ack.Chips
		}
		fmt.Printf("[下注结果广播]\n")
		if ack.Pos == p.Pos {
			p.Action = ack.Action
			p.Bet = ack.Bet
			p.Chips = ack.Chips
			fmt.Printf("\t(You)   玩家%d %s, 当前下注筹码 %d, 手上筹码 %d\n",
				p.Table.Players[ack.Pos - 1].Id,
				Actions[p.Table.Players[ack.Pos - 1].Action],
				p.Table.Players[ack.Pos - 1].Bet,
				p.Table.Players[ack.Pos - 1].Chips,
			)
		} else {
			fmt.Printf("\t玩家%d %s, 当前下注筹码 %d, 手上筹码 %d\n",
				p.Table.Players[ack.Pos - 1].Id,
				Actions[p.Table.Players[ack.Pos - 1].Action],
				p.Table.Players[ack.Pos - 1].Bet,
				p.Table.Players[ack.Pos - 1].Chips,
			)
		}
		//fmt.Print("poker>")
	case 2113: // 站起通知
		ack := &pb.RoomPlayerStandupAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}
		fmt.Printf("\n[玩家站起广播]\n")
		fmt.Printf("\t玩家%d站起\n",ack.PlayerId)

	case 2115: // 坐下通知
		ack := &pb.RoomPlayerSitdownAck{}
		err := proto.Unmarshal(data, ack)
		if err != nil {
			log.Errorf("proto.Unmarshal Error: %s", err)
			return
		}

		fmt.Printf("\n[玩家坐下广播]\n")
		fmt.Printf("\t玩家%d坐下，位置：%d\n",ack.PlayerId, ack.PlayerPos)
	}

}


func (p *Player)CmdLoop() {
	fmt.Println("Welcome to the Texas Hold'em game!!!")
	fmt.Println("指令提示：")
	fmt.Println("    j - 加入牌桌 [ Ex: j1-6/a1-6 ]")
	fmt.Println("    l - 离开牌桌")
	fmt.Println("    u - 站起")
	fmt.Println("    d - 坐下")
	fmt.Println("    c - 查看手牌和公共牌")
	fmt.Println("    q - 退出命令行")
	fmt.Println("    h - 帮助")

	fmt.Printf("用户ID(%d)\n", p.Id)
	fmt.Print("poker>")

	reader := bufio.NewReader(os.Stdin)
	for {

		cmd, _ := reader.ReadString('\n')
		cmd = strings.ToLower(strings.Trim(cmd, " \n"))

		if len(cmd) == 0 {
			continue
		}
		switch cmd[0] {
		case 'a':
			p.autobet = true
			fallthrough
		case 'j': // 加入游戏
			if p.Table == nil {
				var roomid int32
				roomid =1
				if len(cmd) >1 {
					id,e := strconv.ParseInt(cmd[1:2], 10, 32)
					if e!=nil {
						fmt.Println(e)
					}
					roomid = int32(id)
					//fmt.Println(cmd[1:2])
					//fmt.Printf("join room: %d,%d\n",roomid,id)
				}
				p.SendMessage(2101, &pb.RoomPlayerJoinReq{
					BaseReq: &pb.BaseReq{AppFrom:"CMD"},
					RoomId: roomid,  // 房间id
					TableId: "", // 牌桌id为空，自动选择
				})
			} else {
				fmt.Print("poker>")
			}
		case 'u': // 站起
			if p.Table != nil {
				p.SendMessage(2112, &pb.RoomPlayerStandupReq{
					BaseReq: &pb.BaseReq{AppFrom:"CMD"},
					TableId: p.Table.Id,
				})
			}
		case 'd': // 坐下
			if p.Table != nil {
				p.SendMessage(2114, &pb.RoomPlayerSitdownReq{
					BaseReq: &pb.BaseReq{AppFrom:"CMD"},
					TableId: p.Table.Id,
				})
			}
		case 'l': // 离开游戏
			if p.Table != nil {
				p.SendMessage(2103, &pb.RoomPlayerGoneReq{
					BaseReq: &pb.BaseReq{AppFrom:"CMD"},
					TableId: p.Table.Id,
				})
			}
			p.Pos = 0
			p.Table = nil
		case 'c':
			if p.Table != nil {
				cards := []*pb.CardInfo{}
				cards = append(cards, p.Cards...)
				cards = append(cards, p.Table.Cards...)
				fmt.Println("2张手牌+公共牌信息:")
				for  _,v := range cards {
					fmt.Println(SUITNAME[v.Suit]+RANKNAME[v.Val])
				}

				//玩家信息
				for _, v := range p.Table.Players {
					if v.Id == p.Id {
						p.Pos = v.Pos
						p.Nickname  = v.Nickname
						p.Avatar  = v.Avatar
						p.Level  = v.Level
						p.Chips   = v.Chips
						p.Bet     = v.Bet
						p.Action  = v.Action
					}
					if v.Id != 0 {
						fmt.Printf("玩家%d: 位置:%d 筹码:%d  \n", v.Id, v.Pos, v.Chips)
					}
				}
			}
			fmt.Print("poker>")
		case 'q':
			fmt.Println("Goodbye!")
			return
		case 'h':
			fmt.Println("指令提示：")
			fmt.Println("    j - 加入牌桌 [ Ex: j1-6/a1-6 ]")
			fmt.Println("    l - 离开牌桌")
			fmt.Println("    u - 站起")
			fmt.Println("    d - 坐下")
			fmt.Println("    c - 查看手牌和公共牌")
			fmt.Println("    q - 退出命令行")
			fmt.Println("    h - 帮助")
			fmt.Print("poker>")
		default: // 下注
			if p.Table != nil {
				bet, err := strconv.ParseInt(cmd, 10, 32)
				if err != nil {
					fmt.Println("下注数额有误！")
					fmt.Println("poker>")
					continue
				}

				p.SendMessage(2105, &pb.RoomPlayerBetReq{
					BaseReq: &pb.BaseReq{AppFrom:"CMD"},
					TableId: p.Table.Id,
					Bet: int32(bet),
				})
			} else {
				fmt.Print("poker>")
			}

		}

		//time.Sleep(1*time.Second)
	}
}