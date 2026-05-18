package app

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	setDefaults(v)
	bindEnvs(v)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config failed: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config failed: %w", err)
	}

	return &cfg, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.name", "video-feed")
	v.SetDefault("app.port", 8080)
	v.SetDefault("app.mode", "debug")

	v.SetDefault("mysql.host", "127.0.0.1")
	v.SetDefault("mysql.port", 13306)
	v.SetDefault("mysql.user", "root")
	v.SetDefault("mysql.password", "")
	v.SetDefault("mysql.dbname", "video_feed")
	v.SetDefault("mysql.max_open_conns", 20)
	v.SetDefault("mysql.max_idle_conns", 10)

	v.SetDefault("jwt.secret", "")
	v.SetDefault("jwt.expire_hours", 72)

	v.SetDefault("log.level", "info")

	v.SetDefault("redis.addr", "127.0.0.1:16379")
	v.SetDefault("redis.password", "")
	v.SetDefault("redis.db", 0)
	v.SetDefault("redis.video_detail_ttl_seconds", 300)
}

func bindEnvs(v *viper.Viper) {
	_ = v.BindEnv("app.mode", "APP_MODE")
	_ = v.BindEnv("app.port", "APP_PORT")

	_ = v.BindEnv("mysql.host", "MYSQL_HOST")
	_ = v.BindEnv("mysql.port", "MYSQL_PORT")
	_ = v.BindEnv("mysql.user", "MYSQL_USER")
	_ = v.BindEnv("mysql.password", "MYSQL_PASSWORD")
	_ = v.BindEnv("mysql.dbname", "MYSQL_DBNAME")

	_ = v.BindEnv("jwt.secret", "JWT_SECRET")
	_ = v.BindEnv("jwt.expire_hours", "JWT_EXPIRE_HOURS")

	_ = v.BindEnv("redis.addr", "REDIS_ADDR")
	_ = v.BindEnv("redis.password", "REDIS_PASSWORD")
	_ = v.BindEnv("redis.db", "REDIS_DB")
}
