package cfilters

import (
	"github.com/Squirrel-Network/gobotapi/filters"
	"github.com/Squirrel-Network/gobotapi/types"
	"reflect"
)

func Sticker() filters.FilterOperand {
	return func(values *filters.DataFilter) bool {
		if reflect.TypeOf(values.RawUpdate).String() == "types.Message" {
			message := values.RawUpdate.(types.Message)
			return message.Sticker != nil
		}
		return false
	}
}
