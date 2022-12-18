package parse

import (
	"crawl/engine"
	"log"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^<]+)</a>`

func ParseContent(content []byte) engine.ParseResult {
	// <a href="/tag/小说" class="tag">小说</a>
	//								匹配除了 " 之外的 可以理解为都要，直到结尾 +代表匹配一次以上
	re := regexp.MustCompile(regexpStr)
	result := engine.ParseResult{}
	match := re.FindAllSubmatch(content, -1)
	// FindAllSubmatch 返回个三维数组，第一个是整体，后面 2 个是子查询，如果只有一个括号不能索引 2，所以推断子查询个数为 0 - 2
	for _, m := range match {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:       "https://book.douban.com/" + string(m[1]),
			ParseFunc: engine.NilParser,
		})
		log.Printf("fetch url: %s", "https://book.doubancom/" + string(m[1]))

	}

	return result

}
