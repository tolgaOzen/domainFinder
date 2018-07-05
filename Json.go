package main

//

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/ugurethemaydin/gopre/src/gopre"
	"bytes"
	"encoding/gob"
)

type demoStuct struct {
	//Number  int    `json:"number"`
	Message string `json:"host"`
}
type JsonOperation struct {
	downloadedJsonFileURL string
	convertedJsonFile     interface{}
}

//func yazdir(t interface{})  string  {
//	//val := reflect.TypeOf(t)
//	return .Message
//}

//
//func inito() {	//jsonOpe := JsonOperation{}
//
////jsonOpe.setDownloadedJsonFileURL("http://api.open-notify.org/astros.json")
//jsonOpe.setDownloadedJsonFileURL("https://raw.githubusercontent.com/weppos/whois/master/data/tld.json")
//returnByte := jsonOpe.getFileToInternet()
//
//jsonObj := jsonOpe.getJson(returnByte, demoStuct{})
//_ = jsonOpe.getJsonStruct()
//jsonObj["host"] = string("")
////yazdir(jsonObj["com"])
//gopre.Pre(jsonObj["com"])
//for a, x := range  {
//	fmt.Println(a, x)
//}
//gopre.Pre(jsonObj["people"])

//spaceClient := http.Client{
//	Timeout: time.Second * 2, // Maximum of 2 secs
//}
//
//req, err := http.NewRequest(http.MethodGet, url, nil)
//if err != nil {
//	log.Fatal(err)
//}
//
//req.Header.Set("User-Agent", "spacecount-tutorial")
//
//res, getErr := spaceClient.Do(req)
//if getErr != nil {
//	log.Fatal(getErr)
//}
//
//body, readErr := ioutil.ReadAll(res.Body)
//if readErr != nil {
//	log.Fatal(readErr)
//}

//people1 := people{}
//jsonErr := json.Unmarshal(body, &people1)
//if jsonErr != nil {
//	log.Fatal(jsonErr)
//}
//
//gopre.Pre(people1.Number)
//}

func (operation *JsonOperation) getJson(data []byte, myType interface{}) map[string]interface{} {
	inType := myType
	jsonErr := json.Unmarshal(data, &inType)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	operation.convertedJsonFile = inType
	//gopre.Pre(operation.convertedJsonFile)
	return inType.(map[string]interface{})
}
func (operation *JsonOperation) getJsonStruct(autoGoPre ... bool) interface{} {
	if autoGoPre != nil && autoGoPre[0] == true {
		gopre.Pre(operation.convertedJsonFile)
	}
	//if val[0] {
	//
	//}
	//gopre.Pre(operation.convertedJsonFile)
	return operation.convertedJsonFile
}
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (operation *JsonOperation) setDownloadedJsonFileURL(url string) {
	operation.downloadedJsonFileURL = url
}
func (operation *JsonOperation) getFileToInternet(url ... string) []byte {

	httpOperation := http.Client{Timeout: time.Second * 2}
	//gopre.Pre(httpOperation)
	var _url string
	if url != nil && len(url[0]) > 10 {
		_url = url[0]
	} else {
		_url = operation.downloadedJsonFileURL
	}

	httpRequest, err := http.NewRequest(http.MethodGet, _url, nil)
	if err != nil {

		log.Fatal(err)

	}

	httpResponse, getErr := httpOperation.Do(httpRequest)
	if getErr != nil {


		//


		log.Fatal("URL eklenmemis \n", getErr)
		//Get https://raw.githubusercontent.com/weppos/whois/master/data/tld.json: net/http: request canceled (Client.Timeout exceeded while awaiting headers)

	}

	body, readErr := ioutil.ReadAll(httpResponse.Body)
	if readErr != nil {

		//2018/06/29 03:12:27 net/http: request canceled (Client.Timeout exceeded while reading body)
		log.Fatal(readErr)
	}

	io := IoController{}
	io.setFolderPath("downloaded/")
	io.setFileName("astros.json")
	io.openLogFile(false)
	io.appendByte(body)

	return body
}
