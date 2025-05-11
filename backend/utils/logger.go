package utils

// logger.go 提供日志初始化和配置功能。
// Example:
// global.Logger.Info("some message")
// global.Logger.Infof("some message %v",err)
// global.Logger.Warn("some message")
// global.Logger.Warnf("some message %v",err)
// global.Logger.Error("some message")
// global.Logger.Errorf("some message %v",err)
// global.Logger.Panic("some message")
// global.Logger.Panic("some message %v",err)

import (
	"FashOJ_Backend/config"
	"FashOJ_Backend/global"
	"fmt"

	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger 初始化全局日志记录器。
// 该函数用于初始化 zap 日志库，并设置日志级别为 DebugLevel。
// 它创建了一个新的日志核心，并将其与编码器和写入同步器关联起来。
// 最后，它使用 zap.SugaredLogger 将日志记录器赋值给全局变量 global.Logger。
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel)
	global.Logger = zap.New(core)
	global.Logger.Info("Init Logger")
}

// getEncoder 返回一个编码器，用于格式化日志条目。
// 该函数配置了生产环境下的编码器，并设置了时间格式和日志级别格式。
// 它返回一个控制台编码器，该编码器可以将日志条目格式化为易于阅读的文本。
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter 返回一个写入同步器，用于将日志写入文件。
// 该函数根据当前时间生成日志文件名，并打开或创建该文件。
// 它以追加模式打开文件，并返回一个写入同步器，该同步器可以将日志条目写入文件。

type writer struct {
	file *os.File
}

func newWriter() ( *writer){
	w := new(writer)
	w.file = nil
	return w 
}

func (w *writer) Write(p []byte) (n int, err error) {
	fmt.Println(time.Now())
	timeStr := time.Now().Format("2006-01-02")

	if w.file == nil || timeStr != w.file.Name() {
		if w.file != nil {
			w.file.Close()
		}
		w.file, err = os.OpenFile(
			path.Join(config.FashOJConfig.FashOJApp.LogPath, timeStr),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)
		if err != nil {
			return 0, err
		}
	}

	return w.file.Write(p)
}

func getLogWriter() zapcore.WriteSyncer {

	logwriter := newWriter()

	return zapcore.AddSync(logwriter)
}
