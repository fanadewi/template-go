package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fanadewi/template-go/config"
	u "github.com/fanadewi/template-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	u.EnsureFolderExist("tmp")
	u.EnsureFolderExist("logs")
	// WIP set log output to logs/yyyy-mm-dd.log
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Sprintf("Error loading .env file: %s", err.Error()))
	}
	log.Println("Init success")
}

func main() {
	conf := config.New()
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"reply": "pong"})
	})

	r.Run(fmt.Sprintf(":%d", conf.Port))
}
