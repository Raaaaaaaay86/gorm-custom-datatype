package datatype

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type Owner []string

func (o *Owner) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}

	*o = strings.Split(string(bytes), ",")

	return nil
}

func (o Owner) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}

	return strings.Join(o, ","), nil
}
