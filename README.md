# サーバーエンジニア向け 2026新卒採用事前課題

あなたは歌手とアルバムを管理するAPIの機能開発にたずさわることになりました。

次の課題に順に取り組んでください。

できない課題があっても構いません。

面接中に課題に関して質問をしますので、分かる範囲で説明してください。

## 課題1
プログラムのコードを読み、中身を把握しましょう。

## 課題2
Docker と Go をインストールし(各自で調べてください)、歌手を管理するAPIの動作を確認しましょう。

```
# (ターミナルを開いて)
# Docker コンテナを起動する
docker compose up -d
# HTTP サーバーを起動する
go run main.go
```

```
# (別のターミナルを開いて)
# 歌手の一覧を取得する
curl http://localhost:8888/singers

# 指定したIDの歌手を取得する
curl http://localhost:8888/singers/1

# 歌手を追加する
curl -X POST -d '{"id":10,"name":"John"}' http://localhost:8888/singers

# 歌手を削除する
curl -X DELETE http://localhost:8888/singers/10
```

## 課題3
アルバムを管理するAPIを新規作成しましょう。

### 3-1
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer_id":1},{"id":2,"title":"Alice's 2nd Album","singer_id":1},{"id":3,"title":"Bella's 1st Album","singer_id":2}]
```

### 3-2
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer_id":1}
```

### 3-3
アルバムを追加するAPI
```
curl -X POST -d '{"id":10,"title":"Chris 1st","singer_id":3}' http://localhost:8888/albums

# このようなレスポンスを期待しています
{"id":10,"title":"Chris 1st","singer_id":3}

# そして、アルバムを取得するAPIでは、追加したものが存在するように
curl http://localhost:8888/albums/10
```

### 3-4
アルバムを削除するAPI
```
curl -X DELETE http://localhost:8888/albums/10
```

## 課題4
アルバムを取得するAPIでは、歌手の情報も付加するように改修しましょう。

### 4-1
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}}
```

### 4-2
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}},{"id":2,"title":"Alice's 2nd Album","singer":{"id":1,"name":"Alice"}},{"id":3,"title":"Bella's 1st Album","singer":{"id":2,"name":"Bella"}}]
```

## 課題5
歌手とそのアルバムを管理するという点で、現状の実装の改善点を検討し思いつく限り書き出してください。

実装をする必要はありません。

- 誤ってデータが削除される場合を想定して、複数DBを用いた管理方法（ブロックチェーンなど）やバックアップ（差分、増分、フルなど手法は仕様に合わせる）機能の実装
- データの修正・更新用のUpdateのAPIを実装
- データの追加や削除が簡単にできてしまうため、セキュリティ強化を実装する。（アクセストークンを生成するAPIを実装し、アクセストークンとデータ追加・削除などのAPIを合わせて使用するようにする）
- サイバー攻撃対策（SQLインジェクションなど）
- 高トラフィック対策として、キャッシュサーバーなどを追加し、一度検索された情報に対する処理速度を向上させる。
- 歌手やアルバムを名前で検索・フィルタリングできるようにする。
- 価値を増やすために歌手やアルバムのカラムを増やし、情報を増やす。（歌手の年齢や性別、アルバムの値段など）
- サンプルコードから大きな変化を加えずに実装しようとしたためgorilla/muxを使用したが、GinやEchoなどの別のフレームワークを使用することを検討する。
