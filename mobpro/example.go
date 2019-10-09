package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"math/rand"
	"net/http"
	"io/ioutil"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": myrand(),
			"qiita": qiita(),
		})
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func myrand() int {
    //rand.Seed(time.Now().UnixNano())
    my_rand := rand.New(rand.NewSource(time.Now().UnixNano()))
    return my_rand.Intn(6) + 1
}

func qiita() string {
	resp, _ := http.Get("https://qiita.com/api/v2/items")
	return execute(resp)
}

func execute(response *http.Response) string {
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return ""
	}
	return string(body)
}
