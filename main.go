// infoHiroki Website Go版 - ピクセルパーフェクト移植
// 既存のVanilla HTML/CSS/JSサイトをGo + Gin + GORMで完全再現

package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"infohiroki-go/src/models"
)

var db *gorm.DB

func main() {
	// データベース接続
	var err error
	db, err = gorm.Open(sqlite.Open("database/infohiroki.db"), &gorm.Config{})
	if err != nil {
		panic("データベース接続に失敗しました: " + err.Error())
	}

	// テーブル自動マイグレーション
	db.AutoMigrate(&models.Page{}, &models.BlogPost{})

	// マイグレーション実行
	runMigration()

	// Gin ルーター設定
	r := gin.Default()

	// 静的ファイルの配信
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/images", "./static/images")

	// テンプレート読み込み
	r.LoadHTMLGlob("templates/*")

	// Routes - infoHirokiサイト構造
	r.GET("/", homePage)
	r.GET("/blog", blogList)
	r.GET("/blog/:slug", handleBlogPost)
	r.GET("/services", servicesPage)
	r.GET("/products", productsPage)
	r.GET("/results", resultsPage)
	r.GET("/about", aboutPage)
	r.GET("/faq", faqPage)
	r.GET("/contact", contactPage)

	// API endpoints
	r.GET("/api/search", searchBlogPosts)

	// サーバー起動
	r.Run(":8080")
}

// ホームページ
func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "infoHiroki - 福岡の生成AI導入支援専門家",
		"page":  "home",
	})
}

// ブログ一覧
func blogList(c *gin.Context) {
	var posts []models.BlogPost
	query := c.Query("q")
	tag := c.Query("tag")

	dbQuery := db.Where("published = ?", true)

	// 検索機能
	if query != "" {
		dbQuery = dbQuery.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	// タグフィルタ
	if tag != "" {
		dbQuery = dbQuery.Where("tags LIKE ?", "%"+tag+"%")
	}

	dbQuery.Order("created_date DESC").Find(&posts)

	c.HTML(http.StatusOK, "blog.html", gin.H{
		"title": "ブログ | infoHiroki",
		"page":  "blog",
		"posts": posts,
		"query": query,
		"tag":   tag,
	})
}

// ブログ記事詳細（拡張子対応）
func handleBlogPost(c *gin.Context) {
	slug := c.Param("slug")

	// 拡張子をチェック
	if strings.HasSuffix(slug, ".md") {
		// .mdの場合、Markdown形式で返す
		slugWithoutExt := strings.TrimSuffix(slug, ".md")
		showBlogPostMarkdown(c, slugWithoutExt)
	} else if strings.HasSuffix(slug, ".json") {
		// .jsonの場合、JSON形式で返す
		slugWithoutExt := strings.TrimSuffix(slug, ".json")
		showBlogPostJSON(c, slugWithoutExt)
	} else {
		// 拡張子なしの場合、HTML形式で返す
		showBlogPost(c, slug)
	}
}

// ブログ記事詳細（HTML）
func showBlogPost(c *gin.Context, slug string) {
	post := getBlogPostBySlug(c, slug)
	if post == nil {
		return
	}

	// HTMLコンテンツをそのまま表示
	c.HTML(http.StatusOK, "blog_detail.html", gin.H{
		"title": post.Title + " | infoHiroki",
		"page":  "blog",
		"post":  post,
		"html":  template.HTML(post.Content), // HTMLエスケープを回避
	})
}

// ブログ記事詳細（Markdown）
func showBlogPostMarkdown(c *gin.Context, slug string) {
	post := getBlogPostBySlug(c, slug)
	if post == nil {
		return
	}

	c.Data(http.StatusOK, "text/markdown; charset=utf-8", []byte(post.ToMarkdown()))
}

// ブログ記事詳細（JSON）
func showBlogPostJSON(c *gin.Context, slug string) {
	post := getBlogPostBySlug(c, slug)
	if post == nil {
		return
	}
	c.JSON(http.StatusOK, post)
}

// 共通処理：スラッグでブログ記事を取得
func getBlogPostBySlug(c *gin.Context, slug string) *models.BlogPost {
	var post models.BlogPost
	if err := db.Where("slug = ? AND published = ?", slug, true).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "記事が見つかりません"})
		return nil
	}
	return &post
}

// 固定ページ処理（サービス、製品、実績、等）
func servicesPage(c *gin.Context) {
	renderPage(c, "services", "生成AI導入支援サービス | infoHiroki")
}

func productsPage(c *gin.Context) {
	renderPage(c, "products", "開発製品 | infoHiroki")
}

func resultsPage(c *gin.Context) {
	renderPage(c, "results", "実績 | infoHiroki")
}

func aboutPage(c *gin.Context) {
	renderPage(c, "about", "スキルスタック | infoHiroki")
}

func faqPage(c *gin.Context) {
	renderPage(c, "faq", "FAQ | infoHiroki")
}

func contactPage(c *gin.Context) {
	renderPage(c, "contact", "お問い合わせ | infoHiroki")
}

// 固定ページ共通処理
func renderPage(c *gin.Context, slug string, title string) {
	var page models.Page
	if err := db.Where("slug = ?", slug).First(&page).Error; err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "ページが見つかりません | infoHiroki",
		})
		return
	}

	c.HTML(http.StatusOK, slug+".html", gin.H{
		"title": title,
		"page":  slug,
		"data":  page,
	})
}

// ブログ検索API
func searchBlogPosts(c *gin.Context) {
	query := c.Query("q")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	var posts []models.BlogPost
	dbQuery := db.Where("published = ?", true)

	if query != "" {
		dbQuery = dbQuery.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	dbQuery.Order("created_date DESC").Limit(limit).Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
		"total": len(posts),
		"query": query,
	})
}