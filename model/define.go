package model

// Users ...
type Users struct {
	Id            uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	Phone         int    `gorm:"column:phone" db:"phone" json:"phone" form:"phone"`
	Nickname      string `gorm:"column:nickname" db:"nickname" json:"nickname" form:"nickname"`
	Password      string `gorm:"column:password" db:"password" json:"password" form:"password"`
	AvatarUrl     string `gorm:"column:avatar_url" db:"avatar_url" json:"avatar_url" form:"avatar_url"`
	Status        int    `gorm:"column:status" db:"status" json:"status" form:"status"`
	MarkedCount   uint   `gorm:"column:marked_count" db:"marked_count" json:"marked_count" form:"marked_count"`
	QuestionCount uint   `gorm:"column:question_count" db:"question_count" json:"question_count" form:"question_count"`
	AnswerCount   uint   `gorm:"column:answer_count" db:"answer_count" json:"answer_count" form:"answer_count"`
	CreateTime    string `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	ModifyTime    string `gorm:"column:modify_time" db:"modify_time" json:"modify_time" form:"modify_time"`
}

// Questions ...
type Questions struct {
	Id           uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId       uint   `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	Title        string `gorm:"column:title" db:"title" json:"title" form:"title"`
	Status       int    `gorm:"column:status" db:"status" json:"status" form:"status"`
	Detail       string `gorm:"column:detail" db:"detail" json:"detail" form:"detail"`
	AnswerCount  uint   `gorm:"column:answer_count" db:"answer_count" json:"answer_count" form:"answer_count"`
	CommentCount uint   `gorm:"column:comment_count" db:"comment_count" json:"comment_count" form:"comment_count"`
	CreateTime   string `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	ModifyTime   string `gorm:"column:modify_time" db:"modify_time" json:"modify_time" form:"modify_time"`
}

// Answers ...
type Answers struct {
	Id           uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	QuestionId   uint   `gorm:"column:question_id" db:"question_id" json:"question_id" form:"question_id"`
	UserId       uint   `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	Content      string `gorm:"column:content" db:"content" json:"content" form:"content"`
	MarkedCount  uint   `gorm:"column:marked_count" db:"marked_count" json:"marked_count" form:"marked_count"`
	CommentCount uint   `gorm:"column:comment_count" db:"comment_count" json:"comment_count" form:"comment_count"`
	Status       int    `gorm:"column:status" db:"status" json:"status" form:"status"`
	CreateTime   string `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	ModifyTime   string `gorm:"column:modify_time" db:"modify_time" json:"modify_time" form:"modify_time"`
}

// AnswerComments ...
type AnswerComments struct {
	Id         uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId     uint   `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	AnswerId   uint   `gorm:"column:answer_id" db:"answer_id" json:"answer_id" form:"answer_id"`
	CreateTime string `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	ModifyTime string `gorm:"column:modify_time" db:"modify_time" json:"modify_time" form:"modify_time"`
	Content    string `gorm:"column:content" db:"content" json:"content" form:"content"`
}

// MemberFollowers ...
type MemberFollowers struct {
	MemberId   uint   `gorm:"column:member_id" db:"member_id" json:"member_id" form:"member_id"`
	FollowerId uint   `gorm:"column:follower_id" db:"follower_id" json:"follower_id" form:"follower_id"`
	CreateTime string `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	ModifyTime string `gorm:"column:modify_time" db:"modify_time" json:"modify_time" form:"modify_time"`
}
