# infoHiroki Website Go移植プロジェクト - Claude Code設定

## 📋 プロジェクト概要

### システム名
infoHiroki Website Go版（ピクセルパーフェクト移植）

### 目的
既存のinfoHirokiウェブサイト（Vanilla HTML/CSS/JS）をGoで完全移植し、ブログシステムをMarkdown化

### コア機能
- **ピクセルパーフェクト移植**: 既存デザインの完全再現
- **Markdownブログ**: HTML記事（94件）をMarkdown化
- **高速検索**: サーバーサイド全文検索
- **SEO完全対応**: 構造化データ、OGP、サイトマップ

## 🏗️ 技術アーキテクチャ

### 基本構成
```
Go + Gin + GORM + SQLite + blackfriday v2
```

### 技術スタック
| 層 | 技術 | バージョン | 理由 |
|---|------|-----------|------|
| **Backend** | Go | 1.21+ | 型安全性・高速性・シンプル設計 |
| **Web Framework** | Gin | 1.9+ | 軽量・高速・豊富なミドルウェア |
| **Database** | SQLite | 3.x | ファイルベース・管理不要・FTS対応 |
| **Template** | Go標準template | - | サーバーサイドレンダリング |
| **Markdown** | blackfriday v2 | 2.1+ | 高速・GitHub互換・カスタマイズ可能 |
| **Frontend** | 既存CSS/JS移植 | - | ピクセルパーフェクト維持 |

