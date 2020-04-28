package log

type OrmLogger struct {
}

func (o *OrmLogger) Print(v ...interface{}) {
	format := v[0].(string)
	caller := v[1].(string)
	v = v[2:] // 打印sql语句
	logger.With("caller", caller).Infof(format+" %v  ", v)
}

func NewGormLogger() *OrmLogger {
	return &OrmLogger{}
}
