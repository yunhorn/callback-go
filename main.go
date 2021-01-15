package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FrontReq struct {
	Data string `json:"data"`
}

func main() {
	r := gin.Default()
	var path string
	var port int64
	flag.Int64Var(&port, "p", 8080, "服务器地址")
	flag.StringVar(&path, "path", "/", "数据接收路径")
	flag.Parse()

	r.POST(path, func(c *gin.Context) {
		data := c.PostForm("data")
		log.Println("data:", data)
		c.JSON(200, nil)
	})

	srv := &http.Server{
		Addr:        ":" + strconv.FormatInt(port, 10),
		Handler:     r,
		ReadTimeout: time.Second * 3,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
	select {}
}
