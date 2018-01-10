package main

import "fmt"

type DLogger struct {
	tag string
}

func NewDLogger(tag string) (logger *DLogger, err error) {

	result := DLogger{}
	result.tag = tag

	return &result, nil
}

func (t *DLogger) Debug(message string) error {
	return t.log("Debug", message)
}

func (t *DLogger) Information(message string) error {
	return t.log("Information", message)
}

func (t *DLogger) Warning(message string) error {
	return t.log("Warning", message)
}

func (t *DLogger) Error(message string) error {
	return t.log("Error", message)
}

func (t *DLogger) Fatal(message string) error {
	return t.log("Fatal", message)
}

func (t *DLogger) log(level string, message string) error {

	fmt.Println("[", level, "]", "[", t.tag, "]", " - ", message)

	return nil
}
