package user

type NewUser struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Roles           []string `json:"roles"`
	Password        string   `json:"password"`
	PasswordConfirm string   `json:"password_confirm"`
}

type UserUpdate struct {
	Name            *string  `json:"name"`
	Email           *string  `json:"email"`
	Roles           []string `json:"roles"`
	Password        *string  `json:"password"`
	PasswordConfirm *string  `json:"password_confirm"`
}

type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}
