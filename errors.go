package redis

import (
	"strings"

	"github.com/juju/errors"
)

var (
	// ErrNoSuchEntity is returned from a Get operation when there is not a model
	// that matches the query
	ErrNoSuchEntity = errors.New("database: no such entity")
)

type MultiError []error

func (merr MultiError) Error() string {
	var msg []string
	for _, err := range merr {
		if err == nil {
			msg = append(msg, "<nil>")
		} else {
			msg = append(msg, err.Error())
		}
	}

	return strings.Join(msg, "; ")
}

func (merr MultiError) HasError() bool {
	for _, err := range merr {
		if err != nil {
			return true
		}
	}

	return false
}