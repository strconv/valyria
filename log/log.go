package log

import (
	"io"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func InitLog(level string) {
	config := *logConf()
	encoder := zapcore.NewConsoleEncoder(config)

	infoLvl, warnLvl := setLevel(level)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter("info")), infoLvl),
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter("error")), warnLvl),
		// console
		//zapcore.NewCore(zapcore.NewConsoleEncoder(config),
		//	zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel),
	)
	logger = zap.New(
		core,
		zap.AddCaller(), // 文件名和行数
		zap.AddStacktrace(zap.WarnLevel),
	).Sugar()
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

func logConf() *zapcore.EncoderConfig {
	return &zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		CallerKey:   "file",
		TimeKey:     "ts",
		EncodeLevel: zapcore.CapitalLevelEncoder, // 级别转换成大写
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}
}

func setLevel(level string) (zap.LevelEnablerFunc, zap.LevelEnablerFunc) {
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
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.WarnLevel && level >= logLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel && level >= logLevel
	})
	return infoLevel, warnLevel
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
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
