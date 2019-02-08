//not mine, mostly

package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/labstack/echo"
    "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)
type Product struct {
  gorm.Model
  Code string
  Price uint
}
func getPokemon(c echo.Context) error {
	pokemonName := c.QueryParam("name")
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemonName)
	if err != nil {
		fmt.Println("error")
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error")
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	return c.JSON(http.StatusOK, result)
}
func main() {
    myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", getPokemon).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()
    db.AutoMigrate(&Product{})

    db.Create(&Product{Code: "L1212", Price: 1000})
    var product Product
    db.First(&product, 1) // find product with id 1
    db.First(&product, "code = ?", "L1212") // find product with code l1212

    // Update - update product's price to 2000
    db.Model(&product).Update("Price", 2000)

    // Delete - delete product
    db.Delete(&product)
}
