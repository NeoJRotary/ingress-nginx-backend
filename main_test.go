package main

import (
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func TestFileList(t *testing.T) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		t.Error(err)
	}

	bucket := client.Bucket(os.Getenv("GCS_BUCKET"))

	prefix := os.Getenv("CONFIG_FOLDER")
	it := bucket.Objects(ctx, &storage.Query{
		Prefix: prefix,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Error(err)
		}
		fmt.Println(attrs.Name)
	}
}

func TestDownload(t *testing.T) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		t.Error(err)
	}

	bucket := client.Bucket(os.Getenv("GCS_BUCKET"))

	prefix := os.Getenv("CONFIG_FOLDER")
	it := bucket.Objects(ctx, &storage.Query{
		Prefix: prefix,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Error(err)
		}
		download(bucket, "./testDir/", prefix, attrs.Name)
	}
}
