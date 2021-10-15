package data

import (
    "context"
    "fmt"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm/utils"
    "time"
)

// GormLogger 封装Gorm可用的日志
type GormLogger struct {
    *log.Helper
}

// Trace print sql message
func (gl *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {


    if l.LogLevel <= Silent {
        return
    }

    elapsed := time.Since(begin)
    switch {
    case err != nil && l.LogLevel >= Error && (!errors.Is(err, ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
        sql, rows := fc()
        if rows == -1 {
            l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
        } else {
            l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
        }
    case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= Warn:
        sql, rows := fc()
        slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
        if rows == -1 {
            l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
        } else {
            l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
        }
    case l.LogLevel == Info:
        sql, rows := fc()
        if rows == -1 {
            l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
        } else {
            l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
        }
    }
}

func NewGormLogger(logger log.Logger) *GormLogger {
    logInstance := log.NewHelper(logger)

}
