package main

import (
	"io/ioutil"
	"strings"
)

var artCount = 0

type Blog struct {
	articles []string
}

// separation of concerns
// antipattern is God Object like object who is doing everyting
// like keeping state of the blog, reading a file, writing to file...
func (b *Blog) AddArticle(art string) int {
	artCount++
	b.articles = append(b.articles, art)
	return artCount
}

func (b *Blog) RemoveArticle(index int) int {
	artCount--
	b.articles = append(b.articles[:index], b.articles[index+1:]...)
	return artCount
}

// separation of concern
type Persistence struct {
}

func (*Persistence) SaveToFile(b Blog) {
	_ = ioutil.WriteFile("filename",
		[]byte(strings.Join(b.articles, " - ")), 0644)
}

func main() {

}
