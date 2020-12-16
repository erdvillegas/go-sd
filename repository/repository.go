package repository

import (
	"camanager/caservice"
	"camanager/gosoap"
	"errors"
	"log"
)

//URL del CA Services Desk
const URL string = "http://cscmsdpb2.csc.fmx:8080/axis/services/USD_R11_WebService?WSDL"

//Login Inicia sesion y obtiene el SSID del webservices
func Login(user string, pass string) (int32, error) {

	login := &caservice.Login{
		Username: user,
		Password: pass,
	}

	var loginResult LoginResponse
	err := gosoap.SoapCallHandleResponse(URL, LOGIN, login, &loginResult)

	if err != nil {
		return 0, err
	}

	if loginResult.Body.LoginResponse.LoginReturn <= 0 {
		return 0, errors.New("No se pudo ejecutar la consulta")
	}

	return loginResult.Body.LoginResponse.LoginReturn, nil
}

//Logout Finaliza sesion con el servicio
func Logout(sid int32) error {
	if sid == 0 {
		return errors.New("SID no debe estar vacio")
	}
	sidRq := &caservice.Logout{Sid: sid}

	err := gosoap.SoapCallHandleResponse(URL, LOGOUT, sidRq, nil)
	if err != nil {
		return err
	}

	return nil
}

//Select ejecuta una consulta hacia el servidor
//sid ID de la sesion
func Select(sid int32, objeto string, consulta string, limite int, atributos []string) (*DoSelectResponse, error) {
	if sid <= 0 {
		return nil, errors.New("SID no valido")
	} else if objeto == "" {
		return nil, errors.New("Se debe elegir almenos un objeto a consultar")
	}

	consultaReq := &caservice.DoSelect{
		Sid:         sid,
		ObjectType:  objeto,
		WhereClause: consulta,
		Attributes: &caservice.ArrayOfString{
			String: atributos,
		},
	}

	var consultaRes DoSelectResponse
	err := gosoap.SoapCallHandleResponse(URL, DOSELECT, consultaReq, &consultaRes)
	log.Printf("#%s", &consultaRes)

	if err != nil {
		return nil, err
	}

	return &consultaRes, nil
}
