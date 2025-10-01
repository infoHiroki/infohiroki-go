# infoHiroki Website Go移植プロジェクト - Claude Code設定

## 📋 プロジェクト概要

### システム名
infoHiroki Website Go版（ピクセルパーフェクト移植）

### 目的
既存のinfoHirokiウェブサイト（Vanilla HTML/CSS/JS）をGoで完全移植し、ブログシステムをMarkdown化

### コア機能
- **ピクセルパーフェクト移植**: 既存デザインの完全再現
- **Markdownブログ**: HTML記事（95件）をMarkdown化
- **高速検索**: サーバーサイド全文検索
- **SEO完全対応**: 構造化データ、OGP、サイトマップ

## 🏗️ 技術アーキテクチャ

### 基本構成
```
Go + Gin + blackfriday v2（ファイルベース・シンプル構成）
```

### 技術スタック
| 層 | 技術 | バージョン | 理由 |
|---|------|-----------|------|
| **Backend** | Go | 1.21+ | 型安全性・高速性・シンプル設計 |
| **Web Framework** | Gin | 1.9+ | 軽量・高速・豊富なミドルウェア |
| **Storage** | File-based | - | Markdownファイル直接読み込み・シンプル |
| **Template** | Go標準template | - | サーバーサイドレンダリング |
| **Markdown** | blackfriday v2 | 2.1+ | 高速・GitHub互換・カスタマイズ可能 |
| **Frontend** | 既存CSS/JS移植 | - | ピクセルパーフェクト維持 |

### 移植元サイト分析
- **HTMLページ**: 9ページ（index, blog, services等）
- **ブログ記事**: 95記事（完全Markdown化済み）
- **CSS**: 1,958行の完全デザインシステム
- **画像**: 49個のアイコン・画像ファイル
- **JavaScript**: ハンバーガーメニュー、検索、クリップボード機能

## 📁 プロジェクト構造

```
go-learning-project/
├── CLAUDE.md                   # このファイル
├── README.md                   # プロジェクト説明
├── go.mod                      # Go依存関係
├── go.sum                      # 依存関係ハッシュ
├── main.go                     # Gin メインアプリケーション（単一ファイル構成）
├── archive/                    # 開発初期資料・参考ファイル
│   ├── go_for_solo_developers.md
│   ├── go_kiss_yagni_dry_example.go
│   ├── info.md
│   └── quick-start.md
├── src/                        # ソースコード
│   └── models/                 # データモデル
│       ├── page.go            # 固定ページモデル
│       └── blog_post.go       # ブログ記事モデル
├── templates/                 # Go テンプレート
│   ├── index.html             # ホームページ
│   ├── blog.html              # ブログ一覧
│   ├── blog_detail.html       # ブログ記事詳細
│   ├── services.html          # サービス
│   ├── products.html          # 開発製品
│   ├── results.html           # 実績
│   ├── about.html             # スキルスタック
│   ├── faq.html               # FAQ
│   ├── contact.html           # お問い合わせ
│   └── 404.html               # エラーページ
├── static/                    # 静的ファイル（完全移植）
│   ├── css/
│   │   └── style.css          # 1,958行CSS完全コピー
│   ├── js/
│   │   └── main.js            # JavaScript機能移植
│   └── images/                # 49個の画像ファイル
│       ├── logo.svg
│       ├── hero.svg
│       └── note/              # ブログ画像
└── articles/                  # Markdown記事（95記事）
    ├── 2024-03-27-notion-save-to-notion-extension.md
    ├── 2024-03-29-chatgpt-reskilling-guide.md
    ├── 2025-09-20-go-complete-history.md
    └── ... (95個のMarkdownファイル)
```

## 🎯 開発原則

### 設計思想
- **Pixel Perfect**: 既存デザインの完全再現
- **KISS (Keep It Simple, Stupid)**:
  - 複雑なデータベース設計を避け、ファイルベースで実装
  - 必要最小限の機能のみ実装
  - 理解しやすく保守しやすいコード
