package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/patino8308/twittor/bd"
	"github.com/patino8308/twittor/models"
)

/*Email valor del email usado en todos los endpoints*/
var Email string

/*IDUsuario id devuelto del modelo que se  usara en todos los endpoints*/
var IDUsuario string

/*ProcesoToken proceso token para extraer valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	miClave := []byte("MasterdelDesarrollo_grupoFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token Inválido")
	}

	return claims, false, string(""), err
}
