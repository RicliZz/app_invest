package startUpModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/RicliZz/app_invest/pkg"
	"time"
)

type StartUp struct {
	ID     int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID int `gorm:"column:userid" json:"userid"`

	//Информация о стартапе
	Title             string              `gorm:"column:title"             json:"title" validate:"required"`
	Topic             string              `gorm:"column:topic"             json:"topic" validate:"required"`
	Idea              string              `gorm:"column:idea"              json:"idea" validate:"required"`
	Strategy          string              `gorm:"column:strategy"          json:"strategy"`
	HistoryOfCreation string              `gorm:"column:historyofcreation" json:"historyOfCreation"`
	Status            enums.StartUpStatus `gorm:"column:status"            json:"status"`
	Stage             enums.StartUpStage  `gorm:"column:stage"             json:"stage"`
	//цель по финансированию, сколько нужно собрать
	FundingGoal float64 `gorm:"column:fundinggoal"       json:"fundingGoal" validate:"required"`
	//предлагаемый процент инвестору
	OfferedPercent float64 `gorm:"column:offeredpercent"    json:"offeredPercent" validate:"required"`

	//Информация о создателе
	FounderFullName string         `gorm:"column:founderfullname" json:"founderFullName"`
	FounderEmail    string         `gorm:"column:founderemail"    json:"founderEmail"`
	FounderSocials  FounderSocials `gorm:"column:foundersocials;type:jsonb"  json:"founderSocials"`

	//Информация о дате создания и апдейта
	CreatedAt time.Time `gorm:"column:createdat" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedat" json:"updatedAt"`
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

func (StartUp) TableName() string {
	return "startup"
}
