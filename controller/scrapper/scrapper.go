package scrapper

import (
	"encoding/json"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// Scrape : to scrape a website texts
func Scrape(url string) []byte {
	collyIns := colly.NewCollector()
	var wordsCountMap map[string]int
	var collyError error

	collyIns.OnHTML("html", func(e *colly.HTMLElement) {
		e.DOM.Find("script,style").Each(func(index int, item *goquery.Selection) {
			item.Remove()
		})
		wordsCountMap = wordCounter(e.DOM.Contents().Text())
	})

	collyIns.OnError(func(r *colly.Response, err error) {
		collyError = err
	})

	collyIns.Visit(url)
	if collyError != nil {
		panic(collyError)
	}
	byteArr, err := json.Marshal(wordsCountMap)
	if err != nil {
		panic(err)
	}

	return byteArr
}

func wordCounter(text string) map[string]int {
	wordsCountMap := make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words {
		trimmedWord := trimWord(word)
		if len(trimmedWord) > 0 {
			wordsCountMap[word] = wordsCountMap[word] + 1
		}
	}
	return wordsCountMap
}

func trimWord(text string) string {
	startIndex := 0
	charArr := []rune(text)
	lastIndex := len(charArr)

	for index, char := range charArr {
		if isAlphaNumeric(char) {
			startIndex = index
			break
		}
	}

	for lastIndex > 0 && !isAlphaNumeric(charArr[lastIndex-1]) {
		lastIndex = lastIndex - 1
	}

	return text[startIndex:lastIndex]
}

func isAlphaNumeric(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsNumber(char)
}
