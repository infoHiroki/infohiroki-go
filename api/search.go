package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Search(w http.ResponseWriter, r *http.Request) {
	// データが初期化されていない場合は初期化
	if len(allPosts) == 0 {
		InitializeData()
	}

	query := r.URL.Query().Get("q")
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, _ := strconv.Atoi(limitStr)

	// ファイルベースでの検索
	posts := FilterPosts(allPosts, query, "")
	if len(posts) > limit {
		posts = posts[:limit]
	}

	response := map[string]interface{}{
		"posts": posts,
		"total": len(posts),
		"query": query,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}