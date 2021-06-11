package domain

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AuthID    string `json:"auth_id"`
	UserRole  `json:"user_role"`
}

type UserRole string

const (
	UserRoleCustomer UserRole = "customer"
	UserRoleEmployee UserRole = "employee"
	UserRoleAdmin    UserRole = "admin"
)
