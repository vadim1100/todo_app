package main

import (
	"context"
	"fmt"

	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_postgres_pool "github.com/vadim1100/todo_app/internal/core/repository/pool"
	core_service_hash "github.com/vadim1100/todo_app/internal/core/service/hash"
	core_service_jwt "github.com/vadim1100/todo_app/internal/core/service/jwt"
	core_http_middleware "github.com/vadim1100/todo_app/internal/core/transport/http/middleware"
	core_http_server "github.com/vadim1100/todo_app/internal/core/transport/http/server"
	users_repository "github.com/vadim1100/todo_app/internal/features/users/repository/postgres"
	users_service "github.com/vadim1100/todo_app/internal/features/users/service"
	users_transport_http "github.com/vadim1100/todo_app/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	logConfig := core_logger.NewConfigMust()
	logger, err := core_logger.NewLogger(logConfig)

	if err != nil {
		fmt.Println("failed to make logger: %w", err)
	}

	serverConfig := core_http_server.NewConfigMust()

	postgresConfig := core_postgres_pool.NewConfigMust()

	jwtConfig := core_service_jwt.NewConfigMust()

	ctx := context.Background()

	logger.Debug("initializing postgres connection server")
	pool, err := core_postgres_pool.NewConnectionPool(ctx, postgresConfig)

	if err != nil {
		logger.Error("failed to create database pool", zap.Error(err))
		panic(err)
	}
	defer pool.Close()

	hasher := core_service_hash.NewBCryptHasher(5)

	jwtManager := core_service_jwt.NewManager(jwtConfig.JWTSecret, jwtConfig.JWTTTL)

	usersRepo := users_repository.NewUsersRepository(pool)

	usersService := users_service.NewUsersService(usersRepo, hasher, jwtManager)

	usersHandler := users_transport_http.NewUsersHTTPHandler(usersService)

	usersRoutes := usersHandler.Routes()

	logger.Debug("initializing http server")
	httpServer := core_http_server.NewHTTPServer(
		serverConfig,
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	httpRouter := core_http_server.NewRouter()
	httpRouter.RegisterRoutes(usersRoutes...)
	httpServer.RegisterRouters(httpRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("http server run error", zap.Error(err))
	}
}