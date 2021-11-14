package config

import (
	"gopkg.in/ini.v1"
	"log"
)

func NewConfigFromFile(file string) Config {
	var section *ini.Section

	cfgFile, err := ini.Load(file)
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}

	// #####################################################################
	// MySQL section
	// #####################################################################
	section, err = cfgFile.GetSection("mysql")
	if err != nil {
		log.Fatalf("Error reading section 'mysql' from config file: %v", err)
	}

	config.MySQL.Host = section.Key("host").String()
	if config.MySQL.Host == "" {
		log.Fatal("Config setting mysql.host is mandatory")
	}

	config.MySQL.Port = section.Key("port").MustInt(3306)
	if config.MySQL.Port == 0 {
		log.Fatal("Config setting mysql.port is mandatory")
	}

	config.MySQL.DbName = section.Key("dbname").String()
	if config.MySQL.DbName == "" {
		log.Fatal("Config setting mysql.dbname is mandatory")
	}

	config.MySQL.User = section.Key("user").String()
	config.MySQL.Password = section.Key("password").String()

	// #####################################################################
	// Redis section
	// #####################################################################
	section, err = cfgFile.GetSection("redis")
	if err != nil {
		log.Fatalf("Error reading section 'redis' from config file: %v", err)
	}

	config.Redis.Host = section.Key("host").String()
	if config.Redis.Host == "" {
		log.Fatal("Config setting redis.host is mandatory")
	}

	config.Redis.Port = section.Key("port").MustInt(6379)
	if config.Redis.Port == 0 {
		log.Fatal("Config setting redis.port is mandatory")
	}

	config.Redis.Password = section.Key("password").String()

	// #####################################################################
	// App section
	// #####################################################################
	section, err = cfgFile.GetSection("app")
	if err != nil {
		log.Fatalf("Error reading section 'app' from config file: %v", err)
	}

	config.App.BookingLogChannel = section.Key("booking_log_chanel").String()
	if config.App.BookingLogChannel == "" {
		log.Fatal("Config setting app.booking_log_chanel is mandatory")
	}

	return config
}

type Config struct {
	MySQL MySQLConfig
	Redis RedisConfig
	App   AppConfig
}
