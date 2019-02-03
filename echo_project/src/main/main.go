package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
    "github.com/labstack/echo"
)
type Cat struct {
    Name    string  `json:"name"`
    Type    string  `json:"type"`
}
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
func addKittyCat(c echo.Context) error {
    cat := Cat{}
    defer c.Request().Body.Close()
    b, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        log.Printf("Failed reading: %s", err)
        return c.String(http.StatusInternalServerError, "")
    }
    err = json.Unmarshal(b, &cat)
    if err != nil {
        log.Printf("Failed inmarshalling: %s", err)
        return c.String(http.StatusInternalServerError, "")
    }
    log.Printf("Here is your kitty cat, Sir:", cat)
    return c.String(http.StatusOK, "We have received your kitty cat, Sir.")
}
func main() {
    fmt.Printf("Mornin', Starshine!")
    e := echo.New()
    e.GET("/", greetingWeb)
    e.GET("/cats/:data", getCats)
    e.POST("/cats", addKittyCat)
    e.Start(":8000")
}
