# 🚀 infohiroki.com デプロイ手順

## 📊 構成図

```
Internet
  ↓
Cloudflare (DNS / CDN / SSL)
  ↓
Railway (Go App)
  ↓
GitHub (ソースコード)
```

---

## 🎯 事前準備

### 必要なもの

- [x] お名前.comで取得したドメイン (`infohiroki.com`)
- [x] Cloudflareアカウント
- [x] Railwayアカウント
- [x] GitHubリポジトリ (`infoHiroki/infohiroki-go`)

---

## ステップ1: Cloudflareでドメインを追加

### 1-1. サイトの追加

1. [Cloudflareダッシュボード](https://dash.cloudflare.com/)にログイン
2. 「サイトを追加」をクリック
3. `infohiroki.com` を入力
4. プランを選択（Freeで十分）
5. 「サイトを追加」をクリック

### 1-2. ネームサーバーを確認

Cloudflareが表示する**2つのネームサーバー**をメモ：

```
例:
clyde.ns.cloudflare.com
nagali.ns.cloudflare.com
```

> ⚠️ **重要**: この2つのネームサーバーは次のステップで使います

---

## ステップ2: お名前.comでネームサーバーを変更

### 2-1. ネームサーバー設定画面を開く

1. [お名前.com Navi](https://navi.onamae.com/)にログイン
2. 左メニュー「ドメイン」→「ネームサーバー/DNS」を選択
3. 「ネームサーバーの設定」を選択

### 2-2. ドメインを選択

1. 対象ドメイン `infohiroki.com` にチェック
2. 「その他のネームサーバーを使う」を選択

### 2-3. Cloudflareのネームサーバーを入力

1. ネームサーバー1: `clyde.ns.cloudflare.com`
2. ネームサーバー2: `nagali.ns.cloudflare.com`
3. 既存のネームサーバー（01.dnsv.jp〜04.dnsv.jp）を削除
4. 「確認」→「設定する」

> ⚠️ **注意**:
> - **DNSレコード設定画面ではありません**
> - **ネームサーバー設定画面**で変更してください
> - 反映まで最大24時間かかる場合があります

---

## ステップ3: Cloudflareでアクティブ化を確認

### 3-1. 状態確認

1. Cloudflareダッシュボードに戻る
2. `infohiroki.com` のステータスを確認
3. 「**アクティブ**」になるまで待つ（通常1〜24時間）

### 3-2. アクティブ化後の確認

ステータスが「アクティブ」になったら：
- ✅ ネームサーバー変更が完了
- ✅ CloudflareがDNSを管理開始
- ✅ 次のステップに進める

---

## ステップ4: RailwayでGoアプリをデプロイ

### 4-1. プロジェクト作成

1. [Railway](https://railway.app/)にログイン
2. 「New Project」をクリック
3. 「Deploy from GitHub repo」を選択
4. `infoHiroki/infohiroki-go` を選択
5. 自動的にビルド・デプロイが開始される

### 4-2. 環境変数の設定（必要な場合）

1. プロジェクトの「Variables」タブを開く
2. 必要な環境変数を追加：

```bash
PORT=8080
GIN_MODE=release
```

> 💡 **Note**: `PORT`はRailwayが自動設定するので通常不要

### 4-3. デプロイ完了確認

1. Deploymentsタブで「Success」を確認
2. 自動生成されたURL（例: `infohiroki-go-production.up.railway.app`）をメモ

---

## ステップ5: CloudflareでDNSレコードを設定

### 5-1. DNS設定画面を開く

1. Cloudflareダッシュボードで `infohiroki.com` を選択
2. 左メニュー「DNS」→「レコード」を選択

### 5-2. CNAMEレコードを追加

| 項目 | 値 |
|------|-----|
| **タイプ** | CNAME |
| **名前** | `@` （ルートドメイン） |
| **ターゲット** | `infohiroki-go-production.up.railway.app` |
| **プロキシステータス** | オレンジ（プロキシ有効） |
| **TTL** | 自動 |

「保存」をクリック

> 💡 **Tip**: wwwサブドメインも追加する場合は、同様にCNAMEレコードを追加（名前を`www`にする）

---

## ステップ6: Railwayでカスタムドメインを設定

### 6-1. ドメイン設定画面を開く

1. Railwayのプロジェクトページ
2. 「Settings」タブ
3. 「Domains」セクション

### 6-2. カスタムドメインを追加

1. 「Add Custom Domain」をクリック
2. `infohiroki.com` を入力
3. 「Add Domain」をクリック
4. Railwayが自動的にCloudflareと連携

### 6-3. SSL証明書の確認

- Railwayが自動的にSSL証明書を発行
- 通常1〜5分で完了
- ステータスが「Active」になれば完了

---

## ステップ7: 動作確認

### 7-1. アクセステスト

以下のURLにアクセスして確認：

```
https://infohiroki.com
https://infohiroki.com/health
https://infohiroki.com/blog
```

### 7-2. 確認項目

- [x] HTTPSで正常にアクセスできる
- [x] `/health` が `{"status":"ok"}` を返す
- [x] 静的ファイル（CSS/JS/画像）が正常に表示される
- [x] ブログ記事一覧が表示される
- [x] 個別記事が正常に表示される

---

## 🔄 更新・デプロイフロー

### 通常の更新作業

```bash
# 1. コードを変更
vim main.go

# 2. ローカルでテスト
go run main.go

# 3. コミット
git add .
git commit -m "✨ 新機能追加"

# 4. プッシュ
git push

# 5. Railwayが自動デプロイ（2〜5分）
# 6. https://infohiroki.com で確認
```

### 新記事の追加

```bash
# 1. 記事を作成
vim articles/2025-09-30-new-article.md

# 2. コミット・プッシュ
git add articles/
git commit -m "✨ 新記事追加: タイトル"
git push

# 3. Railway自動デプロイ
# 4. 記事が公開される
```

---

## 🐛 トラブルシューティング

### 問題1: サイトにアクセスできない

**症状**: `infohiroki.com` にアクセスできない

**確認事項**:
1. Cloudflareのステータスが「アクティブ」か？
2. DNSレコードが正しく設定されているか？
3. Railwayのデプロイが成功しているか？

**解決策**:
```bash
# DNSの伝播を確認
dig infohiroki.com

# CloudflareのCDNキャッシュをクリア
# Cloudflareダッシュボード → キャッシング → すべてをパージ
```

### 問題2: 記事が表示されない

**症状**: ブログ記事が表示されない

**確認事項**:
1. `articles/` ディレクトリにMarkdownファイルがあるか？
2. Dockerfileで `COPY --from=builder /app/articles ./articles` が設定されているか？
3. `main.go` の `initializeData()` が実行されているか？

**解決策**:
```bash
# ローカルで確認
ls articles/
go run main.go

# ブラウザで http://localhost:8080/blog を確認
```

### 問題3: デプロイが失敗する

**症状**: Railwayでデプロイが失敗する

**確認事項**:
1. `railway.toml` が存在するか？
2. `Dockerfile` が正しいか？
3. ビルドログにエラーメッセージがあるか？

**解決策**:
```bash
# ローカルでDockerビルドを試す
docker build -t infohiroki-go .
docker run -p 8080:8080 infohiroki-go

# エラーがなければRailwayに再デプロイ
git push
```

### 問題4: SSL証明書エラー

**症状**: 「接続が安全ではありません」エラー

**確認事項**:
1. CloudflareのSSL設定が「フル」になっているか？
2. Railwayのカスタムドメインが「Active」か？

**解決策**:
1. Cloudflareダッシュボード → SSL/TLS → 概要
2. 暗号化モードを「フル」に設定
3. 数分待ってから再アクセス

---

## 📝 設定ファイル

### railway.toml

```toml
[build]
builder = "DOCKERFILE"
dockerfilePath = "Dockerfile"

[deploy]
startCommand = "./main"
healthcheckPath = "/health"
healthcheckTimeout = 100
restartPolicyType = "ON_FAILURE"
restartPolicyMaxRetries = 10
```

### Dockerfile

```dockerfile
# Go 1.21のマルチステージビルド
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 本番用の軽量イメージ
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/articles ./articles

EXPOSE 8080
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

---

## 🔗 参考リンク

- [Cloudflare ドキュメント](https://developers.cloudflare.com/)
- [Railway ドキュメント](https://docs.railway.app/)
- [お名前.com サポート](https://www.onamae.com/guide/)
- [Go公式ドキュメント](https://go.dev/doc/)

---

## ✅ チェックリスト

デプロイ完了時の確認項目：

- [ ] Cloudflareがアクティブ
- [ ] DNSレコードが設定済み
- [ ] Railwayのデプロイが成功
- [ ] カスタムドメインが有効
- [ ] HTTPSでアクセス可能
- [ ] `/health` が正常応答
- [ ] ブログ記事が表示される
- [ ] 静的ファイルが読み込まれる
- [ ] 全ページが正常動作

---

**更新日**: 2025-09-30
**作成者**: Claude Code