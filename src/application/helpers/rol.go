package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/lea55/BACKJOBIEX/src/domain/entity"
)

type Rol struct{}

func NewRol() *Rol {
	return &Rol{}
}

func (r Rol) ReadListFromFile() []entity.Rol {
	pmList := make([]entity.Rol, 0)

	jsonFile, err := os.Open("assets/app/roles.json")
	if err != nil {
		log.Fatal("Error en lecatura de archivo de roles")
	}
	defer jsonFile.Close()

	bytValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error en conversion de datos encontrados en archivo de roles")
	}
	err = json.Unmarshal(bytValue, &pmList)
	if err != nil {
		log.Fatal("Error en masrshal de roles provenientes de archivo")
	}

	return pmList
}
