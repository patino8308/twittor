package routers

import (
	"net/http"

	"github.com/patino8308/twittor/bd"
	"github.com/patino8308/twittor/models"
)

/*BajaRelacion realiza el borrado de la relacion entre usuariso*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar la relación"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		if err != nil {
			http.Error(w, "No se ha logrado borrar la relación"+err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)

}
