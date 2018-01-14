package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type cellarDTO struct {
	ExceptionText string      `json:"exceptionText"`
	Data          interface{} `json:"data"`
}

// func other1Page(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the Other1Page! :-) ")
// 	fmt.Println("Endpoint Hit: other1Page")
// }

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomeePage! :-) ")
	fmt.Println("Endpoint Hit: homeePage")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the AboutPage! :-) ")
	fmt.Println("Endpoint Hit: aboutPage")
}

func processesHandler(w http.ResponseWriter, r *http.Request) {

	// Create an *exec.Cmd
	cmd := exec.Command("ps", "-ef")

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	data := printOutput(output)

	dataFormatted := strings.Split(data, "\n")
	dto := cellarDTO{
		ExceptionText: "",
		Data:          dataFormatted,
	}

	logger.Information(data)

	processes.ExecuteTemplate(w, "layouttemplate", dto)
}

func processesNgrokHandler(w http.ResponseWriter, r *http.Request) {

	//1
	// c1 := exec.Command("ps", "-ef")
	// c2 := exec.Command("grep", "ngrok")

	// rr, ww := io.Pipe()
	// c1.Stdout = w
	// c2.Stdin = rr

	// var b2 bytes.Buffer
	// c2.Stdout = &b2

	// c1.Start()
	// c2.Start()
	// c1.Wait()
	// ww.Close()
	// c2.Wait()
	// io.Copy(os.Stdout, &b2)

	// data2 := b2.String()

	//2
	// Create an *exec.Cmd
	// cmd := exec.Command("ps", "-ef", "|", "grep", "ngrok")

	// // Combine stdout and stderr
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// data := printOutput(output)

	// dataFormatted := strings.Split(data, "\n")

	// 3

	cccmd := "ps -ef | grep ngrok"
	c5, err := exec.Command("bash", "-c", cccmd).Output()
	if err != nil {
		logger.Error(err.Error())
	}
	data := printOutput(c5)

	dto := cellarDTO{
		ExceptionText: "",
		Data:          data,
	}

	logger.Information(data)

	processes2.ExecuteTemplate(w, "layouttemplate", dto)
}

func actualdirectoryHandler(w http.ResponseWriter, r *http.Request) {

	var (
		cmdOut []byte
		err    error
	)
	cmd := "ls"
	args := []string{"-l"}
	cmdOut, err = exec.Command(cmd, args...).Output()
	if err != nil {
		fmt.Println("can't run command > " + err.Error())
	}
	cmdOutText := string(cmdOut)
	dataFormatted := strings.Split(cmdOutText, "\n")

	dto := cellarDTO{
		ExceptionText: "",
		Data:          dataFormatted,
	}

	logger.Information(cmdOutText)

	actualDirectory.ExecuteTemplate(w, "layouttemplate", dto)
}

// func updatePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the Update! :-) ")
// 	update("stable")
// }

// func restartPage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the Restart! :-) ")
// 	fmt.Println("PID : ", pid)

// 	s := strconv.Itoa(pid)
// 	cmd := exec.Command("kill", "-USR2", s)

// 	// Combine stdout and stderr
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("Restart method ends", output)
// }

// func other2Page(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the Other2Page! :-) ")
// 	fmt.Println("Endpoint Hit: other2Page")
// }

// func webHandler(name string) http.Handler {
// 	mux := http.NewServeMux()

// 	// mux.HandleFunc("/", homePage)
// 	mux.HandleFunc("/other1", other1Page)
// 	mux.HandleFunc("/home", homePage)
// 	mux.HandleFunc("/about", aboutPage)
// 	mux.HandleFunc("/update", updatePage)
// 	mux.HandleFunc("/restart", restartPage)
// 	mux.HandleFunc("/other2", other2Page)

// 	return mux
// }

// func webHandler2(name string) http.Handler {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/home", homePage)
// 	myRouter.HandleFunc("/about", aboutPage)
// 	myRouter.HandleFunc("/update", updatePage)
// 	myRouter.HandleFunc("/restart", restartPage)

// 	return myRouter
// }
