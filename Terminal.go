package main

import (
	"os/exec"
	"fmt"
	"github.com/ugurethemaydin/gopre/src/gopre"
	"bytes"
	"strings"
)

type Program string

const (
	Whois Program = "whois"
	Ping  Program = "ping"
	Tolga Program = "cd"
)

type Terminal struct {
	program   Program
	arg       string
	errorPool error
}

func (terminal Terminal) debugByte(byteText string) {
	fmt.Printf("buf = %s\n", string(byteText))
}

func (terminal Terminal) getErrors() error {
	return terminal.errorPool
}

func (terminal Terminal) Whois(arg string) bool {
	terminal.arg = arg
	terminal.program = Whois
	returnToTerminal := terminal.runTerminal(terminal.program, terminal.arg)
	if cap(returnToTerminal) == 0 {

		gopre.Pre(terminal.errorPool)
	}
	argument := strings.Split(terminal.arg, ".")
	if len(argument) == 1 {
		terminal.errorPool = fmt.Errorf(".com gibi yazim yok.")
		return false
	}
	var currentExtension = argument[1]
	ExtensionSearchWords := DomainExtension(currentExtension).getSearhText()
	//b := ExtensionResultTextsTest[currentExtension]

	var matchText = ExtensionSearchWords
	if bytes.Contains(returnToTerminal, []byte(matchText)) {

		return true
	} else {
		//gopre.Pre("dolu doncek zaten ya")
		return false
	}

}
func (terminal Terminal) runTerminal(program Program, arg string) []byte {
	head := string(program)

	cmdReturn, err := exec.Command(head, arg).Output()
	//gopre.Pre(cmdReturn)
	if err != nil {
		terminal.errorPool = err
		fmt.Printf("Exec.Command Error : %s", err)
		return []byte{}
	}
	return cmdReturn
}

//func (terminal Terminal) Ping(arg string) {
//
//	returnToTerminal := terminal.runTerminal(Ping, arg)
//	//time.Sleep(5 * time.Second)
//	gopre.Pre(string(returnToTerminal))
//	//gopre.Pre(cap(returnToTerminal))
//	//if cap(returnToTerminal) == 0 {
//	//	gopre.Pre(terminal.errorPool)
//	//}
//
//}
