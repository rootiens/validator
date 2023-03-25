package phone

import (
	"regexp"

	"github.com/rootiens/validator/helpers"
)

func IsValid(number string) bool {
    helpers.PerToEng(&number)
	if match, _ := regexp.MatchString(`^(?:98|\+98|0098|0)?9[0-9]{9}$`, number); !match {
		return false
	}
    return true
}
