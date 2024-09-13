package userdtos

type CreateUserRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
}
