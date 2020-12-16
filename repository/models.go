package repository

import "encoding/xml"

const (
	//LOGIN Operacion de Login
	LOGIN string = "urn:login"
	//LOGOUT Operacion de Logout
	LOGOUT string = "urn:logout"
	//DOSELECT Realiza una consulta en linea
	DOSELECT string = "url:doSelect"
)

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

//DoSelectResponse Respuesta de un Do Select
type DoSelectResponse struct {
	XMLName xml.Name
	Body    struct {
		XMLName          xml.Name
		DoSelectResponse struct {
			XMLName        xml.Name
			DoSelectReturn struct {
				XMLName       xml.Name
				UDSObjectList struct {
					XMLName   xml.Name
					UDSObject []struct {
						XMLName    xml.Name
						Handle     string `xml:"Handle"`
						Attributes struct {
							XMLName   xml.Name
							Attribute []struct {
								XMLName   xml.Name
								DataType  string `xml:"DataType,attr"`
								AttrName  string `xml:"AttrName"`
								AttrValue string `xml:"AttrValue"`
							} `xml:"Attribute"`
						} `xml:"Attributes"`
					} `xml:"UDSObject"`
				} `xml:"UDSObjectList"`
			} `xml:"doSelectReturn"`
		} `xml:"doSelectResponse"`
	} `xml:"Body"`
}
