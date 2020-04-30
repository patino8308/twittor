package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/patino8308/twittor/bd"
)

/*leoTweetsSeguidores lee los tweets de todos los seguidores*/
func leoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro paginacomo entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTwetsSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer  los Tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
