package auth

// AuthorizedPath Get all path that not validate with JWT
func AuthorizedPath() []string {
	return []string{
		"/api/auth/login",
		"/api/auth/register",
	}
}
