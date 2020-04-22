package models

/*RespuestaLogin tiene el token que se de vuelve con el login*/
type RespuestaLogin struct {
	Token string `json:"token,omitempry"`
}
