package handler

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// HTMLテンプレートの作成
	tmpl := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>infoHiroki - 福岡の生成AI導入支援専門家</title>
    <meta name="description" content="福岡・九州の企業向け生成AI導入支援 - ChatGPT・Claude・Whisperで業務効率化を実現">
    <meta property="og:title" content="infoHiroki - 福岡の生成AI導入支援専門家">
    <meta property="og:description" content="福岡・九州の企業向け生成AI導入支援 - ChatGPT・Claude・Whisperで業務効率化を実現">
    <meta property="og:type" content="website">
</head>
<body>
    <header>
        <h1>infoHiroki</h1>
        <p>福岡の生成AI導入支援専門家</p>
    </header>
    <nav>
        <ul>
            <li><a href="/blog">ブログ</a></li>
            <li><a href="/services">サービス</a></li>
            <li><a href="/products">開発製品</a></li>
            <li><a href="/results">実績</a></li>
            <li><a href="/about">スキルスタック</a></li>
            <li><a href="/faq">FAQ</a></li>
            <li><a href="/contact">お問い合わせ</a></li>
        </ul>
    </nav>
    <main>
        <section class="hero">
            <h2>福岡・九州の企業向け生成AI導入支援</h2>
            <p>ChatGPT・Claude・Whisperで業務効率化を実現</p>
        </section>
        <section class="services">
            <h3>主なサービス</h3>
            <ul>
                <li>生成AI導入コンサルティング</li>
                <li>カスタムAIソリューション開発</li>
                <li>AI活用トレーニング</li>
                <li>業務プロセス最適化</li>
            </ul>
        </section>
    </main>
    <footer>
        <p>&copy; 2024 infoHiroki. All rights reserved.</p>
    </footer>
</body>
</html>`

	t, err := template.New("index").Parse(tmpl)
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