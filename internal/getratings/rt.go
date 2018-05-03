package getratings

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

//RtScraper function takes the movie name and year(an optional argument) as arguments and returns the rating from rt
func RtScraper(mname string, year string) string {
	movieName := strings.Replace(mname, " ", "_", 9)
	urlis := "https://www.rottentomatoes.com/m/" + movieName
	if len(year) == 4 {
		urlis = urlis + "_" + year
	}
	doc, err := goquery.NewDocument(urlis)
	if err != nil {
		fmt.Println("Error occurred")
	}
	rating := doc.Find(".meter-value.superPageFontColor span").Text()
	if len(rating) == 0 {
		return "-1"
	} else if len(rating) > 4 {
		return rating[:3]
	} else {
		return rating[:2]
	}
}

//RtReviewScraper function also takes in movieName and year(an optional argument) as arguments prints all the reviews
func RtReviewScraper(mname string, year string) {
	movieName := strings.Replace(mname, " ", "_", 9)
	urlis := "https://www.rottentomatoes.com/m/" + movieName
	if len(year) == 4 {
		urlis = urlis + "_" + year
	}
	doc, err := goquery.NewDocument(urlis)
	if err != nil {
		fmt.Println("error Occurred!")
	}
	finder := doc.Find("#reviews .review_quote")
	if len(finder.Nodes) > 0 {
		fmt.Println("Reviews from RT!",)
		doc.Find("#reviews .review_quote").Each(func(i int, s *goquery.Selection) {
			review := s.Find("p").Text()
			fmt.Println(strings.TrimSpace(review))
			fmt.Println("-------------------")
		})
	} else {
		fmt.Println("Looks like Rt also needs the year argument!")
	}
}
