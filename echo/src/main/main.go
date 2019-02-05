package main

import(
    "fmt"
    "github.com/labstack/echo"
)
func checkAddress(c echo.Context) error {
    return c.String(http.StatusOK, someFunc())
}
func main() {
    e := echo.New()
    e.GET("/", checkAddress)
}
