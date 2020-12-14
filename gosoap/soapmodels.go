package gosoap

import "encoding/xml"

//soapRQ Peticion SOAP
type soapRQ struct {
	XMLName   xml.Name `xml:"soap:Envelope"`
	XMLNsSoap string   `xml:"xmlns:soap,attr"`
	XMLNsXSI  string   `xml:"xmlns:xsi,attr"`
	XMLNsXSD  string   `xml:"xmlns:xsd,attr"`
	Body      soapBody
}

//soapBody Cuerpo de la peticion
type soapBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Payload interface{}
}
