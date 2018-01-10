package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/equinox-io/equinox"
	"github.com/gorilla/mux"
)

//************************************************************
//************************************************************
//************************************************************
//************************************************************
// EQUINOX

const appID = "app_h9SyPnPqLpq"

// KEY MUST BE FORMATED EXACTLY AS IS
// NO WHITESPACE ON BEGIN OF LINES ... etc.
var publicKey = []byte(`
-----BEGIN ECDSA PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE5sQO5CKy1teb4m/AFrZ5e4RDKsA613YL
oklmhuQ8MWisY3cQNpNHFstFc1DjDu29/vQYo2ckurYpf7OOjAStPL4qb+3WSFOR
gfj0W1ovPzXas/+elnyuZumyZ1KMJWgL
-----END ECDSA PUBLIC KEY-----
`)

func update(channel string) string {

	fmt.Println("START UPDATING")

	opts := equinox.Options{Channel: channel}
	if err := opts.SetPublicKeyPEM(publicKey); err != nil {
		fmt.Println(err)
		return err.Error()
	}

	fmt.Println("check for the update")

	// check for the update
	resp, err := equinox.Check(appID, opts)
	switch {
	case err == equinox.NotAvailableErr:
		fmt.Println("No update available, already at the latest version!")
		return "NO_UPDATES"
	case err != nil:
		fmt.Println(err)
		return err.Error()
	}

	// fetch the update and apply it
	err = resp.Apply()
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	fmt.Printf("Updated to new version: %s!\n", resp.ReleaseVersion)
	return "OK"
}

//************************************************************
//************************************************************
//************************************************************
//************************************************************

func restartGrace() {
	//RESTART
	s := strconv.Itoa(pid)
	cmd := exec.Command("kill", "-USR2", s)

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
func restart() {
	syscall.Kill(-pid, syscall.SIGKILL)
}
func restart2() {
	proc, _ := os.FindProcess(pid)
	err := proc.Kill()
	if err != nil {
		logger.Error("process can't be killed > " + err.Error())
	}
}

func check() {

	result := update("stable")

	if result == "OK" {

		//RESTART
		restart2()

		fmt.Println("OK - EVERYTHING WAS UPDATED")

	} else if result == "NO_UPDATES" {

		fmt.Println("OK - EVERYTHING UP TO DATE")

	} else {

		fmt.Println("STRANGE")

	}

}

func startChecking() {
	for {
		time.Sleep(1 * time.Minute)
		check()
	}
}

//************************************************************
//************************************************************

var pid = 0

func myHandler(name string) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", homePage)
	mux.HandleFunc("/about", aboutPage)
	mux.HandleFunc("/allprocesses", processesHandler)
	mux.HandleFunc("/ngrokprocesses", processesNgrokHandler)
	mux.HandleFunc("/actualdirectory", actualdirectoryHandler)
	mux.HandleFunc("/api/allprocesses", apiAllProcessesHandler)
	mux.HandleFunc("/api/dockerimages", dockerImagesHandler)
	mux.HandleFunc("/api/dockerpsa", dockerPsaHandler)
	// mux.HandleFunc("/runngrok", apiRunNgrokHandler)

	return mux
}

func myRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/api/test", http.HandlerFunc(apiTestHandler))
	r.Handle("/api/allprocesses", http.HandlerFunc(apiAllProcessesHandler))
	r.Handle("/api/actualdirectory", http.HandlerFunc(apiActualDirectoryHandler))
	r.Handle("/api/checkprocess/{pid}", http.HandlerFunc(testCheckProcessWorkflowHandler))
	r.Handle("/api/killprocess/{id}", http.HandlerFunc(killprocessHandler))
	r.Handle("/api/dockerimages", http.HandlerFunc(dockerImagesHandler))
	r.Handle("/api/dockerpsa", http.HandlerFunc(dockerPsaHandler))
	r.Handle("/api/runngrok/{port}", http.HandlerFunc(apiRunNgrokHandler))
	return r
}

//************************************************************
//************************************************************

var (
	address0 = flag.String("a00", ":10001", "Web1 address to bind to.")
	address1 = flag.String("a11", ":10002", "Web2 address to bind to.")
	address2 = flag.String("a22", ":10003", "Web3 address to bind to.")
	address3 = flag.String("a33", ":10004", "Web4 address to bind to.")
)

var layoutDir = "views/layout"
var processes *template.Template
var processes2 *template.Template
var actualDirectory *template.Template

//Logging
var logger *DLogger
var err error

func init() {
	//set logging
	logger, err = NewDLogger("Cellar.Hub.Manager")
	if err != nil {
		panic(err)
	}

	//run ngrok

}

func main() {
	defer startChecking()
	defer runNgrok()

	logger.Information("Cellarstone manager v0.3.3")
	pid = os.Getpid()
	pidString := strconv.Itoa(pid)
	logger.Information("PID : " + pidString)

	// files := append(layoutFiles(), "views/processes.gohtml")
	// processes, err = template.ParseFiles(files...)
	// if err != nil {
	// 	//low-level exception logging
	// 	fmt.Println(err)
	// }
	// files = append(layoutFiles(), "views/processes2.gohtml")
	// processes2, err = template.ParseFiles(files...)
	// if err != nil {
	// 	//low-level exception logging
	// 	fmt.Println(err)
	// }
	// files = append(layoutFiles(), "views/actualdirectory.gohtml")
	// actualDirectory, err = template.ParseFiles(files...)
	// if err != nil {
	// 	//low-level exception logging
	// 	fmt.Println(err)
	// }

	go startChecking()
	//go runNgrok()

	// flag.Parse()
	// gracehttp.Serve(
	// 	&http.Server{Addr: *address0, Handler: myHandler("Web11")},
	// 	&http.Server{Addr: *address1, Handler: myHandler("Web22")},
	// 	&http.Server{Addr: *address2, Handler: myHandler("Web33")},
	// 	&http.Server{Addr: *address3, Handler: myHandler("Web44")},
	// )

	r := myRouter()
	http.ListenAndServe(":10001", r)
}

//-------------------------------------
//NGROK
//-------------------------------------
func runNgrok() {
	cccmd := "./ngrok/ngrok http 10001"
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

//-------------------------------------
//HELPERS
//-------------------------------------

func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "/*.gohtml")
	if err != nil {
		//low-level exception logging
		logger.Error(err.Error())
	}
	return files
}

func printOutput(outs []byte) string {
	result := ""
	if len(outs) > 0 {
		result += string(outs)
	}
	return result
}
