package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

var ctx = context.Background()

func init() {
	// use local credentials for test
	if _, err := os.Stat("/service_account.json"); os.IsNotExist(err) {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./service_account.json")
	}
}

func main() {
	// get client
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// get bucket handle
	bucket := client.Bucket(os.Getenv("BUCKET_NAME"))

	// list and download files
	prefix := os.Getenv("BUCKET_FOLDER")
	it := bucket.Objects(ctx, &storage.Query{
		Prefix: prefix,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		download(bucket, "/etc/nginx/conf.d/", prefix, attrs.Name)
	}
}

func download(bucket *storage.BucketHandle, toDir, prefix, name string) {
	rc, err := bucket.Object(name).NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fname := strings.Replace(name, prefix, "", 1)
	if fname == "" {
		return
	}

	log.Println("Dowload ", name)

	fullPath := filepath.Join(toDir, fname)
	dir, _ := filepath.Split(fullPath)
	os.MkdirAll(dir, os.ModePerm)

	f, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(f, rc); err != nil {
		log.Fatal(err)
	}
	if err := rc.Close(); err != nil {
		log.Fatal(err)
	}
}
