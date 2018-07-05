package main

import (
	"log"
	"strings"
	"fmt"
)

type DomainQuest struct {
	domains   []string
	count     int
	extention string
}

func (dm DomainQuest) domainQueryAndWrite() {

	io := IoController{}
	io.setFolderPath(outputFolderPath + "/")
	io.openLogFile(true)

	for _, domain := range dm.domains {
		numberOfRow ++
		WriteNumberFile(numberOfRow - 1)

		splitDomain := strings.Split(domain, ".")
		//dm.extention == "" ||
		if dm.extention == "" || dm.extention == "all" {
			dm.findOkDomain(domain, splitDomain, io)
			//else if exist == 5 {
		} else if dm.extention == splitDomain[1] {
			dm.findOkDomain(domain, splitDomain, io)
		}

	}

}

func (dm DomainQuest) findOkDomain(domain string, splitDomain []string, io IoController) {

	whServer := whoisServer{}

	if len(splitDomain[0]) <= dm.count {

		exist, err := whServer.exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		switch exist {
		case 1:
			DomainWrite := Write{OK}
			DomainWrite.spaceWrite(domain)
			io.log(domain, OK)

		case 2:
			DomainWrite := Write{TAKEN}
			DomainWrite.spaceWrite(domain)

		default:

			DomainWrite := Write{UNKNOWN}
			DomainWrite.spaceWrite(domain)
		}

	}

}

func WriteNumberFile(val int) {
	io := IoController{}

	io.fileName = "/errList"
	io.folderName = "errlist"
	io.openLogFile(false)
	io.log(fmt.Sprintf("%d\n", val), TAKEN)

}

func WriteStringFile(val string) {
	io := IoController{}

	io.fileName = "/errList"
	io.folderName = "errlist"
	io.openLogFile(false)
	io.log(val+"\n", TAKEN)

}

//if len(domain) <= dm.count {
//	if iterm.Whois(domain) {
//		//if len(domain) <= 20{
//		//iterm.Ping(domain)
//		//}
//
//		DomainWrite := Write{OK}
//		DomainWrite.spaceWrite(domain)
//		returnVal = append(returnVal, domain)
//		CreateAndWriteDateFile(returnVal)
//		//domain bos
//	} else {
//		//gopre.Pre(iterm.errorPool)
//		DomainWrite := Write{TAKEN}
//		DomainWrite.spaceWrite(domain)
//		//domain dolu
//	}
//}
//if exist == 1 {
//	//iterm.Ping(domain)
//	DomainWrite := Write{OK}
//	DomainWrite.spaceWrite(domain)
//
//	io.log(domain, OK)
//OkDomains = append(OkDomains, domain)
//io.CreateAndWriteDateFile(domain)
//deferla dosyayi kapatman lazim

//} else if exist == 2 {
////gopre.Pre(iterm.errorPool)
//DomainWrite := Write{TAKEN}
//DomainWrite.spaceWrite(domain)

//domain dolu
//} else if exist == 3 {
//

//io := IoController{}
//io.CreateFile("yes", outputName, "unregisteredExtension.txt", []string{"unregistered extensions"})
//unregisteredExtensions = append(unregisteredExtensions,domain)
//io.WriteFile(outputFolderPath+"/"+"unregisteredExtension.txt",unregisteredExtensions)
// cok uzun domain olunca yazdiramiyo ama unregistered extension: yaziyo onceden ariyo cunku duzelt hatayi
//var OkDomains = []string{}
//var unregisteredExtensions = []string{}
//
//var val = []string{}
//iterm := terminal{Whois, "gitarist.com"}
//iterm := Terminal{}
