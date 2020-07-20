# RamenWikiStoreFunction

- [らーめんWik](https://github.com/Cherry0202/RamenWiki)で使用するデータベースramenにgooglePlaceAPIから取得したデータを保存するapiです。
- 今回はサンプルデータのため20件を保存します。(らーめん　新宿でリクエストを投げてます)
- Dockerをつかって、golangとmysql環境を作成します。

## 前提

- [Docker](https://docs.docker.com/get-docker/)がinstallされていること

## USAGE

```shell
mkdir go/src/github.com/Cherry0202 #user直下に作成してください
cd /go/src/github.com/Cherry0202/
git clone https://github.com/Cherry0202/RamenWikiStoreFunction.git
cd RamenWikiStoreFunction
docker-compose build
docker-compose up
docker-compoose up -d #コンテナのログが邪魔だと思う方はこちら
```

## APIサーバ起動

```shell
# 別タブで 
docker exec -it golang-api-container bash #コンテナ内に入ります
go run server.go
# localhost:8080 にアクセスしよう
```

## MySQL

```shell
# 別タブで 
docker exec -it mysql-container bash
mysql -u user名 -p
```

## 起動したMySQLにホストから接続してみる

```shell
mysql --host ホストip --port ポート番号 -u user名 -p 
```
