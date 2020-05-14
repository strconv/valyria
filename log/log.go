package log

import (
	"context"
	"io"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/strconv/valyria/trace"
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
	infoLvl, debugLvl, errorLvl := setLevel(level)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter("info")), infoLvl),
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter("debug")), debugLvl),
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter("error")), errorLvl),
	)
	logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.PanicLevel), // 堆栈跟踪
	).Sugar()
	logger.Info("log init ...")
}

// trace id
func For(ctx context.Context, args ...interface{}) *zap.SugaredLogger {
	tid := trace.ExtraTraceID(ctx, TraceIDKey)
	var fields []interface{}
	if len(tid) != 0 {
		fields = make([]interface{}, 0, len(args)+2)
		fields = append(fields, TraceIDKey, tid)
	} else {
		fields = make([]interface{}, 0, len(args))
	}
	fields = append(fields, args...)
	return logger.With(fields...)
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
	infoLvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	debugLvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if logLevel >= zapcore.InfoLevel {
			return false
		}
		return logLevel.Enabled(lvl)
	})
	errorLvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	return infoLvl, debugLvl, errorLvl
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Debug(v ...interface{}) {
	logger.Warn(v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}
