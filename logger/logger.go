package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
)

type Logger struct {
	*zerolog.Logger
}

var (
	log         = Logger{}
	logPath     = ""
	logFileName = "default"
	logModel    = "debug"
)

const (
	debugModel   = "debug"
	releaseModel = "release"
)

// 格式化打印的文件名
func CallerMarshalFunc(file string, line int) string {
	eidx := strings.LastIndex(file, "/") + 1
	return fmt.Sprintf("%s:%d", file[eidx:], line)
}

//调用初始化
func SetUp(path, defaultFile, model string) {
	logPath = path
	logFileName = defaultFile
	logModel = model

	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.CallerMarshalFunc = CallerMarshalFunc
	zerolog.CallerSkipFrameCount = 3

	log = New(logFileName + ".log")
}

func New(fileName string) Logger {
	file := os.Stdout
	if logModel == releaseModel {
		curPath, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		path := curPath + "/" + logPath
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.Mkdir(path, 0775)
			if err != nil {
				panic(err)
			}
		}
		fileName = path + "/" + fileName
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			file, err = os.Create(fileName)
		} else {
			file, err = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0666)
		}
		if err != nil {
			panic(err)
		}
	}
	logger := zerolog.New(file).With().Timestamp().Caller().Logger()
	log := Logger{&logger}
	return log
}

func (l *Logger) LogDebug(format string, a ...interface{}) {
	l.Debug().Msgf(format, a...)
}

func (l *Logger) LogInfo(format string, a ...interface{}) {
	l.Info().Msgf(format, a...)
}

func (l *Logger) LogError(format string, a ...interface{}) {
	l.Error().Msgf(format, a...)
}

func (l *Logger) LogFatal(format string, a ...interface{}) {
	l.Fatal().Msgf(format, a...)
}

func LogDebug(format string, a ...interface{}) {
	log.Debug().Msgf(format, a...)
}

func LogInfo(format string, a ...interface{}) {
	log.Info().Msgf(format, a...)
}

func LogError(format string, a ...interface{}) {
	log.Error().Msgf(format, a...)
}

func LogFatal(format string, a ...interface{}) {
	log.Fatal().Msgf(format, a...)
}

func init() {
	SetUp(logPath, logFileName, logModel)
}
