package global

import (
	"github.com/fsnotify/fsnotify"
	"github.com/king-tu/mallweb/app/config"
	logger "github.com/king-tu/mallweb/app/log"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
)

var (
	Config *config.App
	Logger logger.Factory
)

const defaultConfigFile = "../../../config.yaml"

const (
	// logFormat
	LOGFORMAT_JSON = "json"

	// EncoderConfig
	TIME_KEY       = "time"
	LEVLE_KEY      = "level"
	NAME_KEY       = "logger"
	CALLER_KEY     = "caller"
	MESSAGE_KEY    = "msg"
	STACKTRACE_KEY = "stacktrace"
)

func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v\n", err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file changed: %v\n", e.Name)
		if err := v.Unmarshal(&Config); err != nil {
			log.Printf("Fail to unmarshals the config into a Struct: %v\n", err)
		}
	})
	if err := v.Unmarshal(&Config); err != nil {
		log.Printf("Fail to unmarshals the config into a Struct: %v\n", err)
	}
}

// 设置日志级别、输出格式和日志文件的路径
func SetLogs(level, format, fileName, srvName string) {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        TIME_KEY,
		LevelKey:       LEVLE_KEY,
		NameKey:        NAME_KEY,
		CallerKey:      CALLER_KEY,
		MessageKey:     MESSAGE_KEY,
		StacktraceKey:  STACKTRACE_KEY,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 大写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器(相对路径+行号)
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志输出格式
	var encoder zapcore.Encoder
	switch strings.ToLower(format) {
	case LOGFORMAT_JSON:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 添加日志切割归档功能
	hook := lumberjack.Logger{
		Filename:   fileName,              // 日志文件路径
		MaxSize:    Config.Log.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: Config.Log.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     Config.Log.MaxAge,     // 文件最多保存多少天
		Compress:   Config.Log.Compress,   // 是否压缩
	}

	// Set the log level
	var logLevel zapcore.Level
	switch strings.ToLower(level) {
	case "trace":
		logLevel = zap.DebugLevel
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	case "fatal":
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}

	core := zapcore.NewCore(
		encoder, // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zap.NewAtomicLevelAt(logLevel), // 日志级别
	)

	opts := []zap.Option{
		// 开启文件及行号
		zap.AddCaller(),
		// 开启开发模式，堆栈跟踪
		zap.Development(),
		zap.AddStacktrace(zapcore.FatalLevel),
		// 避免zap始终将包装器(wrapper)代码报告为调用方。
		zap.AddCallerSkip(1),
	}

	// 构造日志
	l := zap.New(core, opts...)
	zap.ReplaceGlobals(l)

	Logger = logger.NewFactory(zap.L().With(zap.String("service", srvName)))
}
