package errlib

import (
	"errors"
)

// ErrInternal is error to be matched with 500 http code.
var ErrInternal = errors.New("some internal error happened")

// ErrResourceAlreadyExists is error to be matched with 409 http code.
var ErrResourceAlreadyExists = errors.New("resource already exists")

// ErrInternal is error to be matched with 404 http code.
var ErrNotFound = errors.New("resource not found")

var ErrBadRequest = errors.New("bad data input")

var ErrAccessDenied = errors.New("you cannot perform this operation")
