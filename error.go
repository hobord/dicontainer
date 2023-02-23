package gdic

import "errors"

var (

	ErrFactoryNotFound = errors.New("factory not found")
	ErrFactoryExist    = errors.New("factory already exist")
	ErrNilFactory      = errors.New("factory is nil")
	ErrInstanceExist   = errors.New("instance already exist")
)
