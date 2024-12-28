package userModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
)

type User struct {
	ID           int                    `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName    string                 `gorm:"column:firstName"         json:"firstName"`
	LastName     string                 `gorm:"column:lastName"          json:"lastName"`
	MiddleName   string                 `gorm:"column:middleName"        json:"middleName"`
	Email        string                 `gorm:"column:email"             json:"email"`
	AboutMe      string                 `gorm:"column:aboutMe"           json:"aboutMe"`
	Role         string                 `json:"role" gorm:"default:user" gorm:"omitempty"`
	EmailConfirm bool                   `gorm:"column:emailConfirm"      json:"emailConfirm"`
	Password     string                 `gorm:"column:password"          json:"password"`
	Confirmed    bool                   `gorm:"column:confirmed"         json:"confirmed"`
	UserDetails  *UserDetails           `gorm:"foreignKey:UserID" json:"userDetails"`
	Startups     []startUpModel.StartUp `gorm:"foreignKey:UserID" json:"startUps"`
	Socials      FounderSocials         `gorm:"foreignKey:UserID" json:"socials"`
}

type FounderSocials struct {
	WebsiteLink   string `gorm:"column:website"       json:"website"`
	VkLink        string `gorm:"column:vklink"        json:"vkLink"`
	TelegramLink  string `gorm:"column:telegramlink"  json:"telegramLink"`
	WhatsUpLink   string `gorm:"column:whatsuplink"   json:"whatsUpLink"`
	InstagramLink string `gorm:"column:instagramlink" json:"instagramLink"`
}

func (f FounderSocials) Value() (driver.Value, error) {
	return json.Marshal(f)
}

func (f *FounderSocials) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan FounderSocials")
	}
	return json.Unmarshal(bytes, f)
}
