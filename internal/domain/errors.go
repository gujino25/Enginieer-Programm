package domain

import "errors"

var (
	ErrProjectNotFound      = errors.New("project not found")
	ErrProjectAlreadyExists = errors.New("project already exist")
	ErrMediumInvalid        = errors.New("medium is invailid")
	ErrSystemAlreadyExists  = errors.New("system already exist")
	ErrSystemNotFound       = errors.New("system not found")
	ErrShapeInvalid         = errors.New("shape is invalid")
	ErrGeometryInvalid      = errors.New("geometry is invalid")
	ErrLengthInvalid        = errors.New("lenght is invalid")
	ErrSegmentAlreadyExists = errors.New("segment already exist")
	ErrSegmentNotFound      = errors.New("segemnt not found")
)
