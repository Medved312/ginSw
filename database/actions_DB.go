package database

import (
	"errors"
	"fmt"
)

func Add(bean interface{}) error {
	if !GetDB().NewRecord(bean) {
		return errors.New("unable to create")
	}
	fmt.Println(bean)

	return GetDB().Omit("id").Create(bean).Error
}
