package utils

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
)

func FullName(user types.User) string {
	if len(user.LastName) == 0 {
		return user.FirstName
	}
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}
