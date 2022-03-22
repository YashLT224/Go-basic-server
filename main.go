package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Every movie has a director
type Movie struct {
	ID       string    `json:"id,omitempty"`//omitempty means excluded if it is empty
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Movie{})      			//whe you dont find book by id
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
	return
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}
}

func main() {
	//init router
	r := mux.NewRouter() //newRouter is a function inside mux  library

	movies = append(movies, Movie{ID: "1", Isbn: "3445", Title: "Movie one", Director: &Director{Firstname: "yash", Lastname: "verma"}})
	movies = append(movies, Movie{ID: "2", Isbn: "3446", Title: "Movie Two", Director: &Director{Firstname: "Jasmine", Lastname: "verma"}})
	movies = append(movies, Movie{ID: "3", Isbn: "3447", Title: "Movie Three", Director: &Director{Firstname: "Preeti", Lastname: "verma"}})
	movies = append(movies, Movie{ID: "4", Isbn: "3448", Title: "Movie Four", Director: &Director{Firstname: "Uthkarsh", Lastname: "verma"}})
	movies = append(movies, Movie{ID: "5", Isbn: "3449", Title: "Movie Five", Director: &Director{Firstname: "Bhavya", Lastname: "verma"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting at Port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
