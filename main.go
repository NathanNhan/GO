// package main

// import (
//     "encoding/json"
//     "fmt"
//     "io/ioutil"
//     "log"
//     "net/http"
//     "os"
// 	// "reflect"
// 	// "strings"
// )



// // A Response struct to map the Entire Response
// type Response struct {
//     Name    string    `json:"name"`
//     Pokemon []Pokemon `json:"pokemon_entries"`
	
// }

// // A Pokemon Struct to map every pokemon to.
// type Pokemon struct {
//     EntryNo int            `json:"entry_number"`
//     Species PokemonSpecies `json:"pokemon_species"`
// 	Series is_main_series `json:"is_main_series"`
// }

// // A struct to map our Pokemon's Species which includes it's name
// type PokemonSpecies struct {
//     Name string `json:"name"`
// }

// type is_main_series struct {
// 	Series bool `json:"is_main_series"`
// }

// func main() {
//     response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
//     if err != nil {
//         fmt.Print(err.Error())
//         os.Exit(1)
//     }

//     responseData, err := ioutil.ReadAll(response.Body)
//     if err != nil {
//         log.Fatal(err)
//     }

//     var responseObject Response
//     json.Unmarshal(responseData, &responseObject)
//     // fmt.Println(responseObject)
//     fmt.Println(responseObject.Name)
//     fmt.Println(len(responseObject.Pokemon))

//     for i := 0; i < len(responseObject.Pokemon); i++ {
// 		users := responseObject.Pokemon[i].Species.Name == "151"
//     }
 

// }


// // func Filter(vs []string, f func(string) bool) []string {
// //     filtered := make([]string, 0)
// //     for _, v := range vs {
// //         if f(v) {
// //             filtered = append(filtered, v)
// //         }
// //     }
// //     return filtered
// // }