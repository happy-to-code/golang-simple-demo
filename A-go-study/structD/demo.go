package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	bytes, _ := json.Marshal(movies)
	fmt.Println(string(bytes))
	fmt.Println("-------------------------")
	indent, _ := json.MarshalIndent(movies, "", " ")
	fmt.Println(string(indent))

}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}}}
