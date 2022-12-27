package trouble

import (
	"github.com/teguh-satriya/privy-go/library/troublemaker"
	"google.golang.org/grpc/codes"
)

var tm = troublemaker.NewTroubleMaker(
	troublemaker.WithCodes(
		map[string]codes.Code{
			REASON_INTERNAL_SERVER_ERROR: codes.Internal,
			REASON_CAKE_NOT_FOUND:        codes.NotFound,
		},
	),
	troublemaker.WithMessages(
		map[string]string{
			REASON_INTERNAL_SERVER_ERROR: "Internal server error has occured. Please report this to back-end developer.",
			REASON_CAKE_NOT_FOUND:        "Cake Data Not Found! Please Pass correct ID",
		},
	),
)

var (
	INTERNAL_SERVER_ERROR = tm.NewTrouble(REASON_INTERNAL_SERVER_ERROR)
	CAKE_NOT_FOUND        = tm.NewTrouble(REASON_CAKE_NOT_FOUND)
)
