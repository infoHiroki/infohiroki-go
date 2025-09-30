# 🚀 Cloudflare + Railway 完全デプロイガイド：お名前.comドメインでGoアプリを公開する

## 🎯 中心的な主張

お名前.comで取得したドメインをCloudflare経由でRailwayにデプロイされたGoアプリケーションに接続する完全ガイドです。Cloudflareの無料CDN/SSL、Railwayの自動デプロイ機能を活用することで、高速かつセキュアなWebサイトを構築できます。7つのステップで本番環境を構築し、GitHubへのpush後2〜5分で自動的に本番反映される開発フローを実現します。

---

## 📊 システム構成図

```
🌐 Internet（インターネット）
  ↓
☁️ Cloudflare
  ├─ DNS管理
  ├─ CDN（コンテンツ配信）
  ├─ SSL/TLS証明書
  └─ DDoS保護
  ↓
🚂 Railway
  ├─ Goアプリケーション実行
  ├─ 自動ビルド・デプロイ
  └─ ヘルスチェック
  ↓
📦 GitHub
  └─ ソースコード管理
```

### 各サービスの役割

```
Cloudflare:
  ✅ ドメインのDNS管理
  ✅ 無料SSL/TLS証明書
  ✅ CDNによる高速化
  ✅ DDoS攻撃からの保護
  ✅ キャッシュ管理

Railway:
  ✅ Goアプリケーションのホスティング
  ✅ GitHubからの自動デプロイ
  ✅ 環境変数管理
  ✅ ログ・メトリクス監視
  ✅ カスタムドメイン対応

GitHub:
  ✅ ソースコード管理
  ✅ push時の自動デプロイトリガー
  ✅ バージョン管理
```

---

## 🎯 事前準備

### 必要なもの

```
✅ チェックリスト:

1. ドメイン
   └─ お名前.comで取得済み
   └─ 例: infohiroki.com

2. Cloudflareアカウント
   └─ 無料プランで十分
   └─ https://www.cloudflare.com/

3. Railwayアカウント
   └─ GitHub連携で登録
   └─ https://railway.app/

4. GitHubリポジトリ
   └─ Goアプリケーションのコード
   └─ 例: infoHiroki/infohiroki-go
```

### 前提知識

```
必須:
  ├─ 基本的なGit操作
  ├─ Goアプリケーションの構造理解
  └─ ターミナル操作

あると便利:
  ├─ Docker の基礎知識
  └─ DNS の基本概念
```

---

## 📋 デプロイ手順（全7ステップ）

```
所要時間: 約30分〜1時間
難易度: ⭐⭐⭐☆☆（中級）

ステップ一覧:
  1️⃣ Cloudflareでドメインを追加
  2️⃣ お名前.comでネームサーバーを変更
  3️⃣ Cloudflareでアクティブ化を確認
  4️⃣ RailwayでGoアプリをデプロイ
  5️⃣ CloudflareでDNSレコードを設定
  6️⃣ Railwayでカスタムドメインを設定
  7️⃣ 動作確認
```

---

## 1️⃣ ステップ1: Cloudflareでドメインを追加

### 1-1. サイトの追加

```
手順:
  1. Cloudflareダッシュボードにアクセス
     → https://dash.cloudflare.com/

  2. 「サイトを追加」をクリック

  3. ドメイン名を入力
     → 例: infohiroki.com

  4. プランを選択
     → 「Free」プランで十分

  5. 「サイトを追加」をクリック
```

### 1-2. ネームサーバーを確認

```
⚠️ 重要な情報をメモ:

Cloudflareが表示する2つのネームサーバー:
  例:
    ├─ clyde.ns.cloudflare.com
    └─ nagali.ns.cloudflare.com

💡 ポイント:
  └─ この2つは次のステップで使います
  └─ 必ずメモしておいてください
```

---

## 2️⃣ ステップ2: お名前.comでネームサーバーを変更

### 2-1. ネームサーバー設定画面を開く

