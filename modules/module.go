package modules

import (
	"books-management-system/config"
	"books-management-system/internal/controllers"
	"books-management-system/internal/repositories/sqlite"
	"books-management-system/internal/router"
	"books-management-system/internal/services"
	"books-management-system/pkg/cache"
	"books-management-system/pkg/kafka"
	"go.uber.org/fx"
)

func RegisterConfig() fx.Option {
	return fx.Invoke(config.InitConfig)
}

func RegisterCache() fx.Option {
	return fx.Provide(func() cache.Cache {
		return cache.NewRedisCache()
	})
}

func RegisterKafka() fx.Option {
	return fx.Provide(func() (*kafka.Producer, error) {
		return kafka.NewKafkaProducer()
	})
}

// RegisterRepositories registers all repositories
func RegisterRepositories() fx.Option {
	return fx.Options(
		fx.Provide(sqlite.NewSQLiteConnection),
		fx.Provide(sqlite.NewSQLiteBookRepository),
	)
}

// RegisterServices registers all services
func RegisterServices() fx.Option {
	return fx.Options(
		fx.Provide(services.NewBookService),
	)
}

func RegisterControllers() fx.Option {
	return fx.Options(
		fx.Provide(
			controllers.NewBookController,
			controllers.NewSwaggerController,

			//			controllers.NewUserController, // ✅ Add new controllers here
		),
		fx.Provide(func(
			bookController *controllers.BookController,
			swaggerController *controllers.SwaggerController,
			//			userController *controllers.UserController,
		) []controllers.Controller {
			return []controllers.Controller{
				bookController,
				swaggerController,
				//				userController,
			}
		}),
	)
}

// docker run -d --name kafka --network kafka-net -p 9092:9092 -e KAFKA_BROKER_ID=1 -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_CFG_LISTENERS=PLAINTEXT://:9092 -e KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true -e ALLOW_PLAINTEXT_LISTENER=yes bitnami/kafka:latest
var Module = fx.Options(
	RegisterConfig(),
	RegisterCache(),
	RegisterKafka(),

	RegisterRepositories(),
	RegisterServices(),
	RegisterControllers(),

	fx.Provide(
		router.NewRouter,
	),
)
