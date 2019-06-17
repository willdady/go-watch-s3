package main

import (
	"time"
	"log"
	"os"
	"io"
	"strings"
	"github.com/willdady/go-watch-s3/internal/utils"
	"github.com/radovskyb/watcher"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gabriel-vasile/mimetype"
)

func upload(uploader *s3manager.Uploader, path string, bucketName string, key string, storageClass string) {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open file %q, %v\n", path, err)
		return
	}
	defer f.Close()
	// Detect mimetype of file
	mime, _, err := mimetype.DetectReader(f)
	if err != nil {
		log.Printf("Failed to detect mimetype of file %q, %v\n", path, err)
	}
	if _, err = f.Seek(0, io.SeekStart); err != nil {
		log.Printf("Failed to reset seek to 0 for file %q, %v\n", path, err)
	}
	// Upload file to S3
	upParams := &s3manager.UploadInput{
		Bucket: &bucketName,
		Key: &key,
		StorageClass: &storageClass,
		ContentType: &mime,
		Body: f,
	}
	result, err := uploader.Upload(upParams)
	if err != nil {
		log.Printf("Failed to upload file %q, %v\n", path, err)
		return
	}
	log.Printf("Successfully uploaded %v\n", result.Location)

	// TODO: Add option to delete file after successful upload
}

func main() {
	// Get required environment variables
	utils.GetEnvOrPanic("AWS_REGION")
	watchPath := utils.GetEnvOrPanic("WATCH_PATH")
	bucketName := utils.GetEnvOrPanic("AWS_S3_BUCKET")
	// Get optional environment variables
	keyPrefix := utils.GetEnv("AWS_S3_KEY_PREFIX", "")
	storageClass := utils.GetEnv("AWS_S3_STORAGE_CLASS", "STANDARD")
	watchInterval, err := utils.GetEnvAsInt("WATCH_INTERVAL", 100)
	if err != nil {
		log.Panicln("Unable to parse WATCH_INTERVAL as integer")
	}
	// Instantiate uploader
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)
	// Instantiate watcher
	w := watcher.New()
	w.FilterOps(watcher.Create)

	// TODO: Add regex filtering

	go func() {
		for {
			select {
			case event := <-w.Event:
				if event.IsDir() {
					continue
				}
				var keyBuilder strings.Builder
				keyBuilder.WriteString(keyPrefix)
				keyBuilder.WriteString(strings.Replace(event.Path, watchPath, "", 1))
				key := keyBuilder.String()
				go upload(uploader, event.Path, bucketName, key, storageClass)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive(watchPath); err != nil {
		log.Fatalln(err)
	}
	// Start the watching process
	log.Printf("Watching: %v\n", watchPath)
	if err := w.Start(time.Duration(watchInterval * 100000)); err != nil {
		log.Fatalln(err)
	}
}
