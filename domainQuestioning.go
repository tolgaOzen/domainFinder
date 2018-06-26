package main

import (
	"log"
)

type DomainQuest struct {
	domains []string
	count   int
}

func (dm DomainQuest) domainQueryAndWrite()  {
	//var OkDomains = []string{}
	//var unregisteredExtensions = []string{}
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
		io := IoController{}
		if len(domain) <= dm.count {

			if exist == 1 {
				//iterm.Ping(domain)
				DomainWrite := Write{OK}
				DomainWrite.spaceWrite(domain)
				//OkDomains = append(OkDomains, domain)
				io.CreateAndWriteDateFile(domain)
				//deferla dosyayi kapatman lazim

			} else if exist == 2 {
				//gopre.Pre(iterm.errorPool)
				DomainWrite := Write{TAKEN}
				DomainWrite.spaceWrite(domain)
				//domain dolu
			} else {
				DomainWrite := Write{UNKNOWN}
				DomainWrite.spaceWrite(domain)

				//io := IoController{}
				//io.CreateFile("yes", outputName, "unregisteredExtension.txt", []string{"unregistered extensions"})
				//unregisteredExtensions = append(unregisteredExtensions,domain)
				//io.WriteFile(outputFolderPath+"/"+"unregisteredExtension.txt",unregisteredExtensions)
				// cok uzun domain olunca yazdiramiyo ama unregistered extension: yaziyo onceden ariyo cunku duzelt hatayi
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


}
