package webengine

import (
	"database/sql"
	. "github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

var db *sql.DB

func init() {
	//日志初始化
	loggerInit()
	//配置文件初始化
	configInit()
	//web模块初始化
	ginInit()
	//数据库初始化
	dbInit()
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

func configInit() {
	c = make(map[string]map[string]string)
	con, err := LoadConfigFile(configPath)
	if err != nil {
		Logger.Panic("无法打开配置文件：", configPath)
		panic("无法打开配置文件")
	}

	sectionList := con.GetSectionList()
	for i := range sectionList {
		c[sectionList[i]] = make(map[string]string)
		keyList := con.GetKeyList(sectionList[i])
		for j := range keyList {
			c[sectionList[i]][keyList[j]], _ = con.GetValue(sectionList[i], keyList[j])
			Logger.Info(sectionList[i], "/", keyList[j], "=", c[sectionList[i]][keyList[j]])
		}
	}
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
	dbInfo := GetDBInfoString()
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		Logger.Fatal("连接数据库错误:", err.Error())
	} else {
		Logger.Info("数据库连接成功")
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	err = db.Ping()
	CheckFatal(err)
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
