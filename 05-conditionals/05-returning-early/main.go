package conditionals

var Password = "current-password"

func ResetPassword(code int) {
	Password = "new-password"
}
