package resp

import "errors"

var (
	AuthorizationExpired = errors.New("token expired")
)
