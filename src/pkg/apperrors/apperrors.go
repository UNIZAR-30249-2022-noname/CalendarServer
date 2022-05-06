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
	ErrSql = errors.New("sql")
	//ErrInvalidInput is an error returned when some function is called with
	//incorrect arguments or not enought
	ErrInvalidInput = errors.New("invalid_input")

	//ErrInternal is an error returned when the system fails and
	//no reason is known
	ErrInternal = errors.New("internal")

	ErrNoRowsAffected = errors.New("no_rows_affected")

	//ErrInvalidKind is an error returned when the Kind is incompatible
	//with the Group and Week
	ErrInvalidKind = errors.New("invalid_subject_kind")

	//ErrToDo is an error returned when the functionality isn't made yet
	ErrToDo = errors.New("not_done_yet")

	//ErrConnis an error returned when the connection to AMQP fails
	ErrConn = errors.New("amqp_connection_error")

	//ErrorWrongResponse is returned when the server returned a wrong response
	ErrorWrongResponse = errors.New("wrong response")
)
