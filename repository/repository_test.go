package repository

import (
	"log"
	"testing"
)

//Realiza prueba de inicio de sesion
func TestLoginOk(t *testing.T) {
	caRepository := CARepository{
		User:     "",
		Password: "",
	}

	caRepository.InitRepository()
	if caRepository.sID == 0 {
		t.Errorf("No se pudo iniciar sesion, %d", caRepository.sID)
	}

	defer caRepository.Close()
}

func TestLoginFail(t *testing.T) {
	caRepository := CARepository{
		User:     "",
		Password: "",
	}

	caRepository.InitRepository()

	if caRepository.sID != 0 {
		t.Errorf("Si se pudo realizar la consulta, SID: %d", caRepository.sID)
		defer caRepository.Close()
	}
}

//LogOutTest Realiza una prueba
func TestLogOut(t *testing.T) {
	caRepository := CARepository{
		User:     "",
		Password: "",
	}
	caRepository.InitRepository()

	if caRepository.sID == 0 {
		t.Skip()
	}

	err := caRepository.Close()
	if err != nil {
		t.Skip()
	}

	if caRepository.sID != 0 {
		t.Errorf("Fallo al realizar el LogOut, SID: %d", caRepository.sID)
	}

}

func TestDoSelectRaw(t *testing.T) {

	caRepository := CARepository{
		User:     "",
		Password: "",
	}
	caRepository.InitRepository()

	resultados, err := caRepository.SelectRaw("cr", "ref_num='640103179'", 2, []string{"ref_num"})
	defer caRepository.Close()

	if err != nil {
		log.Fatal("No se pudo ejecutar el Login")
	}

	if resultados.DoSelectReturn == "" {
		t.Errorf("Resultados no validos")
	}
}

func TestDoSelect(t *testing.T) {
	caRepository := CARepository{
		User:     "",
		Password: "",
	}

	caRepository.InitRepository()

	resultados, err := caRepository.Select("cr", "ref_num='640103179'", 1, []string{"ref_num", "summary"})
	defer caRepository.Close()

	if err != nil {
		t.Error("No se pudo ejecutar el Login")
	}

	if len(resultados) <= 0 {
		t.Errorf("Resultados no validos")
	}
}
