package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/arschles/go-bindata-html-template"
)

//go:generate go-bindata views/...

//************************************************************
//************************************************************

var pid = 0

func myHandler(name string) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/allprocesses", processesHandler)
	mux.HandleFunc("/ngrokprocesses", processesNgrokHandler)
	mux.HandleFunc("/actualdirectory", actualdirectoryHandler)
	mux.HandleFunc("/dockerimages", dockerimagesHandler)
	mux.HandleFunc("/dockerpsa", dockerpsaHandler)

	mux.HandleFunc("/api/allprocesses", apiAllProcessesHandler)
	// mux.HandleFunc("/api/dockerimages", dockerImagesHandler)
	// mux.HandleFunc("/api/dockerpsa", dockerPsaHandler)
	// mux.HandleFunc("/runngrok", apiRunNgrokHandler)

	return mux
}

func myRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/allprocesses", http.HandlerFunc(processesHandler))
	r.Handle("/ngrokprocesses", http.HandlerFunc(processesNgrokHandler))
	r.Handle("/actualdirectory", http.HandlerFunc(actualdirectoryHandler))
	r.Handle("/dockerimages", http.HandlerFunc(dockerimagesHandler))
	r.Handle("/dockerpsa", http.HandlerFunc(dockerpsaHandler))
	// r.Handle("/api/test", http.HandlerFunc(apiTestHandler))
	// r.Handle("/api/allprocesses", http.HandlerFunc(apiAllProcessesHandler))
	// r.Handle("/api/actualdirectory", http.HandlerFunc(apiActualDirectoryHandler))
	// r.Handle("/api/checkprocess/{pid}", http.HandlerFunc(apiTestCheckProcessWorkflowHandler))
	// r.Handle("/api/killprocess/{id}", http.HandlerFunc(apiKillprocessHandler))
	// r.Handle("/api/dockerimages", http.HandlerFunc(apiDockerImagesHandler))
	// r.Handle("/api/dockerpsa", http.HandlerFunc(apiDockerPsaHandler))
	// r.Handle("/api/runngrok/{port}", http.HandlerFunc(apiRunNgrokHandler))
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
var processesTemplate *template.Template
var ngrokprocessesTemplate *template.Template
var actualDirectoryTemplate *template.Template
var indexTemplate *template.Template
var dockerimagesTemplate *template.Template
var dockerpsaTemplate *template.Template

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
	logger.Information("Cellarstone manager v0.3.23")

	checkCellarDeviceInfo()
	//killAllNgrokProcesses()

	defer startChecking()
	defer killAllNgrokProcesses()

	connectToNgrok()
	authorizeNgrok()
	go runNgrok("http", "10001")
	go runNgrok("tcp", "22")

	go startChecking()

	// NORMAL HTTP TEMPLATES
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

	// GO-BINDATA-TEMPLATES
	files := append(layoutFiles(), "views/index.gohtml")
	indexTemplate, err = template.New("index", Asset).ParseFiles(files...)
	if err != nil {
		fmt.Printf("error parsing template: %s", err)
	}

	files = append(layoutFiles(), "views/processes.gohtml")
	processesTemplate, err = template.New("processes", Asset).ParseFiles(files...)
	if err != nil {
		fmt.Printf("error parsing template: %s", err)
	}

	files = append(layoutFiles(), "views/ngrokprocesses.gohtml")
	ngrokprocessesTemplate, err = template.New("processes2", Asset).ParseFiles(files...)
	if err != nil {
		fmt.Printf("error parsing template: %s", err)
	}

	files = append(layoutFiles(), "views/actualdirectory.gohtml")
	actualDirectoryTemplate, err = template.New("actualDirectory", Asset).ParseFiles(files...)
	if err != nil {
		fmt.Printf("error parsing template: %s", err)
	}

	files = append(layoutFiles(), "views/dockerimages.gohtml")
	dockerimagesTemplate, err = template.New("dockerimages", Asset).ParseFiles(files...)
	if err != nil {
		fmt.Printf("error parsing template: %s", err)
	}

	files = append(layoutFiles(), "views/dockerpsa.gohtml")
	dockerpsaTemplate, err = template.New("dockerpsa", Asset).ParseFiles(files...)
	if err != nil {
		fmt.Printf("error parsing template: %s", err)
	}

	// FACEBOOK GO GRACE
	// flag.Parse()
	// gracehttp.Serve(
	// 	&http.Server{Addr: *address0, Handler: myHandler("Web11")},
	// 	&http.Server{Addr: *address1, Handler: myHandler("Web22")},
	// 	&http.Server{Addr: *address2, Handler: myHandler("Web33")},
	// )

	// NORMAL ROUTER
	r := myRouter()
	http.ListenAndServe(":10001", r)
}

//-------------------------------------
//HELPERS
//-------------------------------------

func layoutFiles() []string {
	// files, err := filepath.Glob(layoutDir + "/*.gohtml")
	// if err != nil {
	// 	//low-level exception logging
	// 	logger.Error(err.Error())
	// }

	files := []string{"views/layout/layout.gohtml"}
	return files
}

func printOutput(outs []byte) string {
	result := ""
	if len(outs) > 0 {
		result += string(outs)
	}
	return result
}
