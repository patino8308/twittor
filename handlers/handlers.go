package handlers

import(
	"log"
	"net/http"
	"os"
	"github.com/patino8308/twittor/middlew"
	"github.com/patino8308/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores de las rutas*/
func Manejadores(){

	router := mux.NewRouter()

	router.HandleFunc("/registro",middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT,handler))
}



