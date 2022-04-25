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
func GetConfig(ginMode string, port string) *Config {

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
