package movie

import (
	"crawler/pkgs/printer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

type MoviePath []string


func (page *Page) StartCollecting() {
	var movies []Movie

	for index:=1; index<1077; index++{
		printer.PrintPurple(fmt.Sprintf("Page: %d", index))
		if index != 1 {
			page.GotoPage(fmt.Sprintf("%s%s%d", mainUrl, additionParameter, index))
		}
		moviesContent, err := page.Page.QuerySelectorAll(".flw-item")
		HandleErrorByPanic(err)
		for number:=0; number<len(moviesContent); number++ {
			moviesContent2, err := page.Page.QuerySelectorAll(".flw-item")
			HandleErrorByPanic(err)
			printer.PrintBlue(fmt.Sprintf("Movie: %d", number+1))
			var movie Movie
			oneMovieContent, err := moviesContent2[number].QuerySelector(".film-poster")
			HandleErrorByPanic(err)
			innerHtml, err := oneMovieContent.InnerHTML()
			HandleErrorByPanic(err)
			movie.SetImage(innerHtml)
			pageUrl := GetHref(innerHtml)
			page.GotoPage(fmt.Sprintf("%s%s", mainUrl, pageUrl))
			movie.PageUrl = fmt.Sprintf("%s%s", mainUrl, pageUrl)
			movie.SetName(page.Page)
			movie.SetDescription(page.Page)
			movie.SetElements(page.Page)
			movie.Type = "Movie"
			movies = append(movies, movie)
			page.Page.GoBack()
		}
	if len(movies) == 50 {
		content, err := json.Marshal(movies)
		HandleErrorByPanic(err)
		path := fmt.Sprintf("./DB/page%d.json", index)
		ioutil.WriteFile(path, content, 0755)
	}
	}

}


func GetHref(value string) string {
	var href string
	 doc, err := html.Parse(strings.NewReader(value))
    HandleErrorByPanic(err)
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, a := range n.Attr {
                if a.Key == "href" {
                    href = a.Val
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
	f(doc)
	return href
}