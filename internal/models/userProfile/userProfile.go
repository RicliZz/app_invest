package userProfileModel

type UserProfile struct {
	FirstName  string  `gorm:"column:firstname"         json:"firstName"`
	LastName   string  `gorm:"column:lastname"          json:"lastName"`
	MiddleName string  `gorm:"column:MiddleName"        json:"middleName"`
	Email      string  `gorm:"column:email"             json:"email"`
	Balance    float64 `gorm:"column:balance"           json:"balance"`
}
