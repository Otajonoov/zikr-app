package utils

import "regexp"

func IsValidUsername(username string) bool {
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)

	return usernameRegex.MatchString(username)
}
