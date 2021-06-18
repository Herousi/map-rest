package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.wodcloud.com/apaas-gis/apaas-map-rest/src/pkg/beagle/res"
	"gitlab.wodcloud.com/apaas-gis/apaas-map-rest/src/util/commonutil"
)

// 健康检查
func Health(c *gin.Context) {
	SendJsonResponse(c, res.OK, gin.H{
		"host":       c.Request.Host,
		"header":     c.Request.Header,
		"serverTime": time.Now(),
		"ip":         commonutil.RemoteIp(c.Request),
	})
}

//发送json响应信息
func SendJsonResponse(c *gin.Context, err error, data interface{}) {
	code, message, data := res.DecodeErr(err, data)
	c.JSON(http.StatusOK, res.BgRes{
		Code: code,
		Msg:  message,
		Data: data,
	})
	c.Abort()
}
