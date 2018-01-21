package main

import (
	"fmt"
	"os/exec"
)

func cmd_hubmanagerstatus() string {
	// Create an *exec.Cmd
	cmd := exec.Command("service", "cellarhubmanager", "status")

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	data := printOutput(output)

	return data
}
