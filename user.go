package testmedods

type User struct {
	Id       int    `json:"-"`
	Username string `json:username`
	Email    string `json:email binding:"required"`
	Password string `json:password binding:"required"`
}
