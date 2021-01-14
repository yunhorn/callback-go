package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FrontReq struct {
	Data string `json:"data"`
}

func main() {
	r := gin.Default()
	r.POST("/callback", func(c *gin.Context) {
		data := c.PostForm("data")
		log.Println("data:", data)
		c.JSON(200, nil)
	})

	srv := &http.Server{
		Addr:        ":8080",
		Handler:     r,
		ReadTimeout: time.Second * 3,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
	select {}
}
