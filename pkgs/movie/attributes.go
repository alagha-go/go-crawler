package movie

import (
	"crawler/pkgs/printer"
	"strings"

	"github.com/mxschmitt/playwright-go"
	"golang.org/x/net/html"
)


func (movie *Movie) SetName(doc playwright.Page) error {
	movieNameContent, err := doc.QuerySelector(".heading-name > a")
	HandleErrorByPanic(err)
	movie.Name, err = movieNameContent.InnerHTML()
	printer.PrintGreen(movie.Name)
	return err
}


func (movie *Movie) SetImage(value string) {
	 doc, err := html.Parse(strings.NewReader(value))
    HandleErrorByPanic(err)
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "img" {
            for _, a := range n.Attr {
                if a.Key == "data-src"{
                    movie.ImageUrl1 = a.Val
                }
				if a.Key == "src" {
					movie.ImageUrl2 = a.Val
				}
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
	f(doc)
}


func (movie *Movie) SetDescription(doc playwright.Page) {
	movieDescContent, err := doc.QuerySelector(".description")
	HandleErrorByPanic(err)
	movie.Description, err = movieDescContent.TextContent()
	HandleErrorByPanic(err)
}


func (movie *Movie) SetElements(doc playwright.Page) {
	movieElements, err := doc.QuerySelector(".elements")
	HandleErrorByPanic(err)
	firstRow, err := movieElements.QuerySelector(".row")
	HandleErrorByPanic(err)
	movie.SetFirstRow(firstRow)
	movie.SetSecondRow(firstRow)
}


func (movie *Movie) SetFirstRow(doc playwright.ElementHandle) {
	elements, err := doc.QuerySelectorAll("div.row-line")
	HandleErrorByPanic(err)
	release, err := elements[0].TextContent()
	release = strings.ReplaceAll(release, "\n", "")
	release = strings.ReplaceAll(release, "  ", "")
	movie.Release = strings.ReplaceAll(release, "Released:  ", "")
	HandleErrorByPanic(err)
	genresData, err := elements[1].QuerySelectorAll("a")
	HandleErrorByPanic(err)
	for _, genreData := range genresData {
		genre, err := genreData.InnerHTML()
		HandleErrorByPanic(err)
		movie.Genre = append(movie.Genre, genre)
	}
	castsData, err := elements[2].QuerySelectorAll("a")
	HandleErrorByPanic(err)
	for _, castData := range castsData {
		cast, err := castData.InnerHTML()
		HandleErrorByPanic(err)
		movie.Casts = append(movie.Casts, cast)
	}
}


func (movie *Movie) SetSecondRow(doc playwright.ElementHandle) {
	elements, err := doc.QuerySelectorAll("div.row-line")
	HandleErrorByPanic(err)
	duration, err := elements[3].TextContent()
	HandleErrorByPanic(err)
	duration = strings.ReplaceAll(duration, "Duration:  ", "")
	duration = strings.ReplaceAll(duration, "\n", "")
	duration = strings.ReplaceAll(duration, "  ", "")
	movie.Duration = strings.ReplaceAll(duration, "min", "")

	country, err := elements[4].TextContent()
	HandleErrorByPanic(err)
	movie.Country = strings.ReplaceAll(country, "Country:  ", "")
	production, err := elements[5].TextContent()
	production = strings.ReplaceAll(production, "\n", "")
	production = strings.ReplaceAll(production, "  ", "")
	movie.Production = strings.ReplaceAll(production, "Production:", "")
	HandleErrorByPanic(err)
}