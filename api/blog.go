package handler

import (
	"html/template"
	"net/http"
	"log"
	"infohiroki-go/src/models"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	// データが初期化されていない場合は初期化
	if len(allPosts) == 0 {
		InitializeData()
	}

	query := r.URL.Query().Get("q")
	tag := r.URL.Query().Get("tag")

	// ファイルベースでのフィルタリング
	posts := FilterPosts(allPosts, query, tag)

	// HTMLテンプレートの作成（templates/blog.htmlの完全移植）
	tmpl := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="infohirokiのブログ - 技術記事とアーカイブ">
    <title>ブログ | infoHiroki</title>

    <!-- OGPタグ -->
    <meta property="og:title" content="ブログ | infohiroki">
    <meta property="og:description" content="infohirokiのブログ - 技術記事とアーカイブ">
    <meta property="og:type" content="website">
    <meta property="og:site_name" content="infohiroki">
    <meta property="og:locale" content="ja_JP">

    <!-- Twitterカード -->
    <meta name="twitter:card" content="summary">
    <meta name="twitter:title" content="ブログ | infohiroki">
    <meta name="twitter:description" content="infohirokiのブログ - 技術記事とアーカイブ">

    <!-- ファビコン -->
    <link rel="icon" type="image/svg+xml" href="/images/logo.svg">

    <!-- Google Fonts -->
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Montserrat:wght@700;800;900&display=swap" rel="stylesheet">

    <link rel="stylesheet" href="/css/style.css">

    <style>
        /* ブログ固有のスタイル - infohirokiデザインシステムに統一 */
        .search-section {
            background-color: var(--color-background);
            padding: var(--spacing-xl) 0;
        }

        .search-container {
            max-width: 800px;
            margin: 0 auto;
            display: flex;
            flex-direction: column;
            gap: var(--spacing-md);
            align-items: center;
        }

        .search-input-wrapper {
            width: 100%;
            max-width: 400px;
            display: flex;
            justify-content: center;
        }

        .search-input {
            width: 100%;
            padding: var(--spacing-md);
            border: 2px solid var(--color-border);
            border-radius: 8px;
            font-family: var(--font-family);
            font-size: 1rem;
            background-color: var(--color-background);
            color: var(--color-text);
            transition: var(--transition);
        }

        .button-row {
            display: flex;
            gap: var(--spacing-lg);
            align-items: center;
            justify-content: center;
            flex-wrap: wrap;
        }

        .ofuse-btn {
            display: inline-block;
            padding: var(--spacing-sm) var(--spacing-md);
            background-color: #8b5cf6;
            color: white;
            text-decoration: none;
            border-radius: 8px;
            font-size: 0.875rem;
            font-weight: 600;
            box-shadow: 0 2px 8px rgba(139, 92, 246, 0.3);
            transition: all 0.2s ease;
            white-space: nowrap;
            min-height: 44px;
            display: flex;
            align-items: center;
            position: relative;
            overflow: hidden;
        }

        .ofuse-btn::before {
            position: absolute;
            content: '';
            display: inline-block;
            top: -180px;
            left: 0;
            width: 30px;
            height: 100%;
            background-color: rgba(255, 255, 255, 0.8);
            animation: ofuse_shine 3s ease-in-out infinite;
        }

        @keyframes ofuse_shine {
            0% {
                transform: scale(0) rotate(45deg);
                opacity: 0;
            }
            80% {
                transform: scale(0) rotate(45deg);
                opacity: 0.5;
            }
            81% {
                transform: scale(4) rotate(45deg);
                opacity: 1;
            }
            100% {
                transform: scale(50) rotate(45deg);
                opacity: 0;
            }
        }

        .ofuse-btn:hover {
            transform: translateY(-2px);
            background-color: #7c3aed;
            box-shadow: 0 6px 25px rgba(139, 92, 246, 0.7),
                       0 0 30px rgba(139, 92, 246, 0.5);
            text-decoration: none;
            color: white;
        }

        .search-input:focus {
            outline: none;
            border-color: var(--color-accent);
            box-shadow: 0 0 0 3px rgba(231, 62, 143, 0.1);
        }



        .sort-controls {
            display: flex;
            gap: var(--spacing-sm);
        }

        .sort-button {
            padding: var(--spacing-xs) var(--spacing-sm);
            background-color: transparent;
            color: var(--color-text-light);
            border: 1px solid var(--color-border);
            border-radius: 4px;
            font-size: 0.875rem;
            cursor: pointer;
            transition: var(--transition);
        }

        .sort-button:hover,
        .sort-button.active {
            background-color: var(--color-accent);
            color: var(--color-background);
            border-color: var(--color-accent);
        }


        /* infohirokiの.blog-gridを使用 */
        .article-grid {
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: var(--spacing-lg);
            margin-bottom: var(--spacing-xl);
        }

        /* infohirokiの.cardスタイルベース */
        .article-card {
            background-color: var(--color-background);
            border: 2px solid var(--color-border);
            border-radius: 8px;
            padding: var(--spacing-lg);
            transition: var(--transition);
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
            cursor: pointer;
        }

        .article-card:hover {
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
            transform: translateY(-2px);
            border-color: var(--color-accent);
        }

        .article-header {
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            margin-bottom: var(--spacing-sm);
        }


        .article-date {
            font-size: 0.875rem;
            color: var(--color-text-light);
        }

        .article-title {
            font-size: 1.25rem;
            font-weight: 600;
            margin-bottom: var(--spacing-sm);
            color: var(--color-text);
            line-height: 1.4;
        }

        .article-title a {
            text-decoration: none;
            color: inherit;
            transition: var(--transition);
        }

        .article-title a:hover {
            color: var(--color-accent);
        }

        .article-description {
            color: var(--color-text-light);
            margin-bottom: var(--spacing-md);
            line-height: 1.6;
            font-size: 0.9rem;
        }

        .article-tags {
            display: flex;
            flex-wrap: wrap;
            gap: var(--spacing-xs);
        }

        .article-tags .tag {
            padding: var(--spacing-xs) var(--spacing-sm);
            background-color: #E3F2FD;
            color: #1565C0;
            border-radius: 12px;
            font-size: 0.75rem;
            font-weight: 500;
            border: 1px solid #BBDEFB;
            transition: var(--transition);
        }

        .article-tags .tag:hover {
            background-color: var(--color-accent);
            color: var(--color-background);
            border-color: var(--color-accent);
        }

        /* ブログカード内アイコンスタイル - 左上配置 */
        .article-header {
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            margin-bottom: var(--spacing-sm);
        }

        .article-icon {
            width: 40px;
            height: 40px;
            flex-shrink: 0;
            border-radius: 6px;
            font-size: 32px;
            line-height: 1;
            display: inline-flex;
            align-items: center;
            justify-content: center;
        }

        .article-icon img {
            width: 100%;
            height: 100%;
            object-fit: contain;
        }

        .no-results {
            text-align: center;
            padding: var(--spacing-xxl);
            color: var(--color-text-light);
        }

        /* レスポンシブ対応 */
        @media (max-width: 1024px) {
            .article-grid {
                grid-template-columns: repeat(2, 1fr);
            }
        }

        @media (max-width: 768px) {
            .search-container {
                gap: var(--spacing-sm);
            }

            .ofuse-btn {
                font-size: 0.8rem;
                padding: var(--spacing-xs) var(--spacing-sm);
            }

            .article-grid {
                grid-template-columns: 1fr;
            }

            .filters {
                justify-content: flex-start;
                margin-bottom: var(--spacing-sm);
            }

            .sort-controls {
                justify-content: center;
            }

            .filter-button {
                font-size: 0.875rem;
                padding: var(--spacing-xs) var(--spacing-sm);
            }

            .search-input {
                font-size: 16px; /* iOS対応 */
            }

            /* モバイル環境でのアイコン調整 */
            .article-icon {
                width: 36px;
                height: 36px;
                font-size: 28px;
            }
        }
    </style>

    <!-- 構造化データ -->
    <script type="application/ld+json">
    {
      "@context": "https://schema.org",
      "@type": "Blog",
      "name": "infohiroki Blog",
      "description": "infohirokiのブログ - 技術記事とアーカイブ",
      "url": "/blog",
      "publisher": {
        "@type": "Organization",
        "name": "infohiroki",
        "logo": {
          "@type": "ImageObject",
          "url": "/images/logo.svg"
        }
      }
    }
    </script>
</head>
<body>
    <div class="site-layout">
        <!-- モバイル用ヘッダー -->
        <header class="mobile-header">
            <div class="mobile-header-content">
                <a href="/" class="mobile-logo">
                    <img src="/images/logo.svg" alt="infoHiroki Logo" width="36" height="36">
                    <span class="mobile-title">infoHiroki</span>
                </a>
                <button class="hamburger-button" aria-label="メニューを開く">
                    <span class="hamburger-line"></span>
                    <span class="hamburger-line"></span>
                    <span class="hamburger-line"></span>
                </button>
            </div>
        </header>

        <!-- デスクトップ用サイドバー / モバイル用オーバーレイメニュー -->
        <aside class="sidebar">
            <div class="sidebar-header">
                <a href="/" class="site-title">
                    <div class="logo">
                        <img src="/images/logo.svg" alt="infoHiroki Logo" width="36" height="36">
                    </div>
                    <div class="title-text">
                        <span class="company-name">infoHiroki</span>
                    </div>
                </a>
            </div>

            <nav class="sidebar-nav">
                <ul class="nav-menu">
                    <li class="nav-item">
                        <a href="/" class="nav-link">ホーム</a>
                    </li>
                    <li class="nav-item active">
                        <a href="/blog" class="nav-link">ブログ</a>
                    </li>
                    <li class="nav-item">
                        <a href="/services" class="nav-link">サービス</a>
                    </li>
                    <li class="nav-item">
                        <a href="/products" class="nav-link">開発製品</a>
                    </li>
                    <li class="nav-item">
                        <a href="/results" class="nav-link">実績</a>
                    </li>
                    <li class="nav-item">
                        <a href="/about" class="nav-link">スキルスタック</a>
                    </li>
                    <li class="nav-item">
                        <a href="/faq" class="nav-link">FAQ</a>
                    </li>
                    <li class="nav-item">
                        <a href="/contact" class="nav-link">お問い合わせ</a>
                    </li>
                </ul>
            </nav>
        </aside>

        <!-- モバイル用オーバーレイ -->
        <div class="mobile-overlay"></div>

        <div class="main-wrapper">
            <main class="site-main">
                <div class="hero-sub">
                    <div class="container">
                        <h1 class="page-title">ブログ</h1>
                    </div>
                </div>

                <!-- 検索セクション -->
                <section class="search-section">
                    <div class="container">
                        <div class="search-container">
                            <div class="search-input-wrapper">
                                <input type="text" id="searchInput" class="search-input" placeholder="記事を検索..." aria-label="記事検索">
                            </div>

                            <div class="button-row">
                                <div class="sort-controls">
                                    <button class="sort-button active" data-sort="date">📅 新着順</button>
                                    <button class="sort-button" data-sort="title">🔤 名前順</button>
                                </div>
                                <a href="https://ofuse.me/ee8863f7" target="_blank" rel="noopener noreferrer" class="ofuse-btn">
                                    ☕ infoHrokiを応援する
                                </a>
                            </div>
                        </div>
                    </div>
                </section>

                <div class="page-content">
                    <div class="container">
                        <div id="articlesContainer" class="article-grid">
                                {{range .Posts}}
                                <article class="article-card" onclick="location.href='/blog/{{.Slug}}'">
                                    <span class="article-icon">
                                        {{if .Icon}}
                                            {{if .IsIconURL}}
                                                <img src="{{.Icon}}" alt="icon">
                                            {{else}}
                                                {{.Icon}}
                                            {{end}}
                                        {{else}}
                                            📝
                                        {{end}}
                                    </span>
                                    <h3 class="article-title">
                                        <a href="/blog/{{.Slug}}">{{.Title}}</a>
                                    </h3>
                                    <p class="article-description">{{.Description}}</p>
                                    <div class="article-tags">
                                        {{$tags := .GetTagsSlice}}
                                        {{range $tags}}
                                            <span class="tag">{{.}}</span>
                                        {{end}}
                                    </div>
                                </article>
                                {{end}}
                            </div>

                            <div id="noResults" class="no-results" style="display: none;">
                                <p>検索条件に一致する記事が見つかりませんでした。</p>
                                <p><small>別のキーワードで検索してみてください</small></p>
                            </div>
                    </div>
                </div>
            </main>

            <footer class="minimal-footer">
                <div class="container">
                    <p>© 2022 infoHiroki. All rights reserved.</p>
                </div>
            </footer>
        </div>
    </div>

    <script src="/js/main.js"></script>
</body>
</html>`

	t, err := template.New("blog").Parse(tmpl)
	if err != nil {
		log.Printf("Template parse error: %v", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Posts []models.BlogPost
		Query string
		Tag   string
	}{
		Posts: posts,
		Query: query,
		Tag:   tag,
	}

	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, data); err != nil {
		log.Printf("Template execution error: %v", err)
		log.Printf("Number of posts: %d", len(posts))
		if len(posts) > 0 {
			log.Printf("First post: %+v", posts[0])
		}
		http.Error(w, "Execution error", http.StatusInternalServerError)
		return
	}
}