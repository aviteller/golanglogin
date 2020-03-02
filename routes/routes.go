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
	s.HandleFunc("/user/{id}/{onlyuser}", controllers.GetUser).Methods("GET", "OPTIONS")
	s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT", "OPTIONS")
	s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE", "OPTIONS")

	// person routes
	s.HandleFunc("/person", controllers.CreatePerson).Methods("POST", "OPTIONS")
	s.HandleFunc("/person/{id}", controllers.DeletePerson).Methods("DELETE", "OPTIONS")
	s.HandleFunc("/person/{id}", controllers.UpdatePerson).Methods("PUT", "OPTIONS")
	// s.HandleFunc("/person/{user_id}", controllers.GetPeople).Methods("GET", "OPTIONS")
	s.HandleFunc("/person/{id}", controllers.GetPerson).Methods("GET", "OPTIONS")

	//Meals
	s.HandleFunc("/meal", controllers.CreateMeal).Methods("POST", "OPTIONS")
	s.HandleFunc("/meal/{id}", controllers.DeleteMeal).Methods("DELETE", "OPTIONS")
	s.HandleFunc("/meal/{id}", controllers.GetMeal).Methods("GET", "OPTIONS")
	s.HandleFunc("/meal/{id}", controllers.UpdateMeal).Methods("PUT", "OPTIONS")
	//EatenMeals
	s.HandleFunc("/eatenmeal", controllers.CreateEatenMeal).Methods("POST", "OPTIONS")
	s.HandleFunc("/eatenmeal/{id}", controllers.DeleteEatenMeal).Methods("DELETE", "OPTIONS")
	s.HandleFunc("/eatenmeal/{id}", controllers.GetEatenMeal).Methods("GET", "OPTIONS")
	//MealRatings
	s.HandleFunc("/rating", controllers.CreateMealRating).Methods("POST", "OPTIONS")
	s.HandleFunc("/rating/{id}", controllers.DeleteMealRating).Methods("DELETE", "OPTIONS")
	s.HandleFunc("/rating/{id}", controllers.GetMealRating).Methods("GET", "OPTIONS")
	s.HandleFunc("/rating/{id}", controllers.UpdateMealRating).Methods("PUT", "OPTIONS")
	//Ingredients
	s.HandleFunc("/ingredient", controllers.CreateIngredient).Methods("POST", "OPTIONS")
	s.HandleFunc("/ingredient/{id}", controllers.DeleteIngredient).Methods("DELETE", "OPTIONS")

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
