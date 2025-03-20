package main

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// IsDuplicatedKeyError This is not necessary if you set TranslateError to true in gorm.Config
func IsDuplicatedKeyError(err error) bool {
	var errT *mysql.MySQLError
	if errors.As(err, &errT) && errT.Number == 1062 {
		return true
	}

	return false
}
