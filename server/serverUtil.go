package server

import ("log"
	"net/http"
	"github.com/mauriliommachado/fogaoAPI/db"
)

type ServerProperties struct {
	Port    string
	Address string
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
}

func badRequest(w http.ResponseWriter, err error) {
	log.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
}

func validAuthHeader(req *http.Request) bool {
	auth := req.Header.Get("Authorization")
	if len(auth) <= 6 {
		return false
	}
	var user db.User
	user.Token = auth[6:]
	if user.FindHash(db.GetUsersCollection()){
		return true
	}else{
		return false
	}
}