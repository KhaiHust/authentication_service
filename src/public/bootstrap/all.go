package bootstrap

import (
	"github.com/KhaiHust/authen_service/adapter/http/client"
	"github.com/KhaiHust/authen_service/adapter/properties"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres"
	service2 "github.com/KhaiHust/authen_service/adapter/service"
	properties2 "github.com/KhaiHust/authen_service/core/properties"
	"github.com/KhaiHust/authen_service/core/usecase"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/KhaiHust/authen_service/public/controller"
	"github.com/KhaiHust/authen_service/public/router"
	"github.com/KhaiHust/authen_service/public/service"
	"github.com/golibs-starter/golib"
	golibdata "github.com/golibs-starter/golib-data"
	golibgin "github.com/golibs-starter/golib-gin"
	golibsec "github.com/golibs-starter/golib-security"
	"go.uber.org/fx"
)

func All() fx.Option {
	return fx.Options(
		golib.AppOpt(),
		golib.PropertiesOpt(),
		golib.LoggingOpt(),
		golib.EventOpt(),
		golib.BuildInfoOpt(Version, CommitHash, BuildTime),
		golib.ActuatorEndpointOpt(),
		golib.HttpRequestLogOpt(),

		// Http security auto config and authentication filters
		golibsec.HttpSecurityOpt(),
		golibsec.JwtAuthFilterOpt(),
		// Provide datasource auto config
		// redis cache instance
		golibdata.RedisOpt(),
		golibdata.DatasourceOpt(),
		// Provide http client auto config with contextual http client by default,
		// Besides, provide an additional wrapper to easy to control security.
		golib.HttpClientOpt(),
		golibsec.SecuredHttpClientOpt(),

		//Provide config
		golib.ProvideProps(properties.NewNotificationServiceProperties),
		golib.ProvideProps(properties2.NewTokenProperties),

		//Provide port implementation
		fx.Provide(postgres.NewDatabaseTransactionAdapter),
		fx.Provide(postgres.NewUserRepositoryAdapter),
		fx.Provide(client.NewNotificationServiceAdapter),
		fx.Provide(service2.NewRedisServiceAdapter),
		fx.Provide(postgres.NewRefreshTokenRepositoryAdapter),
		fx.Provide(postgres.NewGroupRepositoryAdapter),
		fx.Provide(postgres.NewGroupRoleRepositoryAdapter),
		fx.Provide(postgres.NewGroupMemberRepositoryAdapter),
		fx.Provide(postgres.NewUserProfileRepositoryAdapter),
		fx.Provide(postgres.NewShoppingListRepoAdapter),
		fx.Provide(postgres.NewShoppingListGroupRepoAdapter),
		fx.Provide(postgres.NewShoppingTaskRepoAdapter),
		fx.Provide(postgres.NewCategoryRepoAdapter),
		fx.Provide(postgres.NewUnitRepositoryAdapter),
		fx.Provide(postgres.NewFoodRepositoryAdapter),
		fx.Provide(postgres.NewFridgeItemRepositoryAdapter),
		fx.Provide(postgres.NewMealPlanFoodRepoAdapter),
		fx.Provide(postgres.NewMealPlanRepoAdapter),

		//Provide usecase
		fx.Provide(usecase.NewDatabaseTransactionUsecase),
		fx.Provide(usecase.NewCreateUserUsecase),
		fx.Provide(usecase.NewGetUserUsecase),
		fx.Provide(usecase.NewSendOtpUseCase),
		fx.Provide(usecase.NewVerifyOtpUseCase),
		fx.Provide(usecase.NewUpdateUserUseCase),
		fx.Provide(usecase.NewLoginUserUseCase),
		fx.Provide(usecase.NewGetGroupRoleUsecase),
		fx.Provide(usecase.NewCreateGroupUsecase),
		fx.Provide(usecase.NewAddMemberGroupUsecase),
		fx.Provide(usecase.NewGetGroupUseCase),
		fx.Provide(usecase.NewGetGroupMemberUseCase),
		fx.Provide(usecase.NewRemoveMemberUsecase),
		fx.Provide(usecase.NewCreateShoppingListUseCase),
		fx.Provide(usecase.NewUpdateShoppingListUseCase),
		fx.Provide(usecase.NewDeleteShoppingListUseCase),
		fx.Provide(usecase.NewCreateShoppingTaskUsecase),
		fx.Provide(usecase.NewGetTaskUsecase),
		fx.Provide(usecase.NewGetShoppingListUsecase),
		fx.Provide(usecase.NewDeleteTaskUsecase),
		fx.Provide(usecase.NewUpdateTaskUsecase),
		fx.Provide(usecase.NewGetCategoryUsecase),
		fx.Provide(usecase.NewGetUnitUsecase),
		fx.Provide(usecase.NewCreateFoodUsecase),
		fx.Provide(usecase.NewGetFoodUseCase),
		fx.Provide(usecase.NewUpdateFoodUseCase),
		fx.Provide(usecase.NewDeleteFoodUseCase),
		fx.Provide(usecase.NewCreateFridgeItemUsecase),
		fx.Provide(usecase.NewGetFridgeItemUsecase),
		fx.Provide(usecase.NewUpdateFridgeItemUsecase),
		fx.Provide(usecase.NewCreateMealPlanUsecase),
		fx.Provide(usecase.NewUpdateMealPlanUsecase),
		fx.Provide(usecase.NewGetMealPlanUsecase),
		fx.Provide(usecase.NewGetUserProfileUseCase),
		fx.Provide(usecase.NewDeleteFridgeItemUsecase),
		fx.Provide(usecase.NewDeleteMealPlanUsecase),

		//Provide helper
		fx.Provide(apihelper.TSCustomValidator),

		//Provide services
		fx.Provide(service.NewUserService),
		fx.Provide(service.NewOtpService),
		fx.Provide(service.NewGroupService),
		fx.Provide(service.NewShoppingListService),
		fx.Provide(service.NewShoppingTaskService),
		fx.Provide(service.NewCategoryService),
		fx.Provide(service.NewUnitService),
		fx.Provide(service.NewFoodService),
		fx.Provide(service.NewFridgeItemService),
		fx.Provide(service.NewMealPlanService),

		//Provide controller
		fx.Provide(controller.NewBaseController),
		fx.Provide(controller.NewUserController),
		fx.Provide(controller.NewOtpController),
		fx.Provide(controller.NewGroupController),
		fx.Provide(controller.NewShoppingListController),
		fx.Provide(controller.NewShoppingTaskController),
		fx.Provide(controller.NewCategoryController),
		fx.Provide(controller.NewUnitController),
		fx.Provide(controller.NewFoodController),
		fx.Provide(controller.NewFridgeController),
		fx.Provide(controller.NewMealPlanController),

		// Provide gin http server auto config,
		// actuator endpoints and application routers
		golibgin.GinHttpServerOpt(),
		fx.Invoke(router.RegisterGinRouters),

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		golibgin.OnStopHttpServerOpt(),
		//golibmsg.OnStopProducerOpt(),
	)
}
