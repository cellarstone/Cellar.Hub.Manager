package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func apiTestHandler(w http.ResponseWriter, r *http.Request) {

	dto := cellarDTO{
		ExceptionText: "",
		Data:          "TEST OK",
	}

	json.NewEncoder(w).Encode(dto)
}

func apiAllProcessesHandler(w http.ResponseWriter, r *http.Request) {

	// Create an *exec.Cmd
	cmd := exec.Command("ps", "-e", "-o", "pid,time,%cpu,%mem,rss,cmd")

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error(err.Error())
	}
	data := printOutput(output)

	dataFormatted := strings.Split(data, "\n")
	dto := cellarDTO{
		ExceptionText: "",
		Data:          dataFormatted,
	}

	json.NewEncoder(w).Encode(dto)
}

func apiActualDirectoryHandler(w http.ResponseWriter, r *http.Request) {

	var (
		cmdOut []byte
		err    error
	)
	cmd := "ls"
	args := []string{"-l"}
	cmdOut, err = exec.Command(cmd, args...).Output()
	if err != nil {
		logger.Error("can't run command > " + err.Error())
	}
	cmdOutText := string(cmdOut)
	dataFormatted := strings.Split(cmdOutText, "\n")

	dto := cellarDTO{
		ExceptionText: "",
		Data:          dataFormatted,
	}

	json.NewEncoder(w).Encode(dto)
}

func testCheckProcessWorkflowHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid := vars["pid"]

	// 1. check if PID exists and running
	isExist := checkIfProcessRun(pid)

	//SEND RESPONSE
	dtoOut := cellarDTO{
		ExceptionText: "",
		Data:          isExist,
	}

	json.NewEncoder(w).Encode(dtoOut)

}

func killprocessHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idnumber, _ := strconv.Atoi(id)

	proc, _ := os.FindProcess(idnumber)
	err := proc.Kill()
	if err != nil {
		fmt.Println("process can't be killed > " + err.Error())
	}
}

func dockerImagesHandler(w http.ResponseWriter, r *http.Request) {

	// Create an *exec.Cmd
	cmd := exec.Command("sudo", "docker", "images")

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error(err.Error())
	}
	data := printOutput(output)

	dataFormatted := strings.Split(data, "\n")
	dto := cellarDTO{
		ExceptionText: "",
		Data:          dataFormatted,
	}

	json.NewEncoder(w).Encode(dto)
}

func dockerPsaHandler(w http.ResponseWriter, r *http.Request) {

	// Create an *exec.Cmd
	cmd := exec.Command("sudo", "docker", "ps", "-a")

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error(err.Error())
	}
	data := printOutput(output)

	dataFormatted := strings.Split(data, "\n")
	dto := cellarDTO{
		ExceptionText: "",
		Data:          dataFormatted,
	}

	json.NewEncoder(w).Encode(dto)
}

func apiRunNgrokHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	port := vars["port"]

	cmdName := "./ngrok/ngrok"

	//run
	cmd := exec.Command(cmdName, "http", port)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		logger.Error("can't run command > " + err.Error())
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			//low-level exception logging
			logger.Information("workflow process | " + scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		logger.Error("can't start command > " + err.Error())
	}

	asdf := cmd.Process.Pid

	//---------------------------------------------
	logger.Information(strconv.Itoa(asdf))

	//SEND RESPONSE
	dtoOut := cellarDTO{
		ExceptionText: "",
		Data:          asdf,
	}

	json.NewEncoder(w).Encode(dtoOut)
}

// func apiHandler(name string) http.Handler {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/api/allprocesses", allProcessesHandler)
// 	mux.HandleFunc("/api/dockerimages", dockerImagesHandler)
// 	mux.HandleFunc("/api/dockerpsa", dockerPsaHandler)

// 	return mux
// }

// func apiHandler2(name string) http.Handler {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/api/allprocesses", allProcessesHandler)
// 	myRouter.HandleFunc("/api/dockerimages", dockerImagesHandler)
// 	myRouter.HandleFunc("/api/dockerpsa", dockerPsaHandler)
// 	return myRouter
// }

//-------------------------------------
//-------------------------------------
//-------------------------------------
//-------------------------------------
// HELPERS
//-------------------------------------
//-------------------------------------
//-------------------------------------
//-------------------------------------

func checkIfProcessRun(pid string) bool {
	// Process info -------------------------
	// Create an *exec.Cmd
	cmd := exec.Command("ps", "-p", pid, "-o", "rss")

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error(err.Error())
	}
	data := printOutput(output)

	dataFormatted := strings.Split(data, "\n")

	//***********************************
	//***********************************
	// CONTROL
	//***********************************
	//***********************************

	if len(dataFormatted) <= 1 {
		return false
	}

	text := dataFormatted[1]
	processLine := strings.Replace(text, " ", "", -1)

	if processLine == "" {
		return false
	}

	if processLine == "0" {
		return false
	}

	//EVERYTHING SEEMS OK
	return true
}
