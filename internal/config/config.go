package config

type Config struct {
	Database struct {
		DBName string
	}
	App struct {
		Port    string
		GinMode string
	}
}

// Конфигурация приложения ----- debug | release | test
func GetConfig(useLogs bool, port string) *Config {
	ginMode := "release"
	if useLogs {
		ginMode = "debug"
	}
	return &Config{
		Database: struct {
			DBName string
		}{
			DBName: "trbotdatabase",
		},
		App: struct {
			Port    string
			GinMode string
		}{
			Port:    port,
			GinMode: ginMode,
		},
	}
}
