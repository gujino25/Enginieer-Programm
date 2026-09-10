package domain

import "errors"

var (
	ErrProjectNotFound      = errors.New("project not found")
	ErrProjectAlreadyExists = errors.New("project already exist")
	ErrMediumInvailid       = errors.New("medium is invailid")
	ErrSystemAlreadyExists  = errors.New("system already exist")
)
