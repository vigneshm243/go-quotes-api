package main

import (
	"fmt"
	"net/http"
	"regexp"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type quote struct{
	Author string `json:"author"`
	Quote string `json:"quote"`
	Book string `json:"book"`
}

// var quotes = []quote{
// 	{Author: "writer1", Quote: "quote1", Book: "book1"},
// 	{Author: "writer2", Quote: "quote2", Book: "book2"},
// 	{Author: "writer3", Quote: "quote3", Book: "book3"},
// }

// func getQuotes(context *gin.Context){
// 	context.IndentedJSON(http.StatusOK, quotes)
// }

func getQuoteSearch(context *gin.Context){
	search := context.Param("search")
	searchString  := "https://www.goodreads.com/quotes/search?q=" + search + "&commit=Search"
	contentRegexp := regexp.MustCompile("“(.+?)”")
	var quotes []quote

	c := colly.NewCollector(
		colly.AllowedDomains("www.goodreads.com"),
	)
	// extract all the quotes that are on the page
	c.OnHTML(".quoteDetails", func(e *colly.HTMLElement) {
		res := contentRegexp.FindAllStringSubmatch(e.ChildText("div.quoteText"), -1)

		// it's pretty ugly, but, works ( make sure that we can access
		// that slice's slot )
		if len(res) < 1 {
			return
		}

		if len(res[0]) < 1 {
			return
		}

		quotes = append(quotes, quote{
			Quote: res[0][0],
			Author:  e.ChildText(".authorOrTitle"),
		})

		fmt.Print(".")
	})
	fmt.Println("Launching Scraper !")
	c.Visit(searchString)

	fmt.Printf("Scraped %d quotes.\n\n", len(quotes))
	if len(quotes) > 0{
		context.IndentedJSON(http.StatusOK, quotes)
		return 
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "no matching quotes found"})
}

func main(){
	router := gin.Default()
	fmt.Println("quote")
	// router.GET("/quotes", getQuotes) 
	router.GET("/quotes/:search", getQuoteSearch) 
	router.Run("localhost:9091")
}