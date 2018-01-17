package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var path = "./cellarConfig.txt"
var cellarDeviceID = ""
var cellarHostName = ""
var cellarMACaddress = ""

func checkCellarDeviceInfo() {
	createFile()

	cellarDeviceID = readFile()
	cellarHostName, _ = os.Hostname()
	cellarMACaddress = getMacAddr()

	logger.Information("Hostname : " + cellarHostName)
	logger.Information("CellarDeviceID : " + cellarDeviceID)
	logger.Information("MAC address : " + cellarMACaddress)

	if cellarDeviceID == "" {
		cellarDeviceID = randStringBytesMaskImprSrc(12)
		writeFile(cellarDeviceID)
	}
}

func createFile() {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> done creating file", path)
}

func writeFile(id string) {
	// // open file using READ & WRITE permission
	// var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	// if isError(err) {
	// 	return
	// }
	// defer file.Close()

	// // write some text line-by-line to file
	// _, err = file.WriteString(id)
	// if isError(err) {
	// 	return
	// }

	// // save changes
	// err = file.Sync()
	// if isError(err) {
	// 	return
	// }

	// fmt.Println("==> done writing to file")

	d1 := []byte(id)
	err := ioutil.WriteFile(path, d1, 0644)
	if isError(err) {
		fmt.Println(err.Error())
	}
}

func readFile() string {
	// // re-open file
	// var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	// if isError(err) {
	// 	return err.Error()
	// }
	// defer file.Close()

	// // read file, line by line
	// var text = make([]byte, 1024)
	// for {
	// 	_, err = file.Read(text)

	// 	// break if finally arrived at end of file
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	// break if error occured
	// 	if err != nil && err != io.EOF {
	// 		isError(err)
	// 		break
	// 	}
	// }

	// fmt.Println("==> done reading from file")
	// fmt.Println(string(text))
	// return string(text)

	dat, err := ioutil.ReadFile(path)
	if isError(err) {
		return err.Error()
	}
	fmt.Print(string(dat))
	return string(dat)
}

func deleteFile() {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> done deleting file")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
