package main

import(
    "fmt"
    "net/http"
    "github.com/labstack/echo"
)
func greetingWeb(c echo.Context) error {
    return c.String(http.StatusOK, "Hello from the sever side")
}
func main() {
    fmt.Printf("Mornin', Starshine!")
    e := echo.New()
    e.GET("/", greetingWeb)
    e.Start(":8000")
}
