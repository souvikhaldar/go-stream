package downloader

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/souvikhaldar/go-streamer/uploader"
	"github.com/souvikhaldar/mux"
)

func Download(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---Downloading from S3---")
	vars := mux.Vars(r)
	id := vars["ID"]

	key, e := uploader.RedisClient.Get(id).Result()
	if e != nil {
		fmt.Println("Error in fetching from redis: ", e)
		return
	}
	sess, e := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)
	if e != nil {
		fmt.Println("Error in creating session of s3 upload: ", e)
		return
	}
	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create("downloaded-from-s3" + key)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String("go-streamer"),
			Key:    aws.String(key),
		})

	if err != nil {
		fmt.Println("Unable to download item %q, %v", key, err)
		return
	}
	file.Close()
	fmt.Println("Downloaded", key+"downloaded-from-s3", numBytes, "bytes")
	fmt.Fprintln(w, "Download success!")
}
