package models

import (
	"chess/common/log"
	"github.com/Sirupsen/logrus"
)

var UsersTp = new(UsersTpModel)

type UsersTpModel struct {
	Id         int    `json:"id"`
	Type       string `json:"type"`
	UserID     int    `json:"user_id"`
	OpenID     string `json:"openid"`
	WxUnionId  string `json:"wx_union_id"`
	WxH5OpenId string `json:"wx_h5_openid"`
}

// 是否已经登录过
func (m *UsersTpModel) IsReg(openid, Type string) (UserId int, err error) {
	sqlString := `SELECT user_id,openid
				FROM users_tp WHERE openid = ? AND type = ?`
	err = Mysql.Chess.QueryRow(
		sqlString,
		openid,
		Type,
	).Scan(&UserId, &openid)
	return
}

// 检查微信 union id是否存在
func (m *UsersTpModel) CheckWxUnionId(wxUnionId, Type string) (id, UserId int, err error) {
	sqlString := `SELECT id,user_id,openid
				FROM users_tp WHERE wx_union_id = ? AND type = ?`
	err = Mysql.Chess.QueryRow(
		sqlString,
		wxUnionId,
		Type,
	).Scan(&id, &UserId, &wxUnionId)
	return
}

// 插入
func (m *UsersTpModel) Insert(user *UsersTpModel) (int, error) {
	sqlString := `INSERT INTO users_tp
		(type,user_id,openid,wx_union_id,wx_h5_openid)
		VALUES
		(?, ?, ?, ?,?)`
	result, err := Mysql.Chess.Exec(
		sqlString,
		user.Type,
		user.UserID,
		user.OpenID,
		user.WxUnionId,
		user.WxH5OpenId,
	)

	// Debug
	log.Debugf("UsersTpModel.Insert", logrus.Fields{
		"sql":   sqlString,
		"Error": err,
	})

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	return int(id), err
}

// update wx union id
func (m *UsersTpModel) UpdateWxUnionIdByUid(uid int, wxUnionId string) error {
	sqlString := `UPDATE users_tp SET wx_union_id = ? WHERE user_id = ?`
	_, err := Mysql.Chess.Exec(
		sqlString, wxUnionId, uid,
	)

	// Debug
	log.Debugf("UsersTpModelUpdateWxUnionId", logrus.Fields{
		"sql":   sqlString,
		"Error": err,
	})
	return err
}

// update openid by id
func (m *UsersTpModel) UpdateOpenidIdById(id int, openid string) error {
	sqlString := `UPDATE users_tp SET openid = ? WHERE id = ?`
	_, err := Mysql.Chess.Exec(
		sqlString, openid, id,
	)
	// Debug
	log.Debugf("UsersTpModelUpdateWxUnionId", logrus.Fields{
		"sql":   sqlString,
		"Error": err,
	})
	return err
}

func (m *UsersTpModel) GetByUid(uid int, user *UsersTpModel) error {
	sqlString := `SELECT
					id,type,user_id,openid,wx_union_id
				FROM users_tp WHERE user_id = ?`
	return Mysql.Chess.QueryRow(
		sqlString, uid,
	).Scan(
		&user.Id,
		&user.Type,
		&user.UserID,
		&user.OpenID,
		&user.WxUnionId,
	)
}

func (m *UsersTpModel) GetByMobile(mobile string, user *UsersTpModel) error {
	sqlString := `SELECT a.id,a.type,a.user_id,a.openid,a.wx_union_id
				FROM users_tp AS a,users AS b
				WHERE a.user_id = b.id  AND b.mobile_number = ?`
	return Mysql.Chess.QueryRow(
		sqlString, mobile,
	).Scan(
		&user.Id,
		&user.Type,
		&user.UserID,
		&user.OpenID,
		&user.WxUnionId,
	)
}

func (m *UsersTpModel) GetId(id int, user *UsersTpModel) error {
	sqlString := `SELECT
					id,type,user_id,openid,wx_union_id
				FROM users_tp WHERE id = ?`
	return Mysql.Chess.QueryRow(
		sqlString, id,
	).Scan(
		&user.Id,
		&user.Type,
		&user.UserID,
		&user.OpenID,
		&user.WxUnionId,
	)
}

func (m *UsersTpModel) GetOpenid(id int) (openid string, err error) {
	sqlStr := `SELECT openid FROM users_tp WHERE user_id = ?`
	err = Mysql.Chess.QueryRow(
		sqlStr, id,
	).Scan(&openid)

	return
}
