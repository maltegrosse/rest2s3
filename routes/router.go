package routes

import (
	"github.com/CSpecht/rest2s3/controllers"
	"github.com/CSpecht/rest2s3/models"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	if models.CurrentConfig.Mode == "ui"{
		r.GET("/", controllers.ShowHomepage)
	}
	
	r.POST("/upload", controllers.UploadFile)
	r.GET("/ping", controllers.Ping)
	r.GET("/upload/:objectFolder/:objectName", controllers.DownloadFile)
	return r
}
