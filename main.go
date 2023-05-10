package main

import (
	"log"
	"os"

	"github.com/CSpecht/rest2s3/models"
	routes "github.com/CSpecht/rest2s3/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	entrypoint := GetEnv("MINIO_ENDPOINT", "play.min.io")
	port := GetEnv("MINIO_PORT", "9443")
	accesskey := GetEnv("MINIO_ACCESSKEY", "Q3AM3UQ867SPQQA43P2F")
	secretkey := GetEnv("MINIO_SECRETKEY", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG")
	bucket := GetEnv("MINIO_BUCKET", "test")
	publicUrl := "https://" + entrypoint + "/"
	models.CurrentConfig = models.Config{Entrypoint: entrypoint + ":" + port, Bucket: bucket, AccessKey: accesskey, SecretKey: secretkey, PublicUrl: publicUrl}
	log.Printf("Using endpoint: %s ", entrypoint)
	gin.SetMode(gin.ReleaseMode)
	r := routes.Routes()
	r.Run()

}

func GetEnv(name string, def string) string {
	if len(os.Getenv(name)) > 0 {
		return os.Getenv(name)
	} else {
		return def
	}
}
