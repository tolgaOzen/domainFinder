package main

import (
	"net"
	"strings"
	"fmt"
	"github.com/ugurethemaydin/gopre/src/gopre"
	"bufio"
)

type whoisServer struct {
	arg       string
	errorPool error
}

//var stayDomain string
//type returnStatus struct {
//}

func (WhoisServer whoisServer) exists(domain string) (int, error) {
	WhoisServer.arg = domain

	argument := strings.Split(WhoisServer.arg, ".")
	if len(argument) == 1 {
		WhoisServer.errorPool = fmt.Errorf(".com gibi yazim yok.")
		return 3, nil
	}
	var currentExtension = argument[1]

	if CurrentControl(currentExtension) == true {

		ExtensionSearchWords := DomainExtension(currentExtension).getSearhText()

		if ExtensionSearchWords == "" {

			//fmt.Print("unregistered extension: ", argument[1]+" ")
			//ExtensionSearchWords = "NOT"
			return 3, nil

		}

		jsonOpe := JsonOperation{}

		//jsonOpe.setDownloadedJsonFileURL("http://api.open-notify.org/astros.json")

		jsonOpe.setDownloadedJsonFileURL("https://raw.githubusercontent.com/weppos/whois/master/data/tld.json")
		returnByte := jsonOpe.getFileToInternet()

		//if StayDomain == false{
		//
		//	var stayDomain = WhoisServer.arg
		//
		//	return 5 , nil
		//}

		jsonObj := jsonOpe.getJson(returnByte, demoStuct{})
		//_ = jsonOpe.getJsonStruct()

		a := jsonObj[currentExtension].(map[string]interface{})
		//gopre.Pre(a["host"])

		var whoisServer string = a["host"].(string)
		//a["host"].(string)
		//currentExtension + ".whois-servers.net"

		conn, err := net.Dial("tcp", whoisServer+":43")
		if err != nil {
			gopre.Pre(WhoisServer.errorPool)
			return 2, err
		}

		defer conn.Close()
		conn.Write([]byte (WhoisServer.arg + "\r\n"))

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {

			//fmt.Println(scanner.Text())
			if strings.Contains(strings.ToLower(scanner.Text()), ExtensionSearchWords) {
				return 1, nil
			}
			return 2, nil
		}
		return 2, nil
	}
	return 4, nil
}

func CurrentControl(Current string) bool {
	var UnwantedCurrentExStrings = []string{"0", "9", "8", "7", "6", "5", "4", "3", "2", "1"}

	for _, unwantedLetter := range UnwantedCurrentExStrings {

		if !strings.ContainsAny(Current, unwantedLetter) {

			return true
		}
		return false

	}

	return false
}
