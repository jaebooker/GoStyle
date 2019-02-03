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
    dataType := c.Param("data")
    if dataType == "string" {
        return c.String(http.StatusOK, fmt.Sprintf("Your cat's name is %s\nand she or he's a %s\n", catName, catType))
    }
    if dataType == "json" {
        return c.JSON(http.StatusOK, map[string]string{
            "name": catName,
            "type": catType,
        })
    }
    return c.JSON(http.StatusBadRequest, map[string]string{
        "error": "this is wrong",
    })
}
func main() {
    fmt.Printf("Mornin', Starshine!")
    e := echo.New()
    e.GET("/", greetingWeb)
    e.GET("/cats/:data", getCats)
    e.Start(":8000")
}
