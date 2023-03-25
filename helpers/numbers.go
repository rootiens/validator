package helpers

import (
	"strings"
)

func PerToEng(number *string) {
	*number = strings.ReplaceAll(*number, "۱", "1")
	*number = strings.ReplaceAll(*number, "۲", "2")
	*number = strings.ReplaceAll(*number, "۳", "3")
	*number = strings.ReplaceAll(*number, "۴", "4")
	*number = strings.ReplaceAll(*number, "۵", "5")
	*number = strings.ReplaceAll(*number, "۶", "6")
	*number = strings.ReplaceAll(*number, "۷", "7")
	*number = strings.ReplaceAll(*number, "۸", "8")
	*number = strings.ReplaceAll(*number, "۹", "9")
	*number = strings.ReplaceAll(*number, "۰", "0")
}
