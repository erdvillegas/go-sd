package repository

import "encoding/xml"

//UDSObjectList Respuesta de un Do Select
type UDSObjectList struct {
	XMLName   xml.Name `xml:"UDSObjectList"`
	UDSObject struct {
		XMLName    xml.Name
		Handle     string `xml:"Handle"`
		Attributes []struct {
			XMLName   xml.Name `xml:"Attributes"`
			Attribute []struct {
				XMLName   xml.Name `xml:"Attribute"`
				DataType  string   `xml:"DataType,attr"`
				AttrName  string   `xml:"AttrName"`
				AttrValue string   `xml:"AttrValue"`
			} `xml:"Attribute"`
		} `xml:"Attributes"`
	} `xml:"UDSObject"`
}
