package webengine

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

var Logger *log.Logger

var GinInstance *gin.Engine

func init() {
	// Logger Init
	loggerInit()
	ginInit()
}

func loggerInit() {
	Logger = log.New()
	logfile, err := os.OpenFile("./goserver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	CheckFmt(err)
	writer := io.MultiWriter(os.Stdout, logfile)
	Logger.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05,000000",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "[" + frame.Function + "]", "[" + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line) + "]"
		},
	})
	Logger.SetReportCaller(true)
	Logger.SetOutput(writer)
}

func ginInit() {
	//gin.SetMode(gin.ReleaseMode)
	GinInstance = gin.New()
	GinInstance.Use(logger(), gin.Recovery())

	//GinInstance.Any("/test", func(c *gin.Context) {
	//	Logger.Info("收到了test的日志")
	//	c.JSON(http.StatusOK, gin.H{"code": 1})
	//})
}

func dbInit() {

}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		//// Set example variable
		//c.Set("example", "12345")

		//继续执行
		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI

		clientIp := c.GetHeader("X-Real-IP")
		if clientIp == "" {
			clientIp = c.ClientIP()
		}
		statusCode := c.Writer.Status()

		Logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIp,
			reqMethod,
			reqUri,
		)
	}
}
