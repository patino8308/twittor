package models

/*Tweet captura del body, el mensaje que nmos llega*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}