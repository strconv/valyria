package log

import (
	"context"
	"io"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	TraceIDKey = "trace_id"
	//UserIDKey      = "uid"
	//ServiceNameKey = "service_name"
)

var logger *zap.SugaredLogger

var encoderConf = zapcore.EncoderConfig{
	MessageKey:    "msg",
	LevelKey:      "level",
	TimeKey:       "time",
	NameKey:       "logger",
	CallerKey:     "caller",
	StacktraceKey: "stack",
	LineEnding:    zapcore.DefaultLineEnding,
	EncodeLevel:   zapcore.CapitalColorLevelEncoder,
	EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	},
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
	EncodeName:     zapcore.FullNameEncoder,
}

func Init(level string) {
	encoder := zapcore.NewConsoleEncoder(encoderConf)
	infoLvl, debugLvl, warnLvl := setLevel(level)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter("info")), infoLvl),
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter("debug")), debugLvl),
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter("error")), warnLvl),
	)
	logger = zap.New(
		core,
		zap.AddCaller(), // 堆栈跟踪
		zap.AddStacktrace(zap.WarnLevel),
	).Sugar()
	logger.Info("log init ...")
}

// trace id
func For(ctx context.Context, args ...interface{}) *zap.SugaredLogger {
	logger = logger.With(zap.Field{
		Key:    TraceIDKey,
		Type:   zapcore.StringType,
		String: extraTraceID(ctx),
	})
	return logger
}

func getWriter(level string) io.Writer {
	t := time.Now()
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:  "./logs/" + level + "-" + t.Format("2006_01_02") + ".log",
		MaxSize:   512, // MB
		LocalTime: true,
		MaxAge:    7, // 最多保存七天
	})
}

func setLevel(level string) (zap.LevelEnablerFunc, zap.LevelEnablerFunc, zap.LevelEnablerFunc) {
	logLevel := zap.DebugLevel
	switch level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "error":
		logLevel = zap.ErrorLevel
	//case "warn":
	//	logLevel = zap.WarnLevel
	//case "panic":
	//	logLevel = zap.PanicLevel
	//case "fatal":
	//	logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}
	// 实现两个判断日志等级的interface  自定义级别展示
	debugLvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if logLevel >= zapcore.InfoLevel {
			return false
		}
		return logLevel.Enabled(lvl)
	})
	infoLvl := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel && level >= logLevel
	})
	errorLvl := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel && level >= logLevel
	})
	return infoLvl, debugLvl, errorLvl
}

func extraTraceID(ctx context.Context) string {
	v := ctx.Value(TraceIDKey)
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
