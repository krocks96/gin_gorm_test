package main

import (
        "app/server"
        "app/db"
        )

func main() {
    // DB起動
    db.Init()
    // サーバー起動
    server.StartServer()
    // DBクローズ
    db.Close()
}