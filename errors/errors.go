package _errors

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}
