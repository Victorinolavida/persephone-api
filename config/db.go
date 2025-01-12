package config

type DBConfig struct {
	Port     int    `koanf:"port"`
	Host     string `koanf:"host"`
	DbName   string `koanf:"dbname"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
}
