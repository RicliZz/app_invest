package UserRepository

func (rep *UserRepositoryImpl) DeleteUser(userID int64) error {
	user, err := rep.GetUserById(userID)
	if err != nil {
		return err
	}
	result := rep.db.Delete(&user)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}
