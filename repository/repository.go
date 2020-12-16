package repository

import (
	"camanager/caservice"
	"encoding/xml"
	"errors"
	"strings"

	"github.com/hooklift/gowsdl/soap"
)

//URL del CA Services Desk
const URL string = "http://cscmsdpb2.csc.fmx:8080/axis/services/USD_R11_WebService?WSDL"

//getServices obtiene una sesion con el cliente
func getServices() caservice.USD_WebServiceSoap {
	client := soap.NewClient(URL)
	service := caservice.NewUSD_WebServiceSoap(client)
	return service
}

//Login Inicia sesion y obtiene el SSID del webservices
func Login(user string, pass string) (int32, error) {

	service := getServices()
	res, err := service.Login(&caservice.Login{Username: user, Password: pass})
	if err != nil {
		return 0, err
	}

	return res.LoginReturn, nil
}

//Logout Finaliza sesion con el servicio
func Logout(sid int32) error {
	if sid == 0 {
		return errors.New("SID no debe estar vacio")
	}

	service := getServices()
	service.Logout(&caservice.Logout{Sid: sid})

	return nil
}

//SelectRaw ejecuta una consulta hacia el servidor y regresa un XML de los resultados
//sid ID de la sesion
func SelectRaw(sid int32, objeto string, consulta string, limite int32, atributos []string) (*caservice.DoSelectResponse, error) {
	if sid <= 0 {
		return nil, errors.New("SID no valido")
	} else if objeto == "" {
		return nil, errors.New("Objeto no debe ser vacio")
	}

	selectReq := &caservice.DoSelect{
		Sid:         sid,
		ObjectType:  objeto,
		WhereClause: consulta,
		MaxRows:     limite,
		Attributes: &caservice.ArrayOfString{
			String: atributos,
		},
	}

	service := getServices()
	res, er := service.DoSelect(selectReq)
	if er != nil {
		return nil, er
	}

	return res, nil
}

//Select ejecuta una consulta hacia el servidor y regresa una lista de objetos
//sid ID de la sesion
func Select(sid int32, objeto string, consulta string, limite int32, atributos []string) ([]*UDSObjectList, error) {
	resultados, err := SelectRaw(sid, objeto, consulta, limite, atributos)

	if err != nil {
		return nil, err
	}

	if strings.Contains(resultados.DoSelectReturn, "Error fetching") {
		return nil, errors.New(resultados.DoSelectReturn)
	}

	var listaResultados []*UDSObjectList
	err = xml.Unmarshal([]byte(resultados.DoSelectReturn), &listaResultados)

	if err != nil {
		return nil, err
	}
	return listaResultados, nil
}
