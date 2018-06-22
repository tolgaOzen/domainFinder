package main

import (
	"net"
	"bufio"
	"strings"
	"fmt"
	"github.com/ugurethemaydin/gopre/src/gopre"
)

type whoisServer struct {
	arg       string
	errorPool error
}

func (WhoisServer whoisServer) exists(domain string) (bool, error) {
	WhoisServer.arg = domain

	argument := strings.Split(WhoisServer.arg, ".")
	if len(argument) == 1 {
		WhoisServer.errorPool = fmt.Errorf(".com gibi yazim yok.")
		return false, nil
	}
	var currentExtension = argument[1]
	ExtensionSearchWords := DomainExtension(currentExtension).getSearhText()

	var whoisServer string = currentExtension + ".whois-servers.net"
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		gopre.Pre(WhoisServer.errorPool)
		return false, err
	}

	defer conn.Close()
	conn.Write([]byte(WhoisServer.arg + "\r\n"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), ExtensionSearchWords) {
			return true, nil
		}
	}
	return false, nil
}
