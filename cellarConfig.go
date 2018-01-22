package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//*****************************************************
// VERSION
//*****************************************************
var cellarVersion = "0.4.6"
var cellarDeviceInfo = ""

var path = "./cellarConfig.txt"
var cellarDeviceID = ""
var cellarHostName = ""
var cellarMACaddress = ""
var cellarIPaddress = ""

// var cellarIPaddressOutside = ""
var cellarMachineID = ""
var cellarDeviceHardware = ""
var cellarHubManagerPID = 0

var googleCloudProjectID = "cellarstone-1488228226623"
var googleCloudPubsubTopic = ""

func checkCellarDeviceInfo() {
	createFile()

	configInfo := readFileLineByLine()

	if len(configInfo) == 0 {
		cellarDeviceID = randStringBytesMaskImprSrc(12)
	} else {
		cellarDeviceID = configInfo[0]
	}

	cellarHostName, _ = os.Hostname()
	cellarMACaddress = getMacAddr()
	cellarMachineID = getMachineID()
	// cellarIPaddressOutside = GetOutboundIP()
	cellarIPaddress = GetLocalIP()
	cellarDeviceHardware = getCpuInfo()
	cellarDeviceHardware += getHDDInfo()
	cellarDeviceHardware += getNetworkInfo()
	cellarHubManagerPID = os.Getpid()
	pidString := strconv.Itoa(cellarHubManagerPID)

	logger.Information("CellarDeviceID : " + cellarDeviceID)
	logger.Information("Hostname : " + cellarHostName)
	logger.Information("MAC address : " + cellarMACaddress)
	logger.Information("IP address : " + cellarIPaddress)
	// logger.Information("IP out address : " + cellarIPaddressOutside)
	logger.Information("MachineID : " + cellarMachineID)
	logger.Information("PID : " + pidString)
	logger.Information("Hardware : " + cellarDeviceHardware)

	cellarDeviceInfo = ""
	cellarDeviceInfo += fmt.Sprintf("%v\n", cellarDeviceID)
	cellarDeviceInfo += fmt.Sprintf("%v\n", cellarHostName)
	cellarDeviceInfo += fmt.Sprintf("%v\n", cellarMACaddress)
	cellarDeviceInfo += fmt.Sprintf("%v\n", cellarIPaddress)
	// cellarDeviceInfo += fmt.Sprintf("%v\n", cellarIPaddressOutside)
	cellarDeviceInfo += fmt.Sprintf("%v\n", cellarMachineID)
	cellarDeviceInfo += fmt.Sprintf("%v\n", pidString)
	cellarDeviceInfo += fmt.Sprintf("%v\n", cellarDeviceHardware)

	deleteFile()
	writeFile()
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

		fmt.Println("==> done creating file", path)
	} else {
		fmt.Println("==> config file already exist", path)
	}

}

func readFileLineByLine() []string {
	result := []string{}

	inFile, _ := os.Open(path)

	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func writeAllToFile(text string) {
	d1 := []byte(text)
	err := ioutil.WriteFile(path, d1, 0644)
	if isError(err) {
		fmt.Println(err.Error())
	}
}

func writeFile() {
	d1 := []byte(cellarDeviceInfo)
	err := ioutil.WriteFile(path, d1, 0644)
	if isError(err) {
		fmt.Println(err.Error())
	}
}

func readFile() string {
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