```
手順:
  1. お名前.com Naviにログイン
     → https://navi.onamae.com/

  2. 左メニューから選択:
     「ドメイン」→「ネームサーバー/DNS」

  3. 「ネームサーバーの設定」を選択
```

### 2-2. ドメインを選択

```
操作:
  1. 対象ドメインにチェック
     → infohiroki.com

  2. 「その他のネームサーバーを使う」を選択
```

### 2-3. Cloudflareのネームサーバーを入力

```
設定値:
  ネームサーバー1: clyde.ns.cloudflare.com
  ネームサーバー2: nagali.ns.cloudflare.com

削除:
  既存のネームサーバー:
    ├─ 01.dnsv.jp
    ├─ 02.dnsv.jp
    ├─ 03.dnsv.jp
    └─ 04.dnsv.jp
  → これらは全て削除

最後に:
  「確認」→「設定する」
```

### ⚠️ 重要な注意点

```
よくある間違い:
  ❌ DNSレコード設定画面で変更
  ✅ ネームサーバー設定画面で変更

場所を間違えないこと！

反映時間:
  └─ 最大24時間かかる場合があります
  └─ 通常は1〜3時間で完了
```

---

## 3️⃣ ステップ3: Cloudflareでアクティブ化を確認

### 3-1. 状態確認

```
手順:
  1. Cloudflareダッシュボードに戻る

  2. infohiroki.com のステータスを確認

  3. 「アクティブ」になるまで待つ
     └─ 通常1〜24時間
```

### 3-2. アクティブ化後の確認

```
✅ ステータスが「アクティブ」になったら:

確認事項:
  ├─ ✅ ネームサーバー変更が完了
  ├─ ✅ CloudflareがDNSを管理開始
  └─ ✅ 次のステップに進める状態

表示:
  「アクティブ」← 緑色のチェックマーク
```

---

## 4️⃣ ステップ4: RailwayでGoアプリをデプロイ

### 4-1. プロジェクト作成

```
手順:
  1. Railwayにログイン
     → https://railway.app/

  2. 「New Project」をクリック

  3. 「Deploy from GitHub repo」を選択

  4. リポジトリを選択
     → infoHiroki/infohiroki-go

  5. 自動的にビルド・デプロイが開始
     └─ Dockerfileを検出して自動ビルド
```

### 4-2. 環境変数の設定（必要な場合）

```
設定画面:
  プロジェクト → 「Variables」タブ

環境変数の例:
  PORT=8080
  GIN_MODE=release

💡 Note:
  └─ PORTはRailwayが自動設定するため
     通常は設定不要
```

### 4-3. デプロイ完了確認

```
確認手順:
  1. 「Deployments」タブを開く

  2. ステータスを確認
     ✅ Success → デプロイ成功

  3. 自動生成されたURLをメモ
     例: infohiroki-go-production.up.railway.app

💡 このURLは次のステップで使用
```

---

## 5️⃣ ステップ5: CloudflareでDNSレコードを設定

### 5-1. DNS設定画面を開く

```
手順:
  1. Cloudflareダッシュボードで
     infohiroki.com を選択

  2. 左メニュー:
     「DNS」→「レコード」
```

### 5-2. CNAMEレコードを追加

```
設定値:

| 項目                 | 値                                        |
|---------------------|-------------------------------------------|
| タイプ               | CNAME                                     |
| 名前                 | @（ルートドメイン）                         |
| ターゲット            | infohiroki-go-production.up.railway.app   |
| プロキシステータス    | 🟠 オレンジ（プロキシ有効）                 |
| TTL                 | 自動                                       |

「保存」をクリック
```

### 💡 追加設定（オプション）

```
wwwサブドメインも設定する場合:

| 項目                 | 値                                        |
|---------------------|-------------------------------------------|
| タイプ               | CNAME                                     |
| 名前                 | www                                       |
| ターゲット            | infohiroki-go-production.up.railway.app   |
| プロキシステータス    | 🟠 オレンジ（プロキシ有効）                 |
| TTL                 | 自動                                       |
```

