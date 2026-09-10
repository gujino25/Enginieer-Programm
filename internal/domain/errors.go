package domain

import "errors"

var (
	ErrProjectAlreadyExist = errors.New("project already exist")
	ErrProjectNotFound     = errors.New("project not found")
)
