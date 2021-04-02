package util

import (
	"regexp"
	"strings"
)

func ConcatString(separator string, strs ...string) string {
	var res []string
	for i := range strs {
		if s := strings.TrimSpace(strs[i]); s != "" {
			res = append(res, s)
		}
	}

	return strings.Join(res, separator)
}

func ValidateEmail(email string) bool {
	var emailPattern = regexp.MustCompile(`^[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[a-zA-Z0-9](?:[\w-]*[\w])?$`)
	return emailPattern.MatchString(email)
}
