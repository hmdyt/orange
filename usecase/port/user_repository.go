package port

import "github.com/hmdyt/orange/domain"

type UserRepository interface {
	Save(user *domain.User) error
}
