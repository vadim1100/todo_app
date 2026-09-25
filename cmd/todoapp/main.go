package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_postgres_pool "github.com/vadim1100/todo_app/internal/core/repository/pool"
	core_service_hash "github.com/vadim1100/todo_app/internal/core/service/hash"
	core_service_jwt "github.com/vadim1100/todo_app/internal/core/service/jwt"
	core_http_middleware "github.com/vadim1100/todo_app/internal/core/transport/http/middleware"
	core_http_server "github.com/vadim1100/todo_app/internal/core/transport/http/server"
	tasks_postgres_repository "github.com/vadim1100/todo_app/internal/features/tasks/repository/postgres"
	tasks_service "github.com/vadim1100/todo_app/internal/features/tasks/service"
	tasks_http_transport "github.com/vadim1100/todo_app/internal/features/tasks/transport"
	users_postgres_repository "github.com/vadim1100/todo_app/internal/features/users/repository/postgres"
	users_service "github.com/vadim1100/todo_app/internal/features/users/service"
	users_transport_http "github.com/vadim1100/todo_app/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	logConfig := core_logger.NewConfigMust()
	logger, err := core_logger.NewLogger(logConfig)

	if err != nil {
		fmt.Println("failed to make logger: %w", err)
		os.Exit(1)
	}

	serverConfig := core_http_server.NewConfigMust()

	postgresConfig := core_postgres_pool.NewConfigMust()

	jwtConfig := core_service_jwt.NewConfigMust()

	ctx, ctxCancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer ctxCancel()

	logger.Debug("initializing postgres connection server")
	pool, err := core_postgres_pool.NewConnectionPool(ctx, postgresConfig)

	if err != nil {
		logger.Fatal("failed to create database pool", zap.Error(err))
	}
	defer pool.Close()

	hasher := core_service_hash.NewBCryptHasher(5)

	jwtManager := core_service_jwt.NewManager(jwtConfig.JWTSecret, jwtConfig.JWTTTL)

	usersRepo := users_postgres_repository.NewUsersRepository(pool)

	usersService := users_service.NewUsersService(usersRepo, hasher, jwtManager)

	usersHandler := users_transport_http.NewUsersHTTPHandler(usersService)

	usersRoutes := usersHandler.Routes()


	tasksRepo := tasks_postgres_repository.NewTasksRepository(pool)

	tasksService := tasks_service.NewTasksService(tasksRepo)

	tasksTransport := tasks_http_transport.NewTasksHandler(tasksService)

	tasksRoutes := tasksTransport.Routes()

	logger.Debug("initializing http server")
	httpServer := core_http_server.NewHTTPServer(
		serverConfig,
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	httpRouter := core_http_server.NewRouter(core_http_middleware.Auth(jwtManager))
	httpRouter.RegisterRoutes(usersRoutes...)
	httpRouter.RegisterRoutes(tasksRoutes...)
	httpServer.RegisterRouters(httpRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("http server run error", zap.Error(err))
	}
}