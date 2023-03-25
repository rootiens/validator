package idcard

import (
	"errors"
	"strconv"
	"strings"

	"github.com/rootiens/validator/helpers"
)

func IsValid(id string) (bool, error) {
	helpers.PerToEng(&id)

	idLen := len(id)

	if idLen < 8 || idLen > 10 {
		return false, errors.New("Invalid length. National ID must be between 8 and 10 characters.")
	}

	// add leading 0 to make it 10 characters
	if idLen >= 8 && idLen < 10 {
		id = strings.Repeat("0", 10-idLen) + id
	}

	sum := calculateSum(id)

	sum %= 11
	controlCode, _ := strconv.Atoi(string(id[9]))

	if (sum < 2 && sum != controlCode) || (sum >= 2 && (controlCode != 11-sum)) {
		return false, errors.New("Invalid ID")
	}

	return true, nil
}

func calculateSum(id string) (sum int) {
	for i := 10; i > 1; i-- {
		int1, _ := strconv.Atoi(string(id[10-i]))
		sum += int1 * i
	}
	return sum
}
