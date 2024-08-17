package middlewares

type tokenService interface {
	Decrypt(token string) (userId uint, err error)
}
