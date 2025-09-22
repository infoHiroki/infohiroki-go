package main

import (
	"github.com/gin-gonic/gin"
	handler "infohiroki-go/api"
)

func main() {
	r := gin.Default()

	// Vercel Functions をローカルでテスト
	r.GET("/", func(c *gin.Context) {
		handler.Index(c.Writer, c.Request)
	})

	r.GET("/blog", func(c *gin.Context) {
		handler.Blog(c.Writer, c.Request)
	})

	r.GET("/blog/:slug", func(c *gin.Context) {
		// スラッグをパスに含めてリダイレクト
		slug := c.Param("slug")
		c.Request.URL.Path = "/blog/" + slug
		handler.Post(c.Writer, c.Request)
	})

	r.GET("/api/search", func(c *gin.Context) {
		handler.Search(c.Writer, c.Request)
	})

	r.GET("/services", func(c *gin.Context) {
		c.Request.URL.RawQuery = "page=services"
		handler.Pages(c.Writer, c.Request)
	})

	r.GET("/products", func(c *gin.Context) {
		c.Request.URL.RawQuery = "page=products"
		handler.Pages(c.Writer, c.Request)
	})

	r.GET("/results", func(c *gin.Context) {
		c.Request.URL.RawQuery = "page=results"
		handler.Pages(c.Writer, c.Request)
	})

	r.GET("/about", func(c *gin.Context) {
		c.Request.URL.RawQuery = "page=about"
		handler.Pages(c.Writer, c.Request)
	})

	r.GET("/faq", func(c *gin.Context) {
		c.Request.URL.RawQuery = "page=faq"
		handler.Pages(c.Writer, c.Request)
	})

	r.GET("/contact", func(c *gin.Context) {
		c.Request.URL.RawQuery = "page=contact"
		handler.Pages(c.Writer, c.Request)
	})

	// 静的ファイル配信
	r.Static("/css", "./public/css")
	r.Static("/js", "./public/js")
	r.Static("/images", "./public/images")

	r.Run(":8081")
}