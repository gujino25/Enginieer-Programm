package domain

import "errors"

var (
	ErrProjectNotFound     = errors.New("project not found")
	ErrProjectAlreadyExist = errors.New("project already exist")
)
