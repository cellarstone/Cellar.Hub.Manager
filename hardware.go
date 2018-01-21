package main

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/jaypipes/ghw"
)

// getMacAddr gets the MAC hardware
// address of the host machine
func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

func getMachineID() (id string) {
	cccmd := "cat /etc/machine-id"

	c5 := exec.Command("bash", "-c", cccmd)

	c6, err := c5.Output()
	if err != nil {
		logger.Error("read machine ID error")
		logger.Error(err.Error())
	}
	data := printOutput(c6)

	fmt.Println(data)

	return strings.TrimSpace(data)
}

func getCpuInfo() string {
	result := ""

	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	fmt.Printf("%v\n", cpu)
	result += fmt.Sprintf("%v\n", cpu)

	for _, proc := range cpu.Processors {
		fmt.Printf(" %v\n", proc)
		for _, core := range proc.Cores {
			fmt.Printf("  %v\n", core)
			result += fmt.Sprintf("%v\n", core)
		}
	}

	return result
}

func getHDDInfo() string {
	result := ""

	block, err := ghw.Block()
	if err != nil {
		fmt.Printf("Error getting block storage info: %v", err)
	}

	fmt.Printf("%v\n", block)
	result += fmt.Sprintf("%v\n", block)

	for _, disk := range block.Disks {
		fmt.Printf(" %v\n", disk)
		result += fmt.Sprintf("%v\n", disk)
		for _, part := range disk.Partitions {
			fmt.Printf("  %v\n", part)
			result += fmt.Sprintf("%v\n", part)
		}
	}

	return result
}

func getNetworkInfo() string {
	result := ""

	net, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting network info: %v", err)
	}

	fmt.Printf("%v\n", net)
	result += fmt.Sprintf("%v\n", net)

	for _, nic := range net.NICs {
		fmt.Printf(" %v\n", nic)
		result += fmt.Sprintf("%v\n", nic)
	}

	return result
}
