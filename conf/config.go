package conf

var (
	ArticleFormat = "hasharticle:%d"  //redis 文章key
	CatetagFormt =  "catetag:%d" //分类key
	ArticletocatetagFormt = "catetag:*:article:%d" //通过文章查询所有的分类和标签
	CatetagtoarticleFormt = "catetag:%d:article:*"//通过分类查找文章
	WriteArticleAngTag = "catetag:%d:article:%d" //写入key
	ArticlMail   = "article:%d:mail:*"
	ArticleZset = "zsetarticle" //文章zet 用来分页

	DataSource  = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4"


)
