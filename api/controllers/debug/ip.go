package c_debug

import (
	"github.com/gin-gonic/gin"
	"chess/common/define"
	"chess/common/helper"
)

type IpResult struct {
	define.BaseResult
	Ip string `json:"ip"`
}

func IP(c *gin.Context) {
	var result IpResult

	clientIp := helper.ClientIP(c)

	result.Ret = 1
	result.Ip = clientIp

	helper.EchoResult(c, result)
}

func Config(c *gin.Context) {
	if config, ok := c.Get("config"); ok {
		helper.EchoResult(c, config)
	}
}
