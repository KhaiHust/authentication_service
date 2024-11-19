package router

import (
	"github.com/KhaiHust/authen_service/public/controller"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib/web/actuator"
	"go.uber.org/fx"
)

type RegisterRoutersIn struct {
	fx.In
	App            *golib.App
	Engine         *gin.Engine
	Actuator       *actuator.Endpoint
	UserController *controller.UserController
	OtpController  *controller.OtpController
}

func RegisterGinRouters(p RegisterRoutersIn) {
	group := p.Engine.Group(p.App.Path())
	group.GET("/actuator/health", gin.WrapF(p.Actuator.Health))
	group.GET("/actuator/info", gin.WrapF(p.Actuator.Info))

	router := p.Engine.Group(p.App.Path())
	userV1 := router.Group("/public/v1/user")
	{
		userV1.POST("", p.UserController.RegisterUser)
		userV1.POST("/send-verification-code", p.OtpController.SendOtpForRegistration)
		userV1.POST("/verify-email", p.OtpController.VerifyOtpForRegistration)
	}

}
