// Go + Gin: KISS・YAGNI・DRY の完璧な体現
// ファイル1つで完全なマークダウンパーサーWebアプリ

package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Post モデル - 必要最小限
type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// ToMarkdown - DRY: 1箇所で定義、どこでも使用
func (p *Post) ToMarkdown() string {
	return "# " + p.Title + "\n\n" + p.Content
}

var db *gorm.DB

func main() {
	// KISS: 設定は最小限
	var err error
	db, err = gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&Post{})

	// サンプルデータ - YAGNI: 今必要な分だけ
	createSampleData()

	// KISS: ルーターも最小設定
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// RESTful エンドポイント - DRY: パターンを統一
	r.GET("/", listPosts)
	r.GET("/posts/:id", showPost)
	r.GET("/posts/:id.md", showPostMarkdown) // Rails 8.1と同じ機能
	r.GET("/posts/:id.json", showPostJSON)

	// KISS: ポート固定、設定なし
	r.Run(":8080")
}

// KISS: 関数は1つのことだけ
func listPosts(c *gin.Context) {
	var posts []Post
	db.Find(&posts)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"posts": posts,
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
		"html":  string(html),
	})
}

// YAGNI: Rails 8.1の.md機能だけ実装、他は後回し
func showPostMarkdown(c *gin.Context) {
	post := getPostByID(c)
	if post == nil {
		return
	}

	// Content-Typeヘッダーを設定してマークダウンを返す
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return nil
	}

	var post Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return nil
	}

	return &post
}

// YAGNI: 今必要な最小限のサンプルデータだけ
func createSampleData() {
	var count int64
	db.Model(&Post{}).Count(&count)
	if count > 0 {
		return // 既にデータがあるなら何もしない
	}

	posts := []Post{
		{
			Title: "Go言語の哲学",
			Content: `## KISS (Keep It Simple, Stupid)

Goは**シンプルさ**を最重要視します。

- 25個のキーワードのみ
- 明確な構文
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
` + "```",
		},
		{
			Title: "Ginフレームワークの美学",
			Content: `# 最小限で最大効果

## 1行でWebサーバー

` + "```go" + `
r := gin.Default()
r.Run(":8080")
` + "```" + `

## ルーティングもシンプル

` + "```go" + `
r.GET("/posts/:id", showPost)
r.GET("/posts/:id.md", showPostMarkdown)
r.GET("/posts/:id.json", showPostJSON)
` + "```" + `

## エラーハンドリングも明確

` + "```go" + `
if err != nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
}
` + "```

**Goは哲学が言語に体現されています！**`,
		},
	}

	for _, post := range posts {
		db.Create(&post)
	}
}

/*
必要なファイル構成（最小限）:

go.mod:
module markdown-parser
go 1.21
require (
    github.com/gin-gonic/gin v1.9.1
    github.com/russross/blackfriday/v2 v2.1.0
    gorm.io/driver/sqlite v1.5.4
    gorm.io/gorm v1.25.5
)

templates/index.html:
<!DOCTYPE html>
<html>
<head><title>Go Markdown Parser</title></head>
<body>
<h1>Posts</h1>
{{range .posts}}
<div>
    <h2><a href="/posts/{{.ID}}">{{.Title}}</a></h2>
    <p>Created: {{.CreatedAt.Format "2006-01-02"}}</p>
    <a href="/posts/{{.ID}}.md">Markdown</a> |
    <a href="/posts/{{.ID}}.json">JSON</a>
</div>
{{end}}
</body>
</html>

templates/show.html:
<!DOCTYPE html>
<html>
<head><title>{{.post.Title}}</title></head>
<body>
<h1>{{.post.Title}}</h1>
<div>{{.html}}</div>
<p><a href="/">← Back</a></p>
</body>
</html>

起動方法:
go mod init markdown-parser
go mod tidy
go run main.go
*/