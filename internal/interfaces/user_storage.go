package interfaces

import (
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type UserStorage interface {
	Find(login string) *ormmodels.User
	Create(u *ormmodels.User) (*ormmodels.User, error)
}
