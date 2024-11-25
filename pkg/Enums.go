package enums

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
