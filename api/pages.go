package handler

import (
	"html/template"
	"net/http"
)

// 汎用ページハンドラー（完全にtemplates/services.htmlベースに書き換え）
func Page(w http.ResponseWriter, r *http.Request, pageName, title, description string) {
	// 各ページごとの完全なHTMLテンプレート（templates/services.htmlの完全移植）
	tmpl := getPageTemplate(pageName, title, description)

	t, err := template.New("page").Parse(tmpl)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, "Execution error", http.StatusInternalServerError)
		return
	}
}

// 各ページごとの完全なHTMLテンプレートを返す
func getPageTemplate(pageName, title, description string) string {
	baseTemplate := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="` + description + `">
    <title>` + title + `</title>

    <!-- OGPタグ -->
    <meta property="og:title" content="` + title + `">
    <meta property="og:description" content="` + description + `">
    <meta property="og:type" content="article">
    <meta property="og:site_name" content="infohiroki">
    <meta property="og:locale" content="ja_JP">

    <!-- Twitterカード -->
    <meta name="twitter:card" content="summary">
    <meta name="twitter:title" content="` + title + `">
    <meta name="twitter:description" content="` + description + `">

    <!-- ファビコン -->
    <link rel="icon" type="image/svg+xml" href="/images/logo.svg">

    <!-- Google Fonts -->
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Montserrat:wght@700;800;900&display=swap" rel="stylesheet">

    <link rel="stylesheet" href="/css/style.css">

    <!-- 構造化データ -->
    <script type="application/ld+json">
    {
      "@context": "https://schema.org",
      "@type": "Article",
      "headline": "` + pageName + `",
      "description": "` + description + `",
      "author": {
        "@type": "Organization",
        "name": "infohiroki"
      },
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
                    <li class="nav-item">
                        <a href="/blog" class="nav-link">ブログ</a>
                    </li>
                    <li class="nav-item` + getActiveClass(pageName, "services") + `">
                        <a href="/services" class="nav-link">サービス</a>
                    </li>
                    <li class="nav-item` + getActiveClass(pageName, "products") + `">
                        <a href="/products" class="nav-link">開発製品</a>
                    </li>
                    <li class="nav-item` + getActiveClass(pageName, "results") + `">
                        <a href="/results" class="nav-link">実績</a>
                    </li>
                    <li class="nav-item` + getActiveClass(pageName, "about") + `">
                        <a href="/about" class="nav-link">スキルスタック</a>
                    </li>
                    <li class="nav-item` + getActiveClass(pageName, "faq") + `">
                        <a href="/faq" class="nav-link">FAQ</a>
                    </li>
                    <li class="nav-item` + getActiveClass(pageName, "contact") + `">
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
                        <h1 class="page-title">` + getPageTitle(pageName) + `</h1>
                    </div>
                </div>

                <div class="page-content">
                    <div class="container">
                        ` + getPageContent(pageName) + `
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

	return baseTemplate
}

// アクティブなナビゲーションクラスを返す
func getActiveClass(currentPage, targetPage string) string {
	if currentPage == targetPage {
		return " active"
	}
	return ""
}

// ページタイトルを返す
func getPageTitle(pageName string) string {
	switch pageName {
	case "services":
		return "サービス"
	case "products":
		return "開発製品"
	case "results":
		return "実績"
	case "about":
		return "スキルスタック"
	case "faq":
		return "FAQ"
	case "contact":
		return "お問い合わせ"
	default:
		return pageName
	}
}

func getPageContent(pageName string) string {
	switch pageName {
	case "services":
		return `
                        <section class="section">
                            <h2>infohiroki 生成AI導入支援サービス</h2>
                            <p>ChatGPT・Claude・Geminiを活用した企業の生成AI導入を総合的に支援します🤖</p>
                        </section>

                        <!-- サービス価格プラン -->
                        <section class="section">
                            <h2 class="section-title">サービス・料金体系</h2>
                            <div class="service-plans">
                                <div class="plan-card recommended">
                                    <div class="plan-badge">人気</div>
                                    <h3>技術顧問サービス</h3>
                                    <div class="plan-price">月額 5万円</div>
                                    <p>月15時間の継続的な生成AI活用技術支援</p>
                                </div>
                                <div class="plan-card">
                                    <h3>生成AI導入プロジェクト</h3>
                                    <div class="plan-price">20〜500万円</div>
                                    <p>企業規模・内容に応じた生成AI導入プロジェクト</p>
                                </div>
                                <div class="plan-card spot">
                                    <h3>AI導入相談</h3>
                                    <div class="plan-price">1回 2万円</div>
                                    <p>1.5時間の相談で現状分析レポート＋具体的改善提案書を提供</p>
                                </div>
                            </div>
                        </section>

                        <!-- サービス詳細 -->
                        <section class="section">
                            <h2 class="section-title">各サービス詳細</h2>

                            <div class="service-detail">
                                <h3>🚀 生成AI導入プロジェクト（20〜500万円）</h3>
                                <ul class="service-features">
                                    <li><strong>フェーズ1：</strong>現状分析・課題特定・AI選定</li>
                                    <li><strong>フェーズ2：</strong>システム設計・API連携・テスト実装</li>
                                    <li><strong>フェーズ3：</strong>本格導入・社員研修・運用定着</li>
                                    <li><strong>期間：</strong>2〜6ヶ月・企業規模に応じた最適提案</li>
                                </ul>
                            </div>

                            <div class="service-detail">
                                <h3>🤝 技術顧問サービス（月額5万円）</h3>
                                <ul class="service-features">
                                    <li><strong>月間稼働時間：</strong>15時間</li>
                                    <li><strong>対応範囲：</strong>生成AI活用相談・技術実装・運用支援</li>
                                    <li><strong>契約期間：</strong>6ヶ月または12ヶ月</li>
                                    <li><strong>サポート：</strong>月次レポート・緊急対応含む</li>
                                </ul>
                            </div>

                            <div class="service-detail">
                                <h3>💡 AI導入相談（1回2万円）</h3>
                                <ul class="service-features">
                                    <li><strong>時間：</strong>1.5時間の詳細ヒアリング</li>
                                    <li><strong>成果物：</strong>現状分析レポート＋具体的導入提案書</li>
                                    <li><strong>内容：</strong>最適AI選定・費用見積・期待効果算出</li>
                                    <li><strong>対象：</strong>生成AI導入を検討中の企業様</li>
                                </ul>
                            </div>
                        </section>

                        <!-- 対応可能技術・分野 -->
                        <section class="section">
                            <h2 class="section-title">対応可能な生成AI・技術分野</h2>
                            <div class="card-grid">
                                <div class="card">
                                    <div class="card-header">
                                        <h3 class="card-title">🤖 主要生成AI</h3>
                                    </div>
                                    <div class="card-body">
                                        <p><strong>ChatGPT（OpenAI）</strong> - 汎用的な文書作成・質問応答<br>
                                        <strong>Claude（Anthropic）</strong> - 長文分析・複雑な推論<br>
                                        <strong>Gemini（Google）</strong> - Google Workspace連携</p>
                                    </div>
                                </div>
                                <div class="card">
                                    <div class="card-header">
                                        <h3 class="card-title">🎤 音声・画像AI</h3>
                                    </div>
                                    <div class="card-body">
                                        <p><strong>Whisper（OpenAI）</strong> - 音声認識・議事録自動化<br>
                                        <strong>DALL-E / Midjourney</strong> - 画像生成・デザイン支援<br>
                                        <strong>音声合成</strong> - 自動アナウンス・読み上げ</p>
                                    </div>
                                </div>
                                <div class="card">
                                    <div class="card-header">
                                        <h3 class="card-title">⚙️ システム連携</h3>
                                    </div>
                                    <div class="card-body">
                                        <p><strong>API統合</strong> - 既存システムとAIの連携<br>
                                        <strong>Notion AI連携</strong> - データベース×生成AI<br>
                                        <strong>Excel/Google Sheets</strong> - スプレッドシート自動化</p>
                                    </div>
                                </div>
                                <div class="card">
                                    <div class="card-header">
                                        <h3 class="card-title">🏢 業務領域</h3>
                                    </div>
                                    <div class="card-body">
                                        <p><strong>議事録・文書作成</strong> - 会議効率化・資料自動生成<br>
                                        <strong>顧客対応</strong> - チャットボット・メール自動返信<br>
                                        <strong>データ分析</strong> - レポート生成・傾向分析</p>
                                    </div>
                                </div>
                                <div class="card">
                                    <div class="card-header">
                                        <h3 class="card-title">🎯 特化支援</h3>
                                    </div>
                                    <div class="card-body">
                                        <p><strong>医療・介護</strong> - 議事録、カルテ支援（桜十字病院実績）<br>
                                        <strong>製造・建設</strong> - 作業記録、安全管理<br>
                                        <strong>サービス業</strong> - 顧客対応、業務マニュアル</p>
                                    </div>
                                </div>
                            </div>
                        </section>

                        <!-- LINE連絡セクション -->
                        <section class="line-contact-footer">
                            <div class="container">
                                <h3>💬 サービスについてご相談ください</h3>
                                <p>技術顧問サービスの詳細やお見積もりについて、LINEでお気軽にご相談ください</p>
                                <a href="https://lin.ee/8ymv2Nw" target="_blank" rel="noopener noreferrer" class="line-button-footer">
                                    LINEで相談する
                                </a>
                            </div>
                        </section>
		`
	case "products":
		return `
<section class="section">
    <p>infohirokiでは、実際の業務で使える効率化ツールを開発・提供しています。すべて実際の導入実績があり、現場で使われている製品です。</p>
</section>

<section class="section">
    <div class="card-grid">
        <div class="card" id="koemoji">
            <div class="card-header">
                <h3 class="card-title">Koemoji-Go</h3>
                <p class="card-subtitle">オールインワン音声処理システム</p>
            </div>
            <div class="card-body">
                <p><strong>機能：</strong> 録音→文字起こし→AI要約の完全自動化</p>
                <p><strong>特徴：</strong> GUI/TUI対応、フォルダ監視、シングルバイナリ配布</p>
                <p><strong>技術：</strong> Go、Fyne、FasterWhisper、OpenAI API</p>
                <p><strong>対応：</strong> Windows、macOS（Apple Silicon対応）</p>
                <div class="card-links">
                    <a href="https://github.com/infoHiroki/KoeMoji-Go" target="_blank" class="link-button">GitHub (Go版)</a>
                    <a href="https://github.com/infoHiroki/KoeMojiAuto-cli" target="_blank" class="link-button">GitHub (Python版)</a>
                    <a href="https://koemoji.hmtc.jp/index.html" target="_blank" class="link-button">紹介ページ</a>
                </div>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">YouTube MojiCopy</h3>
                <p class="card-subtitle">Chrome拡張機能</p>
            </div>
            <div class="card-body">
                <p><strong>機能：</strong> プロンプト保存機能付きYouTube文字起こしコピー</p>
                <p><strong>特徴：</strong> 「要約して」等のプロンプト＋文字起こしでLLMに直接ペースト可能</p>
                <p><strong>技術：</strong> JavaScript、Chrome Extension API</p>
                <p><strong>対応：</strong> ChatGPT、Claude等のLLMとの連携</p>
                <div class="card-links">
                    <a href="https://chromewebstore.google.com/detail/youtubemojicopy/ejeafnfdgeipigfackgkhcgfbjiijbgf" target="_blank" class="link-button">Chrome Store</a>
                </div>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">NotionTasker</h3>
                <p class="card-subtitle">Notion連携Chrome拡張</p>
            </div>
            <div class="card-body">
                <p><strong>機能：</strong> WebページからNotionへ直接タスク・メモ追加</p>
                <p><strong>技術：</strong> JavaScript、Notion API、Chrome Extension</p>
                <p><strong>対応：</strong> Notion データベースと連携</p>
                <div class="card-links">
                    <a href="https://chromewebstore.google.com/detail/notiontasker/pkbibgfhgicoahenebmkhbklkffjdfea" target="_blank" class="link-button">Chrome Store</a>
                </div>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">Webサイト制作</h3>
                <p class="card-subtitle">コーポレートサイト・LP制作</p>
            </div>
            <div class="card-body">
                <p><strong>機能：</strong> 企業サイト、ランディングページ、ブログサイト</p>
                <p><strong>特徴：</strong> レスポンシブ対応、SEO最適化、高速表示</p>
                <p><strong>技術：</strong> HTML/CSS、JavaScript、静的サイトジェネレーター</p>
                <p><strong>対応：</strong> お問い合わせフォーム、アナリティクス連携</p>
                <div class="card-links">
                    <a href="/contact" class="link-button">お問い合わせ</a>
                </div>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">Webアプリケーション</h3>
                <p class="card-subtitle">カスタム業務システム開発</p>
            </div>
            <div class="card-body">
                <p><strong>機能：</strong> 顧客管理、在庫管理、予約システム等</p>
                <p><strong>特徴：</strong> 完全カスタマイズ、データベース連携</p>
                <p><strong>技術：</strong> JavaScript、React、Node.js、PostgreSQL</p>
                <p><strong>対応：</strong> Vercel、Firebase等のクラウドデプロイ</p>
                <div class="card-links">
                    <a href="/contact" class="link-button">お問い合わせ</a>
                </div>
            </div>
        </div>
    </div>
</section>

<!-- LINE連絡セクション -->
<section class="line-contact-footer">
    <div class="container">
        <h3>💬 製品についてご相談ください</h3>
        <p>開発製品の導入や詳細について、LINEでお気軽にご相談ください</p>
        <a href="https://lin.ee/8ymv2Nw" target="_blank" rel="noopener noreferrer" class="line-button-footer">
            LINEで相談する
        </a>
    </div>
</section>
		`
	case "results":
		return `
<section class="section">
    <div class="card-grid">
        <div class="card">
            <div class="card-header">
                <h3 class="card-title">桜十字福岡病院様</h3>
                <p class="card-subtitle">医療業界 - Whisper活用議事録自動化システム導入</p>
            </div>
            <div class="card-body">
                <p><strong>課題：</strong> 会議・打ち合わせの議事録作成に月40時間費やしていた</p>
                <p><strong>解決策：</strong> OpenAI Whisper + <a href="/products#koemoji" style="color: #E73E8F; text-decoration: none;">Koemoji</a>システムによる音声認識自動化</p>
                <p><strong>成果：</strong> 議事録作成80%時短・月15万円コスト削減・転記ミス0件達成</p>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">地域密着型歯科医院様</h3>
                <p class="card-subtitle">医療業界 - Notion導入支援＋業務自動化</p>
            </div>
            <div class="card-body">
                <p><strong>課題：</strong> 患者管理とメール対応業務の効率化</p>
                <p><strong>解決策：</strong> Notion導入支援＋GAS活用でメール自動下書き機能</p>
                <p><strong>成果：</strong> 継続案件として機能追加、業務フロー改善</p>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">訪問看護事業所様</h3>
                <p class="card-subtitle">医療業界 - ITコンサル・デジタル化支援</p>
            </div>
            <div class="card-body">
                <p><strong>課題：</strong> WebマーケティングとNotion活用</p>
                <p><strong>解決策：</strong> WebマーケティングへのAI技術導入・Notion運用支援</p>
                <p><strong>成果：</strong> ITコンサルによる業務デジタル化推進</p>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">建築事務所様</h3>
                <p class="card-subtitle">建築業界 - 報告書業務効率化</p>
            </div>
            <div class="card-body">
                <p><strong>課題：</strong> 点検調査データ管理と報告書作成</p>
                <p><strong>解決策：</strong> マクロ自動化・ドローン解析結果効率化</p>
                <p><strong>成果：</strong> 報告書作成時間短縮、データ管理システム構築</p>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">建設コンサル様</h3>
                <p class="card-subtitle">建設業界 - 現場記録自動化</p>
            </div>
            <div class="card-body">
                <p><strong>課題：</strong> 現場会議の記録業務効率化</p>
                <p><strong>解決策：</strong> AI音声文字起こしシステム導入</p>
                <p><strong>成果：</strong> 現場会議の記録業務を効率化</p>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">MEO対策企業様</h3>
                <p class="card-subtitle">IT業界 - 大規模言語モデル活用支援</p>
            </div>
            <div class="card-body">
                <p><strong>課題：</strong> コンテンツ制作の効率化とクオリティ向上</p>
                <p><strong>解決策：</strong> ChatGPT・Claude活用フロー構築</p>
                <p><strong>成果：</strong> 制作スピード3倍向上、品質の標準化</p>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">個人飲食店様</h3>
                <p class="card-subtitle">飲食業界 - 業務管理システム統合</p>
            </div>
            <div class="card-body">
                <p><strong>課題：</strong> 在庫管理・日報・売上管理の一元化</p>
                <p><strong>解決策：</strong> 統合管理システム構築</p>
                <p><strong>成果：</strong> 業務管理の効率化、データ一元化</p>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">制作会社様</h3>
                <p class="card-subtitle">エンタメ業界 - Notion導入支援</p>
            </div>
            <div class="card-body">
                <p><strong>課題：</strong> プロジェクト管理の統一とチーム連携</p>
                <p><strong>解決策：</strong> Notionデータベース設計・運用支援</p>
                <p><strong>成果：</strong> 情報共有効率化、プロジェクト進行の可視化</p>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3 class="card-title">AI研修受講企業様</h3>
                <p class="card-subtitle">企業研修 - AI活用スキル習得支援</p>
            </div>
            <div class="card-body">
                <p><strong>課題：</strong> 実務に直結するAI活用スキル習得</p>
                <p><strong>解決策：</strong> 4時間×4回の企業向けAI研修プログラム</p>
                <p><strong>成果：</strong> 各業界でのAIツール導入・活用指導</p>
            </div>
        </div>
    </div>
</section>

<!-- LINE連絡セクション -->
<section class="line-contact-footer">
    <div class="container">
        <h3>💬 導入についてご相談ください</h3>
        <p>同様の成果を実現するソリューションについて、LINEでお気軽にご相談ください</p>
        <a href="https://lin.ee/8ymv2Nw" target="_blank" rel="noopener noreferrer" class="line-button-footer">
            LINEで相談する
        </a>
    </div>
</section>
		`
	case "about":
		return `
<section class="section">
    <h2>プログラミング言語</h2>
    <div class="skill-category">
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=python" class="tech-icon" alt="Python">Python</h3>
            <p>データ分析、AI/ML、Webアプリケーション開発</p>
        </div>
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=js" class="tech-icon" alt="JavaScript">JavaScript/TypeScript</h3>
            <p>フロントエンド・バックエンド開発</p>
        </div>
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=go" class="tech-icon" alt="Go">Go</h3>
            <p>マイクロサービス、API開発</p>
        </div>
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=postgresql" class="tech-icon" alt="SQL">SQL</h3>
            <p>PostgreSQL、MySQL、データ分析</p>
        </div>
    </div>
</section>

<section class="section">
    <h2>フレームワーク・ライブラリ</h2>
    <div class="skill-category">
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=react" class="tech-icon" alt="React">React/Next.js</h3>
            <p>Webアプリケーション開発、SSG/SSR</p>
        </div>
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=flask" class="tech-icon" alt="Flask">FastAPI/Flask</h3>
            <p>軽量Python Webフレームワーク、REST API開発</p>
        </div>
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=nodejs" class="tech-icon" alt="Node.js">Node.js/Express</h3>
            <p>バックエンドAPI、リアルタイム通信</p>
        </div>
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=tensorflow" class="tech-icon" alt="TensorFlow">TensorFlow/PyTorch</h3>
            <p>機械学習、深層学習モデル開発</p>
        </div>
    </div>
</section>

<section class="section">
    <h2>AI・機械学習</h2>
    <div class="skill-category">
        <div class="skill-item">
            <h3><span class="tech-icon">🤖</span>LLM活用</h3>
            <p>OpenAI API、Claude、カスタムGPT開発</p>
        </div>
        <div class="skill-item">
            <h3><span class="tech-icon">📈</span>データ分析</h3>
            <p>pandas、numpy、scikit-learn</p>
        </div>
        <div class="skill-item">
            <h3><span class="tech-icon">🧠</span>自然言語処理</h3>
            <p>テキスト分析、感情分析、文書生成</p>
        </div>
        <div class="skill-item">
            <h3><span class="tech-icon">🔧</span>モデル運用</h3>
            <p>APIデプロイ、バージョン管理、スクリプト自動化</p>
        </div>
    </div>
</section>

<section class="section">
    <h2>インフラ・ツール</h2>
    <div class="skill-category">
        <div class="skill-item">
            <h3><span class="tech-icon">☁️</span>クラウドサービス</h3>
            <p>Vercel、Supabase、Firebase</p>
        </div>
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=docker" class="tech-icon" alt="Docker">コンテナ</h3>
            <p>Docker、ローカル開発環境</p>
        </div>
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=githubactions" class="tech-icon" alt="GitHub Actions">自動化</h3>
            <p>GitHub Actions、デプロイ自動化</p>
        </div>
        <div class="skill-item">
            <h3><span class="tech-icon">📉</span>分析・モニタリング</h3>
            <p>Google Analytics、Sentry、エラートラッキング</p>
        </div>
    </div>
</section>

<section class="section">
    <h2>業務効率化・ツール</h2>
    <div class="skill-category">
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=notion" class="tech-icon" alt="Notion">Notion</h3>
            <p>ワークスペース設計、API活用、自動化</p>
        </div>
        <div class="skill-item">
            <h3><span class="tech-icon">⚡</span>Zapier/Make</h3>
            <p>ワークフロー自動化、システム連携</p>
        </div>
        <div class="skill-item">
            <h3><img src="https://skillicons.dev/icons?i=github" class="tech-icon" alt="GitHub">Git/GitHub</h3>
            <p>バージョン管理、チーム開発、OSS貢献</p>
        </div>
        <div class="skill-item">
            <h3><span class="tech-icon">📊</span>プロジェクト管理</h3>
            <p>アジャイル開発、スクラム、リーン手法</p>
        </div>
    </div>
</section>

<!-- LINE連絡セクション -->
<section class="line-contact-footer">
    <div class="container">
        <h3>💬 お気軽にご相談ください</h3>
        <p>技術的なお悩みや業務効率化について、LINEでお気軽にご相談ください</p>
        <a href="https://lin.ee/8ymv2Nw" target="_blank" rel="noopener noreferrer" class="line-button-footer">
            LINEで相談する
        </a>
    </div>
</section>
		`
	case "faq":
		return `
<section class="section">
    <div class="faq-simple">
        <div class="faq-item">
            <div class="faq-q">Q. 🏢 どのような業界・規模の企業が対象ですか？</div>
            <div class="faq-a">A. 業界問わず、小規模事業者から大企業まで幅広く対応しています。特に<strong>医療・建設・IT・製造・サービス業</strong>での生成AI導入実績が豊富です。桜十字病院でのWhisper活用から中小企業のChatGPT導入まで、多種多様な業界での生成AI導入を支援しています。</div>
        </div>

        <div class="faq-item">
            <div class="faq-q">Q. 📋 サービス開始までの流れを教えてください</div>
            <div class="faq-a">A. ①<strong>LINEでお問い合わせ</strong> → ②<strong>AI導入相談（2万円・1.5時間）</strong> → ③<strong>現状分析レポート＋導入提案書の提出</strong> → ④<strong>契約（プロジェクトまたは顧問）</strong> → ⑤<strong>サービス開始</strong>の流れです。まずはAI導入相談で具体的な効果と費用をご確認ください。</div>
        </div>

        <div class="faq-item">
            <div class="faq-q">Q. 💰 料金プランと追加費用について教えてください</div>
            <div class="faq-a">A. <strong>技術顧問サービス：月額5万円</strong>（6-12ヶ月契約）、<strong>生成AI導入プロジェクト：20-500万円</strong>（企業規模に応じて）、<strong>AI導入相談：1回2万円</strong>（1.5時間）です。基本的に追加費用はかかりませんが、特別なソフトウェアライセンスが必要な場合は事前にご相談いたします。</div>
        </div>

        <div class="faq-item">
            <div class="faq-q">Q. ⏰ 契約期間の縛りはありますか？</div>
            <div class="faq-a">A. <strong>生成AI導入プロジェクトは2-6ヶ月</strong>（規模により）、<strong>技術顧問サービスは6ヶ月または12ヶ月の契約期間</strong>があります。<strong>AI導入相談は単発のため契約期間の縛りはありません。</strong></div>
        </div>

        <div class="faq-item">
            <div class="faq-q">Q. 💻 オンラインでの対応は可能ですか？</div>
            <div class="faq-a">A. はい、全国どこでもオンラインで対応可能です。Zoom、Teams、Google Meet等のツールを使用し、画面共有やリモート操作でサポートします。</div>
        </div>

        <div class="faq-item">
            <div class="faq-q">Q. 🔒 データの安全性は保障されますか？</div>
            <div class="faq-a">A. お客様のデータは厳重に管理し、機密保持契約（NDA）の締結も可能です。作業終了後はお客様のデータを完全削除いたします。</div>
        </div>

        <div class="faq-item">
            <div class="faq-q">Q. 🤖 どのようなAIツールに対応していますか？</div>
            <div class="faq-a">A. <strong>ChatGPT、Claude、Gemini</strong>等の大規模言語モデル、<strong>音声認識AI（Whisper等）</strong>、<strong>Google Apps Script、Excel VBA</strong>、<strong>Notion、音声文字起こしシステム</strong>等、幅広いツールに対応しています。実際の導入実績に基づいて最適なツールをご提案します。</div>
        </div>

        <div class="faq-item">
            <div class="faq-q">Q. ❓ 技術的な知識がなくても大丈夫ですか？</div>
            <div class="faq-a">A. はい、技術的な知識は不要です。業務の課題や改善したい点をお聞かせいただければ、技術的な部分は全てお任せください。操作方法も丁寧にレクチャーいたします。</div>
        </div>
    </div>
</section>

<!-- LINE連絡セクション -->
<section class="line-contact-footer">
    <div class="container">
        <h3>💬 他にご質問はありませんか？</h3>
        <p>FAQで解決しないご質問があれば、LINEでお気軽にお問い合わせください</p>
        <a href="https://lin.ee/8ymv2Nw" target="_blank" rel="noopener noreferrer" class="line-button-footer">
            LINEで質問する
        </a>
    </div>
</section>
		`
	case "contact":
		return `
<section class="section">
    <p>infohirokiへのお問い合わせ方法をご案内いたします。お客様のご都合に合わせてお選びください。</p>
</section>

<section class="section">
    <div class="line-contact">
        <h2>🔥 推奨：LINEでお問い合わせ</h2>
        <p><strong>最も迅速で確実</strong>にご対応できます。以下のボタンからLINE公式アカウントを友だち追加して、お気軽にメッセージをお送りください。</p>

        <div class="contact-benefits">
            <h3>LINEでお問い合わせのメリット</h3>
            <ul>
                <li>✅ <strong>即座に対応</strong> - リアルタイムでやり取り可能</li>
                <li>✅ <strong>気軽に相談</strong> - チャット感覚で質問できます</li>
                <li>✅ <strong>ファイル共有</strong> - 画像や資料も簡単に送信</li>
                <li>✅ <strong>無料相談</strong> - 初回は完全無料でご相談承ります</li>
            </ul>
        </div>

        <div class="line-button-container">
            <a href="https://lin.ee/8ymv2Nw" target="_blank" rel="noopener noreferrer">
                <img src="https://scdn.line-apps.com/n/line_add_friends/btn/ja.png" alt="友だち追加" height="36" border="0">
            </a>
        </div>
        <div class="line-qr-container">
            <p>または、こちらのQRコードをLINEアプリで読み取ってください：</p>
            <img src="/images/LineQR.png" alt="LINE QRコード" class="line-qr">
        </div>
    </div>
</section>

<section class="section">
    <h2>📧 メールでのお問い合わせ</h2>
    <div class="contact-info">
        <p>LINEをご利用でない場合は、メールでもお問い合わせいただけます。</p>
        <p><strong>Email:</strong> <a href="mailto:info.hirokitakamura@gmail.com">info.hirokitakamura@gmail.com</a></p>
        <p><strong>営業時間:</strong> 平日 9:00-18:00（土日祝休み）</p>

        <div class="email-tips">
            <h3>メールでお問い合わせの際は</h3>
            <ul>
                <li>📝 ご相談内容を具体的にお書きください</li>
                <li>🏢 会社名・お名前をご記載ください</li>
                <li>📱 お急ぎの場合はLINEをご利用ください</li>
            </ul>
        </div>
    </div>
</section>

<section class="section">
    <h2>❓ お問い合わせ前のお役立ち情報</h2>
    <div class="helpful-links">
        <p>お問い合わせの前に、以下もご確認ください：</p>
        <ul>
            <li>🔍 <a href="/faq">よくある質問（FAQ）</a> - 一般的な疑問にお答えしています</li>
            <li>💼 <a href="/services">サービス詳細</a> - 技術顧問サービスの詳細情報</li>
            <li>📊 <a href="/results">導入実績</a> - 他社様での成功事例</li>
            <li>🛠️ <a href="/products">開発製品</a> - 既存の開発ツール一覧</li>
        </ul>
    </div>
</section>
		`
	default:
		return "<p>ページコンテンツが見つかりません。</p>"
	}
}

// Vercel用メインハンドラー
func Pages(w http.ResponseWriter, r *http.Request) {
	pageName := r.URL.Query().Get("page")

	switch pageName {
	case "services":
		Page(w, r, "services", "生成AI導入支援サービス | infoHiroki", "福岡・九州企業向け生成AI導入支援サービス - ChatGPT・Claude・Whisper活用で業務効率化")
	case "products":
		Page(w, r, "products", "開発製品 | infoHiroki", "infoHirokiが開発した製品・ツール・アプリケーション一覧")
	case "results":
		Page(w, r, "results", "実績 | infoHiroki", "infoHirokiの開発実績・導入事例・お客様の声")
	case "about":
		Page(w, r, "about", "スキルスタック | infoHiroki", "infoHirokiの技術スタック・経歴・スキル")
	case "faq":
		Page(w, r, "faq", "FAQ | infoHiroki", "よくある質問と回答 - infoHirokiサービスについて")
	case "contact":
		Page(w, r, "contact", "お問い合わせ | infoHiroki", "infoHirokiへのお問い合わせ・ご相談はこちら")
	default:
		http.Error(w, "ページが見つかりません", http.StatusNotFound)
	}
}