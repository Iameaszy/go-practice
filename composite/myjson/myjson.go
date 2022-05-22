package myjson

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"example.com/practice/packages/issuesearch"
)


type Movie struct {
  Title string
  Year int `json:"released"`
  Color bool `json:"color,omitempty"`
  Actors []string
}

var movies = []Movie{
{Title: "Casablanca", Year: 1942, Color: false,
Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
{Title: "Cool Hand Luke", Year: 1967, Color: true,
Actors: []string{"Paul Newman"}},
{Title: "Bullitt", Year: 1968, Color: true,
Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}



func init() {
  fmt.Printf("JSON")
}

func Marshalling() {
  data, err := json.Marshal(movies)
  if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
  }
  fmt.Printf("%s\n", data)
}

func MarshallingIndent() {
  data, err := json.MarshalIndent(movies, "", " ")
  if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
  }
  fmt.Printf("%s\n", data)
}

func IssueSearch() {
  result, err := issuesearch.SearchIssues(os.Args[1:])
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%d issues:\n", result.TotalCount)
  for _, item := range result.Items {
    fmt.Printf("#%-5d %9.9s %.55s\n",
    item.Number, item.User.Login, item.Title)
  }
}