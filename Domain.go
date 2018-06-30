package main

type DomainExtension string

const (
	Com  DomainExtension = "com"
	Net  DomainExtension = "net"
	IO   DomainExtension = "io"
	Org  DomainExtension = "org"
	Biz  DomainExtension = "biz"
	Info DomainExtension = "info"
	Pro  DomainExtension = "pro"
	Name DomainExtension = "name"
	Edu  DomainExtension = "edu"
	Gov  DomainExtension = "gov"
	Us   DomainExtension = "us"
	In   DomainExtension = "in"
	//Xyz DomainExtension = "xyz"
)

type Domain struct {
	extension        DomainExtension
	searchResultWord string
}

func (m DomainExtension) getSearhText() string {

	return ExtensionResultTexts[m].searchResultWord
}

var ExtensionResultTexts = map[DomainExtension]Domain{
	Com:  Domain{Com, "no match",},
	Net:  Domain{Net, "No match for domain"},
	IO:   Domain{IO, "NOT FOUND"},
	Org:  Domain{Org, "NOT FOUND"},
	Biz:  Domain{Biz, "No Data Found"},
	Info: Domain{Info, "NOT FOUND"},
	Pro:  Domain{Pro, "NOT FOUND"},
	Name: Domain{Name, "No match for"},
	Edu:  Domain{Edu, "NO MATCH"},
	Gov:  Domain{Gov, "No match for"},
	Us:   Domain{Us, "No Data Found"},
	In:   Domain{In, "NOT FOUND"},
	//Xyz: Domain{Xyz,"DOMAIN NOT FOUND"},
}

//var ExtensionResultTextsTest = map[string]string{
//	"com": "No match for domain",
//	"net":  "No match for domain",
//	"io": "NOT FOUND",
//}
