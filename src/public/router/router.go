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
	CategoryController     *controller.CategoryController
	UnitController         *controller.UnitController
	FoodController         *controller.FoodController
	FridgeController       *controller.FridgeController
	MealPlanController     *controller.MealPlanController
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
		userV1.POST("/refresh-token", p.UserController.RefreshToken)
		userV1.POST("/logout", middleware.GetInfoFromToken(p.SecurityProperties.Jwt), p.UserController.Logout)
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
		shoppingV1.GET("/:shoppingListId/tasks", p.ShoppingTaskController.GetShoppingTasksByShoppingListID)
		shoppingV1.DELETE("/:shoppingListId/tasks/:taskId", p.ShoppingTaskController.DeleteTaskByID)
		shoppingV1.PUT("/:shoppingListId/tasks/:taskId", p.ShoppingTaskController.UpdateTaskByID)
	}
	foodV1 := router.Group("/public/v1/food", middleware.GetInfoFromToken(p.SecurityProperties.Jwt))
	{
		foodV1.GET("/category", p.CategoryController.GetAllCategory)
		foodV1.GET("/unit", p.UnitController.GetAllUnits)
		foodV1.POST("", p.FoodController.CreateFood)
		foodV1.PUT("/:foodId", p.FoodController.UpdatedFood)
		foodV1.DELETE("/:foodId", p.FoodController.DeleteFood)
		foodV1.GET("", p.FoodController.GetAllFood)
	}
	fridgeV1 := router.Group("/public/v1/fridge", middleware.GetInfoFromToken(p.SecurityProperties.Jwt))
	{
		fridgeV1.POST("", p.FridgeController.CreateFridgeItem)
		fridgeV1.PUT("/:itemId", p.FridgeController.UpdateFridgeItem)
		fridgeV1.DELETE("/:itemId", p.FridgeController.DeleteFridgeItem)
		fridgeV1.GET("/:itemId", p.FridgeController.DeleteFridgeItem)
		fridgeV1.GET("", p.FridgeController.GetAllFridgeItems)
	}
	mealPlanV1 := router.Group("/public/v1/meal-plan", middleware.GetInfoFromToken(p.SecurityProperties.Jwt))
	{
		mealPlanV1.POST("", p.MealPlanController.CreateNewMealPlan)
		mealPlanV1.PUT("/:mealId", p.MealPlanController.UpdateMealPlan)
		mealPlanV1.DELETE("/:mealId", p.MealPlanController.DeleteMealPlan)
		mealPlanV1.GET("", p.MealPlanController.GetPlanByDate)
		mealPlanV1.GET("/:mealId", p.MealPlanController.GetDetailMealPlan)
	}
}
