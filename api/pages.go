package handler

import (
	"html/template"
	"net/http"
	"strings"
)

// 汎用ページハンドラー
func Page(w http.ResponseWriter, r *http.Request, pageName, title, description string) {
	// HTMLテンプレートの作成
	tmpl := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <meta name="description" content="{{.Description}}">
    <meta property="og:title" content="{{.Title}}">
    <meta property="og:description" content="{{.Description}}">
    <meta property="og:type" content="website">
</head>
<body>
    <header>
        <h1>{{.PageTitle}}</h1>
        <nav>
            <a href="/">ホーム</a> |
            <a href="/blog">ブログ</a> |
            <a href="/services">サービス</a> |
            <a href="/products">開発製品</a> |
            <a href="/results">実績</a> |
            <a href="/about">スキルスタック</a> |
            <a href="/faq">FAQ</a> |
            <a href="/contact">お問い合わせ</a>
        </nav>
    </header>
    <main>
        <h2>{{.PageTitle}}</h2>
        <p>{{.Description}}</p>
        {{.Content}}
    </main>
    <footer>
        <p>&copy; 2024 infoHiroki. All rights reserved.</p>
    </footer>
</body>
</html>`

	content := getPageContent(pageName)

	t, err := template.New("page").Parse(tmpl)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title       string
		PageTitle   string
		Description string
		Content     template.HTML
	}{
		Title:       title,
		PageTitle:   strings.Title(pageName),
		Description: description,
		Content:     template.HTML(content),
	}

	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Execution error", http.StatusInternalServerError)
		return
	}
}

func getPageContent(pageName string) string {
	switch pageName {
	case "services":
		return `
		<h3>生成AI導入支援サービス</h3>
		<ul>
			<li>ChatGPT・Claude・Whisper活用コンサルティング</li>
			<li>カスタムAIソリューション開発</li>
			<li>AI活用トレーニング・研修</li>
			<li>業務プロセス最適化支援</li>
		</ul>
		`
	case "products":
		return `
		<h3>開発製品・ツール</h3>
		<ul>
			<li>AI活用業務効率化ツール</li>
			<li>カスタムChatbotソリューション</li>
			<li>データ分析・可視化システム</li>
		</ul>
		`
	case "results":
		return `
		<h3>導入実績・お客様の声</h3>
		<ul>
			<li>企業向けAI導入プロジェクト実績</li>
			<li>業務効率化事例</li>
			<li>ROI改善実績</li>
		</ul>
		`
	case "about":
		return `
		<h3>技術スタック・経歴</h3>
		<ul>
			<li>プログラミング言語: Go, Python, JavaScript</li>
			<li>AI/ML: OpenAI API, Claude API, LangChain</li>
			<li>インフラ: AWS, Docker, Kubernetes</li>
		</ul>
		`
	case "faq":
		return `
		<h3>よくある質問</h3>
		<dl>
			<dt>Q: 導入期間はどの程度ですか？</dt>
			<dd>A: プロジェクトの規模により1週間〜3ヶ月程度です。</dd>
			<dt>Q: 料金体系を教えてください。</dt>
			<dd>A: プロジェクトベースまたは月額サポートをご用意しています。</dd>
		</dl>
		`
	case "contact":
		return `
		<h3>お問い合わせ</h3>
		<p>生成AI導入支援に関するご相談は、以下からお気軽にお問い合わせください。</p>
		<form>
			<p><label>お名前: <input type="text" name="name" required></label></p>
			<p><label>メール: <input type="email" name="email" required></label></p>
			<p><label>件名: <input type="text" name="subject" required></label></p>
			<p><label>メッセージ: <textarea name="message" required></textarea></label></p>
			<p><button type="submit">送信</button></p>
		</form>
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