---

## 6️⃣ ステップ6: Railwayでカスタムドメインを設定

### 6-1. ドメイン設定画面を開く

```
手順:
  1. Railwayのプロジェクトページを開く

  2. 「Settings」タブを選択

  3. 「Domains」セクションを探す
```

### 6-2. カスタムドメインを追加

```
操作:
  1. 「Add Custom Domain」をクリック

  2. ドメイン名を入力
     → infohiroki.com

  3. 「Add Domain」をクリック

  4. 自動連携
     └─ RailwayがCloudflareと自動的に連携
```

### 6-3. SSL証明書の確認

```
自動処理:
  ├─ Railwayが自動的にSSL証明書を発行
  ├─ 通常1〜5分で完了
  └─ ステータスが「Active」になれば完了

確認:
  Domains → infohiroki.com → Status: Active
```

---

## 7️⃣ ステップ7: 動作確認

### 7-1. アクセステスト

```
以下のURLにアクセスして確認:

1. メインページ
   → https://infohiroki.com

2. ヘルスチェック
   → https://infohiroki.com/health
   → 期待値: {"status":"ok"}

3. ブログページ
   → https://infohiroki.com/blog
```

### 7-2. 確認項目チェックリスト

```
✅ HTTPSで正常にアクセスできる
   └─ 鍵マークが表示されているか

✅ `/health` が正常応答
   └─ {"status":"ok"} が返ってくるか

✅ 静的ファイルが正常に表示
   ├─ CSS が適用されているか
   ├─ JavaScript が動作しているか
   └─ 画像が表示されているか

✅ ブログ記事一覧が表示される
   └─ articles/ 内の記事が表示されるか

✅ 個別記事が正常に表示される
   └─ 記事ページが開けるか
```

---

## 🔄 更新・デプロイフロー

### 通常の更新作業

```bash
# 1. コードを変更
vim main.go
# または
code main.go

# 2. ローカルでテスト
go run main.go
# ブラウザで http://localhost:8080 を確認

# 3. コミット
git add .
git commit -m "✨ 新機能追加"

# 4. プッシュ
git push

# 5. Railwayが自動デプロイ（2〜5分）
# 6. https://infohiroki.com で確認
```

### 新記事の追加フロー

```bash
# 1. 記事を作成
vim articles/2025-09-30-new-article.md

# Markdown形式で執筆:
# タイトル（# で開始）
# 本文

# 2. コミット・プッシュ
git add articles/
git commit -m "✨ 新記事追加: タイトル"
git push

# 3. Railway自動デプロイ（2〜5分）
# 4. 記事が公開される
#    → https://infohiroki.com/blog で確認
```

### デプロイの流れ

```
Git push
  ↓
GitHub (Webhookトリガー)
  ↓
Railway
  ├─ コード取得
  ├─ Dockerビルド
  ├─ ヘルスチェック
  └─ デプロイ完了
  ↓
本番公開
  └─ 2〜5分で反映
```

---

## 🐛 トラブルシューティング

### 問題1: サイトにアクセスできない

#### 症状

```
❌ https://infohiroki.com にアクセスできない
❌ DNS_PROBE_FINISHED_NXDOMAIN エラー
❌ ERR_NAME_NOT_RESOLVED エラー
```

#### 確認事項

```
1. Cloudflareのステータス
   → 「アクティブ」になっているか？

2. DNSレコード
   → CNAMEレコードが正しく設定されているか？
   → タイプ: CNAME
   → 名前: @
   → ターゲット: infohiroki-go-production.up.railway.app

3. Railwayのデプロイ
   → ステータスが「Success」か？
```

#### 解決策

```bash
# DNSの伝播を確認
dig infohiroki.com

# 期待される結果:
# infohiroki.com. IN CNAME infohiroki-go-production.up.railway.app.

# CloudflareのCDNキャッシュをクリア
# 手順:
# 1. Cloudflareダッシュボード
# 2. キャッシング → すべてをパージ
```

