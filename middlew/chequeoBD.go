
package middlew

import (
	"net/http"
	"github.com/patino8308/twittor/bd"
)
/*ChequeoBD valida si esta conectado a la bd*/
func  ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if bd.ChequeoConnection() == 0{
			http.Error(w,"Conexion Perdida con la Base de Datos",500)
		}
		next.ServeHTTP(w, r)
	}

}