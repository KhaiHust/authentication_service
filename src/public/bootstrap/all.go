package bootstrap

import (
	"github.com/KhaiHust/authen_service/adapter/http/client"
	"github.com/KhaiHust/authen_service/adapter/properties"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres"
	service2 "github.com/KhaiHust/authen_service/adapter/service"
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

		//Provide port implementation
		fx.Provide(postgres.NewDatabaseTransactionAdapter),
		fx.Provide(postgres.NewUserRepositoryAdapter),
		fx.Provide(client.NewNotificationServiceAdapter),
		fx.Provide(service2.NewRedisServiceAdapter),

		//Provide usecase
		fx.Provide(usecase.NewDatabaseTransactionUsecase),
		fx.Provide(usecase.NewCreateUserUsecase),
		fx.Provide(usecase.NewGetUserUsecase),
		fx.Provide(usecase.NewSendOtpUseCase),
		fx.Provide(usecase.NewVerifyOtpUseCase),
		fx.Provide(usecase.NewUpdateUserUseCase),

		//Provide helper
		fx.Provide(apihelper.TSCustomValidator),

		//Provide services
		fx.Provide(service.NewUserService),
		fx.Provide(service.NewOtpService),

		//Provide controller
		fx.Provide(controller.NewBaseController),
		fx.Provide(controller.NewUserController),
		fx.Provide(controller.NewOtpController),

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
