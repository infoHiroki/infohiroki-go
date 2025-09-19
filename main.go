// infoHiroki Website Go版 - ピクセルパーフェクト移植
// 既存のVanilla HTML/CSS/JSサイトをGo + Gin + GORMで完全再現

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"infohiroki-go/src/models"
)

var db *gorm.DB

type FileMetadata struct {
	ID          string   `json:"id"`
	Path        string   `json:"path"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Icon        string   `json:"icon"`
	Tags        []string `json:"tags"`
	Created     string   `json:"created"`
}

type FilesJSON struct {
	Files []FileMetadata `json:"files"`
}

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
	r.LoadHTMLGlob("templates/*.html")

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

	// 管理画面
	r.GET("/admin", adminDashboard)
	r.GET("/admin/posts", adminPostList)
	r.GET("/admin/posts/new", adminNewPost)
	r.POST("/admin/posts", adminCreatePost)

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

// データベース初期化・マイグレーション機能
func runMigration() {
	fmt.Println("🔄 データベース初期化中...")

	// 固定ページデータ投入
	seedPages()

	// ブログ記事データ投入
	seedBlogPosts()

	fmt.Println("✅ データベース初期化完了！")
}

func seedPages() {
	pages := []models.Page{
		{
			Slug:            "home",
			Title:           "infoHiroki - 福岡の生成AI導入支援専門家",
			Content:         "ホームページコンテンツ",
			Template:        "home",
			MetaDescription: "福岡・九州の企業向け生成AI導入支援 - ChatGPT・Claude・Whisperで業務効率化を実現",
		},
		{
			Slug:            "services",
			Title:           "生成AI導入支援サービス",
			Content:         "サービスページコンテンツ",
			Template:        "services",
			MetaDescription: "福岡・九州企業向け生成AI導入支援サービス - ChatGPT・Claude・Whisper活用で業務効率化",
		},
		{
			Slug:            "products",
			Title:           "開発製品",
			Content:         "開発製品ページコンテンツ",
			Template:        "products",
			MetaDescription: "infoHirokiが開発した製品・ツール・アプリケーション一覧",
		},
		{
			Slug:            "results",
			Title:           "実績",
			Content:         "実績ページコンテンツ",
			Template:        "results",
			MetaDescription: "infoHirokiの開発実績・導入事例・お客様の声",
		},
		{
			Slug:            "about",
			Title:           "スキルスタック",
			Content:         "スキルスタックページコンテンツ",
			Template:        "about",
			MetaDescription: "infoHirokiの技術スタック・経歴・スキル",
		},
		{
			Slug:            "faq",
			Title:           "FAQ",
			Content:         "FAQページコンテンツ",
			Template:        "faq",
			MetaDescription: "よくある質問と回答 - infoHirokiサービスについて",
		},
		{
			Slug:            "contact",
			Title:           "お問い合わせ",
			Content:         "お問い合わせページコンテンツ",
			Template:        "contact",
			MetaDescription: "infoHirokiへのお問い合わせ・ご相談はこちら",
		},
	}

	for _, page := range pages {
		var existingPage models.Page
		if err := db.Where("slug = ?", page.Slug).First(&existingPage).Error; err != nil {
			if err := db.Create(&page).Error; err != nil {
				log.Printf("ページ作成エラー %s: %v", page.Slug, err)
			} else {
				fmt.Printf("✅ ページ作成: %s\n", page.Title)
			}
		}
	}
}

func seedBlogPosts() {
	// files.jsonを読み込み
	jsonData, err := ioutil.ReadFile("files.json")
	if err != nil {
		log.Printf("files.json読み込みエラー: %v", err)
		return
	}

	var filesData FilesJSON
	if err := json.Unmarshal(jsonData, &filesData); err != nil {
		log.Printf("JSON解析エラー: %v", err)
		return
	}

	fmt.Printf("📚 %d件のブログ記事を処理中...\n", len(filesData.Files))

	count := 0
	for _, file := range filesData.Files {
		// 全記事を処理（制限を削除）

		// HTMLファイルパス
		htmlPath := filepath.Join("markdown", file.Path)

		// HTMLファイルを読み込み
		htmlContent, err := ioutil.ReadFile(htmlPath)
		if err != nil {
			log.Printf("HTMLファイル読み込みエラー %s: %v", file.Path, err)
			continue
		}

		// HTMLからコンテンツを簡単に抽出
		content := extractContentFromHTML(string(htmlContent))

		// 作成日をパース
		createdDate, err := time.Parse("2006-01-02", file.Created)
		if err != nil {
			createdDate = time.Now()
		}

		// タグをJSON文字列に変換
		tagsJSON, _ := json.Marshal(file.Tags)

		// スラッグ生成
		slug := strings.TrimSuffix(file.Path, ".html")

		blogPost := models.BlogPost{
			Slug:        slug,
			Title:       file.Title,
			Content:     content,
			Description: file.Description,
			Tags:        string(tagsJSON),
			Icon:        file.Icon,
			CreatedDate: createdDate,
			Published:   true,
		}

		// 既存記事チェック
		var existingPost models.BlogPost
		if err := db.Where("slug = ?", blogPost.Slug).First(&existingPost).Error; err != nil {
			if err := db.Create(&blogPost).Error; err != nil {
				log.Printf("記事作成エラー %s: %v", blogPost.Slug, err)
			} else {
				fmt.Printf("✅ 記事作成: %s\n", blogPost.Title)
				count++
			}
		}
	}
}

// HTMLからコンテンツを簡単に抽出
func extractContentFromHTML(html string) string {
	// <body>タグ内のコンテンツを取得
	bodyRegex := regexp.MustCompile(`<body[^>]*>(.*?)</body>`)
	matches := bodyRegex.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1]
	}
	return html
}

// 管理画面ダッシュボード
func adminDashboard(c *gin.Context) {
	var postCount int64
	db.Model(&models.BlogPost{}).Count(&postCount)

	c.HTML(http.StatusOK, "admin_dashboard.html", gin.H{
		"title": "管理画面 | infoHiroki",
		"postCount": postCount,
	})
}

// 管理画面 - 投稿一覧
func adminPostList(c *gin.Context) {
	var posts []models.BlogPost
	db.Order("created_date DESC").Find(&posts)

	c.HTML(http.StatusOK, "admin_posts.html", gin.H{
		"title": "投稿管理 | infoHiroki",
		"posts": posts,
	})
}

// 管理画面 - 新規投稿画面
func adminNewPost(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_new_post.html", gin.H{
		"title": "新規投稿 | infoHiroki",
	})
}

// 管理画面 - 投稿作成
func adminCreatePost(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	description := c.PostForm("description")
	tags := c.PostForm("tags")

	// スラッグ生成（日付 + タイトル）
	now := time.Now()
	slug := fmt.Sprintf("%s-%s", now.Format("2006-01-02"), strings.ToLower(strings.ReplaceAll(title, " ", "-")))

	post := models.BlogPost{
		Slug:        slug,
		Title:       title,
		Content:     content,
		Description: description,
		Tags:        tags,
		CreatedDate: now,
		Published:   true,
	}

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/admin/posts")
}