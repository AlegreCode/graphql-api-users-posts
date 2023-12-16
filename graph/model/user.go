package model

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Lastname string  `json:"lastname"`
	Posts    []*Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"posts"`
}
