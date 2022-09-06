
# gin_gorm_test  

ginとgormをdocker環境上で動かしてみるテスト用リポジトリです  

## 起動方法  

* コンテナを起動して入る  

> docker-compose -f ./docker-compose-dev.yml up -d  
> docker-compose -f ./docker-compose-dev.yml exec go bash  

* go.mainの起動  

> go run main.go  

* API実行
POSTMANあたりでPOSTしてください

* DBコンテナの操作  

> docker-compose exec db bash  
> mysql -u root -p  
> Enter password: root  
> mysql> use backend;

## その他
* 環境変数
{ENV_GO}.iniから読み込んでいる
ENV_GOはdocker-compose内で指定
src/app/config直下ini_sampleがあるのでそれを参考に作成する