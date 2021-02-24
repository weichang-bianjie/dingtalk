package common

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
	"strings"
)

func HttpPost(url string, reqBody interface{}) (Response *http.Response, bz []byte, err error) {
	reqBz := MarshalJsonIgnoreErr(reqBody)
	resp, respBody, errs := gorequest.New().Post(url).Send(string(reqBz)).End()

	if len(errs) != 0 {
		return nil, nil, errs[0]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("StatusCode != 200")
	}
	return resp, []byte(respBody), nil
}

func MarshalJsonIgnoreErr(v interface{}) []byte {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		log.Println("marshal json fail", err.Error())
	}

	return jsonBytes
}

func UnMarshalJsonIgnoreErr(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Println("unmarshal json fail", err.Error())
	}
}

func FilterContent(content []byte) string {
	//var data types.AlertsContent
	//common.UnMarshalJsonIgnoreErr(content, &data)
	//data.Receiver = ""
	//data.CommonLabels = nil
	//data.CommonAnnotations = types.Annotation{}
	data := strings.ReplaceAll(string(content), "\",", "\",\n")
	data = strings.ReplaceAll(data, "},", "\n},\n")
	return strings.ReplaceAll(data, "{\"", "\n{\n\"")
	//return string(content)
}
