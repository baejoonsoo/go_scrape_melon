package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	cleanstring "github.com/baejoonsoo/Melon/cleanString"
	"github.com/labstack/echo"
)

const BASE_URL string="https://www.melon.com/chart/"

type( 
	songData struct{
		rank string
		title string
		singer string
		album string
	}

	songDatas struct{
		songs []songData
	}
)


func getData(card *goquery.Selection) []songData{
	songDataArr := []songData{}
	
	card.Each(func(i int, card *goquery.Selection) {
		song_info:=card.Find(".wrap_song_info")


		rank:=(cleanstring.CleanString(card.Find(".t_center  .rank ").Text()))
		title:=(cleanstring.CleanString(song_info.Find(".rank01").Text()))
		singer:=(cleanstring.CleanString(song_info.Find(".rank02 > a").Text()))
		album:=(cleanstring.CleanString(song_info.Find(".rank03 > a").Text()))

		// fmt.Println("["+rank+"] ["+title+"] ["+singer+"] ["+album+"]")

		newData := songData{
			rank: rank,
			title: title,
			singer: singer,
			album: album,
		}
		songDataArr = append(songDataArr, newData)
	})

	return songDataArr
}

func scpape() []songData{
	res,_:=http.Get(BASE_URL)
	
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	
	searchCards := doc.Find("table").Find("tbody")

	lst50:=getData(searchCards.Find(".lst50"))
	lst100:=getData(searchCards.Find(".lst100"))

	lists:=append(lst50,lst100...)
	
	defer res.Body.Close()

	return lists
}

func handleScrape(c echo.Context)(err error) {
	// lists := scpape()
	u := new(songDatas)

	if err = c.Bind(u); err != nil {
		return
	}
	if err = c.Validate(u); err != nil {
		return
	}

	// u.songs=lists

	// jsonData.songs=lists

	// fmt.Println(jsonData)
	// fmt.Println(len(u.songs))
	

	
	// return c.JSON(http.StatusOK,u)
	return c.String(http.StatusOK,"0")
}


func main(){

	e := echo.New()


	e.GET("/songRank",handleScrape)


	e.Logger.Fatal(e.Start(":8000"))
}