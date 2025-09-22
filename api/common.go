package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"infohiroki-go/src/models"
)

// データはファイルベースで管理
var allPosts []models.BlogPost
var allPages []models.Page

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

// データ初期化（ファイルベース）
func InitializeData() {
	// メモリを初期化
	allPosts = []models.BlogPost{}
	allPages = []models.Page{}

	// metadata.jsonからブログ記事を読み込み
	loadFromFilesJSON()

	// Markdownファイルの読み込み
	loadMarkdownFiles()

	fmt.Printf("✅ データ初期化完了: %d件の記事を読み込み\n", len(allPosts))
}

// ファイルベースでのフィルタリング関数
func FilterPosts(posts []models.BlogPost, query, tag string) []models.BlogPost {
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

		// タグフィルタ
		if tag != "" {
			if !strings.Contains(strings.ToLower(post.Tags), strings.ToLower(tag)) {
				continue
			}
		}

		result = append(result, post)
	}

	// 作成日の降順でソート
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].CreatedDate.Before(result[j].CreatedDate) {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}

// スラッグでブログ記事を取得
func GetBlogPostBySlug(slug string) *models.BlogPost {
	for _, post := range allPosts {
		if post.Slug == slug && post.Published {
			return &post
		}
	}
	return nil
}

// 全ブログ記事を取得
func GetAllPosts() []models.BlogPost {
	return allPosts
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
			Slug:         file.ID,
			Title:        file.Title,
			Description:  file.Description,
			Tags:         string(tagsJSON),
			Icon:         file.Icon,
			CreatedDate:  createdDate,
			Published:    true,
			MarkdownPath: file.Path, // metadata.jsonのpathを保存
		}

		// 重複チェック
		exists := false
		for _, existing := range allPosts {
			if existing.Slug == file.ID {
				exists = true
				break
			}
		}
		if !exists {
			allPosts = append(allPosts, blogPost)
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

		// HTMLファイルの場合、既存記事にコンテンツを追加
		if ext == ".html" {
			return loadHTMLContent(path)
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

	allPosts = append(allPosts, blogPost)

	fmt.Printf("✅ Markdown記事を追加: %s\n", slug)
	return nil
}

// HTMLファイルのコンテンツを既存記事に読み込み
func loadHTMLContent(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// ファイル名からスラッグを推定
	fileName := filepath.Base(filePath)

	// metadata.jsonのパスと一致する記事を探す
	for i := range allPosts {
		if allPosts[i].ContentType == "" { // まだコンテンツが読み込まれていない
			// metadata.jsonのpathとファイル名を照合
			if fileName == filepath.Base(allPosts[i].MarkdownPath) || strings.Contains(fileName, allPosts[i].Slug) {
				allPosts[i].Content = string(content)
				allPosts[i].ContentType = "html"
				allPosts[i].MarkdownPath = filePath
				fmt.Printf("✅ HTML記事コンテンツを読み込み: %s\n", allPosts[i].Slug)
				return nil
			}
		}
	}

	// 一致する記事が見つからない場合、ファイル名から新規作成
	ext := filepath.Ext(fileName)
	slug := strings.TrimSuffix(fileName, ext)

	// 日付プレフィックスを除去してより良いスラッグを作成
	if len(slug) > 11 && slug[10] == '-' {
		slug = slug[11:] // "2025-08-26-" を除去
	}

	var createdDate time.Time
	if len(fileName) >= 10 && fileName[4] == '-' && fileName[7] == '-' {
		dateStr := fileName[:10]
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err == nil {
			createdDate = parsedDate
		}
	}

	if createdDate.IsZero() {
		createdDate = time.Now()
	}

	blogPost := models.BlogPost{
		Slug:         slug,
		Title:        "HTML記事: " + slug,
		Content:      string(content),
		ContentType:  "html",
		MarkdownPath: filePath,
		CreatedDate:  createdDate,
		Published:    true,
		Description:  "HTMLで作成された記事",
		Tags:         `["HTML","ブログ"]`,
		Icon:         "📄",
	}

	allPosts = append(allPosts, blogPost)
	fmt.Printf("✅ 新規HTML記事を追加: %s\n", slug)
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