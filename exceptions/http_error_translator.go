package exceptions

type HTTPErrorTranslator struct {
	directories    map[string]*HTTPError
	grpcStatusCode map[uint32]uint32
}

func NewHTTPErrorTranslator() *HTTPErrorTranslator {
	return &HTTPErrorTranslator{
		directories: map[string]*HTTPError{
			"SCOPE_REPOSITORY.DUPLICATE_NAME": &HTTPError{
				StatusCode: 400,
				Message:    "duplicate name value",
			},
			"SCOPE_REPOSITORY.SCOPE_NOT_FOUND": &HTTPError{
				StatusCode: 404,
				Message:    "Scope not found",
			},
			"SCOPE_USE_CASE.SCOPE_TIED_TO_ROLES": &HTTPError{
				StatusCode: 409,
				Message:    "unable to delete Scope while tied to a Role / Roles",
			},
			"SCOPE_USE_CASE.USER_NOT_AUTHORIZED": &HTTPError{
				StatusCode: 403,
				Message:    "unautorized",
			},
			"SCOPE_USE_CASE.ALREADY_EXISTS": &HTTPError{
				StatusCode: 409,
				Message:    "Scope with such name already exists",
			},
			"VALIDATOR.INVALID_DATA": &HTTPError{
				StatusCode: 400,
				Message:    "validation error: missing request bodies",
			},
			"VALIDATOR.UNKNOWN": &HTTPError{
				StatusCode: 500,
				Message:    "validation error: unknown",
			},
		},

		grpcStatusCode: map[uint32]uint32{
			400: 3, // HTTP code 400 or GRPC code 3 loosely means BadRequest
			403: 7, // HTTP code 403 or GRPC code 7 means Unauthorized
			404: 5, // HTTP code 404 or GRPC code 5 means NotFound
			409: 6, // HTTP code 409 or GRPC code 6 loosely means Conflict
			500: 2, // HTTP code 500 or GRPC code 2 loosely means Server/Unknown Error
		},
	}
}

func (t *HTTPErrorTranslator) Translate(err error) *HTTPError {
	httpError := t.directories[err.Error()]
	if httpError == nil {
		return &HTTPError{
			StatusCode: 500,
			Message:    err.Error(),
		}
	}

	return httpError
}

func (t *HTTPErrorTranslator) TranslateMessage(err error) string {
	httpError := t.directories[err.Error()]
	if httpError == nil {
		return err.Error()
	}

	return httpError.Message
}

func (t *HTTPErrorTranslator) TranslateGRPCStatusCode(statusCode uint32) uint32 {
	return t.grpcStatusCode[statusCode]
}
