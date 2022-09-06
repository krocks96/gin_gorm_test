package db

import (
	"app/config"

	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

// データベースの初期化
func Init() {
	// 接続情報
	config := config.Get()
	dbConnectInfo := buildDBConfig(config)
	// データベース作成
	db, err := sql.Open("mysql", dbConnectInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.DbName)
	_, err = db.Exec(sql)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	// configから読み込んだ情報を元に、データベースに接続する
	Db, err = gorm.Open(mysql.Open(dbConnectInfo))
	if err != nil {
		panic(err)
	}
	// テーブルの作成
	autoMigration(config, dbConnectInfo)
}

func GetDB() *gorm.DB {
	return Db
}

func Close() {
	db, _ := Db.DB()
	defer db.Close()
}

func buildDBConfig(config config.ConfigList) string {
	dbConnectInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUserName,
		config.DbUserPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)
	return dbConnectInfo
}

func autoMigration(configs config.ConfigList, dbConnectInfo string) {
	// golang-migrateのインスタンス生成
	log.Println(dbConnectInfo)
	log.Println(configs.DbType + "://" + dbConnectInfo)
	m, err := migrate.New(configs.MigrateFilePath, configs.DbType+"://"+dbConnectInfo)
	if err != nil {
		log.Println("--- Migration Initialize Failed ---")
		panic(err)
	}
	// Migration実行
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		log.Println("--- Migration Failed ---")
		panic(err)
	}
}
