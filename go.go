package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"fmt"
	"encoding/json"
    "io/ioutil"
    "log"
 
    "os"
)

// type user struct {
//     Id        int     `json:"id"`
//     Username  string  `json:"username"`
// }
type Response struct {
    Name    string    `json:"name"`
    Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
    EntryNo int            `json:"entry_number"`
    Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
    Name string `json:"name"`
}


func getUsers(c *gin.Context) {
	  response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject Response
    json.Unmarshal(responseData, &responseObject)

    fmt.Println(responseObject.Name)
    fmt.Println(len(responseObject.Pokemon))

    for i := 0; i < len(responseObject.Pokemon); i++ {
        fmt.Println(responseObject.Pokemon[i].Species.Name)
    }
     c.IndentedJSON(http.StatusOK, responseObject)
}

func main() {
    router := gin.Default()

    // new `GET /users` route associated with our `getUsers` function
    router.GET("/users", getUsers)

    router.Run("localhost:8080")
}