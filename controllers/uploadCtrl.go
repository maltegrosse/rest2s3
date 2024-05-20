package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/CSpecht/rest2s3/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func Ping(c *gin.Context) {

	c.JSON(http.StatusOK, "pong")
}
func DownloadFile(c *gin.Context) {
	ctx := context.Background()
	minioClient, err := models.MinioConnection()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	objectName := c.Param("objectFolder") + "/" + c.Param("objectName")
	fmt.Println(objectName)
	object, err := minioClient.GetObject(ctx, models.CurrentConfig.Bucket, objectName, minio.GetObjectOptions{})

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	stat, err := object.Stat()
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	splitted := strings.Split(objectName, "/")

	extraHeaders := map[string]string{
		"Content-Disposition": "attachment; filename=" + fmt.Sprint(splitted[:1]),
	}
	c.DataFromReader(http.StatusOK, stat.Size, stat.ContentType, object, extraHeaders)

}
func UploadFile(c *gin.Context) {
	ctx := context.Background()
	file, err := c.FormFile("fileUpload")

	if err != nil {
		c.AbortWithError(http.StatusConflict, err)
		return
	}

	buffer, err := file.Open()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer buffer.Close()

	minioClient, err := models.MinioConnection()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	uuid := uuid.New()
	objectName := uuid.String() + "/" + file.Filename
	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	info, err := minioClient.PutObject(ctx, models.CurrentConfig.Bucket, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	type result struct {
		PublicUrl  string
		Size       int64
		Expiration time.Time
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	c.JSON(http.StatusCreated, result{PublicUrl: models.CurrentConfig.PublicUrl+"/upload/" + info.Key, Size: info.Size, Expiration: info.Expiration})

}
