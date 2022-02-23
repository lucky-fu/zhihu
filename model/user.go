package model

// CreateUser ...
func CreateUser(phone int, nickName, passport, avatorURL string) (int, error) {
	var user = &Users{
		Phone:     phone,
		Nickname:  nickName,
		Password:  passport,
		AvatarUrl: avatorURL,
	}

	tx := GetWriteDB()
	err := tx.Table("users").Create(user).Error

	return int(user.Id), err
}