---

### 問題2: 記事が表示されない

#### 症状

```
❌ ブログページは開くが記事が表示されない
❌ 記事一覧が空
❌ 個別記事ページが404エラー
```

#### 確認事項

```
1. articles/ ディレクトリ
   └─ Markdownファイルが存在するか？
   └─ ファイル名が正しいか？
      例: 2025-09-30-article-title.md

2. Dockerfile
   └─ articles/ のコピーが設定されているか？
   COPY --from=builder /app/articles ./articles

3. main.go
   └─ initializeData() が実行されているか？
   └─ articlesディレクトリを読み込んでいるか？
```

#### 解決策

```bash
# ローカルで確認
ls articles/
# Markdownファイルが表示されるか確認

# ローカルで実行
go run main.go

# ブラウザで確認
open http://localhost:8080/blog

# 記事が表示されれば、Dockerfileの問題
# Dockerfile を確認:
vim Dockerfile

# この行があるか確認:
# COPY --from=builder /app/articles ./articles

# 修正後、再デプロイ:
git add Dockerfile
git commit -m "🐛 Fix: articles directory not copied"
git push
```

---

### 問題3: デプロイが失敗する

#### 症状

```
❌ Railwayのデプロイステータスが「Failed」
❌ ビルドエラーが発生
❌ アプリケーションが起動しない
```

#### 確認事項

```
1. railway.toml
   └─ 存在するか？
   └─ 設定が正しいか？

2. Dockerfile
   └─ 構文エラーがないか？
   └─ 必要なファイルがコピーされているか？

3. ビルドログ
   └─ Railwayのログにエラーメッセージがあるか？
```

#### 解決策

```bash
# ローカルでDockerビルドを試す
docker build -t infohiroki-go .

# ビルドが成功したら実行
docker run -p 8080:8080 infohiroki-go

# ブラウザで確認
open http://localhost:8080

# エラーがなければRailwayに再デプロイ
git push

# それでもエラーが出る場合:
# Railwayのログを確認
# → Deploymentsタブ → ログを読む
# → エラーメッセージに従って修正
```

---

### 問題4: SSL証明書エラー

#### 症状

```
❌ 「接続が安全ではありません」エラー
❌ NET::ERR_CERT_AUTHORITY_INVALID
❌ SSL証明書の警告
```

#### 確認事項

```
1. CloudflareのSSL設定
   └─ 暗号化モードは何か？

2. Railwayのカスタムドメイン
   └─ ステータスが「Active」か？

3. 待機時間
   └─ 設定後、十分な時間が経過したか？
```

#### 解決策

```
手順:
  1. Cloudflareダッシュボードを開く

  2. SSL/TLS → 概要

  3. 暗号化モードを確認
     ❌ Off
     ❌ Flexible
     ✅ Full（推奨）
     ✅ Full (strict)

  4. 「Full」に設定

  5. 数分待ってから再アクセス

注意:
  └─ SSL証明書の発行には最大15分かかる場合があります
```

---

## 📝 必要な設定ファイル

### railway.toml

```toml
# Railway設定ファイル

[build]
builder = "DOCKERFILE"
dockerfilePath = "Dockerfile"

[deploy]
# 起動コマンド
startCommand = "./main"

# ヘルスチェック
healthcheckPath = "/health"
healthcheckTimeout = 100

# 再起動ポリシー
restartPolicyType = "ON_FAILURE"
restartPolicyMaxRetries = 10
```

### Dockerfile

```dockerfile
# Go 1.21のマルチステージビルド
FROM golang:1.21-alpine AS builder

# 作業ディレクトリ設定
WORKDIR /app

# Go modulesファイルをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 本番用の軽量イメージ
FROM alpine:latest

# セキュリティ更新とCA証明書をインストール
RUN apk --no-cache add ca-certificates

# 作業ディレクトリ設定
WORKDIR /root/

# ビルドしたバイナリをコピー
COPY --from=builder /app/main .

# 静的ファイル・テンプレート・記事をコピー
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/articles ./articles

# ポート設定
EXPOSE 8080

# アプリケーション実行
CMD ["./main"]
```

