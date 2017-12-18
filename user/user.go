package user

// User struct
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

// PrivateUserDetails struct
type PrivateUserDetails struct {
	ID       int64
	Password string
	Salt     string
}
