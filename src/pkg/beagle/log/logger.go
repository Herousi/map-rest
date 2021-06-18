package log

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	Config zapcore.EncoderConfig
}

func NewLog(config zapcore.EncoderConfig) *Log {
	log := new(Log)
	log.Config = config
	log.initLogger()
	return log
}

// InitLogger 初始化Logger
func (l *Log) initLogger() {
	var core zapcore.Core
	encoder := zapcore.NewJSONEncoder(l.Config) // 日志的格式相关
	core = zapcore.NewTee(
		// 创建一个将debug级别以上的日志输出到终端的配置信息
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
	)
	lg := zap.New(core, zap.AddCaller()) // 根据上面的配置创建logger
	zap.ReplaceGlobals(lg)               // 替换zap库里全局的logger
	return
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info("visit",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func (l *Log) Debugf(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	zap.L().Debug(s, f...)
}

func (l *Log) Infof(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	zap.L().Info(s, f...)
}

func (l *Log) Warnf(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	zap.L().Warn(s, f...)
}

func (l *Log) Errorf(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	zap.L().Error(s, f...)
}

func (l *Log) Panicf(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	zap.L().Panic(s, f...)
}

func (l *Log) Fatalf(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	zap.L().Fatal(s, f...)
}

//判断其他类型--start
func getOtherFileds(format string, args ...interface{}) (string, []zap.Field) {
	//判断是否有context
	l := len(args)
	if l > 0 {
		return fmt.Sprintf(format, args[:l]...), []zap.Field{}
	}
	return format, []zap.Field{}
}

func WriteErr(err error) {
	zap.L().Error("错误打印", zap.Error(err))
}
