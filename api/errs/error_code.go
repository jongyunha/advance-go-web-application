package errs

type ErrorCode int

const (
	// 500
	InternalServerError ErrorCode = 500000

	// 404
	NotFound ErrorCode = 404000

	// 400
	InvalidRequest ErrorCode = 400000

	// 403
	Forbidden ErrorCode = 403000

	// 401
	Unauthorized ErrorCode = 401000
)
