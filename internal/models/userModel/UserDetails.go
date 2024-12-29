package userModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type UserDetails struct {
	ID      int     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID  int64   `gorm:"column:user_id" json:"userId"`
	Balance float64 `gorm:"column:balance" json:"balance"`
	Socials Socials `gorm:"column:socials;type:jsonb"  json:"socials"`
}

type Socials struct {
	WebsiteLink   string `gorm:"column:website"       json:"website"`
	VkLink        string `gorm:"column:vklink"        json:"vkLink"`
	TelegramLink  string `gorm:"column:telegramlink"  json:"telegramLink"`
	WhatsUpLink   string `gorm:"column:whatsuplink"   json:"whatsUpLink"`
	InstagramLink string `gorm:"column:instagramlink" json:"instagramLink"`
}

func (f Socials) Value() (driver.Value, error) {
	return json.Marshal(f)
}

func (f *Socials) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan FounderSocials")
	}
	return json.Unmarshal(bytes, f)
}
