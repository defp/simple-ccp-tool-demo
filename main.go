package main

import (
	"ccp-tool/ccp_utils"
	"log"
	"net/http"
	"os"
	"path/filepath"

	ccpsdk "github.com/alibabacloud-go/ccppath-sdk/client"
	"github.com/gin-gonic/gin"
)

var akClient *ccpsdk.Client

func init() {
	var akConfig = new(ccpsdk.Config).
		SetDomainId(os.Getenv("DOMAIN_ID")).
		SetProtocol("https").
		SetAccessKeyId(os.Getenv("ACCESS_KEY_ID")).
		SetAccessKeySecret(os.Getenv("ACCESS_KEY_SECRET")).
		SetEndpoint(os.Getenv("ENDPOINT"))

	var err error
	akClient, err = ccpsdk.NewClient(akConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/files", func(c *gin.Context) {
		if body, err := ccp_utils.ListFiles(akClient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err_msg": err.Error()})
		} else {
			c.JSON(http.StatusOK, body)
		}
	})

	r.POST("/files/create", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "need file params"})
			return
		}

		filename := filepath.Base(fileHeader.Filename)
		file, err := fileHeader.Open()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file open error"})
			return
		}

		if createFileBody, err := ccp_utils.CreateFile(akClient, filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			if completeFileBody, err := ccp_utils.CompleteFile(akClient, createFileBody, file); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, completeFileBody)
			}
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
