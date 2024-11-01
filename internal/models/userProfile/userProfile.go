package userProfileModel

type UserProfile struct {
	FirstName  string  `gorm:"column:firstname"         json:"firstName"`
	LastName   string  `gorm:"column:lastname"          json:"lastName"`
	Patronymic string  `gorm:"column:patronymic"        json:"patronymic"`
	Email      string  `gorm:"column:email"             json:"email"`
	Balance    float64 `gorm:"column:balance"           json:"balance"`
}
