package middlewares

type tokenService interface {
	Decrypt(token string) (userID uint, err error)
}
