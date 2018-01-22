package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

// STORAGE -> check newer version of docker-stack.yml

func checkDockerStackFile() {

	fmt.Println("TEST1")

	bucket := "cellarhub-dockerstack-files"
	file := "docker-stack.yml"

	ctx := context.Background()

	// [START setup]
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	// [END setup]

	fmt.Println("TEST2")

	isChangedVariable := isChanged(client, bucket, file)

	fmt.Println("TEST3")

	if isChangedVariable {

		// list buckets from the project
		data, err := read(client, "cellarhub-dockerstack-files", "docker-stack.yml")
		if err != nil {
			log.Fatalf("Cannot read object: %v", err)
		}
		fmt.Printf("Object contents: %s\n", data)

		//Write to file
		err2 := ioutil.WriteFile("./docker-stack.yml", data, 0644)
		if err2 != nil {
			log.Fatalf("cannor write into file: %v", err2)
		}

	}

}

//************************************************************
//************************************************************
//************************************************************
//************************************************************
// GOOGLE CLOUD STORAGE

func isChanged(client *storage.Client, bucket, object string) bool {
	ctx := context.Background()
	// [START download_file]
	rc, err := client.Bucket(bucket).Object(object).Attrs(ctx)
	if err != nil {
		log.Fatalf("cannor get attribute for file: %v", err)
	}

	var lastUpdateInCloud = rc.Updated

	fmt.Println(lastUpdated)
	fmt.Println(lastUpdateInCloud)

	if lastUpdated.Before(lastUpdateInCloud) {
		fmt.Print("IN CLOUD IS NEWER VERSION")
		lastUpdated = lastUpdateInCloud
		return true
	}

	return false
}

func read(client *storage.Client, bucket, object string) ([]byte, error) {
	ctx := context.Background()
	// [START download_file]
	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
	// [END download_file]
}
