package main

import(
    "fmt"
    "net/http"
    "github.com/labstack/echo"
)
func main() {
    fmt.Printf("Mornin', Starshine!")
    e := echo.New()
    e.GET("/", func (c echo.Context) error {
        return c.String(http.StatusOK, "Hello from the sever side")
    })
    e.Start(":8000")
}
