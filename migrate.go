package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"infohiroki-go/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

func runMigration() {
	// データベース接続
	db, err := gorm.Open(sqlite.Open("database/infohiroki.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("データベース接続失敗:", err)
	}

	// テーブル作成
	err = db.AutoMigrate(
		&models.Page{},
		&models.BlogPost{},
	)
	if err != nil {
		log.Fatal("マイグレーション失敗:", err)
	}

	fmt.Println("✅ データベーステーブル作成完了")

	// 固定ページデータ投入
	seedPages(db)

	// ブログ記事データ投入
	seedBlogPosts(db)

	fmt.Println("🎉 データベース初期化完了！")
}

func seedPages(db *gorm.DB) {
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
			MetaDescription: "infoHirokiが開発した生成AI活用ツールとソリューションの紹介",
		},
		{
			Slug:            "results",
			Title:           "実績",
			Content:         "実績ページコンテンツ",
			Template:        "results",
			MetaDescription: "生成AI導入支援の実績と成功事例をご紹介",
		},
		{
			Slug:            "about",
			Title:           "スキルスタック",
			Content:         "スキルスタックページコンテンツ",
			Template:        "about",
			MetaDescription: "Hiroki Takamuraのスキルセットと経歴について",
		},
		{
			Slug:            "faq",
			Title:           "FAQ",
			Content:         "FAQページコンテンツ",
			Template:        "faq",
			MetaDescription: "生成AI導入支援に関するよくある質問",
		},
		{
			Slug:            "contact",
			Title:           "お問い合わせ",
			Content:         "お問い合わせページコンテンツ",
			Template:        "contact",
			MetaDescription: "生成AI導入支援のご相談・お問い合わせはこちら",
		},
	}

	for _, page := range pages {
		var existingPage models.Page
		if err := db.Where("slug = ?", page.Slug).First(&existingPage).Error; err != nil {
			// ページが存在しない場合は作成
			if err := db.Create(&page).Error; err != nil {
				log.Printf("ページ作成エラー %s: %v", page.Slug, err)
			} else {
				fmt.Printf("✅ ページ作成: %s\n", page.Title)
			}
		}
	}
}

func seedBlogPosts(db *gorm.DB) {
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

	for _, file := range filesData.Files {
		// HTMLファイルパス
		htmlPath := filepath.Join("markdown", file.Path)

		// HTMLファイルを読み込み
		htmlContent, err := ioutil.ReadFile(htmlPath)
		if err != nil {
			log.Printf("HTMLファイル読み込みエラー %s: %v", file.Path, err)
			continue
		}

		// HTMLからコンテンツを抽出（簡単な処理）
		content := extractContentFromHTML(string(htmlContent))

		// 作成日をパース
		createdDate, err := time.Parse("2006-01-02", file.Created)
		if err != nil {
			createdDate = time.Now()
		}

		// タグをJSON文字列に変換
		tagsJSON, _ := json.Marshal(file.Tags)

		// スラッグ生成（ファイル名から拡張子を除去）
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
			// 記事が存在しない場合は作成
			if err := db.Create(&blogPost).Error; err != nil {
				log.Printf("記事作成エラー %s: %v", blogPost.Slug, err)
			} else {
				fmt.Printf("✅ 記事作成: %s\n", blogPost.Title)
			}
		}
	}
}

// HTMLからコンテンツを簡単に抽出（本格的な処理は後で実装）
func extractContentFromHTML(html string) string {
	// <body>タグ内のコンテンツを取得
	bodyRegex := regexp.MustCompile(`<body[^>]*>(.*?)</body>`)
	matches := bodyRegex.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1]
	}

	// <body>タグがない場合は全体を返す
	return html
}