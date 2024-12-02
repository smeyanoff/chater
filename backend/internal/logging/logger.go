package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	// Настраиваем encoder для вывода логов
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// Настраиваем вывод логов в стандартный поток (os.Stdout)
	consoleWriteSyncer := zapcore.AddSync(os.Stdout)

	// Настраиваем core для логирования в консоль
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // Формат вывода
		consoleWriteSyncer,                       // Куда выводить
		zapcore.DebugLevel,                       // Уровень логирования (Debug и выше)
	)

	// Создаем логгер с core и экспортируем его
	Logger = zap.New(core)
}
