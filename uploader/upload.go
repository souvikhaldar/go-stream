package uploader

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---Uploading to s3---")
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error in reading from request file: ", err)
		return
	}
	fmt.Println("File details:", header.Filename)
	reader, writer := io.Pipe()
	// compressing the file before upload
	go func() {
		io.Copy(writer, file)
		file.Close()
		writer.Close()
	}()
	sess, e := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)
	if e != nil {
		fmt.Println("Error in creating session of s3 upload: ", e)
		return
	}

	up := s3manager.NewUploader(sess)
	result, er := up.Upload(&s3manager.UploadInput{
		Body:   reader,
		Bucket: aws.String("go-streamer"),
		Key:    aws.String(header.Filename),
	})
	if er != nil {
		fmt.Println("Failed to upload to s3: ", er)
		return
	}
	fmt.Println("Successfully uploaded: ", result.Location)
	u := uuid.NewV4()
	RedisClient.Set(u.String(), header.Filename, 0)
	fmt.Fprintf(w, "ID: %s URL: %s \n", u.String(), result.Location)
	return
}
