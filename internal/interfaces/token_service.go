package interfaces

type TokenService interface {
	Decrypt(token string) (userId uint, err error)
	Encrypt(userID uint) (token string, err error)
}
