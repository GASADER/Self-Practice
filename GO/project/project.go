package main

import (
	"fmt"

	movie "github.com/GASADER/project/movie"
	ticket "github.com/GASADER/project/ticket"
)

func main() {
	movieName := movie.FindMovieName("tt4154796")
	fmt.Println("hello", movieName)
	movie.ReviewMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
