package rooms

import "errors"

var (
	ERR_ROOMS_CREATE_ROOM_INVALID_OPTIONS = errors.New("invalid options")
	ERR_ROOMS_NIL_OPTIONS                 = errors.New("option is nil")
	ERR_ROOMS_OPERATION_FAILED            = errors.New("operation failed")
)
