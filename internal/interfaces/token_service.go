package interfaces

type TokenService interface {
	Decrypt(token string) (userID uint, err error)
	Encrypt(userID uint) (token string, err error)
}
