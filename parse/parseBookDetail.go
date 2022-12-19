package parse

import (
	"crawl/engine"
	"crawl/model"
	"regexp"
)

var authRe = regexp.MustCompile(`<span class="pl"> 作者</span>[\s\S]+?<a.*?>([^<]+)</a>[\s\S]+?<br>`)

// <span class="pl">出版社:</span><a href="https://book.douban.com/press/2146">湖南文艺出版社</a><br>
var publisherRe = regexp.MustCompile(`<span class="pl">出版社:</span>[\s\S]+?<a.*?>([^!]+?)</a>[\s\S]+?<br>`)

// <span class="pl">页数:</span> "224"<br>
var pagesRe = regexp.MustCompile(`<span class="pl">页数:</span>([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var descRe = regexp.MustCompile(`<div class="intro">[\s\S]+?<p>(.*?)</p>`)

func ParseBookDetail(content []byte, bookName string) engine.ParseResult {
	bookdetail := model.BookDetail{}
	// 传参测试，防止那些，进入页面前有的数据，但是进入页面后没有的，比如一些描述摘要啥的
	bookdetail.Name = bookName
	bookdetail.Author = ExtraString(content, authRe)
	bookdetail.Publisher = ExtraString(content, publisherRe)
	bookdetail.BookPages = ExtraString(content, pagesRe)
	bookdetail.Price = ExtraString(content, priceRe)
	bookdetail.Score = ExtraString(content, scoreRe)
	bookdetail.Desc = ExtraString(content, descRe)
	result := engine.ParseResult{
		Items: []interface{}{bookdetail},
	}
	return result
}
func ExtraString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""

}
