package startUpModel

import (
	"github.com/RicliZz/app_invest/pkg"
	"time"
)

type StartUp struct {
	ID     int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID int `gorm:"column:userid" json:"userid"`

	//Информация о стартапе
	Title             string              `gorm:"column:title"             json:"title" validate:"required"`
	Topic             enums.StartUpTopic  `gorm:"column:topic"             json:"topic" validate:"required"`
	Idea              string              `gorm:"column:idea"              json:"idea"`
	Strategy          string              `gorm:"column:strategy"          json:"strategy"`
	HistoryOfCreation string              `gorm:"column:historyofcreation" json:"historyOfCreation"`
	Status            enums.StartUpStatus `gorm:"column:status"            json:"status" validate:"required"`
	Stage             enums.StartUpStage  `gorm:"column:stage"             json:"stage"`
	//цель по финансированию, сколько нужно собрать
	FundingGoal float64 `gorm:"column:fundinggoal"       json:"fundingGoal" validate:"required"`
	//предлагаемый процент инвестору
	OfferedPercent float64 `gorm:"column:offeredpercent"    json:"offeredPercent" validate:"required"`

	//Информация о создателе
	FounderFullName string         `gorm:"-" json:"founderFullName"`
	FounderEmail    string         `gorm:"-" json:"founderEmail"`
	FounderSocials  FounderSocials `gorm:"-" json:"founderSocials"`

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

func (StartUp) TableName() string {
	return "startup"
}
