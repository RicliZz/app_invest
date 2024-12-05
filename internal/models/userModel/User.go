package userModel

type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName    string `gorm:"column:firstName"         json:"firstName"`
	LastName     string `gorm:"column:lastName"          json:"lastName"`
	MiddleName   string `gorm:"column:middleName"        json:"middleName"`
	Email        string `gorm:"column:email"             json:"email"`
	Role         string `json:"role" gorm:"default:user" gorm:"omitempty"`
	EmailConfirm bool   `gorm:"column:emailConfirm"      json:"emailConfirm"`
	Password     string `gorm:"column:password"          json:"password"`
	UserDetails  UserDetails
}
