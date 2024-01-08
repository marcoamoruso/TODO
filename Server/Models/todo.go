package Models

type TODO struct {
	Id       int    `json:"id" gorm:"primaryKey;type:INT UNSIGNED NOT NULL AUTO_INCREMENT"`
	Title    string `json:"title" gorm:"uniqueIndex:title_deadline_unique;not null;type:VARCHAR(255)" validate:"required"`
	Deadline string `json:"deadline" gorm:"uniqueIndex:title_deadline_unique;not null;type:DATE" validate:"required"`
	Done     bool   `json:"done" gorm:"not null"`
}

// Title and Deadline share uniqueIndex:title_deadline_unique since it doesn't make sense to have 2 TODO elements with the same title and the same deadline

// Return table name
func (todoElement *TODO) TableName() string {
	return "TODO"
}
