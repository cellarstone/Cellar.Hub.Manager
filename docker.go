package main

import (
	"fmt"
	"os/exec"
)

func cmd_dockerstack_deploy() string {
	cccmd := "docker stack deploy -c docker-stack.yml cellarhub --with-registry-auth"
	c5, err := exec.Command("bash", "-c", cccmd).Output()
	if err != nil {
		logger.Error(err.Error())
	}
	data := printOutput(c5)

	return data
}

func cmd_dockerimages() string {
	// Create an *exec.Cmd
	cmd := exec.Command("docker", "images")

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	data := printOutput(output)

	return data
}

func cmd_dockerpsa() string {
	// Create an *exec.Cmd
	cmd := exec.Command("docker", "ps", "-a")

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	data := printOutput(output)

	return data
}
