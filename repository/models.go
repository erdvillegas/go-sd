package repository

import "encoding/xml"

//LoginResponse Respuesta del Login
type LoginResponse struct {
	XMLName xml.Name
	Body    struct {
		XMLName       xml.Name
		LoginResponse struct {
			XMLName     xml.Name
			LoginReturn int32 `xml:"loginReturn"`
		} `xml:"loginResponse"`
	} `xml:"Body"`
}

//DoSelectResponse Respuesta de la consulta
type DoSelectResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Body    struct {
		Text             string `xml:",chardata"`
		DoSelectResponse struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			DoSelectReturn struct {
				Text  string `xml:",chardata"`
				Xmlns string `xml:"xmlns,attr"`
			} `xml:"doSelectReturn"`
		} `xml:"doSelectResponse"`
	} `xml:"Body"`
}
