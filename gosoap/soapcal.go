package gosoap

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

//Fuente: https://fale.io/blog/2018/12/03/calling-a-soap-service-in-go/

//SoapCallHandleResponse ejecuta la llamada al servicio y obtiene el resultado
func SoapCallHandleResponse(ws string, action string, payloadInterface interface{}, result interface{}) error {
	body, err := call(ws, action, payloadInterface)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	return nil
}

//call Ejecuta la llamada al WebServices
func call(ws string, action string, payloadInterface interface{}) ([]byte, error) {
	v := soapRQ{
		XMLNsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		XMLNsXSD:  "http://www.w3.org/2001/XMLSchema",
		XMLNsXSI:  "http://www.w3.org/2001/XMLSchema-instance",
		Body: soapBody{
			Payload: payloadInterface,
		},
	}
	payload, err := xml.MarshalIndent(v, "", "  ")

	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", ws, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/xml, multipart/related")
	req.Header.Set("SOAPAction", action)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	// dump, err := httputil.DumpRequestOut(req, true)
	// if err != nil {
	// 	return nil, err
	// }
	// _ := dump

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(bodyBytes))
	defer response.Body.Close()
	return bodyBytes, nil
}
