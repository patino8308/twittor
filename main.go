package main
import (
	"log"
	"github.com/patino8308/twittor/bd"
	"github.com/patino8308/twittor/handlers"
	

)

func main() {
	if bd.ChequeoConnection()==0{
		log.Fatal("Sin conexion a la BD")
	}
	handlers.Manejadores()
}
