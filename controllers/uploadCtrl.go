package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/CSpecht/rest2s3/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func Ping(c *gin.Context) {

	c.JSON(http.StatusOK, "pong")
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

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	c.JSON(http.StatusCreated, info)

}
