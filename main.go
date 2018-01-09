package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/equinox-io/equinox"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gorilla/mux"
)

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

// type Article struct {
// 	Id      int    `json:"Id"`
// 	Title   string `json:"Title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// type Articles []Article

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	articles := Articles{
// 		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
// 	}
// 	fmt.Println("Endpoint Hit: returnAllArticles")

// 	json.NewEncoder(w).Encode(articles)
// }

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	fmt.Fprintf(w, "Key: "+key)
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// func aboutPage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the AboutPage!")
// 	fmt.Println("Endpoint Hit: aboutPage")
// }

// func updatePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the Update!")
// 	update("stable")
// }

// //***********************************************
// // GORILLA MUX package
// //***********************************************

// func handleRequests() {

// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/", homePage)
// 	myRouter.HandleFunc("/home", homePage)
// 	myRouter.HandleFunc("/about", aboutPage)
// 	myRouter.HandleFunc("/update", updatePage)
// 	myRouter.HandleFunc("/all", returnAllArticles)
// 	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
// 	log.Fatal(http.ListenAndServe(":10000", myRouter))
// }

// func main() {
// 	fmt.Println("Rest API v2.3 - Mux Routers")
// 	handleRequests()
// }

//************************************************************
//************************************************************

//************************************************************
//************************************************************

func check() {

	result := update("stable")

	if result == "OK" {

		//RESTART
		s := strconv.Itoa(pid)
		cmd := exec.Command("kill", "-USR2", s)

		// Combine stdout and stderr
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)

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
		go check()
	}
}

//************************************************************
//************************************************************

var pid = 0

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage! :-) ")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the AboutPage! :-) ")
	fmt.Println("Endpoint Hit: aboutPage")
}

func updatePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Update! :-) ")
	update("stable")
}

func restartPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Restart! :-) ")
	fmt.Println("PID : ", pid)

	s := strconv.Itoa(pid)
	cmd := exec.Command("kill", "-USR2", s)

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Restart method ends", output)
}

//************************************************************
//************************************************************

var (
	address0 = flag.String("a0", ":10001", "Zero address to bind to.")
	address1 = flag.String("a1", ":10002", "First address to bind to.")
	address2 = flag.String("a2", ":10003", "Second address to bind to.")
	now      = time.Now()
)

func main() {
	fmt.Println("Cellarstone manager v0.2.4")

	pid = os.Getpid()
	fmt.Println("PID : ", pid)

	go startChecking()

	flag.Parse()
	gracehttp.Serve(
		&http.Server{Addr: *address0, Handler: newHandler("Zero  ")},
		&http.Server{Addr: *address1, Handler: newHandler("First ")},
		&http.Server{Addr: *address2, Handler: newHandler("Second")},
	)
}

func newHandler(name string) http.Handler {
	mux := http.NewServeMux()
	// mux.HandleFunc("/sleep/", func(w http.ResponseWriter, r *http.Request) {
	// 	duration, err := time.ParseDuration(r.FormValue("duration"))
	// 	if err != nil {
	// 		http.Error(w, err.Error(), 400)
	// 		return
	// 	}
	// 	time.Sleep(duration)
	// 	fmt.Fprintf(
	// 		w,
	// 		"%s started at %s slept for %d nanoseconds from pid %d.\n",
	// 		name,
	// 		now,
	// 		duration.Nanoseconds(),
	// 		os.Getpid(),
	// 	)
	// })

	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/home", homePage)
	mux.HandleFunc("/about", aboutPage)
	mux.HandleFunc("/update", updatePage)
	mux.HandleFunc("/restart", restartPage)
	mux.HandleFunc("/all", returnAllArticles)
	mux.HandleFunc("/article/{id}", returnSingleArticle)

	return mux
}
