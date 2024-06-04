package main

import "fmt"

type Blog struct {
	title   string
	content string
  }

func GetFullArticle(b Blog) string {
	return b.title + "\n" + "------------" + "\n" + b.content
  }

func main() {
	var b Blog
	b.title = "title!"
	b.content = "content!"
	s := GetFullArticle(b)
	fmt.Println(s)
}
