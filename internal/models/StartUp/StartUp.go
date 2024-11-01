package startUpModel

import "time"

type StartUpStatus int
type StartUpStage int

const (
	CLOSE StartUpStatus = iota
	ACTIVE
)

const (
	IDEA StartUpStage = iota
	PROTOTYPE
	PRODUCT
)

type StartUp struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`

	//Информация о стартапе
	Title             string        `gorm:"column:title"             json:"title"`
	Topic             string        `gorm:"column:topic"             json:"topic"`
	Idea              string        `gorm:"column:idea"              json:"idea"`
	Strategy          string        `gorm:"column:strategy"          json:"strategy"`
	HistoryOfCreation string        `gorm:"column:historyOfCreation" json:"historyOfCreation"`
	Status            StartUpStatus `gorm:"column:status"            json:"status"`
	Stage             StartUpStage  `gorm:"column:stage"             json:"stage"`
	//цель по финансированию, сколько нужно собрать
	FundingGoal float64 `gorm:"column:fundingGoal"       json:"fundingGoal"`
	//предлагаемый процент инвестору
	OfferedPercent float64 `gorm:"column:offeredPercent"    json:"offeredPercent"`

	//Информация о создателе
	FounderFullName string         `gorm:"column:founderFullName" json:"founderFullName"`
	FounderEmail    string         `gorm:"column:founderEmail"    json:"founderEmail"`
	FounderSocials  FounderSocials `gorm:"column:founderSocials"  json:"founderSocials"`

	//Ссылки

	//Информация о дате создания и апдейта
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

type FounderSocials struct {
	WebsiteLink   string `gorm:"column:website"       json:"website"`
	VkLink        string `gorm:"column:vkLink"        json:"vkLink"`
	TelegramLink  string `gorm:"column:telegramLink"  json:"telegramLink"`
	WhatsUpLink   string `gorm:"column:whatsUpLink"   json:"whatsUpLink"`
	InstagramLink string `gorm:"column:instagramLink" json:"instagramLink"`
}
