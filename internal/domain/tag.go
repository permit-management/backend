package domain

// Model for Tag.
type TagModel struct {
	Model
	// ID   string `gorm:"column:tag_id;->;primaryKey" json:"tag_id,readonly"`
	Name string `json:"name"`
}

func (TagModel) TableName() string {
	return "blog_tag"
}
