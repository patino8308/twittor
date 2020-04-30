package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/patino8308/twittor/middlew"
	"github.com/patino8308/twittor/routers"
	"github.com/rs/cors"
)

/*Manejadores de las rutas*/
func Manejadores() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminartweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("get")

	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("get")

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(routers.AltaRelacion)).Methods("post")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(routers.BajaRelacion)).Methods("delete")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(routers.ConsultaRelacion)).Methods("get")

	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(routers.ListaUsuarios)).Methods("get")
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoBD(routers.LeoTweets)).Methods("get")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
