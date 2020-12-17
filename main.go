package main

import (
	"log"

	"github.com/Jebus01/curso-rj-go-mongo/bd"
	"github.com/Jebus01/curso-rj-go-mongo/handlers"
)

func main() {

	if bd.VerificaConexionBD() == 0 {
		log.Fatal("No hay conexion a la base de datos")
		return
	}

	handlers.Manejadores()
}
