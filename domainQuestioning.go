package main

import (
	"log"
)

type DomainQuest struct {
	domains []string
	count   int
}

func (dm DomainQuest) domainQueryAndWrite() []string {
	var returnVal = []string{}
	//
	//var val = []string{}
	//iterm := terminal{Whois, "gitarist.com"}
	//iterm := Terminal{}
	whServer := whoisServer{}

	for _, domain := range dm.domains {

		//var marks = map[bool]string{true: "◯ ", false: "☓"}
		exist, err := whServer.exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Println(marks[exist])
		//time.Sleep(1 * time.Second)
		if len(domain) <= dm.count {

			if exist == 1 {
				//iterm.Ping(domain)
				DomainWrite := Write{OK}
				DomainWrite.spaceWrite(domain)
				returnVal = append(returnVal, domain)
				CreateAndWriteDateFile(returnVal)

			} else if exist == 2 {
				//gopre.Pre(iterm.errorPool)
				DomainWrite := Write{TAKEN}
				DomainWrite.spaceWrite(domain)
				//domain dolu
			} else {
				DomainWrite := Write{UNKNOWN}
				DomainWrite.spaceWrite(domain)
				// cok uzun domain olunca yazdiramiyo kontrol et
			}

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
	}

	return returnVal
}
