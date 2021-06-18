package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.wodcloud.com/apaas-gis/apaas-map-rest/src/common/conf"
	"gitlab.wodcloud.com/apaas-gis/apaas-map-rest/src/controller"
	"gitlab.wodcloud.com/apaas-gis/apaas-map-rest/src/router/middleware/header"
)

// 加载路由
func Load(r *gin.Engine, middleware ...gin.HandlerFunc) {
	r.Use()
	r.Use(gin.Recovery())
	r.Use(header.NoCache)
	r.Use(header.Options)
	r.Use(header.Secure)
	r.Use(middleware...)
	base := r.Group(conf.Options.Prefix)
	{
		base.GET("/health", controller.Health) // 健康检查
	}
}
