package domain

type ArticleReader interface {
	GetRecent(n int) []*Article
	GetById(id string) (*Article, error)
}

type Storage interface {
	ArticleReader
	AddArticles(articles []*Article) error
}
