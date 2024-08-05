package utils

import _errors "github.com/Many-Men/crowdfund_backend/errors"

func MapErrorToStatusCode(err error) int {
	switch err.(type) {
	case *_errors.NotFoundError:
		return 404
	case *_errors.UnauthorizedError:
		return 401
	case *_errors.BadRequestError:
		return 400
	case *_errors.InternalServerError:
		return 500
	default:
		return 500
	}
}
