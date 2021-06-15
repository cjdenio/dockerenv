package main

import (
	"log"
	"net/http"

	"github.com/cjdenio/dockerenv/pkg/images"
	"github.com/gin-gonic/gin"
)

func main() {
	imgs, err := images.LoadImages("images/")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/api/image/:image", func(c *gin.Context) {
		image, err := imgs.FindOne(c.Param("image"))
		if err != nil {
			c.String(http.StatusNotFound, "not found :(")
			return
		}

		c.JSON(http.StatusOK, image)
	})

	r.GET("/api/images", func(c *gin.Context) {
		c.JSON(http.StatusOK, imgs)
	})

	r.Run(":3000")
}
