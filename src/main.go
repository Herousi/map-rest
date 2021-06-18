package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Herousi/map-rest/src/common/client"
	"github.com/Herousi/map-rest/src/common/conf"
	"github.com/Herousi/map-rest/src/pkg/beagle/log"
	"github.com/Herousi/map-rest/src/router"
	"github.com/Herousi/map-rest/src/util/commonutil"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// run args
var (
	port       = pflag.Int("port", 8085, "server port")                                                                                              // 端口
	prefix     = pflag.String("prefix", "/map/rest", "url prefix")                                                                                   // 访问地址
	dbURL      = pflag.String("dbURL", "host=localhost port=1119 user=postgres password=password dbname=apaasgis sslmode=disable", "db connect url") // mysql 数据库连接
	dbDriver   = pflag.String("dbDriver", "postgres", "db driver: mysql postgres oci8")                                                              // mysql 数据库驱动
	redisURL   = pflag.String("redisURL", "localhost:6379", "redis url")                                                                             //  redis 连接
	redisDB    = pflag.Int("redisDB", 0, "redis db")                                                                                                 //  redis 驱动
	pgDbURL    = pflag.String("pgDbURL", "host=localhost port=1119 user=postgres password=password dbname=apaasgis sslmode=disable", "postgis 数据库")  // postgresql 数据库连接
	pgDbDriver = pflag.String("pgDbDriver", "postgres", "postgis 数据库驱动")                                                                             // postgresql 数据库连接
)

// main start
func main() {
	// init start args
	pflag.Parse()
	initConfig()
	// init log config
	cfg := initLogConfig()
	log.NewLog(cfg)
	// pg db client
	go client.GetPgDbClient()
	// redis client
	go client.GetRedisClient()
	// server start...
	zap.L().Error("server start err", zap.Error(newServer().ListenAndServe()))
}

// init commonutil config value
func initConfig() {
	conf.Options = &conf.Config{
		Prefix:     *prefix,
		DbURL:      commonutil.SetEnvStr(*dbURL, "DB_URL"),
		DbDriver:   commonutil.SetEnvStr(*dbDriver, "DB_DRIVER"),
		PgDbURL:    commonutil.SetEnvStr(*pgDbURL, "PG_DB_URL"),
		PgDbDriver: commonutil.SetEnvStr(*pgDbDriver, "PG_DB_DRIVER"),
		RedisURL:   commonutil.SetEnvStr(*redisURL, "REDIS_URL"),
		RedisDB:    *redisDB,
		RedisTag:   "apaasgis",
	}
}

// init log config
func initLogConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func newServer() *http.Server {
	// set run mod
	gin.SetMode(os.Getenv("GIN_MODE"))
	// load gin router
	r := gin.New()
	router.Load(r, log.GinLogger(), log.GinRecovery(true))
	zap.L().Info("server is starting...", zap.Int("port", *port))
	return &http.Server{
		Addr:           fmt.Sprintf(":%d", *port),
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
