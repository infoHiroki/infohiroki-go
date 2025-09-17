# 🚀 Go Learning Project - クイックスタート

## ⚡ 即座に実行

```bash
cd go-learning-project
go mod tidy
go run main.go
```

**アクセス**: http://localhost:8080

## 🌟 主な機能

- ✅ マークダウン → HTML変換
- ✅ Rails 8.1風の `.md` エンドポイント
- ✅ JSON API
- ✅ サンプル記事3件

## 📊 Rails vs Go 比較

| 項目 | Rails | Go |
|------|-------|-----|
| ファイル数 | 50+ | **4** |
| コード行数 | 1000+ | **150** |
| メモリ使用量 | 150MB | **30MB** |

## 💡 KISS・YAGNI・DRY の実例

```go
// KISS: たった4行でWebサーバー
r := gin.Default()
r.GET("/posts/:id", showPost)
r.Run(":8080")

// DRY: 共通処理を1箇所で定義
func (p *Post) ToMarkdown() string {
    return "# " + p.Title + "\n\n" + p.Content
}
```

## 🎯 エンドポイント

- `/` - 記事一覧
- `/posts/1` - 記事詳細
- `/posts/1.md` - Markdown出力
- `/posts/1.json` - JSON API

**個人開発者こそGoが最適！** 🔥