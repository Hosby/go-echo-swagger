package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"

	_ "gosampleswagger/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Error struct {
	Errors []Errors `json:"errors"`
}

type Errors struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// Swagger 문서용 주석
// @title Go Swagger Sample API
// @version 1.0
// @description This is a sample gorilla mux API with Swagger documentation.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic  XAuth

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

	r.HandleFunc("/board/{id}", getBoardId).Methods("GET")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.Handle("/", r)
	http.ListenAndServe(":1323", nil)
}

// Path Param : path
// @Summary 유저를 가져온다.
// @Description 유저의 이름을 가져온다.
// @Tags user
// @Accept json
// @Produce json
// @Param name path string true "name of the user"
// @Success 200 string name
// @Failure 400 json Error
// @Failure 404 json Error
// @Failure 500 json Error
// @Router /user/{name} [get]
func getUserName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["name"]
	fmt.Fprintf(w, "user: %s", user)
}

// Query Param : query
// @Summary 유저를 가져온다.
// @Description 유저의 ID를 가져온다.
// @Tags user
// @Accept json
// @Produce json
// @Param id query string true "user ID"
// @Success 200 string id
// @Failure 400 json Error
// @Failure 404 json Error
// @Failure 500 json Error
// @Router /user [get]
func getUserId(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	fmt.Fprintf(w, "userId: %s", userId)
}

// @Summary 게시판을 가져온다.
// @Description 게시판의 ID를 가져온다.
// @Tags board
// @Accept json
// @Produce json
// @Param id path string true "board ID"
// @Success 200 string id
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /board/{id} [get]
func getBoardId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardId := vars["id"]
	fmt.Fprintf(w, "boardId: %s", boardId)
}

// LoggerMiddleware는 Mux에서 사용할 로거 미들웨어입니다.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
