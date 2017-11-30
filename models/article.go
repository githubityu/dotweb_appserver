package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/initialization"
)

// Article 文章模型
type Article struct {
	Id           int64  `json:"id"`
	Userid       int64  `json:"userid"` //作者id
	Picture      string `json:"picture"`
	Title        string `json:"title"`            //标题
	Content      string `json:"content" `         //内容
	Thumbscount  int    `json:"thumbscount"`      //点赞数
	Viewcount    int    `json:"viewcount"`        //阅读次数
	Commentcount int    `json:"commentcount"`     //评论次数
	Updatedat    int64  `json:"-" xorm:"updated"` //更新时间
	ReleaseStr   string `json:"releasestr" xorm:"created"`    //发布的时间
	Year         int    `json:"year"`             //发布的年
	Month        int    `json:"month"`            //发布的日期
	Day          int    `json:"day"`              //发布的天
	ReleaseTime  int64  `json:"-"`                //发布时间
	Copyright    string `json:"Copyright"`        //文章底部版权
}

func init() {
	orm.GetEngine().CreateTables(new(Article))
}

 func (*Article) TableName() string {
 	return "article"
 }

func (article *Article) Inset() (newId int64, err error) {
	engine := orm.GetEngine()

	newId, err = engine.Insert(article)
	return
}

func (article *Article) Delete() (delId int64, err error) {
	engine := orm.GetEngine()

	delId, err = engine.Delete(article)
	return
}

func (article *Article) Update() (updId int64, err error) {
	engine := orm.GetEngine()

	updId, err = engine.Id(article.Id).Update(article)
	return
}

// UpdateViewCount 专门更新viewCount
func (article *Article) UpdateViewCount() (err error) {
	engine := orm.GetEngine()

	_, err = engine.Id(article.Id).Update(article)
	return
}

func (article *Article) GetOne(id int64) (ok bool, err error) {
	engine := orm.GetEngine()
	ok, err = engine.Id(id).Get(article)
	return
}

// TopN 获取文章列表
func (article *Article) TopN(n int) (articles []*Article, err error) {
	engine := orm.GetEngine()
	err = engine.Desc("id").Limit(n).Find(&articles)
	return
}

// PageUser 分页的文章数
func (article *Article) PageArticle(offset, limit int) (articles []*Article, err error) {
	engine := orm.GetEngine()

	err = engine.Limit(limit, offset).Find(&articles)
	return
}

// PageArticleCid 分类id分页
// func (article *Article) PageArticleCid(offset, limit int,cid int64) (articles []*Article, err error) {
// 	engine := orm.GetEngine()
// 	err = engine.Where("cid=?",cid).Limit(limit, offset).Find(&articles)
// 	return
// }

// GetArticleByTag 通过标签列表
// func (article *Article) GetArticleByTag(ctid int64) (articles []*Article, err error) {
// 	// err=engine.Join("INNER","tag_article_rel","tag_article_rel").Find(&articles)
// 	err = engine.Where("Ctid =?", ctid).Join("INNER", "tag_article_rel", "tag_article_rel.cid=article.id").Desc("id").Find(&articles)

// 	return
// }

// GetArticleByTag 通过标签列表
// func (article *Article) PageArticleByTag(ctid, offset, limit int64) (articles []*Article, err error) {
// 	// err=engine.Join("INNER","tag_article_rel","tag_article_rel").Find(&articles)
// 	err = engine.Where("Ctid =?", ctid).Join("INNER", "tag_article_rel", "tag_article_rel.cid=article.id").Desc("id").Limit(int(limit), int(offset)).Find(&articles)
// 	return
// }

// // Count 统计所有文章的数量
func (article *Article) Total() (count int64, err error) {
	engine := orm.GetEngine()

	count, err = engine.Count(article)
	return
}

func (article *Article) Hot() (articles []*Article, err error) {
	engine := orm.GetEngine()
	err = engine.Desc("viewcount").Limit(10).Find(&articles)
	if err != nil {
		return
	}
	return
}

//通过
