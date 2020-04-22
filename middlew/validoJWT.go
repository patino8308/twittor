package middlew

import(
	"net/http"
	"github.com/patino8308/twittor/routers"
)

/*ValidoJWT permite  validar el jwt*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		_, _, _, err := routes.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil{
			http.Error(w,"Error en el Token !"+err.Error()), http.StatusBadRequest
			return
		}
		next.ServeHTTP(w, r)
	}
}