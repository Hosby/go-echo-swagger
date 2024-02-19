package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"

	_ "gosampleswagger/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Swagger 문서용 주석
// @title MyAPI
// @version 1.0
// @description This is a sample Echo API with Swagger documentation.
// @host localhost:1323
// @BasePath /
func main() {
	r := mux.NewRouter()

	r.Use(LoggerMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	r.HandleFunc("/user/{name}", getUserName).Methods("GET")
	r.HandleFunc("/user", getUserId).Methods("GET")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.Handle("/", r)
	http.ListenAndServe(":1323", nil)
}

// @Summary 유저를 가져온다.
// @Description 유저의 이름을 가져온다.
// @Accept json
// @Produce json
// @Param name path string true "name of the user"
// @Success 200 string name
// @Router /user/{name} [get]
func getUserName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["name"]
	fmt.Fprintf(w, "user: %s", user)
}

// @Summary 유저를 가져온다.
// @Description 유저의 ID를 가져온다.
// @Accept json
// @Produce json
// @Param id query string true "user ID"
// @Success 200 string id
// @Router /user [get]
func getUserId(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	fmt.Fprintf(w, "userId: %s", userId)
}

// LoggerMiddleware는 Mux에서 사용할 로거 미들웨어입니다.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
