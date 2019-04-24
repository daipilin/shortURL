package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortURL/utils"
)

func main() {
	mysqlConn, largestKeyword := utils.NewDB("root","dssc", "10.112.119.224", 3306, "shortUrl")
	redisConn := utils.NewClient("127.0.0.1:6379", "", 0, largestKeyword)

	router := gin.Default()
	router.POST("/create", func(context *gin.Context) {
		url := context.PostForm("url")
		keyword := mysqlConn.GetKeyword(url)
		if "" == keyword {
			keyword = redisConn.NextKeyword()
			if !mysqlConn.Update(keyword, url) {
				keyword = mysqlConn.GetKeyword(url)
			}
		}
		redisConn.SetCache(keyword, url)
		context.JSON(http.StatusOK, gin.H{
			"Code": 0,
			"ShortUrl" : keyword,
			"LongUrl": url,
			"ErrMsg": "",
		})
	})

	router.POST("/query", func(context *gin.Context) {
		keyword := context.PostForm("shortUrl")
		url, found := redisConn.GetUrlFromCache(keyword)
		if !found {
			url = mysqlConn.GetUrlFromDB(keyword)
		}
		redisConn.SetCache(keyword, url)
		if "" != url {
			context.JSON(http.StatusOK, gin.H{
				"Code": 0,
				"ShortUrl" : keyword,
				"LongUrl": url,
				"ErrMsg": "",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"Code": -2,
				"ShortUrl" : keyword,
				"ErrMsg": "short url dose not exist",
			})
		}
	})

	router.Run("127.0.0.1:8000")
}