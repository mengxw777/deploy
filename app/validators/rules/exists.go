package rules

import (
	"deploy/database"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

var db = database.DB

func Exists(fl validator.FieldLevel) bool {
	var count int
	param := strings.Split(fl.Param(), `:`)
	table := param[0]
	field := param[1]

	if err := db.Table(table).Where(fmt.Sprintf("%s = ?", field), fmt.Sprintf("%v", fl.Field())).Count(&count).Error; err != nil {
		return false
	}

	if count > 0 {
		return true
	}

	return false
}
