package requests

import enums "github.com/RicliZz/app_invest/pkg"

type StartupRequest struct {
	Title             *string              `json:"title"` // required
	Topic             *enums.StartUpTopic  `json:"topic"` // required
	Idea              *string              `json:"idea"`
	Strategy          *string              `json:"strategy"`
	HistoryOfCreation *string              `json:"historyOfCreation"`
	Status            *enums.StartUpStatus `json:"status"` // required
	Stage             *enums.StartUpStage  `json:"stage"`

	FundingGoal    *float64 `json:"fundingGoal"`    // required
	OfferedPercent *float64 `json:"offeredPercent"` // required
}
