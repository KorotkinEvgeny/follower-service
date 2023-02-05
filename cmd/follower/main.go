package main

import (
	"fmt"
	"github.com/follower-service/pkg/api"
	"github.com/follower-service/pkg/api/handler"
	"github.com/follower-service/pkg/config/viper"
	"github.com/follower-service/pkg/repository/postgres"
	"github.com/follower-service/pkg/server"
	"github.com/follower-service/pkg/service/follow"
	"github.com/follower-service/pkg/service/user"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	cfg, err := viper.ReadAndReturn()
	if err != nil {
		log.Fatalf("Can't read config %s", err.Error())
	}

	apiHealthHandler := handler.NewAPIHealthHandler()
	validate := validator.New()

	connMaxLifetime := time.Duration(cfg.GetInt("db.conn-max-lifetime-s")) * time.Second

	dbConnection, err := makeConnection(&connectionCfg{
		User:            cfg.GetString("db.username"),
		Pass:            cfg.GetString("db.password"),
		Host:            cfg.GetString("db.host"),
		Port:            cfg.GetInt("db.port"),
		DB:              cfg.GetString("db.name"),
		MaxOpenConns:    cfg.GetInt("db.max-open-conns"),
		MaxIdleConns:    cfg.GetInt("db.max-idle-conns"),
		ConnMaxLifeTime: connMaxLifetime,
	})

	if err != nil {
		panic(err.Error())
	}

	if cfg.GetBool("debug") {
		err = postgres.Migrate(dbConnection)
		if err != nil {
			panic(err)
		}
	}

	userRepo := postgres.NewUserRepository(dbConnection)
	userService := user.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, validate)

	followerRepo := postgres.NewFollowRepository(dbConnection)
	followService := follow.NewFollowService(followerRepo)
	followHandler := handler.NewFollowHandler(followService, validate, userService)

	newRouter := api.NewChiRouter(apiHealthHandler, followHandler, userHandler)

	host := cfg.GetString("server.host")
	port := cfg.GetString("server.port")
	readTimeout := cfg.GetDuration("server.read_timeout")
	writeTimeout := cfg.GetDuration("server.write_timeout")
	address := fmt.Sprintf("%v:%v", host, port)

	serverConfig := server.Config{
		Addr:         address,
		ReadTimeOut:  readTimeout,
		WriteTimeOut: writeTimeout,
	}
	srv := server.NewHttpServer(newRouter, serverConfig)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error %s", err.Error())
	}
}

type connectionCfg struct {
	User            string
	Pass            string
	Host            string
	Port            int
	DB              string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifeTime time.Duration
}

func makeConnection(cfg *connectionCfg) (*sqlx.DB, error) {
	dbConnectString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	db, err := sqlx.Open("postgres", dbConnectString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifeTime)
	log.Infof("max open conns: %s, max idle conns: %s", cfg.MaxOpenConns, cfg.MaxIdleConns)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
