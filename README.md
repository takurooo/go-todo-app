# go-todo-app

TODO Web Application with AUTH by Go.

[詳解Go言語Webアプリケーション開発](https://www.c-r.com/book/detail/1462)を読みながら実装したもの

下記URLが本家のリポジトリ  
<https://github.com/budougumi0617/go_todo_app>

## 開発環境作成

```shell
# 開発環境用のコンテナを起動
make up

# DBのテーブル作成
make migrate
```

## デバッグ

```shell
# コンテナのログ出力
make logs

# DBに接続
./bin/connect_mysql.sh 
```

## 開発環境削除

```shell
# 開発環境用のコンテナを削除
make down
```

## サンプルリクエスト

```shell
# ヘルスチェック
curl -XGET localhost:18000/health

# タスク登録
curl -i -XPOST localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden

# タスク一覧取得
curl -i -XGET localhost:18000/tasks

# ユーザー登録
curl -X POST localhost:18000/register -d '{"name": "john2", "password":"test", "role":"user"}'
```
