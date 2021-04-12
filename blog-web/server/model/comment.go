package model

// Comment 评论结构
type Comment struct {
	baseEntity
	ArticleId string
	Pid       string
	Nickname  string
	Email     string
	Content   string
}

func (_ *Comment) TableName() string {
	return "tb_comment"
}
