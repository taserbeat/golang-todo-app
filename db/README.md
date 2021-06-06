# Sqlite

sqlite を使って簡易的にデータベースを使用する。  
sqlite は、[sql フォルダ](./sql)内にある`.sql`ファイルがデータベースを表すファイルとなる。

## 操作コマンド

- コマンドラインツールで sqlite のデータベースに接続

```bash
sqlite3 ./sql/example.sql
```

- テーブルの一覧

```bash
.table
```

- SQL を実行

```bash
SELECT * FROM persons;
```

- コマンドラインツールの接続を解除

```bash
.exit
```

# PostgreSQL

Docker で DB コンテナを起動する。  
ホストから DB コンテナに`psql`コマンドでログインはできるようにしておく。

```bash
# 起動
docker-compose up -d --build

# 終了
docker-compose down --rmi all --volumes
```

## 操作コマンド

```bash
# 接続
psql -h localhost -p 15432 -U root -d test_db

# データベースの一覧
\l

# 指定のデータベースに接続し直す
\c {dbname}

# 現在のデータベースのテーブル一覧
\dt
```

## Docker からコンテナに入るコマンド

```bash
# bashでログイン
docker-compose exec db bash
```
