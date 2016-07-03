package userprofile

import (
    "regexp"
    "strings"
)

func ValidateEmail(email string) bool {
    loweredEmail := strings.ToLower(email)
    pattern := "^[a-z0-9.%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$"
    matched, _ := regexp.MatchString(pattern, loweredEmail)
    return matched
}

func ValidateUsername(username string) bool {
    pattern := "^[a-zA-Z0-9.].{2,12}$"
    matched, _ := regexp.MatchString(pattern, username)
    return matched
}
