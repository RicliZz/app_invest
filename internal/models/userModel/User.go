package userModel

type User struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string `gorm:"column:firstname"         json:"firstName"`
	LastName  string `gorm:"column:lastname"          json:"lastName"`
	Email     string `gorm:"column:email"             json:"email"`
	Password  string `gorm:"column:password"          json:"password"`
}
