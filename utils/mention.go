package utils

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
)

func Mention(user types.User) string {
	return fmt.Sprintf("<a href='tg://user?id=%d'>%s</a>", user.ID, FullName(user))
}
