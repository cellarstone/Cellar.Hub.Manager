package main

import "os/exec"

func cmd_filterProcesses(programname string) string {
	cccmd := "ps -ef | grep " + programname
	c5, err := exec.Command("bash", "-c", cccmd).Output()
	if err != nil {
		logger.Error(err.Error())
	}
	data := printOutput(c5)

	return data
}

func cmd_run(command string) string {
	c5, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		logger.Error(err.Error())
	}
	data := printOutput(c5)

	return data
}
