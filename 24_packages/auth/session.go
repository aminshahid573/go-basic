package auth

//private
func extractSession() string {
	return "LoogedIn"
}

//public due to capital letter at first
func GetSession() string {
	return extractSession()
}
