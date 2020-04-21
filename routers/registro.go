package routers

import(
	"encoding/json"
	"net/http"
	"github.com/patino8308/twittor/bd"
	"github.com/patino8308/twittor/models"
)

/*Registro es la funcion para crear en la bd el registro de usuarios*/
func Registro(w http.ResponseWriter, r *http.Request)  {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil{
		http.Error(w, "Error en los datos Recibidos " + err.Error(), 400)
		return
	}

	if len(t.Email) == 0{
		http.Error(w, "El email es requerido " , 400)
		return
	}
	if len(t.Password) < 6{
		http.Error(w, "Debe especificar una contraseña de almenos 6 caracteres " , 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true{
		http.Error(w, "Ya existe un Usuario Registrado con ese Email " , 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error() , 400)
		return
	}

	if status == false{
		http.Error(w, "No se ha logrado insertar el registro del Usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}