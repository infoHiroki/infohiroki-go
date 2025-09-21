// infoHiroki Website Go版 - ピクセルパーフェクト移植
// 既存のVanilla HTML/CSS/JSサイトをGo + Gin + GORMで完全再現

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
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

	// Markdownファイルの読み込み
	loadMarkdownFiles()
}

// content/metadata.jsonからブログ記事を読み込み
func loadFromFilesJSON() {
	fmt.Println("📚 ブログ記事を読み込み中...")

	jsonData, err := os.ReadFile("content/metadata.json")
	if err != nil {
		fmt.Printf("content/metadata.json読み込みエラー: %v\n", err)
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

// content/articlesディレクトリから記事ファイルを読み込み（HTML/Markdown両対応）
func loadMarkdownFiles() {
	fmt.Println("📝 記事ファイルを読み込み中...")

	postsDir := "content/articles"
	if _, err := os.Stat(postsDir); os.IsNotExist(err) {
		fmt.Println("postsディレクトリが存在しません")
		return
	}

	err := filepath.Walk(postsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		ext := filepath.Ext(path)
		if ext != ".md" && ext != ".html" {
			return nil
		}

		// HTMLファイルはmetadata.jsonで既に処理済みなのでスキップ
		if ext == ".html" {
			return nil
		}

		fmt.Printf("処理中: %s\n", path)
		return loadMarkdownFile(path)
	})

	if err != nil {
		fmt.Printf("Markdownファイル読み込みエラー: %v\n", err)
	}
}

// 個別のMarkdownファイルを読み込み
func loadMarkdownFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// ファイル名からスラッグを生成
	fileName := filepath.Base(filePath)
	ext := filepath.Ext(fileName)
	slug := strings.TrimSuffix(fileName, ext)

	// 既存記事の確認
	var existingPost models.BlogPost
	result := db.Where("slug = ?", slug).First(&existingPost)
	if result.Error == nil {
		// 既に存在する場合はスキップ
		return nil
	}

	// ファイル名から日付を抽出
	var createdDate time.Time
	if len(slug) >= 10 && slug[4] == '-' && slug[7] == '-' {
		dateStr := slug[:10]
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err == nil {
			createdDate = parsedDate
		}
	}

	if createdDate.IsZero() {
		createdDate = time.Now()
	}

	// ファイル名に基づいてメタデータを設定
	title, description, tags, icon := generateMetadataFromSlug(slug)

	blogPost := models.BlogPost{
		Slug:         slug,
		Title:        title,
		Content:      string(content),
		ContentType:  "markdown",
		MarkdownPath: filePath,
		CreatedDate:  createdDate,
		Published:    true,
		Description:  description,
		Tags:         tags,
		Icon:         icon,
	}

	if err := db.Create(&blogPost).Error; err != nil {
		return fmt.Errorf("データベース保存エラー: %v", err)
	}

	fmt.Printf("✅ Markdown記事を追加: %s\n", slug)
	return nil
}

// ファイル名（スラッグ）からメタデータを生成
func generateMetadataFromSlug(slug string) (title, description, tags, icon string) {
	switch {
	case strings.Contains(slug, "go-complete-history"):
		return "Go言語完全史：クラウドネイティブ時代を切り開いた革新言語の18年間",
			"2007年から2025年まで：Google三巨頭が創造した言語が、いかにしてDocker・Kubernetesの基盤となり、現代インフラを支配するに至ったか",
			`["Go","プログラミング言語","歴史","Docker","Kubernetes","Google","クラウドネイティブ","技術史","コンテナ","DevOps"]`,
			"🏛️"
	case strings.Contains(slug, "golang"):
		return "Go言語の歴史と技術革新",
			"Go言語（Golang）の開発歴史と現代への影響を詳しく解説",
			`["Go","Golang","プログラミング言語","歴史","Google"]`,
			"🏛️"
	case strings.Contains(slug, "markdown-test"):
		return "Markdownテスト記事",
			"Markdownシステムのテスト記事です",
			`["Markdown","テスト","ブログシステム"]`,
			"📝"
	default:
		// デフォルトの場合はファイル名から生成
		title = strings.ReplaceAll(slug, "-", " ")
		title = strings.Title(title)
		return title,
			"Markdownで作成された記事",
			`["Markdown","ブログ"]`,
			"📝"
	}
}