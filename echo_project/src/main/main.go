package main

import(
    "fmt"
    "net/http"
    "github.com/labstack/echo"
)
func greetingWeb(c echo.Context) error {
    return c.String(http.StatusOK, "Hello from the sever side")
}
func getCats(c echo.Context) error {
    catName := c.QueryParam("catname")
    catType := c.QueryParam("catspecies")
    return c.String(http.StatusOK, fmt.Sprintf("Your cat's name is %s\nand she or he's a %s\n", catName, catType))
}
func main() {
    fmt.Printf("Mornin', Starshine!")
    e := echo.New()
    e.GET("/", greetingWeb)
    e.GET("/cats", getCats)
    e.Start(":8000")
}
