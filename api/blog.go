package handler

import (
	"html/template"
	"net/http"
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

	// HTMLテンプレートの作成
	tmpl := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ブログ | infoHiroki</title>
    <meta name="description" content="infoHirokiのブログ - 生成AI・技術・開発に関する記事を配信中">
    <meta property="og:title" content="ブログ | infoHiroki">
    <meta property="og:description" content="福岡・九州の生成AI導入支援専門家による技術ブログ">
    <meta property="og:type" content="website">
</head>
<body>
    <h1>ブログ</h1>
    <div class="search-form">
        <form method="GET">
            <input type="text" name="q" value="{{.Query}}" placeholder="記事を検索...">
            <input type="text" name="tag" value="{{.Tag}}" placeholder="タグ検索...">
            <button type="submit">検索</button>
        </form>
    </div>
    <div class="blog-posts">
        {{range .Posts}}
        <article class="blog-post">
            <h2><a href="/blog/{{.Slug}}">{{.Title}}</a></h2>
            <p>{{.Description}}</p>
            <div class="meta">
                <span class="date">{{.CreatedDate.Format "2006年01月02日"}}</span>
                {{if .Icon}}<span class="icon">{{.Icon}}</span>{{end}}
            </div>
        </article>
        {{end}}
    </div>
    <p>検索結果: {{len .Posts}}件</p>
</body>
</html>`

	t, err := template.New("blog").Parse(tmpl)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Posts []interface{}
		Query string
		Tag   string
	}{
		Posts: make([]interface{}, len(posts)),
		Query: query,
		Tag:   tag,
	}

	for i, post := range posts {
		data.Posts[i] = post
	}

	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Execution error", http.StatusInternalServerError)
		return
	}
}