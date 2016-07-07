package userprofile

import (
    "regexp"
    "strings"
    "github.com/ip4368/go-password"
)

func ValidateEmail(email string) bool {
    loweredEmail := strings.ToLower(email)
    pattern := "^[a-z0-9.%+-]+@([a-z0-9-].)+[a-z]{2,4}$"
    compiledp, _ := regexp.Compile(pattern)
    matched := compiledp.MatchString(loweredEmail)
    return matched
}

func ValidateUsername(username string) bool {
    pattern := "^[a-zA-Z0-9.]{2,12}$"
    compiledp, _ := regexp.Compile(pattern)
    matched := compiledp.MatchString(username)
    return matched
}

func ValidateName(name string) bool {
    length := len(name)
    if length <= 2 { return false }
    return true
}

func ProcessNewUser(username string, email string, pass string) (string, string, bool) {
    if !ValidateUsername(username) { return "", "", false }
    if !ValidateEmail(email) { return "", "", false }
    hashed, salt, valid := password.HashAutoSalt(pass)
    if !valid { return "", "", false }
    return hashed, salt, true
}
