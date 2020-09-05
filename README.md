# twitter-bulk-silencer

特定の検索ワードで抽出されたツイート投稿者をまとめてブロック／ミュートするコマンドラインアプリケーションです。

Goで実装されています。バイナリが必要なら適当にビルドしてください。

## usage

### twitter developer account

Twitterの開発者用アカウントおよびアプリ登録が必要なので、各自申請してください。

アプリ権限は「Read and write」が必要です。

API keyとAPI secret keyを控え、Access tokenおよびAccess token secretを発行してください（自分用のtokenはダッシュボードから発行可能）。

### setting .env file

.env.sampleをもとに、必要事項を追記してください。

ファイルは実行ディレクトリに設置します。

```
$ cp .env.sample .env
```

* BASE_DIR
  * アプリのプロジェクトルートディレクトリ
  * リストファイルの参照に利用する
* ACCESS_TOKEN
  * アプリのAPI key
* ACCESS_TOKEN_SECRET
  * アプリのAPI secret key
* CONSUMER_KEY
  * ユーザごとのAccess token
* CONSUMER_SECRET
  * ユーザごとのAccess token secret
* ENABLE_STANDARD_QUERY
  * true/false
  * 検索対象ワードを自由に受け入れるかどうかの設定
  * trueの場合は任意の検索ワードを有効に、falseの場合はハッシュタグのみを有効にする
* BLOCK_MODE
  * true/false
  * 抽出された処理対象ユーザをブロックするかの設定
  * trueの場合はブロック、falseの場合はミュートする
* PROTECT_FOLLOWER
  * true/false
  * 抽出された処理対象ユーザから「フォロワー」に該当するユーザを除外する設定
  * trueの場合は保護される
* PROTECT_FOLLOWEE
  * true/false
  * 抽出された処理対象ユーザから「フォロー」に該当するユーザを除外する設定
  * trueの場合は保護される

### initialize

フォロー、フォロワー、ミュート、ブロック者を取得しておく必要があります。

対応するAPIの制限に準じるので、APIの利用状況や、取得するユーザ数によっては長時間処理待ちで停止することがあります。

詳細はTwitter APIの公式ドキュメントを参照ください。

```
$ go run cmd/prepare/main.go [followee/folower/mute/block]
```

まとめて実行するスクリプトも同梱しています。

```
$ ./setup.sh
```

### execute

ENABLE_STANDARD_QUERYがfalseの場合、ハッシュタグ（半角シャープから始まる文字列）のみを指定することができます。

ドライラン

```
$ go run cmd/silencer/main.go [検索ワード]
```

実行

```
$ go run cmd/silencer/main.go [検索ワード] execute
```

## note

検索処理はStandard Search APIを利用しており、PremiumおよびEnterpriseのAPIには対応していません。

検索対象は過去7日分ですが、Twitter APIの挙動に依存します。

TwitterのAPIが渋いのと、過激なアプリなので開発API申請ができてアプリを実行できる人に暗黙的に利用を限定しています。自己責任でお願いします。

これをベースに改造したりWebサービスとして公開するなどお好きにどうぞ。