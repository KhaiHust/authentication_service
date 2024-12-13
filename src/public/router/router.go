package router

import (
	"github.com/KhaiHust/authen_service/public/controller"
	"github.com/KhaiHust/authen_service/public/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib-security/web/config"
	"github.com/golibs-starter/golib/web/actuator"
	"go.uber.org/fx"
)

type RegisterRoutersIn struct {
	fx.In
	App                    *golib.App
	Engine                 *gin.Engine
	Actuator               *actuator.Endpoint
	UserController         *controller.UserController
	OtpController          *controller.OtpController
	SecurityProperties     *config.HttpSecurityProperties
	GroupController        *controller.GroupController
	ShoppingListController *controller.ShoppingListController
	ShoppingTaskController *controller.ShoppingTaskController
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
		userV1.POST("/login", p.UserController.LoginUser)
	}
	userV1.Use(middleware.GetInfoFromToken(p.SecurityProperties.Jwt))
	{
		userV1.POST("/group", p.GroupController.CreateGroup)
		userV1.POST("/group/add", p.GroupController.AddMember)
		userV1.GET("/group/:groupID/members", p.GroupController.GetListMember)
		userV1.DELETE("/group", p.GroupController.RemoveMember)
	}
	shoppingV1 := router.Group("/public/v1/shopping", middleware.GetInfoFromToken(p.SecurityProperties.Jwt))
	{
		shoppingV1.POST("", p.ShoppingListController.CreateNewShoppingList)
		shoppingV1.PUT("/:shoppingListId", p.ShoppingListController.UpdateShoppingList)
		shoppingV1.DELETE("/:shoppingListId", p.ShoppingListController.DeleteShoppingList)
		shoppingV1.POST("/:shoppingListId/tasks", p.ShoppingTaskController.CreateNewShoppingTask)
	}
}
