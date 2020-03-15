package serializers

// UserSerializer serializes users struct to json
type UserSerializer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// GetUserSerializer generates serializer for user model
func GetUserSerializer() UserSerializer {
	return UserSerializer{}
}
