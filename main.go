package main

import (
	"ccp-tool/ccp_utils"
	"log"
	"os"

	ccpsdk "github.com/alibabacloud-go/ccppath-sdk/client"
	"github.com/gin-gonic/gin"
)

var akClient  *ccpsdk.Client

func init()  {
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
		c.String(200, "pong")
	})

	r.GET("/files", func(c *gin.Context) {
		if body, err := ccp_utils.ListFiles(akClient); err != nil {
			c.JSON(400, gin.H{ "err_msg": err.Error()})
		} else {
			c.JSON(200, body)
		}
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
