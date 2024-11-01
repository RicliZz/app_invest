package userModel

type UserDetails struct {
	ID      int     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID  int64   `gorm:"column:user_id" json:"userId"`
	Balance float64 `gorm:"column:balance" json:"balance"`
}
