package routers

import (
	"net/http"

	"github.com/patino8308/twittor/bd"
)

/*EliminarTweet elimina untweet*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar parametro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intesntar borrar el Twett"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

}
