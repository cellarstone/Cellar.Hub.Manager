package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

// STORAGE -> check newer version of docker-stack.yml

func checkDockerStackFile(filename string) {

	bucket := "cellarhub-dockerstack-files"
	//file := filename

	ctx := context.Background()

	// [START setup]
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// [END setup]

	isChangedVariable := isChanged(client, bucket, filename)

	if isChangedVariable {

		// list buckets from the project
		data, err := read(client, "cellarhub-dockerstack-files", filename)
		if err != nil {
			fmt.Printf("Cannot read object: %v", err)
		}
		//fmt.Printf("Object contents: %s\n", data)

		//remove file
		err = os.Remove("./" + filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("delete file: ", "./", filename)
		}

		//Write to file
		err2 := ioutil.WriteFile("./"+filename, data, 0755)
		if err2 != nil {
			fmt.Println("cannor write into file: ", err2)
		} else {
			fmt.Println("new file: ", "./", filename)
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
		fmt.Println("IN CLOUD IS NEWER VERSION")
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
