package wechat

type UserInfo struct {
	OpenId   string `json:"openid"`   // 用户的唯一标识
	Nickname string `json:"nickname"` // 用户昵称
	Sex      int    `json:"sex"`      // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	City     string `json:"city"`     // 普通用户个人资料填写的城市
	Province string `json:"province"` // 用户个人资料填写的省份
	Country  string `json:"country"`  // 国家, 如中国为CN

	// 用户头像, 最后一个数值代表正方形头像大小(有0, 46, 64, 96, 132数值可选, 0代表640*640正方形头像),
	// 用户没有头像时该项为空
	HeadImageURL string `json:"headimgurl"`

	// 用户特权信息, json 数组, 如微信沃卡用户为(chinaunicom)
	Privilege []string `json:"privilege"`

	// 用户统一标识. 针对一个微信开放平台帐号下的应用, 同一用户的unionid是唯一的.
	UnionId string `json:"unionid"`
}

type JsApiTicket struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	Ticket   string `json:"ticket"`
	ExpireIn int    `json:"expires_in"`
}
