package main
//
//import (
//	"net/http"
//	"io/ioutil"
//	"encoding/json"
//	"fmt"
//)
//
//type hos struct {
//	//Number  int    `json:"number"`
//	host string
//}
//type JsonOprtn struct {
//	downloadedJsonFileURL string
//}
//func ini(){
//
//	jsonO := JsonOprtn{}
//	jsonO.setDownloadedJsonURL("https://raw.githubusercontent.com/weppos/whois/master/data/tld.json")
//	jsonByte := jsonO.getInternet()
//
//	jsonParse(jsonByte)
//}
//func (operation *JsonOprtn) setDownloadedJsonURL(url string) {
//	operation.downloadedJsonFileURL = url
//}
//
//func (operation *JsonOprtn) getInternet(url ... string) []byte {
//
//	httpAnsver, _ := http.Get(operation.downloadedJsonFileURL)
//
//	jsonAnswer, _ := ioutil.ReadAll(httpAnsver.Body)
//
//	defer httpAnsver.Body.Close()
//	return jsonAnswer
//}
//func jsonParse(val []byte) {
//	var hosts []hos
//	json.Unmarshal(val, &hosts)
//
//	//gopre.Pre(hosts)
//
//	for _, h := range hosts {
//
//		fmt.Print(h)
//	}
//
//}