### 移植元サイト分析
- **HTMLページ**: 9ページ（index, blog, services等）
- **ブログ記事**: 94記事（html-files/*.html）
- **CSS**: 1,958行の完全デザインシステム
- **画像**: 49個のアイコン・画像ファイル
- **JavaScript**: ハンバーガーメニュー、検索、クリップボード機能

## 📁 プロジェクト構造

```
infohiroki-go/
├── CLAUDE.md                   # このファイル
├── README.md                   # プロジェクト説明
├── go.mod                      # Go依存関係
├── go.sum                      # 依存関係ハッシュ
├── main.go                     # Gin メインアプリケーション
├── config.py                   # アプリケーション設定
├── migrate.go                  # データベース初期化・マイグレーション
├── convert_html_to_md.go       # HTML→Markdown変換ツール
├── src/                        # ソースコード
│   ├── models/                 # データモデル
│   │   ├── page.go            # 固定ページモデル
│   │   └── blog_post.go       # ブログ記事モデル
│   ├── handlers/              # ハンドラー（コントローラー）
│   │   ├── home.go            # ホームページ
│   │   ├── blog.go            # ブログ機能
│   │   ├── pages.go           # 固定ページ
│   │   └── api.go             # JSON API
│   ├── services/              # ビジネスロジック
│   │   ├── blog_service.go    # ブログ検索・フィルタリング
│   │   └── markdown_service.go # Markdown処理
│   └── utils/                 # ユーティリティ
│       ├── template.go        # テンプレートヘルパー
│       └── seo.go             # SEO対応ヘルパー
├── templates/                 # Go テンプレート
│   ├── base.html              # ベーステンプレート
│   ├── index.html             # ホームページ
│   ├── blog/                  # ブログテンプレート
│   │   ├── list.html          # 記事一覧
│   │   └── detail.html        # 記事詳細
│   └── pages/                 # 固定ページテンプレート
│       ├── services.html      # サービス
│       ├── products.html      # 開発製品
│       ├── results.html       # 実績
│       ├── about.html         # スキルスタック
│       ├── faq.html           # FAQ
│       └── contact.html       # お問い合わせ
├── static/                    # 静的ファイル（完全移植）
│   ├── css/
│   │   └── style.css          # 1,958行CSS完全コピー
│   ├── js/
│   │   └── main.js            # JavaScript機能移植
│   └── images/                # 49個の画像ファイル
│       ├── logo.svg
│       ├── hero.svg
│       └── icons/             # 技術アイコン群
├── markdown/                  # Markdownブログ記事
│   ├── 2024-03-27-notion-save-to-notion-extension.md
│   ├── 2024-03-29-chatgpt-reskilling-guide.md
│   └── ... (94記事)
├── database/                  # データベース関連
│   ├── infohiroki.db          # SQLiteデータベース
│   └── migrations/            # マイグレーションファイル
├── tools/                     # 開発ツール
│   └── html_to_markdown.go    # 変換スクリプト
└── docs/                      # ドキュメント
    ├── migration_guide.md     # 移行ガイド
    └── deployment.md          # デプロイ手順
```

## 🎯 開発原則

### 設計思想
- **Pixel Perfect**: 既存デザインの完全再現
- **KISS**: Keep It Simple, Stupid - シンプル設計
- **YAGNI**: You Aren't Gonna Need It - 必要な機能のみ
- **DRY**: Don't Repeat Yourself - コード重複排除

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
module infohiroki-go

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/russross/blackfriday/v2 v2.1.0
    gorm.io/driver/sqlite v1.5.4
    gorm.io/gorm v1.25.5
    github.com/microcosm-cc/bluemonday v1.0.25  // HTMLサニタイズ
)
```

## 📊 データベース設計

### メインテーブル
```sql
-- 固定ページ
CREATE TABLE pages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    content TEXT,
    template TEXT,
    meta_description TEXT,
    meta_keywords TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- ブログ記事
CREATE TABLE blog_posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    description TEXT,
    tags TEXT,  -- JSON形式
    icon TEXT,
    created_at DATE,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    published BOOLEAN DEFAULT 1
);

-- FTS（全文検索）テーブル
CREATE VIRTUAL TABLE blog_posts_fts USING fts5(
    title, content, description, tags,
    content='blog_posts',
    content_rowid='id'
);

-- 設定管理
CREATE TABLE settings (
    id INTEGER PRIMARY KEY,
    key TEXT UNIQUE NOT NULL,
    value TEXT NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## 🔄 実行フロー

### HTMLからMarkdown変換フロー
1. **HTML解析**: 94記事のHTMLを解析
2. **メタデータ抽出**: title, description, tags, created_at
3. **コンテンツ変換**: HTML → Markdownに自動変換
4. **データベース投入**: 変換済みデータをSQLiteに格納

### ブログ表示フロー
1. **リクエスト受信**: `/blog` または `/blog/:slug`
2. **データベース検索**: SQLite FTSで高速検索
3. **Markdown処理**: blackfriday でHTMLレンダリング
4. **テンプレート適用**: Go templateで最終HTML生成
5. **レスポンス**: ピクセルパーフェクトなHTMLを返却

## 🚀 デプロイ環境

### 推奨VPS仕様
- **CPU**: 1コア以上
- **メモリ**: 512MB以上
- **ストレージ**: 10GB以上
- **OS**: Ubuntu 22.04 LTS

### 本番環境構成
```
Internet → Nginx → Go App → SQLite
                    ↓
            Static Files (CSS/JS/Images)
```

## 🔒 セキュリティ対策

### 基本セキュリティ
- **HTTPS**: SSL/TLS証明書
- **CSRF**: Gin CSRFミドルウェア
- **XSS**: HTMLサニタイズ（bluemonday）
- **SQLインジェクション**: GORM ORM使用

### パフォーマンス対策
- **静的ファイルキャッシュ**: Nginx設定
- **データベースインデックス**: 検索性能最適化
- **Gzip圧縮**: レスポンス圧縮

## 📝 運用・保守

### 開発・デバッグ
```bash
# ホットリロード開発
air  # air をインストール後

# ログ出力確認
tail -f logs/app.log

# データベース確認
sqlite3 database/infohiroki.db
```

### 新記事追加
```bash
# Markdownファイル作成
touch markdown/2025-01-01-new-article.md

# データベース再読み込み
go run migrate.go --reload-posts
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

## 📚 開発ガイドライン

### 新機能追加時
1. 既存デザインとの整合性確認
2. モバイル対応確認
3. SEO対応（メタタグ等）確認
4. パフォーマンステスト実行

### コミットルール
```
✨ 新機能追加
🐛 バグ修正
♻️ リファクタリング
📝 ドキュメント更新
🎨 デザイン調整
🚀 パフォーマンス改善
```

### 品質基準
- **デザイン**: ピクセル単位での一致
- **機能**: 既存JavaScript機能の完全再現
- **SEO**: 検索エンジン対応維持
- **パフォーマンス**: 既存サイト比2倍以上高速

---

このCLAUDE.mdにより、infoHirokiウェブサイトの完全移植プロジェクトを効率的に進めることができます。