package models

// gorm:"index" 用于声明这个字段为索引,Tag Tag是嵌套的tab
type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	// Preload就是一个预加载器，它会执行两条 SQL，分别是SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);，那么在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，会特别方便，并且避免了循环查询
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	//Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id= ?", id).Find(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func AddArticle(data map[string]interface{}) bool {
	//这里是将"data["tag_id"]"的值转换为整数类型（int
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}

func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}

// func (tag *Article) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("CreatedOn", time.Now().Unix())
// 	return nil
// }

// func (tag *Article) BeforeUpdate(scope *gorm.Scope) error {
// 	scope.SetColumn("ModifiedOn", time.Now().Unix())
// 	return nil
// }
