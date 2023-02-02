package core

import (
	"LibraryManagementSys/global"
	"LibraryManagementSys/utils"
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// GormZapLogger
// zap与gorm整合
type GormZapLogger struct {
	ZapLogger     *zap.SugaredLogger
	SlowThreshold time.Duration
	gormlogger.Config
}

func NewGormZapLogger() GormZapLogger {
	logger := GormZapLogger{
		ZapLogger:     global.LMS_LOG,
		SlowThreshold: 200 * time.Millisecond,
	}
	logger.LogLevel = translationGormLogLevel(global.LMS_CONFIG.Zap.GormLevel)
	logger.Colorful = true
	return logger
}

// LogMode 实现 gormlogger.Interface 的 LogMode 方法
func (l GormZapLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	gormZapLogger := l
	gormZapLogger.LogLevel = level
	return gormZapLogger
}

// Info 实现 gormlogger.Interface 的 Info 方法
func (l GormZapLogger) Info(ctx context.Context, str string, args ...interface{}) {
	l.logger().Debugf(str, args...)
}

// Warn 实现 gormlogger.Interface 的 Warn 方法
func (l GormZapLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	l.logger().Warnf(str, args...)
}

// Error 实现 gormlogger.Interface 的 Error 方法
func (l GormZapLogger) Error(ctx context.Context, str string, args ...interface{}) {
	l.logger().Errorf(str, args...)
}

// Trace 实现 gormlogger.Interface 的 Trace 方法
func (l GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	// 获取运行时间
	elapsed := time.Since(begin)
	// 获取 SQL 请求和返回条数
	sql, rows := fc()

	// 通用字段
	timeStr := utils.MicrosecondsStr(elapsed)

	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.logger().Warnf("Database ErrRecordNotFound, sql:%s, time:%s, rows:%d", sql, timeStr, rows)
		} else {
			// 其他错误使用 error 等级
			l.logger().Errorf("Database Error, sql:%s, time:%s, rows:%d, err:%s", sql, timeStr, rows, zap.Error(err))
		}
	}

	// 慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.logger().Warnf("Database Slow Log, sql:%s, time:%s, rows:%d", sql, timeStr, rows)
	}

	// 记录所有 SQL 请求
	l.logger().Debugf("Database Query, sql:%s, time:%s, rows:%d", sql, timeStr, rows)
}

// logger 内用的辅助方法，确保 Zap 内置信息 Caller 的准确性（如 paginator/paginator.go:148）
func (l GormZapLogger) logger() *zap.SugaredLogger {

	// 跳过 gorm 内置的调用
	var (
		gormPackage    = filepath.Join("gorm.io", "gorm")
		zapgormPackage = filepath.Join("moul.io", "zapgorm2")
	)

	// 减去一次封装，以及一次在 logger 初始化里添加 zap.AddCallerSkip(1)
	clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-2))

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			// 返回一个附带跳过行号的新的 zap logger
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger
}

// translationGormLogLevel
// 日志级别转换
func translationGormLogLevel(level string) gormlogger.LogLevel {
	switch level {
	case "info":
		return gormlogger.Info
	case "warn":
		return gormlogger.Warn
	case "error":
		return gormlogger.Error
	case "silent":
		return gormlogger.Silent
	default:
		return gormlogger.Info
	}
}
