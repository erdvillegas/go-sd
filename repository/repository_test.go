package repository

import (
	"testing"
	"time"
)

//Realiza prueba de inicio de sesion
func TestLoginOk(t *testing.T) {
	ssid, err := Login("", "")
	if err != nil {
		t.Errorf("Ocurrio el error al generar la consulta, %s", err.Error())
	}

	if ssid == 0 {
		t.Errorf("Error al realizar la consulta")
	}
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
