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

type people struct {
	Number  int    `json:"number"`
	Message string `json:"message"`
}
type JsonOperation struct {
	downloadedJsonFileURL string
	convertedJsonFile     interface{}
}

func init() {
	jsonOpe := JsonOperation{}

	jsonOpe.setDownloadedJsonFileURL("http://api.open-notify.org/astros.json")
	//jsonOpe.setDownloadedJsonFileURL("https://raw.githubusercontent.com/weppos/whois/master/data/tld.json")
	returnByte := jsonOpe.getFileToInternet()

	jsonObj := jsonOpe.getJson(returnByte, people{})
	_ = jsonOpe.getJsonStruct()
	//gopre.Pre(ret)
	gopre.Pre(jsonObj)
	//
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
}
func (operation *JsonOperation) getJson(data []byte, myType interface{}) map[string]interface{} {
	inType := myType
	jsonErr := json.Unmarshal(data, &inType)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	operation.convertedJsonFile = inType
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
		log.Fatal("URL eklenmemis \n",getErr)
	}

	body, readErr := ioutil.ReadAll(httpResponse.Body)
	if readErr != nil {
		//log.Fatal(readErr)
	}

	return body
}
