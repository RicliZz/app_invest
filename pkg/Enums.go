package enums

type StartUpStatus int
type StartUpStage int
type StartUpTopic int

const (
	CLOSE StartUpStatus = iota
	ACTIVE
)

const (
	IDEA StartUpStage = iota
	PROTOTYPE
	PRODUCT
)

const (
	SPORT StartUpTopic = iota
	TEACHING
	CAR
	//все категории обсуждать с парнями
)
