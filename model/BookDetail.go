package model

type BookDetail struct {
	Author    string
	Publisher string
	BookPages string
	Price     string
	Score     string
	Desc      string
}

func (b BookDetail) String() string {
	return "作者:" + b.Author + " 出版社:" + b.Publisher + " 书籍页数:" + b.BookPages + " 价格:" + b.Price + " 得分：" + b.Score + "\n简介:" + b.Desc
}
