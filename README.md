
# gin_gorm_test  

ginとgormをdocker環境上で動かしてみるテスト用リポジトリです  

## 起動方法  

* コンテナを起動して入る  

> docker-compose up -d  
> docker-compose exec go bash  

* go.mainの起動  

> cd app  
> go run main.go  

* API実行
POSTMANあたりでPOSTしてください
GETは未実装

* DBコンテナの操作  

> docker-compose exec mysql bash  
> mysql -u root -p  
> Enter password: root  
> mysql> use backend;
