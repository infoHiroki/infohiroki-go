// Go Learning Project - Markdown Parser Web App
// KISS・YAGNI・DRY の実践例

package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Post モデル - 必要最小限の構造
type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// ToMarkdown - DRY: 共通処理を1箇所で定義
func (p *Post) ToMarkdown() string {
	return "# " + p.Title + "\n\n" + p.Content
}

var db *gorm.DB

func main() {
	// データベース初期化 - KISS: 設定は最小限
	var err error
	db, err = gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})
	if err != nil {
		panic("データベース接続に失敗しました")
	}
	db.AutoMigrate(&Post{})

	// サンプルデータ作成 - YAGNI: 今必要な分だけ
	createSampleData()

	// Gin ルーター設定 - KISS: 最小構成
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Routes - RESTful設計
	r.GET("/", listPosts)
	r.GET("/posts/:id", handlePost)  // 拡張子に応じて処理を分岐

	// サーバー起動
	r.Run(":8080")
}

// handlePost - 拡張子に応じて処理を分岐
func handlePost(c *gin.Context) {
	id := c.Param("id")

	// 拡張子をチェック
	if strings.HasSuffix(id, ".md") {
		// .mdの場合、IDから拡張子を除去
		idStr := strings.TrimSuffix(id, ".md")
		c.Params[0].Value = idStr
		showPostMarkdown(c)
	} else if strings.HasSuffix(id, ".json") {
		// .jsonの場合、IDから拡張子を除去
		idStr := strings.TrimSuffix(id, ".json")
		c.Params[0].Value = idStr
		showPostJSON(c)
	} else {
		// 拡張子なしの場合、HTMLとして表示
		showPost(c)
	}
}

// KISS: 各関数は1つのことだけを担当
func listPosts(c *gin.Context) {
	var posts []Post
	db.Find(&posts)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"posts": posts,
		"title": "Go Markdown Parser",
	})
}

func showPost(c *gin.Context) {
	post := getPostByID(c)
	if post == nil {
		return
	}

	// マークダウンをHTMLに変換
	html := blackfriday.Run([]byte(post.Content))

	c.HTML(http.StatusOK, "show.html", gin.H{
		"post": post,
		"html": template.HTML(html), // HTMLエスケープを回避
	})
}

// Rails 8.1の .md 機能と同等 - YAGNI: 必要な機能のみ
func showPostMarkdown(c *gin.Context) {
	post := getPostByID(c)
	if post == nil {
		return
	}

	c.Data(http.StatusOK, "text/markdown; charset=utf-8", []byte(post.ToMarkdown()))
}

func showPostJSON(c *gin.Context) {
	post := getPostByID(c)
	if post == nil {
		return
	}
	c.JSON(http.StatusOK, post)
}

// DRY: 共通処理を関数化
func getPostByID(c *gin.Context) *Post {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID"})
		return nil
	}

	var post Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "記事が見つかりません"})
		return nil
	}

	return &post
}

// YAGNI: 今必要な最小限のサンプルデータ
func createSampleData() {
	var count int64
	db.Model(&Post{}).Count(&count)
	if count > 0 {
		return // 既にデータがあるなら作成しない
	}

	posts := []Post{
		{
			Title: "Go言語とKISS・YAGNI・DRY",
			Content: `## KISS (Keep It Simple, Stupid)

Goは**シンプルさ**を最重要視した言語設計です。

- キーワードは25個のみ
- 明確で読みやすい構文
- 余計な機能は排除

## YAGNI (You Aren't Gonna Need It)

必要になってから実装する：

` + "```go" + `
// 必要最小限の構造体
type Post struct {
    ID      uint   ` + "`json:\"id\"`" + `
    Title   string ` + "`json:\"title\"`" + `
    Content string ` + "`json:\"content\"`" + `
}
` + "```" + `

## DRY (Don't Repeat Yourself)

共通処理は関数化：

` + "```go" + `
func (p *Post) ToMarkdown() string {
    return "# " + p.Title + "\n\n" + p.Content
}
` + "```" + `

**Goはプログラミング哲学が言語に完璧に反映されています！**`,
		},
		{
			Title: "個人開発者にGoが最適な理由",
			Content: `# 個人開発者こそGoを選ぶべき

## ⚡ 開発速度が速い

` + "```go" + `
// たった数行でWebサーバー
r := gin.Default()
r.GET("/", handler)
r.Run(":8080")
` + "```" + `

## 💰 運用コストが安い

- メモリ使用量: 10-50MB
- CPU使用率: 最小
- サーバー代: 月額$5のVPSで十分

## 🚀 デプロイが簡単

` + "```bash" + `
# 1コマンドで完了
go build -o app main.go
./app  # 依存関係なし！
` + "```" + `

## 📈 将来性抜群

- GitHub、Docker、Kubernetesの言語
- 求人数4倍成長
- 平均年収700-900万円

**個人開発者なら迷わずGo！**`,
		},
		{
			Title: "実際のGoコード例",
			Content: `# このアプリ自体がGoの実例

## 150行で完全なWebアプリ

このマークダウンパーサー自体が、Goの威力を証明しています：

` + "```go" + `
func main() {
    // データベース初期化
    db, _ := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})

    // ルーター設定
    r := gin.Default()
    r.GET("/posts/:id.md", showPostMarkdown)

    // サーバー起動
    r.Run(":8080")
}
` + "```" + `

## 特徴

- **KISS**: シンプルで理解しやすい
- **YAGNI**: 必要な機能のみ
- **DRY**: 重複コードなし

## パフォーマンス

- 同時接続: 10,000+
- レスポンス時間: 1ms以下
- メモリ使用量: 30MB以下

**Rails の1/10のリソースで10倍のパフォーマンス！**`,
		},
	}

	for _, post := range posts {
		db.Create(&post)
	}
}