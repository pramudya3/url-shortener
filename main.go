package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pramudya3/url-shortener/helper"
)

type ReqURL struct {
	URL string `json:"url" binding:"required"`
}

var mapURL = make(map[string]string)

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	//url shortener
	g := r.Group("api/v1/urls")
	g.POST("/shorten", func(c *gin.Context) {
		payload := &ReqURL{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, helper.ValidationError(err))
			return
		}
		// unixStr := helper.RandString(16)
		// url := fmt.Sprintf("localhost:8080/%s", unixStr)
		randCode := helper.RandString(16)
		mapURL[payload.URL] = randCode
		c.JSON(http.StatusCreated, randCode)
	})

	//go to url already shorten
	g.GET("/:url", func(c *gin.Context) {
		urlShorten := c.Param("url")
		// urlShorten := c.Query("url")
		if urlShorten == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
			return
		}

		url := ""
		for urlRaw, urlShort := range mapURL {
			if urlShorten == urlShort {
				url = urlRaw
				break
			}
		}

		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url not found"})
			return
		}

		c.Redirect(http.StatusPermanentRedirect, url)
	})

	//show all urls
	g.GET("/", func(c *gin.Context) {
		if len(mapURL) == 0 {
			c.JSON(http.StatusOK, mapURL)
			return
		}

		c.JSON(http.StatusOK, mapURL)
	})
	srv := &http.Server{
		Addr:    ":1234",
		Handler: r,
	}

	log.Fatalln(srv.ListenAndServe())
}
