package routes

import (
	"net/http"

	"../controllers"
	"../utils/auth"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	r.Use(CommonMiddleware)
	p := r.PathPrefix("/api").Subrouter()

	p.HandleFunc("/test1", controllers.TestAPI).Methods("GET")
	p.HandleFunc("/test2", controllers.TestAPI).Methods("GET")
	p.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	p.HandleFunc("/login", controllers.Login).Methods("POST")

	// Auth route
	s := r.PathPrefix("/api").Subrouter()
	s.Use(CommonMiddleware)
	s.Use(auth.JwtVerify)
	s.HandleFunc("/user", controllers.FetchUsers).Methods("GET", "OPTIONS")
	s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET", "OPTIONS")
	s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT", "OPTIONS")
	s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE", "OPTIONS")

	s.HandleFunc("/person", controllers.CreatePerson).Methods("POST", "OPTIONS")
	s.HandleFunc("/person/{id}", controllers.DeletePerson).Methods("DELETE", "OPTIONS")

	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header, x-access-token")

		if r.Method == "OPTIONS" {
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
