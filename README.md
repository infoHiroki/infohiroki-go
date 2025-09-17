# infoHiroki Website Go版

既存のinfoHirokiウェブサイトをピクセルパーフェクトレベルでGoに移植し、ブログシステムをMarkdown化するプロジェクト。

## 🎯 プロジェクト目標

- **ピクセルパーフェクト移植**: 既存デザインの完全再現
- **Markdownブログ**: HTML記事（94件）をMarkdown化
- **高速検索**: サーバーサイド全文検索
- **SEO完全対応**: 構造化データ、OGP、サイトマップ維持

## 🏗️ 技術スタック

- **Backend**: Go 1.21 + Gin + GORM + SQLite
- **Frontend**: 既存CSS/JS完全移植（1,958行CSS + JavaScript）
- **Markdown**: blackfriday v2
- **Search**: SQLite FTS（全文検索）
- **Security**: bluemonday（HTMLサニタイズ）

## 📊 移植元サイト分析

- **HTMLページ**: 9ページ（index, blog, services, products, results, about, faq, contact, scenarios）
- **ブログ記事**: 94記事（html-files/*.html）
- **CSS**: 1,958行の完全デザインシステム
- **画像**: 49個のアイコン・画像ファイル
- **JavaScript**: ハンバーガーメニュー、検索、クリップボード機能

## 🚀 クイックスタート

```bash
# 依存関係インストール
go mod tidy

# 開発サーバー起動
go run main.go

# ブラウザでアクセス
open http://localhost:8080
```

## 📁 プロジェクト構造

```
infohiroki-go/
├── CLAUDE.md                   # 開発ガイド
├── README.md                   # このファイル
├── main.go                     # メインアプリケーション
├── src/                        # ソースコード
│   ├── models/                 # データモデル
│   ├── handlers/               # ハンドラー
│   ├── services/               # ビジネスロジック
│   └── utils/                  # ユーティリティ
├── templates/                  # Goテンプレート
├── static/                     # 静的ファイル（CSS/JS/画像）
├── markdown/                   # Markdownブログ記事
└── database/                   # SQLiteデータベース
```

## 🔧 開発コマンド

```bash
# データベース初期化
go run migrate.go

# HTML→Markdown変換
go run tools/html_to_markdown.go

# 新記事追加
touch markdown/2025-01-01-new-article.md
go run migrate.go --reload-posts
```

## 📝 詳細ドキュメント

詳細な開発ガイド、アーキテクチャ、設定については [CLAUDE.md](./CLAUDE.md) を参照してください。

## 🎨 デザイン

既存のinfoHirokiウェブサイトのデザインをピクセル単位で完全移植。レスポンシブデザイン、モバイルナビゲーション、アニメーション効果もすべて維持。

## 🔍 検索機能

SQLite FTS（Full-Text Search）による高速な記事検索とタグベースフィルタリング機能を実装。

---

© 2024 infoHiroki. All rights reserved.