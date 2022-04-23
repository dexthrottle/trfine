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
func GetConfig() *Config {

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
			Port:    "10000",
			GinMode: "debug",
		},
	}
}
