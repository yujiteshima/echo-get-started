package main

import(
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    // create instance
    e := echo.New()

    // set middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // set root
    e.GET("/", hello)

    // execute server port 1234
    e.Logger.Fatal(e.Start(":1234"))

}

// define hundler
func hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}
