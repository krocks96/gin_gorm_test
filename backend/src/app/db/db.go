package db

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "log"
    "app/config"
	"app/models/entity"
)

var (
	Db *gorm.DB
	err error
)

// データベースの初期化
func Init() {
    dbConnectInfo := fmt.Sprintf(
        `%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
        config.Config.DbUserName,
        config.Config.DbUserPassword,
        config.Config.DbHost,
        config.Config.DbPort,
        config.Config.DbName,
    )

    println(dbConnectInfo)

    // configから読み込んだ情報を元に、データベースに接続する
    Db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)
    if err != nil {
        log.Fatalln(err)
    } else {
        fmt.Println("Successfully connect database..")
    }

    // テーブルの作成します
	autoMigration()
}

func GetDB() *gorm.DB {
    return Db
}

func Close(){
    db := Db.DB()
    defer db.Close()
}

func autoMigration() {
	Db.AutoMigrate(&entity.User{})
}