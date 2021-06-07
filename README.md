# golang-todo-app

Go で TODO 管理する Web アプリ

# セットアップ (Mac)

## sqlite3

1. Homebrew をインストール

https://qiita.com/zaburo/items/29fe23c1ceb6056109fd

2. brew から sqlite3 をインストール

```bash
brew install sqlite
```

3. gcc をインストール

https://developer.apple.com/xcode

4. Go でアプリケーションのモジュールを作成

```bash
go mod init main
```

5. Go で sqlite のドライバをインストール

```bash
go get github.com/mattn/go-sqlite3
```

6. 不要らしいが一応

```bash
export CGO_ENABLED=1
```

7. PostgreSQL のクライアント(コマンドライン)のみをインストール

```bash
brew update
brew install libpq
```

# 実行方法

1. Docker でデータベースを起動する

```bash
docker-compose up -d --build
```

2. アプリケーションを実行する

```bash
go run main.go
```

or

```bash
go build -o build/main
./build/main
```

3. アプリケーションの URL にアクセスし、アカウントを登録する

http://localhost:3000/signup

4. ログインする

http://localhost:3000/login

5. Todo データを登録したり、更新したり、削除したりする

6. データベースのデータは専用の Web アプリからでも見れる

http://localhost:8081

以上

# 終了方法

1. アプリケーションの終了

実行したターミナルで`Ctrl + c`を押すことで終了できる。

2. Docker のクリーンアップ

docker-compose を終了し、リソースも削除する。

```bash
docker-compose down --rmi all --volumes
```

以上

# 参考

[Udemy Golang 基礎〜応用](https://www.udemy.com/course/golang-webgosql/learn/lecture/23672858#questions)

[docker-compose で postgres の環境を作る](https://qiita.com/mabubu0203/items/5cdff1caf2b024df1d95)