### .env.example

```bash
# アプリケーション設定
PORT=8080
GIN_MODE=release

# Railway環境では自動的にPORTが設定されます
# ローカル開発時は上記のPORTを使用します
```

### main.go（ヘルスチェックエンドポイント）

```go
// Health check endpoint for Railway/Cloudflare
r.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
})
```

---

## 🔗 参考リンク

### 公式ドキュメント

```
Cloudflare:
  └─ https://developers.cloudflare.com/

Railway:
  └─ https://docs.railway.app/

お名前.com:
  └─ https://www.onamae.com/guide/

Go言語:
  └─ https://go.dev/doc/
```

### 関連サービス

```
GitHub:
  └─ https://docs.github.com/

Docker:
  └─ https://docs.docker.com/

Gin Framework:
  └─ https://gin-gonic.com/docs/
```

---

## ✅ デプロイ完了チェックリスト

```
全体フロー確認:
  □ Cloudflareアカウント作成済み
  □ Railwayアカウント作成済み
  □ GitHubリポジトリ作成済み

Cloudflare設定:
  □ サイト追加完了
  □ ネームサーバー確認済み
  □ ステータスが「アクティブ」
  □ CNAMEレコード設定完了
  □ プロキシ有効（オレンジ）

お名前.com設定:
  □ ネームサーバー変更完了
  □ 既存ネームサーバー削除済み

Railway設定:
  □ プロジェクト作成完了
  □ GitHubリポジトリ連携済み
  □ デプロイ成功（Status: Success）
  □ カスタムドメイン追加済み
  □ SSL証明書発行完了（Active）

動作確認:
  □ HTTPSでアクセス可能
  □ /health が正常応答
  □ 静的ファイル読み込み成功
  □ ブログ記事表示成功
  □ 全ページ正常動作

自動デプロイ確認:
  □ git push でデプロイされる
  □ 2〜5分で反映される
```

---

## 💡 運用のベストプラクティス

### モニタリング

```
監視項目:
  ├─ Railwayのログ確認
  ├─ デプロイ成功率の追跡
  ├─ レスポンスタイムの監視
  └─ エラーログの定期確認

頻度:
  └─ デプロイ後は必ず確認
  └─ 週1回の定期チェック
```

### バックアップ

```
重要ファイル:
  ├─ ソースコード → GitHubで管理
  ├─ 記事ファイル → articles/
  ├─ 設定ファイル → railway.toml, Dockerfile
  └─ 環境変数 → Railway Variables

戦略:
  └─ すべてGit管理下に置く
  └─ 定期的にpush
```

### セキュリティ

```
注意点:
  ├─ 環境変数に秘密情報を保存
  ├─ .gitignore で機密ファイルを除外
  ├─ HTTPS通信を強制
  └─ 定期的な依存関係の更新

コマンド:
  # 依存関係の更新
  go get -u ./...
  go mod tidy
```

---

## ✨ 300字要約

お名前.comで取得したドメインをCloudflare経由でRailwayのGoアプリに接続する完全ガイドです。7つのステップで構成され、①Cloudflareでドメイン追加、②お名前.comでネームサーバー変更、③アクティブ化確認、④RailwayでGoアプリデプロイ、⑤CloudflareでCNAMEレコード設定、⑥Railwayでカスタムドメイン設定、⑦動作確認、という流れで進めます。Cloudflareの無料CDN/SSL、Railwayの自動デプロイ機能により、GitHubへのpush後2〜5分で自動的に本番反映される効率的な開発フローを実現します。トラブルシューティングとして、アクセス不可・記事非表示・デプロイ失敗・SSL証明書エラーの4つの問題と解決策を詳述し、実務で即座に活用できる実践的なガイドとなっています。