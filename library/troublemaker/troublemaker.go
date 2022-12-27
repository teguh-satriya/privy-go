package troublemaker

import (
	"fmt"
	"regexp"
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TroubleMaker interface {
	NewTrouble(reason string) error
}

type TroubleMakerImpl struct {
	codes    map[string]codes.Code
	messages map[string]string
}

func (tmi *TroubleMakerImpl) NewTrouble(reason string) error {
	var cd codes.Code
	var msg string
	var ok bool

	if cd, ok = tmi.codes[reason]; !ok {
		panic(fmt.Sprintf("codes for %s does not exist", reason))
	}

	if msg, ok = tmi.messages[reason]; !ok {
		panic(fmt.Sprintf("message for %s does not exist", reason))
	}

	trouble := &Trouble{
		Code:    cd,
		Reason:  reason,
		Message: msg,
	}

	st := status.New(trouble.Code, trouble.Message)
	st, _ = st.WithDetails(&errdetails.ErrorInfo{
		Reason: trouble.Reason,
	})

	return st.Err()
}

type TroubleMakerSetter func(*TroubleMakerImpl)

type ValidationError interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
}

type ValidationMultiError interface {
	Error() string
	AllErrors() []error
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func FromValidationError(err error) error {
	switch err := err.(type) {
	case ValidationError:
		field := ToSnakeCase(err.Field())

		trouble := &Trouble{
			Code:    codes.InvalidArgument,
			Reason:  "BAD_REQUEST",
			Message: fmt.Sprintf("Request parameter is invalid. Please pass valid %s.\n\t%s: %s", field, field, err.Reason()),
		}

		st := status.New(trouble.Code, trouble.Message)
		st, _ = st.WithDetails(&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       field,
					Description: err.Reason(),
				},
			},
		})

		return st.Err()
	case ValidationMultiError:
		var prefix []string
		trouble := &Trouble{
			Code:   codes.InvalidArgument,
			Reason: "BAD_REQUEST",
		}

		var fv []*errdetails.BadRequest_FieldViolation

		for _, e := range err.AllErrors() {

			if err, ok := e.(ValidationError); ok {
				// TODO: Too lazy to map the string
				prefix = append(prefix, fmt.Sprintf("\t%s: %s", ToSnakeCase(err.Field()), err.Reason()))

				fv = append(fv, &errdetails.BadRequest_FieldViolation{
					Field:       ToSnakeCase(err.Field()),
					Description: err.Reason(),
				})
			}
		}

		trouble.Message = fmt.Sprintf("Request parameter is invalid. Please pass valid request parameters.\n%s", strings.Join(prefix, "\n"))

		// NOTE: Provide rich gRPC error message
		st := status.New(trouble.Code, trouble.Message)
		st, _ = st.WithDetails(&errdetails.BadRequest{
			FieldViolations: fv,
		})

		return st.Err()
	default:
		return err
	}
}

func NewTroubleMaker(setters ...TroubleMakerSetter) TroubleMaker {
	tmi := new(TroubleMakerImpl)

	for _, set := range setters {
		set(tmi)
	}

	return tmi
}

func WithCodes(codes map[string]codes.Code) TroubleMakerSetter {
	return func(tmi *TroubleMakerImpl) {
		tmi.codes = codes
	}
}

func WithMessages(messages map[string]string) TroubleMakerSetter {
	return func(tmi *TroubleMakerImpl) {
		tmi.messages = messages
	}
}
