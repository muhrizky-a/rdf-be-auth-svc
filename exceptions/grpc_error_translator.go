package exceptions

type GRPCErrorTranslator struct {
	directories map[string]*GRPCError
}

func NewGRPCErrorTranslator() *GRPCErrorTranslator {
	// GRPC code 2 similar to HTTP code 500 Server/Unknown Error
	// GRPC code 3 similar to HTTP code 400 BadRequest)
	// GRPC code 5 similar to HTTP code 404 NotFound)
	// GRPC code 6 similar to HTTP code 409 Conflict
	// GRPC code 7 similar to HTTP code 403 Unauthorized)
	return &GRPCErrorTranslator{
		directories: map[string]*GRPCError{
			"SCOPE_REPOSITORY.DUPLICATE_NAME": &GRPCError{
				StatusCode: 3,
				// Message:    "duplicate name value",
				Message: "Scope with such name already exists",
			},
			"SCOPE_REPOSITORY.SCOPE_NOT_FOUND": &GRPCError{
				StatusCode: 5,
				Message:    "Scope not found",
			},
			"SCOPE_USE_CASE.SCOPE_TIED_TO_ROLES": &GRPCError{
				StatusCode: 6,
				Message:    "unable to delete Scope while tied to a Role / Roles",
			},
			"SCOPE_USE_CASE.USER_NOT_AUTHORIZED": &GRPCError{
				StatusCode: 7,
				Message:    "unautorized",
			},
			"VALIDATOR.INVALID_DATA": &GRPCError{
				StatusCode: 3,
				Message:    "validation error: missing request bodies",
			},
			"VALIDATOR.UNKNOWN": &GRPCError{
				StatusCode: 2,
				Message:    "validation error: unknown",
			},
		},
	}
}

func (t *GRPCErrorTranslator) Translate(err error) *GRPCError {
	gRPCError := t.directories[err.Error()]
	if gRPCError == nil {
		return &GRPCError{
			StatusCode: 2,
			Message:    err.Error(),
		}
	}

	return gRPCError
}

func (t *GRPCErrorTranslator) TranslateMessage(err error) string {
	gRPCError := t.directories[err.Error()]
	if gRPCError == nil {
		return err.Error()
	}

	return gRPCError.Message
}
