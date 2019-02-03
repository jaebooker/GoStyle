package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)
type Cat struct {
    Name    string  `json:"name"`
    Type    string  `json:"type"`
}
type Slugs struct {
    Name    string  `json:"name"`
    Type    string  `json:"type"`
    Count    int  `json:"count"`
    SlimeLevel    int  `json:"slime-level"`
}
type Pandas struct {
    Name    string  `json:"name"`
    Type    string  `json:"type"`
    Count    int  `json:"count"`
    CannibalisticLevel    int  `json:"cannibal-level"`
    HumansConsumed    int  `json:"humans-consumed"`
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
func addSlugs(c echo.Context) error {
    slugs := Slugs{}
    defer c.Request().Body.Close()
    err := json.NewDecoder(c.Request().Body).Decode(&slugs)
    if err != nil {
        log.Printf("Failed reading addSlugs request: %s", err)
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    log.Printf("Here are your slugs, Sir. Mind the slime:", slugs)
    return c.String(http.StatusOK, "We have received your slugs, Sir.")
}
func addPandas(c echo.Context) error {
    pandas := Pandas{}
    err := c.Bind(&pandas)
    if pandas.CannibalisticLevel < 1 {
        logger := "A panda has never had a Cannibalistic Level below one."
        log.Printf(logger)
        return echo.NewHTTPError(http.StatusInternalServerError, logger)
    }
    if err != nil {
        log.Printf("Failed reading your request to add pandas: %s", err)
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    log.Printf("Here are your pandas, Sir. Monsterous things, aren't they?", pandas)
    return c.String(http.StatusOK, "We have received your pandas, Sir. I do hope they don't kill anyone.")
}
func mainAdmin(c echo.Context) error {
    return c.String(http.StatusOK, "You twinkle above us, we twinkle below")
}
func main() {
    fmt.Printf("Mornin', Starshine!")
    e := echo.New()
    g := e.Group("/admin", middleware.Logger())
    g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
            Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
    }))
    g.GET("/main", mainAdmin, middleware.Logger())
    e.GET("/", greetingWeb)
    e.GET("/cats/:data", getCats)
    e.POST("/cats", addKittyCat)
    e.POST("/slugs", addSlugs)
    e.POST("/pandas", addPandas)
    e.Start(":8000")
}
