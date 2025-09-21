package handler

import (
	"html/template"
	"net/http"
	"strings"
)

func Post(w http.ResponseWriter, r *http.Request) {
	// データが初期化されていない場合は初期化
	if len(allPosts) == 0 {
		InitializeData()
	}

	// URLパスからスラッグを抽出
	path := strings.TrimPrefix(r.URL.Path, "/blog/")
	slug := path

	// 拡張子をチェック
	if strings.HasSuffix(slug, ".md") {
		// .mdの場合、Markdown形式で返す
		slugWithoutExt := strings.TrimSuffix(slug, ".md")
		showBlogPostMarkdown(w, r, slugWithoutExt)
	} else if strings.HasSuffix(slug, ".json") {
		// .jsonの場合、JSON形式で返す
		slugWithoutExt := strings.TrimSuffix(slug, ".json")
		showBlogPostJSON(w, r, slugWithoutExt)
	} else {
		// 拡張子なしの場合、HTML形式で返す
		showBlogPost(w, r, slug)
	}
}

// ブログ記事詳細（HTML）
func showBlogPost(w http.ResponseWriter, r *http.Request, slug string) {
	post := GetBlogPostBySlug(slug)
	if post == nil {
		http.Error(w, "記事が見つかりません", http.StatusNotFound)
		return
	}

	// SEOメタデータの設定
	metaDescription := post.Description
	if metaDescription == "" {
		metaDescription = "infoHiroki - 福岡・九州の生成AI導入支援専門家のブログ記事"
	}

	// HTMLテンプレートの作成
	tmpl := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}} | infoHiroki</title>
    <meta name="description" content="{{.MetaDescription}}">
    <meta property="og:title" content="{{.Title}} | infoHiroki">
    <meta property="og:description" content="{{.MetaDescription}}">
    <meta property="og:type" content="article">
    <meta name="twitter:card" content="summary">
    <meta name="twitter:title" content="{{.Title}}">
    <meta name="twitter:description" content="{{.MetaDescription}}">
</head>
<body>
    <article class="blog-post-detail">
        <header>
            <h1>{{.Post.Title}}</h1>
            <div class="meta">
                <span class="date">{{.Post.CreatedDate.Format "2006年01月02日"}}</span>
                {{if .Post.Icon}}<span class="icon">{{.Post.Icon}}</span>{{end}}
            </div>
            {{if .Post.Description}}<p class="description">{{.Post.Description}}</p>{{end}}
        </header>
        <div class="content">
            {{.Content}}
        </div>
        <footer>
            <a href="/blog">← ブログ一覧に戻る</a>
        </footer>
    </article>
</body>
</html>`

	t, err := template.New("post").Parse(tmpl)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title           string
		MetaDescription string
		Post            interface{}
		Content         template.HTML
	}{
		Title:           post.Title,
		MetaDescription: metaDescription,
		Post:            post,
		Content:         post.RenderContent(),
	}

	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Execution error", http.StatusInternalServerError)
		return
	}
}

// ブログ記事詳細（Markdown）
func showBlogPostMarkdown(w http.ResponseWriter, r *http.Request, slug string) {
	post := GetBlogPostBySlug(slug)
	if post == nil {
		http.Error(w, "記事が見つかりません", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.Write([]byte(post.ToMarkdown()))
}

// ブログ記事詳細（JSON）
func showBlogPostJSON(w http.ResponseWriter, r *http.Request, slug string) {
	post := GetBlogPostBySlug(slug)
	if post == nil {
		http.Error(w, "記事が見つかりません", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// 簡単なJSON出力（本来はjson.Marshalを使用）
	response := `{"title":"` + post.Title + `","slug":"` + post.Slug + `","description":"` + post.Description + `"}`
	w.Write([]byte(response))
}