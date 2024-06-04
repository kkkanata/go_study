package main

import "fmt"

type Blog struct {
	title   string
	content string
  }

/*func GetFullArticle(b Blog) string {
	return b.title + "\n" + "------------" + "\n" + b.content
  }
  */
  //ポインタレシーバでないため値の変更ができない
  func (b Blog) ValueChange2(s string) {
	b.title += s
	b.content += s
	return
  }

  //ポインタレシーバであるため値の変更が可能
  func (b *Blog) ValueChange(s string) {
	b.title += s
	b.content += s
	return
  }

  func (b Blog) GetFullArticle() string {
	return b.title + "\n" + "------------" + "\n" + b.content
  }

func main() {
	var b Blog
	b.title = "title!"
	b.content = "content!"
	s := "happy!"
	b.ValueChange(s)
	fmt.Println(b.GetFullArticle())
}
