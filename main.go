// infoHiroki Website Go版 - ピクセルパーフェクト移植
// 既存のVanilla HTML/CSS/JSサイトをGo + Gin + GORMで完全再現

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
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

	// データ初期化
	initializeData()

	// Gin ルーター設定
	r := gin.Default()

	// カスタムテンプレート関数を設定
	r.SetFuncMap(template.FuncMap{
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix,
	})

	// 静的ファイルの配信
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/images", "./static/images")

	// テンプレート読み込み（base.htmlを含むすべてのテンプレート）
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

	// API endpoints
	r.GET("/api/search", searchBlogPosts)

	// サーバー起動
	r.Run(":8080")
}

// ホームページ
func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":           "infoHiroki - 福岡の生成AI導入支援専門家",
		"page":            "home",
		"metaDescription": "福岡・九州の企業向け生成AI導入支援 - ChatGPT・Claude・Whisperで業務効率化を実現",
		"ogTitle":         "infoHiroki - 福岡の生成AI導入支援専門家",
		"ogDescription":   "福岡・九州の企業向け生成AI導入支援 - ChatGPT・Claude・Whisperで業務効率化を実現",
		"ogType":          "website",
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
		"title":           "ブログ | infoHiroki",
		"page":            "blog",
		"posts":           posts,
		"query":           query,
		"tag":             tag,
		"metaDescription": "infoHirokiのブログ - 生成AI・技術・開発に関する記事を配信中",
		"ogTitle":         "ブログ | infoHiroki",
		"ogDescription":   "福岡・九州の生成AI導入支援専門家による技術ブログ",
		"ogType":          "website",
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

	// SEOメタデータの設定
	metaDescription := post.Description
	if metaDescription == "" {
		metaDescription = "infoHiroki - 福岡・九州の生成AI導入支援専門家のブログ記事"
	}

	// HTMLコンテンツをそのまま表示
	c.HTML(http.StatusOK, "blog_detail.html", gin.H{
		"title":           post.Title + " | infoHiroki",
		"page":            "blog",
		"post":            post,
		"html":            template.HTML(post.Content), // HTMLエスケープを回避
		"metaDescription": metaDescription,
		"ogTitle":         post.Title + " | infoHiroki",
		"ogDescription":   metaDescription,
		"ogType":          "article",
		"twitterCard":     "summary",
		"twitterTitle":    post.Title,
		"twitterDescription": metaDescription,
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
	renderPageWithMeta(c, "services", "生成AI導入支援サービス | infoHiroki", "福岡・九州企業向け生成AI導入支援サービス - ChatGPT・Claude・Whisper活用で業務効率化")
}

func productsPage(c *gin.Context) {
	renderPageWithMeta(c, "products", "開発製品 | infoHiroki", "infoHirokiが開発した製品・ツール・アプリケーション一覧")
}

func resultsPage(c *gin.Context) {
	renderPageWithMeta(c, "results", "実績 | infoHiroki", "infoHirokiの開発実績・導入事例・お客様の声")
}

func aboutPage(c *gin.Context) {
	renderPageWithMeta(c, "about", "スキルスタック | infoHiroki", "infoHirokiの技術スタック・経歴・スキル")
}

func faqPage(c *gin.Context) {
	renderPageWithMeta(c, "faq", "FAQ | infoHiroki", "よくある質問と回答 - infoHirokiサービスについて")
}

func contactPage(c *gin.Context) {
	renderPageWithMeta(c, "contact", "お問い合わせ | infoHiroki", "infoHirokiへのお問い合わせ・ご相談はこちら")
}

// 固定ページ共通処理（メタデータ付き）
func renderPageWithMeta(c *gin.Context, slug string, title string, description string) {
	var page models.Page
	if err := db.Where("slug = ?", slug).First(&page).Error; err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "ページが見つかりません | infoHiroki",
		})
		return
	}

	c.HTML(http.StatusOK, slug+".html", gin.H{
		"title":           title,
		"page":            slug,
		"data":            page,
		"metaDescription": description,
		"ogTitle":         title,
		"ogDescription":   description,
		"ogType":          "website",
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

// データ初期化
func initializeData() {
	// ブログ記事が空の場合のみデータを読み込み
	var count int64
	db.Model(&models.BlogPost{}).Count(&count)
	if count == 0 {
		loadFromFilesJSON()
	}
}

// files.jsonからブログ記事を読み込み
func loadFromFilesJSON() {
	fmt.Println("📚 ブログ記事を読み込み中...")

	jsonData, err := os.ReadFile("files.json")
	if err != nil {
		fmt.Printf("files.json読み込みエラー: %v\n", err)
		return
	}

	var filesJSON FilesJSON
	if err := json.Unmarshal(jsonData, &filesJSON); err != nil {
		fmt.Printf("JSON解析エラー: %v\n", err)
		return
	}

	for _, file := range filesJSON.Files {
		createdDate, _ := time.Parse("2006-01-02", file.Created)
		tagsJSON, _ := json.Marshal(file.Tags)

		blogPost := models.BlogPost{
			Slug:        file.ID,
			Title:       file.Title,
			Description: file.Description,
			Tags:        string(tagsJSON),
			Icon:        file.Icon,
			CreatedDate: createdDate,
			Published:   true,
		}

		var existingPost models.BlogPost
		result := db.Where("slug = ?", file.ID).First(&existingPost)
		if result.Error != nil {
			db.Create(&blogPost)
		}
	}

	fmt.Printf("✅ %d件のブログ記事を処理完了\n", len(filesJSON.Files))
}