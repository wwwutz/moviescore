package getratings

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

//PrettyPrinter Print function which prints all the information from all the modules
func PrettyPrinter(MovieName string, year string) {
	RtRating := RtScraper(MovieName, year)
	ImdbRatings := GetImdbRatings(MovieName)
	IntRtRatings, err := strconv.Atoi(RtRating)
	if IntRtRatings == -1 && len(ImdbRatings.Title) == 0 {
		fmt.Println("The Movie Does not seem to exist!")
		fmt.Println("Tip: If you are using spaces in your film name, enclose the movie name in double quotes!")
	} else {
		fmt.Println("Movie Name: "+ImdbRatings.Title)
		fmt.Println("Director: "+ImdbRatings.Director)
		fmt.Println("Cast: "+ImdbRatings.Actors)
		fmt.Println("Year: "+ImdbRatings.Year)
		fmt.Println("Released: "+ImdbRatings.Released)
		fmt.Println("Rated: "+ImdbRatings.Rated)
		fmt.Println("Genre: "+ImdbRatings.Genre)
		fmt.Println("Poster: "+ImdbRatings.Poster)
		fmt.Println("Metascore Rated: "+ImdbRatings.Metascore)
		fmt.Println("Awards: "+ImdbRatings.Awards)
		fmt.Println("Plot: "+ImdbRatings.Plot)
		fmt.Println("Movie Trailer: "+GetTrailer(MovieName))
		fmt.Println(" Ratings from IMDB and Rotten Tomatoes---")
		fmt.Println("IMDB Rating: "+ImdbRatings.ImdbRating)
		if IntRtRatings == -1 && err == nil {
			fmt.Println("There seems to be a problem with rt, try with the year argument!")
		} else if IntRtRatings > 60 && err == nil {
			fmt.Println("Rotten Tomatoes Rating: "+RtRating+"% (Certified Fresh!)")
		} else {
			fmt.Println("Rotten Tomatoes Rating: "+RtRating+"% (Rotten!)")
		}
	}
}

//ASCIIPoster generates the ASCII POSTER!!
func ASCIIPoster() {
	fmt.Println("# moviescore #")
}

//GetJSON Function which takes the url and the target as arguments for parsing json
func GetJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
