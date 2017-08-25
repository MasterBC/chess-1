package models

type TaskPriceReceiveModel struct {
	Id         int
	UserId     int
	TaskId     int
	RewardType int
	RewardNum  int
}

var TaskPriceReceive = new(TaskPriceReceiveModel)

func (m *TaskPriceReceiveModel) Insert() error {
	sqlStr := `INSERT INTO task_prize_receive(user_id,task_id,reward_type,reward_num) VALUES(?,?,?,?)`
	_, err := Mysql.Chess.Exec(sqlStr, m.UserId, m.TaskId, m.RewardType, m.RewardNum)
	return err
}
