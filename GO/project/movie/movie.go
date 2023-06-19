package movie

import "fmt"

func init() {
	fmt.Println("init: movie")
}

func ReviewMovie(name string, rating float64) {
	fmt.Printf("i reviewed %s and it's rating is %f \n", name, rating)
}

func FindMovieName(imdbID string) string {
	switch imdbID {
	case "tt4154796":
		return "Avenger: Endgame"
	case "tt1825683":
		return "Black Panther"
	}
	return "not found."
}
