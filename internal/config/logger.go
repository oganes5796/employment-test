package config

type LoggerConfig struct {
	LogLevel      string // debug, info, warn, error
	FilePath      string // путь к файлу логов
	MaxSize       int    // максимальный размер файла (MB)
	MaxBackups    int    // количество старых файлов
	MaxAge        int    // сколько дней хранить логи
	EnableConsole bool   // логировать ли в консоль
}

// DefaultConfig возвращает конфиг по умолчанию
func DefaultLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		LogLevel:      "debug",
		FilePath:      "logs/app.log",
		MaxSize:       10,
		MaxBackups:    5,
		MaxAge:        7,
		EnableConsole: true,
	}
}
