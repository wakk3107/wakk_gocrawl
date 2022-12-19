package parse

import (
	"crawl/engine"
	"regexp"
)

/*
<a href="https://book.douban.com/subject/36104107/" title="长安的荔枝" onclick="moreurl(
	this,{i:'0',query:'',subject_id:'36104107',from:'book_subject_search'})">长安的荔枝 </a>*/
const BookListRe = `<a href="([^"]+)" title="([^"]+)"`

// 获取 tag 页面下的书名
func ParseBookList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(BookListRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	// FindAllSubmatch 返回个三维数组，第一个是整体，后面 2 个是子查询，如果只有一个括号不能索引 2，所以推断子查询个数为 0 - 2
	for _, m := range matches {
		bookName := string(m[2])
		// 结果（书名）
		result.Items = append(result.Items, bookName)
		// 衍生的请求
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseBookDetail(c, bookName)
			},
		})

	}

	return result
}
