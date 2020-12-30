package repository

import (
	"camanager/caservice"
	"encoding/xml"
	"errors"
	"strings"

	"github.com/hooklift/gowsdl/soap"
)

//CARepository Repositorio de CA
type CARepository struct {
	sID      int32
	User     string
	Password string
}

//InitRepository Inicializa una nueva instancia del servicio para comunicarse con el CA Services Desk
func (c *CARepository) InitRepository() error {
	if c.sID == 0 {
		err := login(c)
		if err != nil {
			return err
		}
	}
	return nil
}

//URL del CA Services Desk
const URL string = "http://cscmsdpb2.csc.fmx:8080/axis/services/USD_R11_WebService?WSDL"

//getServices obtiene una sesion con el cliente
func getServices() caservice.USD_WebServiceSoap {
	client := soap.NewClient(URL)
	service := caservice.NewUSD_WebServiceSoap(client)
	return service
}

//Login Inicia sesion y obtiene el SSID del webservices
func login(c *CARepository) error {

	service := getServices()
	res, err := service.Login(&caservice.Login{Username: c.User, Password: c.Password})
	if err != nil {
		c.sID = 0
		return err
	}

	c.sID = res.LoginReturn
	return nil
}

//Close Finaliza sesion con el servicio
func (c *CARepository) Close() error {
	service := getServices()
	_, err := service.Logout(&caservice.Logout{Sid: c.sID})
	if err != nil {
		return err
	}
	c.sID = 0
	return nil
}

//SelectRaw ejecuta una consulta hacia el servidor y regresa un XML de los resultados
//sid ID de la sesion
func (c CARepository) SelectRaw(objeto string, consulta string, limite int32, atributos []string) (*caservice.DoSelectResponse, error) {

	selectReq := &caservice.DoSelect{
		Sid:         c.sID,
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
func (c CARepository) Select(objeto string, consulta string, limite int32, atributos []string) ([]*UDSObjectList, error) {
	resultados, err := c.SelectRaw(objeto, consulta, limite, atributos)

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
