package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync/atomic"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type progressWriter struct {
	written int64
	writer  io.WriterAt
	size    int64
}

func parseFilename(keyString string) (filename string) {
	ss := strings.Split(keyString, "/")
	s := ss[len(ss)-1]
	return s
}

func getFileSize(svc *s3.S3, bucket string, prefix string) (filesize int64, error error) {
	params := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(prefix),
	}

	resp, err := svc.HeadObject(params)
	if err != nil {
		return 0, err
	}

	return *resp.ContentLength, nil
}

func (pw *progressWriter) WriteAt(p []byte, off int64) (int, error) {
	atomic.AddInt64(&pw.written, int64(len(p)))

	percentageDownloaded := float32(pw.written*100) / float32(pw.size)

	fmt.Printf("File size:%d downloaded:%d percentage:%.2f%%\r", pw.size, pw.written, percentageDownloaded)

	return pw.writer.WriteAt(p, off)
}

func main() {
	bucket := "slm-test-bucket-transactional"
	key := "brooklinen_test/brooklinen_txfile_11.1.2019-1.5.20.csv"

	filename := parseFilename(key)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	s3Client := s3.New(sess)
	downloader := s3manager.NewDownloader(sess)
	size, err := getFileSize(s3Client, bucket, key)
	if err != nil {
		panic(err)
	}

	log.Println("Starting download")
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	temp, err := ioutil.TempFile(cwd, "getObjWithProgress-tmp-")
	if err != nil {
		panic(err)
	}
	tempfileName := temp.Name()

	writer := &progressWriter{writer: temp, size: size, written: 0}
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	if _, err := downloader.Download(writer, params); err != nil {
		log.Printf("Download failed! Deleting tempfile: %s", tempfileName)
		os.Remove(tempfileName)
		panic(err)
	}

	if err := temp.Close(); err != nil {
		panic(err)
	}

	if err := os.Rename(temp.Name(), filename); err != nil {
		panic(err)
	}

	fmt.Println()
	log.Println("File downloaded! Avaliable at:", filename)
}
