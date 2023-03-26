package phone

import (
	"strconv"

	"github.com/rootiens/validator/helpers"
)

func IsValid(number string) bool {
	helpers.PerToEng(&number)
	if len(number) == 12 && string(number[0:3]) == "+98" {
		number = string(number[3:])
	} else if len(number) == 10 {
		number = "0" + number
	} else if len(number) == 11 && number[0] == '0' {
		// do nothing
	} else {
		return false
	}
	_, err := strconv.Atoi(number)
	if err != nil {
		return false
	}
	if string(number[0:2]) != "09" {
		return false
	}
	return true
}
