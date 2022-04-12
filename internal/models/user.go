package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Email    string
	Password string
}

func (u *User) EncryptPassword() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)

		if err != nil {
			return err
		}

		u.Password = enc
	}

	return nil
}

func encryptString(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
