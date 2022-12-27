package troublemaker

import (
	"google.golang.org/grpc/codes"
)

type Trouble struct {
	Code    codes.Code `json:"code"`
	Reason  string     `json:"reason"`
	Message string     `json:"message"`
}
