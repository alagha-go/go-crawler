package movie

import (
	"fmt"

	"github.com/mxschmitt/playwright-go"
)

type Movie struct {
	Name							string							`json:"name"`
	Type							string							`json:type"`
	PageUrl							string							`json:"page_url"`
	Release							string							`json:"release"`
	ImageUrl1						string							`json:"image_url1"`
	ImageUrl2						string							`json:"image_url2"`
	Duration						string							`json:"duration"`
	Country							string							`json:"country"`
	Production						string							`json:"production"`
	Description						string							`json:"description"`
	Genre							[]string						`json:"genre"`
	Casts							[]string						`json:"casts"`
}

type Browser struct {
	Browser playwright.Browser
}

type Page struct {
	Page 	playwright.Page
}



var (
	mainUrl = "https://tinyzonetv.to"
	additionParameter = "/movie?page="
	startPos = 1
	browser Browser
)


func StartCrawling() {
	browser = LaunchCrawler()
}



func LaunchCrawler() Browser{
	var mainBrowser Browser
	var headless bool = true
	pw, err := playwright.Run()
	HandleErrorByPanic(err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless})
	HandleErrorByPanic(err)
	mainBrowser.Browser = browser
	page := mainBrowser.NewPage(fmt.Sprintf("%s%s%d", mainUrl, additionParameter, startPos))
	page.StartCollecting()
	return mainBrowser
}


func (browser *Browser)NewPage(url string) Page {
	var MainPage Page
	page, err := browser.Browser.NewPage()
	HandleErrorByPanic(err)
	_, err = page.Goto(url)
	HandleErrorByPanic(err)
	MainPage.Page = page
	return MainPage
}

func (page *Page) GotoPage(url string) {
	_, err := page.Page.Goto(url)
	HandleErrorByPanic(err)
}
