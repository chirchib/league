package logger

import (
	"errors"
	"net"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	Level        zapcore.Level
	UDPAddress   string // если не задан, логи не будут отправляться по udp
	WriteToFile  bool   // если не задан, логи не будут записываться в файл
	LogsFilePath string
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05.000Z07:00]"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func CustomNameEncoder(name string, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + name + "]")
}

func CustomEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

func NewLogger(config *LoggerConfig) (*zap.SugaredLogger, error) {
	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.

	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = SyslogTimeEncoder
	pe.EncodeLevel = CustomLevelEncoder
	pe.EncodeName = CustomNameEncoder
	pe.EncodeCaller = CustomEncodeCaller
	jsonEncoder := zapcore.NewJSONEncoder(pe)

	cores := make([]zapcore.Core, 0)

	stdoutCore := zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), config.Level)
	cores = append(cores, stdoutCore)

	if config.UDPAddress != "" {
		conn, err := net.Dial("udp", config.UDPAddress)
		if err != nil {
			return nil, err
		}

		udpOutWriter := zapcore.AddSync(conn)
		cores = append(cores, zapcore.NewCore(jsonEncoder, udpOutWriter, config.Level))
	}

	if config.WriteToFile {
		if config.LogsFilePath == "" {
			return nil, errors.New("empty logs file path")
		}

		fw, err := NewFileWriter(config.LogsFilePath)
		if err != nil {
			return nil, err
		}

		fileCore := zapcore.AddSync(fw)
		cores = append(cores, zapcore.NewCore(jsonEncoder, fileCore, config.Level))
	}

	core := zapcore.NewTee(cores...)

	logger := zap.New(core, zap.AddCaller())

	return logger.Sugar(), nil
}
