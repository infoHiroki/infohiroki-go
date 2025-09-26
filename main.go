// infoHiroki Website Go版 - ピクセルパーフェクト移植
// 既存のVanilla HTML/CSS/JSサイトをGo + Gin + GORMで完全再現

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"infohiroki-go/src/models"
)

// データはファイルベースで管理
var allPosts []models.BlogPost
var allPages []models.Page


func main() {
	// データ初期化（ファイルベース）
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
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
	query := c.Query("q")

	// ファイルベースでのフィルタリング
	posts := filterPosts(allPosts, query)

	c.HTML(http.StatusOK, "blog.html", gin.H{
		"title":           "ブログ | infoHiroki",
		"page":            "blog",
		"posts":           posts,
		"query":           query,
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
	for _, post := range allPosts {
		if post.Slug == slug && post.Published {
			return &post
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "記事が見つかりません"})
	return nil
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
	// 固定ページはテンプレートのみで処理
	c.HTML(http.StatusOK, slug+".html", gin.H{
		"title":           title,
		"page":            slug,
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

	// ファイルベースでの検索
	posts := filterPosts(allPosts, query)
	if len(posts) > limit {
		posts = posts[:limit]
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
		"total": len(posts),
		"query": query,
	})
}

// ファイルベースでのフィルタリング関数
func filterPosts(posts []models.BlogPost, query string) []models.BlogPost {
	var result []models.BlogPost

	for _, post := range posts {
		if !post.Published {
			continue
		}

		// 検索クエリフィルタ
		if query != "" {
			if !strings.Contains(strings.ToLower(post.Title), strings.ToLower(query)) &&
			   !strings.Contains(strings.ToLower(post.Description), strings.ToLower(query)) {
				continue
			}
		}


		result = append(result, post)
	}

	// 新しい記事が上に来るように作成日の降順でソート
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].CreatedDate.Before(result[j].CreatedDate) {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}

// データ初期化（ファイルベース）
func initializeData() {
	// メモリを初期化
	allPosts = []models.BlogPost{}
	allPages = []models.Page{}

	// Markdownファイルの読み込み
	loadMarkdownFiles()

	fmt.Printf("✅ データ初期化完了: %d件の記事を読み込み\n", len(allPosts))
}

// content/metadata.jsonからブログ記事を読み込み

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
	for _, existing := range allPosts {
		if existing.Slug == slug {
			// 既に存在する場合はスキップ
			return nil
		}
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

	// Markdownファイルからメタデータを動的に抽出
	title := extractTitleFromMarkdown(string(content))
	description := extractDescriptionFromMarkdown(string(content))
	icon := extractIconFromTitle(title)

	blogPost := models.BlogPost{
		Slug:         slug,
		Title:        title,
		Content:      string(content),
		ContentType:  "markdown",
		MarkdownPath: filePath,
		CreatedDate:  createdDate,
		Published:    true,
		Description:  description,
		Icon:         icon,
	}

	allPosts = append(allPosts, blogPost)

	fmt.Printf("✅ Markdown記事を追加: %s\n", slug)
	return nil
}

// Markdownファイルからタイトルを抽出
func extractTitleFromMarkdown(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			title := strings.TrimSpace(strings.TrimPrefix(line, "# "))
			if title != "" {
				return title
			}
		}
	}
	return "タイトル未設定"
}

// Markdownファイルから説明文を抽出
func extractDescriptionFromMarkdown(content string) string {
	lines := strings.Split(content, "\n")

	// 🎯 中心的な主張セクションを探す
	inCentralClaim := false
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// 中心的な主張セクションの開始
		if strings.Contains(line, "🎯 中心的な主張") {
			inCentralClaim = true
			continue
		}

		// 次のセクションに到達したら終了
		if inCentralClaim && strings.HasPrefix(line, "##") {
			break
		}

		// 中心的な主張セクション内の最初の段落を使用
		if inCentralClaim && line != "" && !strings.HasPrefix(line, "#") {
			// Markdownの強調記号を削除してプレーンテキストに
			cleanText := strings.ReplaceAll(line, "**", "")
			cleanText = strings.ReplaceAll(cleanText, "*", "")
			cleanText = strings.ReplaceAll(cleanText, "`", "")

			// 最初の文章のみを取得（。で区切る）
			sentences := strings.Split(cleanText, "。")
			if len(sentences) > 0 && sentences[0] != "" {
				firstSentence := sentences[0]
				if len(firstSentence) > 150 {
					return firstSentence[:150] + "..."
				}
				return firstSentence + "。"
			}

			// 句点がない場合は最初の150文字
			if len(cleanText) > 150 {
				return cleanText[:150] + "..."
			}
			return cleanText
		}
	}

	// 中心的な主張が見つからない場合は従来の方法
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// 空行や見出し、画像、テーブル記号はスキップ
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "![") ||
		   strings.HasPrefix(line, "---") || strings.HasPrefix(line, "|") ||
		   strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
			continue
		}

		// 最初の有効な段落を説明文として使用
		if len(line) > 20 { // 短すぎる行は除外
			// **で囲まれた部分を削除
			cleanText := strings.ReplaceAll(line, "**", "")
			if len(cleanText) > 150 {
				return cleanText[:150] + "..."
			}
			return cleanText
		}
	}
	return "Markdownで作成された記事"
}

// タイトルからアイコンを抽出
func extractIconFromTitle(title string) string {
	if title == "" {
		return "📝"
	}

	titleLower := strings.ToLower(title)

	// 特定のキーワードベースでアイコンを決定
	if strings.Contains(titleLower, "chatgpt") || strings.Contains(titleLower, "ai") || strings.Contains(titleLower, "リスキリング") {
		return "🤖"
	}
	if strings.Contains(titleLower, "notion") {
		return "📝"
	}
	if strings.Contains(titleLower, "go") || strings.Contains(titleLower, "golang") {
		return "🐹"
	}

	// タイトルの最初の文字が絵文字の場合はそれを使用
	runes := []rune(title)
	if len(runes) > 0 {
		firstChar := runes[0]
		// 絵文字の範囲をチェック（簡易版）
		if firstChar >= 0x1F300 && firstChar <= 0x1F9FF {
			return string(firstChar)
		}
		// 基本的な絵文字もチェック
		switch firstChar {
		case '🐹', '📖', '🔖', '📝', '🚀', '💡', '🎯', '⚡', '🌟', '🤖':
			return string(firstChar)
		}
	}
	return "📝" // デフォルト
}

