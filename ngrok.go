package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	resty "gopkg.in/resty.v1"
)

//-------------------------------------
//NGROK
//-------------------------------------
var ngrokID = ""
var ngrokDescription = ""
var ngrokApiToken = "6m5AcBUnnwPQcEyaj4mKR_3Z8N9koigW53vEehtYDNj"
var ngrokAuthToken = ""

var ngrokProcesses []int

type MessageRequest struct {
	Description string   `json:"description"`
	Acl         []string `json:"acl"`
}

type CredentialDTO struct {
	ID          string   `json:"id"`
	Token       string   `json:"token"`
	Description string   `json:"description"`
	Acl         []string `json:"acl"`
	Url         string   `json:"uri"`
}

type GetCredentialsResponse struct {
	Credentials []CredentialDTO `json:"credentials"`
	Url         string          `json:"uri"`
}

func checkNgrok() {
	isAlreadyRunning := checkRunningNgrok("http", "10001")
	if !isAlreadyRunning {
		go runNgrok("http", "10001")
	}

	isAlreadyRunning = checkRunningNgrok("tcp", "22")
	if !isAlreadyRunning {
		go runNgrok("tcp", "22")
	}
}

func checkIfDeviceExists(name string) string {

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(ngrokApiToken).
		Get("https://api.ngrok.com/credentials")

	if err != nil {
		logger.Error(err.Error())
		return "ERROR"
	}

	fmt.Printf("\nResponse Body: %v", resp)

	var m GetCredentialsResponse
	err = json.Unmarshal(resp.Body(), &m)
	if err != nil {
		logger.Error(err.Error())
		return "ERROR 2"
	}
	fmt.Println(m)

	result := ""
	for _, item := range m.Credentials {
		if item.Description == name {
			result = item.Token
		}
	}

	return result
}

func connectToNgrok() {
	getTokenNgrok()
	authorizeNgrok()
}

func getTokenNgrok() {
	deviceName := cellarHostName + " - " + cellarDeviceID + " - " + cellarMACaddress

	authtoken := checkIfDeviceExists(deviceName)

	if authtoken != "" {
		ngrokAuthToken = authtoken
	} else {
		body2 := MessageRequest{Description: deviceName, Acl: nil}

		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body2).
			SetAuthToken(ngrokApiToken).
			Post("https://api.ngrok.com/credentials")

		if err != nil {
			logger.Error(err.Error())
			return
		}

		//fmt.Printf("\nResponse Body: %v", resp)

		var m CredentialDTO
		err = json.Unmarshal(resp.Body(), &m)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		fmt.Println(m)

		ngrokAuthToken = m.Token
	}
}

func authorizeNgrok() {

	if ngrokAuthToken == "" {
		logger.Warning("ngrokAuthToken == ''")
		return
	}

	if ngrokAuthToken == "ERROR" {
		logger.Warning("ngrokAuthToken == 'ERROR'")
		return
	}

	cccmd := "./ngrok/ngrok authtoken " + ngrokAuthToken

	fmt.Println(cccmd)

	c5 := exec.Command("bash", "-c", cccmd)

	c6, err := c5.Output()
	if err != nil {
		logger.Error("ngrok error")
		logger.Error(err.Error())
	}
	data := printOutput(c6)

	fmt.Println(data)
}

func runNgrok(protocol string, port string) {

	if ngrokAuthToken == "" {
		logger.Warning("ngrokAuthToken == ''")
		return
	}

	cccmd := "./ngrok/ngrok " + protocol + " " + port
	c5 := exec.Command("bash", "-c", cccmd)

	c6, err := c5.Output()
	if err != nil {
		logger.Error("ngrok error")
		logger.Error(err.Error())
	}
	data := printOutput(c6)

	fmt.Println(data)

	pidString := strconv.Itoa(c5.Process.Pid)
	fmt.Println("run Ngrok process : " + pidString + "_______________________" + cccmd)
	ngrokProcesses = append(ngrokProcesses, c5.Process.Pid)
}

func checkRunningNgrok(protocol string, port string) bool {
	data := cmd_filterProcesses("ngrok")

	dataFormatted := strings.Split(data, "\n")

	for _, line := range dataFormatted {

		itt := strings.Split(line, "./ngrok/ngrok")
		if len(itt) > 1 {

			ittt := itt[1]
			vals := strings.Split(ittt, " ")

			temp_protocol := vals[1]
			temp_port := vals[2]

			// fmt.Println(temp_protocol)
			// fmt.Println(temp_port)

			if temp_protocol == protocol && temp_port == port {
				return true
			}
		}
	}

	return false
}

func killAllNgrokProcesses() {
	for _, item := range ngrokProcesses {
		pidString := strconv.Itoa(item)
		fmt.Println("killing Ngrok process : " + pidString)
		killProcess(item)
	}
}
