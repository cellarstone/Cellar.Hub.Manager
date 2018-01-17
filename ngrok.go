package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	resty "gopkg.in/resty.v1"
)

//-------------------------------------
//NGROK
//-------------------------------------
var ngrokID = ""
var ngrokDescription = ""
var ngrokApiToken = "6m5AcBUnnwPQcEyaj4mKR_3Z8N9koigW53vEehtYDNj"
var ngrokAuthToken = ""

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
	//Get Information about this machine
	hostname, _ := os.Hostname()
	macaddress := getMacAddr()

	deviceName := hostname + " - " + cellarDeviceID + " - " + macaddress

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

		fmt.Printf("\nResponse Body: %v", resp)

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

	//asdf := c5.Process.Pid
}