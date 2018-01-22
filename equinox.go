package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/equinox-io/equinox"
)

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

func checkEquinox() {

	result := equinoxUpdate()

	if result == "OK" {

		//kill all ngrok
		//killAllNgrokProcesses()

		//RESTART
		//restartBrute()

		fmt.Println("OK - EVERYTHING WAS UPDATED")

	} else if result == "NO_UPDATES" {

		fmt.Println("OK - EVERYTHING UP TO DATE")

	} else {

		fmt.Println(result)

	}

}

//************************************************************
//************************************************************
//************************************************************
//************************************************************
// EQUINOX
func equinoxUpdate() string {
	var opts equinox.Options
	if err := opts.SetPublicKeyPEM(publicKey); err != nil {
		fmt.Println(err)
		return err.Error()
	}

	// check for the update
	resp, err := equinox.Check(appID, opts)
	switch {
	case err == equinox.NotAvailableErr:
		fmt.Println("No update available, already at the latest version!")
		return "NO UPDATES"
	case err != nil:
		fmt.Println("Update failed:", err)
		return err.Error()
	}

	// fetch the update and apply it
	err = resp.Apply()
	if err != nil {
		return err.Error()
	}

	fmt.Printf("Updated to new version: %s!\n", resp.ReleaseVersion)
	return "OK"
}

// func update(channel string) string {

// 	fmt.Println("START UPDATING")

// 	opts := equinox.Options{Channel: channel}
// 	if err := opts.SetPublicKeyPEM(publicKey); err != nil {
// 		fmt.Println(err)
// 		return err.Error()
// 	}

// 	fmt.Println("check for the update")

// 	// check for the update
// 	resp, err := equinox.Check(appID, opts)
// 	switch {
// 	case err == equinox.NotAvailableErr:
// 		fmt.Println("No update available, already at the latest version!")
// 		return "NO_UPDATES"
// 	case err != nil:
// 		fmt.Println(err)
// 		return err.Error()
// 	}

// 	// fetch the update and apply it
// 	err = resp.Apply()
// 	if err != nil {
// 		fmt.Println(err)
// 		return err.Error()
// 	}

// 	fmt.Printf("Updated to new version: %s!\n", resp.ReleaseVersion)
// 	return "OK"
// }

//************************************************************
//************************************************************
//************************************************************
//************************************************************

func restartGrace() {

	text := "PID : " + strconv.Itoa(cellarHubManagerPID) + " -------------------------------------------------------- "
	fmt.Println(text)

	//RESTART
	s := strconv.Itoa(cellarHubManagerPID)
	cmd := exec.Command("kill", "-USR2", s)

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
func restart() {
	syscall.Kill(-cellarHubManagerPID, syscall.SIGKILL)
}
func restartBrute() {
	proc, _ := os.FindProcess(cellarHubManagerPID)
	err := proc.Kill()
	if err != nil {
		logger.Error("process can't be killed > " + err.Error())
	}
}
