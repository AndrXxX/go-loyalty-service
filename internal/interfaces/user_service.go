package interfaces

import (
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type UserService interface {
	Find(login string) *ormmodels.User
	Create(u *ormmodels.User) (*ormmodels.User, error)
}
