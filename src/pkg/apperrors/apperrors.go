package apperrors

import "errors"

//variables defining common error in our system
var (
	//ErrNotFound is an error returned when the resource asked isn't found
	ErrNotFound = errors.New("not_found")

	//ErrIllegalOperation is an error returned when the
	//operation asked isn´t available
	ErrIllegalOperation = errors.New("illegal_operation")
  
	//ErrSql is an error returned from repositorio
	ErrSql         		= errors.New("sql")
	//ErrInvalidInput is an error returned when some function is called with
	//incorrect arguments or not enought
	ErrInvalidInput = errors.New("invalid_input")

	//ErrInternal is an error returned when the system fails and
	//no reason is known
	ErrInternal = errors.New("internal")
)
