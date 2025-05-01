package helper

import (
	"errors"
	"strings"

	"github.com/ryakadev/rdf-be-auth-svc/exceptions"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCValidator struct {
	*Validator
	httpErrorTranslator *exceptions.HTTPErrorTranslator
}

func NewGRPCValidator(
	validator *Validator,
	httpErrorTranslator *exceptions.HTTPErrorTranslator,
) *GRPCValidator {
	return &GRPCValidator{
		validator,
		httpErrorTranslator,
	}
}

func (v *GRPCValidator) CreateGRPCValidationError(err error) error {
	validationError, unknownValidationErr := v.CreateValidationError(err)
	if unknownValidationErr != nil {
		return status.Errorf(
			codes.Unknown,
			v.httpErrorTranslator.TranslateMessage(unknownValidationErr),
		)
	}

	fieldViolations := make(
		[]*errdetails.BadRequest_FieldViolation,
		0,
		len(validationError.Validations),
	)
	for _, v := range validationError.Validations {
		// Create validation error messages
		// Fields are Fiels that doesnt pass validation
		// Tag like required have no Param, thus "required"
		// Example: {
		//   Field: "Age",
		//   Description: "required",
		// }
		//
		// Other Tag like gt have Param, like "gt: 18"
		// Example: {
		//   Field: "Age",
		//   Description: "gt : 18",
		// }
		//
		description := v.Tag + " : " + v.ParamValue
		if strings.TrimSpace(v.ParamValue) == "" {
			description = v.Tag
		}

		validation := &errdetails.BadRequest_FieldViolation{
			Field:       v.Field,
			Description: description,
		}

		fieldViolations = append(fieldViolations, validation)
	}
	detail := &errdetails.BadRequest{FieldViolations: fieldViolations}

	newErr := status.New(
		codes.InvalidArgument,
		v.httpErrorTranslator.TranslateMessage(
			errors.New("VALIDATOR.INVALID_DATA"),
		),
	)

	newErrWithDetails, _ := newErr.WithDetails(detail)
	return newErrWithDetails.Err()
}