- **YAGNI (You Aren't Gonna Need It)**:
  - 将来的に必要になるかもしれない機能は実装しない
  - 現在必要な機能のみに集中
  - 過度な汎用化を避ける
- **DRY (Don't Repeat Yourself)**:
  - 同一の処理を複数箇所に書かない
  - 共通処理の関数化・モジュール化
  - テンプレートの再利用

### コーディング規約
- **Go**: Go標準スタイル準拠
- **命名**: snake_case（DB）、camelCase（Go）
- **型ヒント**: 全関数に型定義必須
- **エラーハンドリング**: 適切なerror処理

### 移植品質基準
- **ピクセル完全一致**: デザイン100%維持
- **機能完全移植**: JavaScript機能すべて再現
- **SEO完全対応**: メタデータ、構造化データ維持
- **パフォーマンス向上**: サーバーサイドレンダリングで高速化

## 🔧 環境・設定

### 開発環境
```bash
# Go依存関係インストール
go mod tidy

# 開発サーバー起動
go run main.go

# データベース初期化
go run migrate.go

# HTML→Markdown変換
go run tools/html_to_markdown.go
```

### 必要な環境変数
```bash
# アプリケーション設定
export GIN_MODE="debug"  # debug/release
export PORT="8080"

# データベース
export DATABASE_PATH="database/infohiroki.db"

# Markdown設定
export MARKDOWN_DIR="markdown"
export STATIC_DIR="static"
```

### go.mod
```go
module go-learning-project

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/russross/blackfriday/v2 v2.1.0
)
```

## 📁 ファイルベース設計

### データ管理方式
**データベースレス設計**: データベースを使わず、ファイルシステムで直接管理

```go
// ブログ記事構造体
type BlogPost struct {
    ID          uint      `json:"id"`
    Slug        string    `json:"slug"`
    Title       string    `json:"title"`
    Content     string    `json:"content"`
    Description string    `json:"description"`
    Icon        string    `json:"icon"`
    MarkdownPath string   `json:"markdown_path"`
    CreatedDate time.Time `json:"created_date"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    Published   bool      `json:"published"`
}
```

### ファイル管理
- **記事**: `articles/*.md` - 95個のMarkdownファイル
- **メタデータ**: ファイル名・内容から自動抽出
- **検索**: メモリ内文字列検索（シンプル・高速）

## 🔄 実行フロー

### アプリケーション起動フロー
1. **記事読み込み**: `articles/`ディレクトリから95個のMarkdownファイル読み込み
2. **メタデータ抽出**: ファイル名・内容から自動抽出
3. **メモリ格納**: 全記事をメモリ上のスライスに格納
4. **サーバー起動**: Ginサーバー起動

### ブログ表示フロー
1. **リクエスト受信**: `/blog` または `/blog/:slug`
2. **メモリ検索**: スライス内検索（高速）
3. **Markdown処理**: blackfriday でHTMLレンダリング
4. **テンプレート適用**: Go templateで最終HTML生成
5. **レスポンス**: ピクセルパーフェクトなHTMLを返却

## 🚀 デプロイ環境

### Railway + Cloudflare構成（本番環境）

#### 全体構成
```
Internet → Cloudflare（CDN/DNS/SSL）
              ↓
          Railway（Go App）
              ↓
      Static Files + Markdown Articles
```

#### 1. お名前.comでのネームサーバー設定

**手順:**
1. お名前.comの管理画面で「**ネームサーバー/DNS**」→「**ネームサーバーの設定**」を選択
2. 対象ドメイン（infohiroki.com）を選択
3. 「**その他のネームサーバーを使う**」を選択
4. Cloudflareのネームサーバーを入力：
   - ネームサーバー1: `clyde.ns.cloudflare.com`
   - ネームサーバー2: `nagali.ns.cloudflare.com`
5. 既存のネームサーバー（01.dnsv.jp〜04.dnsv.jp）を削除
6. 確認して保存

**注意:**
- **DNSレコード設定画面ではなく、ネームサーバー設定画面で変更すること**
- 反映まで最大24時間かかる場合あり

#### 2. Cloudflareでのドメイン設定

**手順:**
1. Cloudflareダッシュボードで「ドメインを追加」
2. ネームサーバーアドレス（2つ）をメモ
3. お名前.comで上記ネームサーバーを設定（前述）
4. Cloudflareのステータスが「**アクティブ**」になるまで待機

**DNS設定:**
```
タイプ: CNAME
名前: @ （またはサブドメイン）
ターゲット: [Railwayのドメイン].railway.app
プロキシステータス: オレンジ（プロキシ有効）
TTL: 自動
```

#### 3. Railwayでのデプロイ設定

**railway.toml:**
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

**Dockerfile:**
```dockerfile
# マルチステージビルド
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

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

**main.goヘルスチェックエンドポイント:**
```go
// Health check endpoint for Railway/Cloudflare
r.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
})
```

**カスタムドメイン設定:**
1. Railwayのプロジェクト → Settings → Domains
2. 「Add Custom Domain」
3. `infohiroki.com`を入力
4. CloudflareのCNAMEレコードと連携

**環境変数（Railway）:**
```
PORT=8080          # Railwayが自動設定
GIN_MODE=release   # 本番モード
```

#### 4. デプロイ確認

**確認URL:**
- メインページ: `https://infohiroki.com`
- ヘルスチェック: `https://infohiroki.com/health`
- ブログ: `https://infohiroki.com/blog`

**確認項目:**
- [ ] HTTPSで正常アクセス可能
- [ ] `/health`が`{"status":"ok"}`を返す
- [ ] 静的ファイル（CSS/JS/画像）が正常に読み込まれる
- [ ] ブログ記事一覧が表示される
- [ ] 個別記事が正常に表示される

#### 5. デプロイ後のキャッシュパージ（重要）

**🚨 問題: デプロイ直後に`ERR_QUIC_PROTOCOL_ERROR`が発生する場合**

Railwayで新バージョンがデプロイされた直後、Cloudflareの古いキャッシュとの不整合で`ERR_QUIC_PROTOCOL_ERROR`が発生することがあります。

**解決手順:**

1. **Cloudflareダッシュボードにアクセス**
   - https://dash.cloudflare.com/ にログイン
   - `infohiroki.com` を選択

2. **キャッシュをパージ**
   - 左メニュー「**キャッシング**」（Caching）をクリック
   - 「**構成**」（Configuration）タブを選択
   - 「**すべてパージ**」（Purge Everything）ボタンをクリック
   - 確認ダイアログで「パージ」を実行

3. **2〜3分待機**
   - キャッシュパージ完了を待つ
   - `https://infohiroki.com` にアクセスして確認

**代替策（パージで解決しない場合）:**

1. **HTTP/3 (QUIC) を無効化**
   - 左メニュー「**ネットワーク**」（Network）をクリック
   - 「**HTTP/3 (with QUIC)**」トグルを**オフ**に切り替え

2. **開発モードを有効化（一時的）**
   - 左メニュー「**概要**」（Overview）をクリック
   - 右側「**開発モード**」（Development Mode）を**オン**
   - 3時間キャッシュが無効化される

**推奨運用:**

```bash
# デプロイフロー
git add .
git commit -m "✨ 新機能追加"
git push

# → Railwayデプロイ完了を待つ（2〜5分）
# → Railway直接URL確認: https://infohiroki-go-production.up.railway.app
# → 大きな更新の場合: Cloudflareでキャッシュパージ実行
# → https://infohiroki.com で最終確認
```

**頻度:**
- 小さな記事追加: パージ不要（5〜10分で自動反映）
- 大きな機能追加・デザイン変更: パージ推奨

### 推奨VPS仕様（代替案）
- **CPU**: 1コア以上
- **メモリ**: 512MB以上
- **ストレージ**: 10GB以上
- **OS**: Ubuntu 22.04 LTS

### VPS本番環境構成（代替案）
```
Internet → Nginx → Go App (File-based)
                    ↓
            Static Files (CSS/JS/Images)
                    ↓
            Markdown Articles (articles/)
```

## 🔒 セキュリティ対策

### 基本セキュリティ
- **HTTPS**: SSL/TLS証明書
- **XSS**: HTMLエスケープ（Go標準template）
- **ファイルアクセス**: 静的ファイルのみ公開

### パフォーマンス対策
- **静的ファイルキャッシュ**: Nginx設定
- **メモリキャッシュ**: 全記事メモリ格納で高速アクセス
- **Gzip圧縮**: レスポンス圧縮

## 📝 運用・保守

### 開発・デバッグ
```bash
# ホットリロード開発
air  # air をインストール後

# ログ出力確認 (Go標準出力)
go run main.go

# 記事確認
ls articles/
```

### 新記事追加
```bash
# Markdownファイル作成
touch articles/2025-01-01-new-article.md

# メタデータ設定（main.goのgenerateMetadataFromSlug関数に追加）
# アイコン画像確認（/images/内）

# サーバー再起動（記事再読み込み）
PORT=8081 go run main.go
```

## 🐛 トラブルシューティング

### よくある問題
1. **テンプレート読み込みエラー**
   - templates/ ディレクトリの権限確認
   - パス設定確認

2. **静的ファイル404エラー**
   - static/ ディレクトリ配置確認
   - Ginの静的ファイルルーティング確認

3. **Markdown表示エラー**
   - blackfriday設定確認
   - HTMLエスケープ設定確認

## 📊 プロジェクト現状

### ✅ 完了済み項目
- **95記事Markdown化**: HTML記事すべてMarkdown変換完了
- **ディレクトリ最適化**: `content/articles/` → `articles/` 移行完了
- **コード整理**: 不要ファイル・バックアップ削除完了
- **ピクセルパーフェクト移植**: 既存デザイン完全再現
- **SEO完全対応**: メタデータ、構造化データ実装済み
- **高速検索機能**: サーバーサイド全文検索実装済み

### 🎯 技術実装詳細
- **ファイルベース**: データベース不使用、シンプル構成
- **メモリキャッシュ**: 全記事をメモリに格納、高速アクセス
- **KISS原則**: 必要最小限の機能のみ実装
- **YAGNI原則**: 将来不要な機能は未実装
- **DRY原則**: コード重複なし、共通処理集約

## 📚 開発ガイドライン

### ✅ HTML→Markdown記事変換完了

#### 🎯 基本原則
- **KISS**: Keep It Simple, Stupid - 必要最小限のシンプル作成
- **元ファイル踏襲**: HTMLファイルの内容・メタデータを完全に踏襲
- **手動作成**: 自動変換ツールは使わず、一件ずつ手動で丁寧に作成
- **品質重視**: 各記事を確実に動作確認してから次へ進む

#### 📝 HTML→Markdown変換完了
**✅ 95記事すべてMarkdown化完了済み**

過去の変換手順（参考）:
1. **HTMLファイル確認**: 元HTMLファイルを読み込み
2. **コンテンツ抽出**: `<article class="article-content">`内の純粋なコンテンツのみ抽出
3. **Markdown作成**: `articles/YYYY-MM-DD-記事名.md`として手動作成
4. **メタデータ設定**: `main.go`内の`generateMetadataFromSlug`関数にケース追加
5. **アイコン設定**: `/images/`内の公式ブランドアイコン画像パスを使用
6. **サーバー再起動**: 変更を反映
7. **ブラウザ確認**: `http://localhost:8081/blog/YYYY-MM-DD-記事名`で動作確認

#### 🏷️ メタデータ設定規則
```go
case strings.Contains(slug, "記事キーワード"):
    return "正確なタイトル",
        "元HTMLファイルと同じ説明文",
        `["元タグ1","元タグ2","元タグ3"]`,
        "/images/公式アイコン.png"
```

#### 🖼️ アイコン使用規則
- **公式ブランドアイコン優先**: `/images/`内の実在するアイコンファイルを使用
- **例**: ChatGPT → `/images/ChatGPT icon.webp`, Notion → `/images/Notion icon.png`
- **避けるもの**: 絵文字アイコン（🤖📝等）は使わない

#### ✅ 品質確認項目
- [ ] タイトルが元HTMLと一致
- [ ] 公式アイコン画像が表示される
- [ ] タグが元HTMLと一致
- [ ] 背景色が白（#ffffff）
- [ ] リンクが正常に動作
- [ ] 画像パスが正確（/images/note/等）
- [ ] Markdownが適切にレンダリング

#### 🚫 禁止事項
- 自動変換ツールの使用
- フロントマターの追加
- 勝手なタイトル・説明文の変更
- 絵文字アイコンの使用
- 一括処理（必ず一件ずつ）

### 新機能追加時
1. 既存デザインとの整合性確認
2. モバイル対応確認
3. SEO対応（メタタグ等）確認
4. パフォーマンステスト実行

### コミットルール
日本語でコミット。文末は動詞が望ましい。
```
✨ 新機能追加
🐛 バグ修正
♻️ リファクタリング
📝 ドキュメント更新
🎨 デザイン調整
🚀 パフォーマンス改善
```
