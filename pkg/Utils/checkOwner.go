package Utils

func ThisOwner(pathID, ID int64) bool {
	if pathID == ID {
		return true
	}
	return false
}
