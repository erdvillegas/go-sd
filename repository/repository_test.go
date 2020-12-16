package repository

import (
	"log"
	"testing"
	"time"
)

//Realiza prueba de inicio de sesion
func TestLoginOk(t *testing.T) {
	sid, err := Login("", "")
	if err != nil {
		t.Errorf("Ocurrio el error al generar la consulta, %s", err.Error())
	}

	if sid == 0 {
		t.Errorf("Error al realizar la consulta")
	}

	Logout(sid)
}

func TestLoginFail(t *testing.T) {
	sid, err := Login("", "")

	time.Sleep(5 * time.Second)
	if sid > 0 || err == nil {
		t.Errorf("Si se pudo realizar la consulta, SID: %d", sid)
	}
}

//LogOutTest Realiza una prueba
func TestLogOut(t *testing.T) {
	sid, err := Login("", "")

	if err != nil {
		err = Logout(sid)
		if err != nil {
			t.Errorf("Fallo al realizar el LogOut, SID: %d", sid)
		}
	}
}

func TestDoSelect(t *testing.T) {

	sid, err := Login("", "")
	if err != nil {
		log.Fatal("No se pudo ejecutar el Login")
	}

	resultados, err := Select(sid, "cr", "ref_num=640103179", 2, []string{"ref_num"})
	if err != nil {
		log.Fatal("No se pudo ejecutar el Login")
	}

	if len(resultados.Body.DoSelectResponse.DoSelectReturn.UDSObjectList.UDSObject) <= 0 {
		t.Errorf("No se encontraron resultados")
	}

	Logout(sid)
}
