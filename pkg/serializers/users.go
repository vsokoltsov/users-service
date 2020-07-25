package serializers

import (
	"time"

	"github.com/vsokoltsov/users-service/pkg/models"
)

// UserSerializer serializes users struct to json
type UserSerializer struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

// GetUserSerializer generates serializer for user model
func GetUserSerializer(user models.User) UserSerializer {
	return UserSerializer{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
