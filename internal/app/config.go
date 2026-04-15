package app

type Config struct {
	App   AppConfig   `mapstructure:"app"`
	MySQL MySQLConfig `mapstructure:"mysql"`
	JWT   JWTConfig   `mapstructure:"jwt"`
	Log   LogConfig   `mapstructure:"log"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type JWTConfig struct {
	Secret      string `mapstructure:"secret"`
	ExpireHours int    `mapstructure:"expire_hours"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}
