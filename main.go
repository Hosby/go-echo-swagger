package main

import (
	_ "gosampleswagger/docs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Swagger 문서용 주석
// @title MyAPI
// @version 1.0
// @description This is a sample Echo API with Swagger documentation.
// @host localhost:1323
// @BasePath /
func main() {
	e := echo.New()

	// Middelware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/user/:name", getUserName)

	e.GET("/user", getUserId)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

// @Summary 유저를 가져온다.
// @Description 유저의 이름을 가져온다.
// @Accept json
// @Produce json
// @Param name path string true "name of the user"
// @Success 200 string name
// @Router /user/{name} [get]
func getUserName(c echo.Context) error {
	user := c.Param("name")
	return c.String(http.StatusOK, "user: "+user)
}

// @Summary 유저를 가져온다.
// @Description 유저의 ID를 가져온다.
// @Accept json
// @Produce json
// @Param id query string true "user ID"
// @Success 200 string id
// @Router /user [get]
func getUserId(c echo.Context) error {
	userId := c.QueryParam("id")
	return c.String(http.StatusOK, "userId: "+userId)
}
