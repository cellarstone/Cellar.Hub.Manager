package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type cellarDTO struct {
	Hostname       string      `json:"hostname"`
	Version        string      `json:"version"`
	DeviceID       string      `json:"deviceid"`
	IPaddress      string      `json:"ipaddress"`
	MACaddress     string      `json:"macaddress"`
	MachineID      string      `json:"machineid"`
	DeviceHardware string      `json:"devicehardware"`
	ExceptionText  string      `json:"exceptionText"`
	Data           interface{} `json:"data"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	dto := cellarDTO{
		Hostname:       cellarHostName,
		Version:        cellarVersion,
		DeviceID:       cellarDeviceID,
		IPaddress:      cellarIPaddress,
		MACaddress:     cellarMACaddress,
		MachineID:      cellarMachineID,
		DeviceHardware: cellarDeviceHardware,
		ExceptionText:  "",
		Data:           "",
	}

	indexTemplate.ExecuteTemplate(w, "layouttemplate", dto)
}

func processesNgrokHandler(w http.ResponseWriter, r *http.Request) {

	data := cmd_filterProcesses("ngrok")

	dataFormatted := strings.Split(data, "\n")
	dto := cellarDTO{
		Hostname:      cellarHostName,
		Version:       cellarVersion,
		ExceptionText: "",
		Data:          dataFormatted,
	}

	// logger.Information(data)

	ngrokprocessesTemplate.ExecuteTemplate(w, "layouttemplate", dto)
}

func dockerimagesHandler(w http.ResponseWriter, r *http.Request) {

	data := cmd_dockerimages()

	dataFormatted := strings.Split(data, "\n")
	dto := cellarDTO{
		Hostname:      cellarHostName,
		Version:       cellarVersion,
		ExceptionText: "",
		Data:          dataFormatted,
	}

	// logger.Information(data)

	dockerimagesTemplate.ExecuteTemplate(w, "layouttemplate", dto)
}

func dockerpsaHandler(w http.ResponseWriter, r *http.Request) {

	data := cmd_dockerpsa()

	dataFormatted := strings.Split(data, "\n")
	dto := cellarDTO{
		Hostname:      cellarHostName,
		Version:       cellarVersion,
		ExceptionText: "",
		Data:          dataFormatted,
	}

	// logger.Information(data)

	dockerpsaTemplate.ExecuteTemplate(w, "layouttemplate", dto)
}

func hubprocessesHandler(w http.ResponseWriter, r *http.Request) {

	data := cmd_filterProcesses("cellarhub")

	dataFormatted := strings.Split(data, "\n")
	dto := cellarDTO{
		Hostname:      cellarHostName,
		Version:       cellarVersion,
		ExceptionText: "",
		Data:          dataFormatted,
	}

	// logger.Information(data)

	cellarhubprocessesTemplate.ExecuteTemplate(w, "layouttemplate", dto)
}

func hubsystemdHandler(w http.ResponseWriter, r *http.Request) {

	data := cmd_hubmanagerstatus()

	dataFormatted := strings.Split(data, "\n")
	dto := cellarDTO{
		Hostname:      cellarHostName,
		Version:       cellarVersion,
		ExceptionText: "",
		Data:          dataFormatted,
	}

	// logger.Information(data)

	cellarhubsystemdTemplate.ExecuteTemplate(w, "layouttemplate", dto)
}

func cliHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		cliTemplate.ExecuteTemplate(w, "layouttemplate", nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		command := r.Form.Get("command")

		fmt.Println(command)

		c5, err := exec.Command("bash", "-c", command).Output()
		if err != nil {
			logger.Error(err.Error())
		}
		data := printOutput(c5)
		dataFormatted := strings.Split(data, "\n")

		dto := cellarDTO{
			Hostname:      cellarHostName,
			Version:       cellarVersion,
			ExceptionText: "",
			Data:          dataFormatted,
		}

		// logger.Information(data)

		cliTemplate.ExecuteTemplate(w, "layouttemplate", dto)
	}

}

func dockerStackHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		dockerStackTemplate.ExecuteTemplate(w, "layouttemplate", nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		command := r.Form.Get("command")

		result := ""

		if command == "Start" {
			cmd_dockerlogin()
			result = cmd_dockerstack_deploy()
		} else if command == "Check" {
			result = cmd_dockerstack_check()
		} else if command == "Stop" {
			result = cmd_dockerstack_stop()
		}

		dataFormatted := strings.Split(result, "\n")

		dto := cellarDTO{
			Data: dataFormatted,
		}

		// logger.Information(data)

		dockerStackTemplate.ExecuteTemplate(w, "layouttemplate", dto)
	}

}
