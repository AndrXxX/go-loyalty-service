package interfaces

import (
	"context"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type UserStorage interface {
	Init(ctx context.Context) error
	Find(login string) *ormmodels.User
	Create(u *ormmodels.User) (*ormmodels.User, error)
}
