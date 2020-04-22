package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/patino8308/twittor/bd"
	"github.com/patino8308/twittor/models"
)

/*GraboTweet permite grabar el tweeet en la bd */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTeewt(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "no se ha logrado insertar el registro Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
