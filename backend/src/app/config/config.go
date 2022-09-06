package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	DbType          string
	DbDriverName    string
	DbName          string
	DbUserName      string
	DbUserPassword  string
	DbHost          string
	DbPort          string
	MigrateFilePath string
}

var config ConfigList

func Get() ConfigList {
	return config
}

func Init() {
	env := os.Getenv("ENV_GO")
	configFile := env + ".ini"
	coonfigPath := filepath.Join("config", configFile)
	cfg, err := ini.Load(coonfigPath)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	config = ConfigList{
		DbType:          cfg.Section("db").Key("db_type").String(),
		DbDriverName:    cfg.Section("db").Key("db_driver_name").String(),
		DbName:          cfg.Section("db").Key("db_name").String(),
		DbUserName:      cfg.Section("db").Key("db_user_name").String(),
		DbUserPassword:  cfg.Section("db").Key("db_user_password").String(),
		DbHost:          cfg.Section("db").Key("db_host").String(),
		DbPort:          cfg.Section("db").Key("db_port").String(),
		MigrateFilePath: cfg.Section("db").Key("migrate_path").String(),
	}
}
