package main

import (
	"fmt"
	handler "infohiroki-go/api"
)

func main() {
	fmt.Println("データ初期化をテスト中...")
	handler.InitializeData()

	posts := handler.GetAllPosts()
	fmt.Printf("読み込まれた記事数: %d\n", len(posts))

	if len(posts) > 0 {
		fmt.Printf("最初の記事: %s\n", posts[0].Title)
	}
}