package core

import (
	"LibraryManagementSys/constant"
	"LibraryManagementSys/global"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

func LmsInitLogger() {
	// debug级别
	writerSynDebug := getLoggerWriter("debug")
	// err级别
	writeSynErr := getLoggerWriter("err")
	encoder := getEncoder()
	// 创建两个core，分别将不同等级的日志输入到对应的日志文件中
	cDebug := zapcore.NewCore(encoder, writerSynDebug, zapcore.DebugLevel)
	cErr := zapcore.NewCore(encoder, writeSynErr, zapcore.ErrorLevel)
	var core zapcore.Core
	// 控制台输出日志
	if global.LMS_CONFIG.Zap.LogInConsole {
		ws := io.MultiWriter(os.Stdout)
		wsSyn := zapcore.AddSync(ws)
		cConfigLevel := zapcore.NewCore(encoder, wsSyn, translationZapLevel(global.LMS_CONFIG.Zap.Level))
		core = zapcore.NewTee(cDebug, cErr, cConfigLevel) // 使用NewTee将cErr、c2和cConfigLevel合并到core，可同时输出到对应的文件中
	} else {
		core = zapcore.NewTee(cDebug, cErr)
	}
	logger := zap.New(core, zap.AddCaller()) // AddCaller可以定位日志打印位置
	// 设置到全局变量中
	global.LMS_LOG = logger.Sugar()
}

// getEncoder
// 时间格式设置
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = consumerTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if global.LMS_CONFIG.Zap.Format == constant.ZapLogFormatJson {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLoggerWriter
// 按时间切割
func getLoggerWriter(level string) zapcore.WriteSyncer {
	logs, _ := rotatelogs.New(global.LMS_CONFIG.Zap.Path+constant.FileFix+global.LMS_CONFIG.System.Name+constant.DASH+level+"-%Y%m%d.log", // 存储位置
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 最长保存30天
		rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
	)
	return zapcore.AddSync(logs)
}

// consumerTimeEncoder
// 自定义格式
func consumerTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.LMS_CONFIG.Zap.Prefix + t.Format(constant.TimeFormatOne))
}

// translationLevel
// 日志级别转换
func translationZapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
