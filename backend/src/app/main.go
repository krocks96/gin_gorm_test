package main

import (
	"app/config"
	"app/db"
	"app/server"
)

func main() {
	// config読込
	config.Init()
	// DB起動
	db.Init()
	// サーバー起動
	server.StartServer()
	// DBクローズ
	db.Close()
}
