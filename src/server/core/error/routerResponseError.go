package error

import "net/http"

type RouterResponseError struct {
	ErrorCode int    `json:"errorCode"`
	Cause     string `json:"cause"`
	Message   string `json:"message"`
	Status    string `json:"status"`
}

func RestErrorSomethingWentWrong(message string) RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusInternalServerError,
		Cause:     message,
		Message:   "Something went wrong",
		Status:    "Something went wrong",
	}
}

func RestErrorResourceNotFoundWithId(resourceName string, resourceId any) RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusNotFound,
		Cause:     ResourceNotFoundWithId(resourceName, resourceId),
		Message:   ResourceNotFoundWithId(resourceName, resourceId),
		Status:    ResourceNotFoundWithId(resourceName, resourceId),
	}
}

func RestErrorParameterMandatory(parameter string) RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusBadRequest,
		Cause:     ParameterMandatory(parameter),
		Message:   ParameterMandatory(parameter),
		Status:    ParameterMandatory(parameter),
	}
}

func RestErrorUnwritableParameter(parameter string) RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusBadRequest,
		Cause:     UnwritableParameter(parameter),
		Message:   UnwritableParameter(parameter),
		Status:    UnwritableParameter(parameter),
	}
}

func RestErrorBadRequestValidationError(message string) RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusBadRequest,
		Cause:     message,
		Message:   message,
		Status:    message,
	}
}

func RestErrorInvalidAuthentication() RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusUnauthorized,
		Cause:     InvalidAuthentication(),
		Message:   InvalidAuthentication(),
		Status:    InvalidAuthentication(),
	}
}

func RestErrorValueMustBeWithinRange(parameter string, lowerBoundary int, upperBoundary int) RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusBadRequest,
		Cause:     ValueMustBeWithinRange(parameter, lowerBoundary, upperBoundary),
		Message:   ValueMustBeWithinRange(parameter, lowerBoundary, upperBoundary),
		Status:    ValueMustBeWithinRange(parameter, lowerBoundary, upperBoundary),
	}
}

func RestErrorValueDoesNotMatchFormat(field string, value string, desiredPattern string) RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusBadRequest,
		Cause:     ValueDoesNotMatchFormat(value, field, desiredPattern),
		Message:   ValueDoesNotMatchFormat(value, field, desiredPattern),
		Status:    ValueDoesNotMatchFormat(value, field, desiredPattern),
	}
}

func RestErrorStringMustBeWithinSize(parameter string, lowerBoundary int, upperBoundary int) RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusBadRequest,
		Cause:     StringMustBeWithinSize(parameter, lowerBoundary, upperBoundary),
		Message:   StringMustBeWithinSize(parameter, lowerBoundary, upperBoundary),
		Status:    StringMustBeWithinSize(parameter, lowerBoundary, upperBoundary),
	}
}

func RestErrorStringMustBeAtLeastSize(parameter string, lowerBoundary int) RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusBadRequest,
		Cause:     StringMustBeAtLeastSize(parameter, lowerBoundary),
		Message:   StringMustBeAtLeastSize(parameter, lowerBoundary),
		Status:    StringMustBeAtLeastSize(parameter, lowerBoundary),
	}
}

func RestErrorParameterMustOneOfFollowingValues[T string | int](parameter string, values []T) RouterResponseError {
	return RouterResponseError{
		ErrorCode: http.StatusBadRequest,
		Cause:     ParameterMustOneOfFollowingValues(parameter, values),
		Message:   ParameterMustOneOfFollowingValues(parameter, values),
		Status:    ParameterMustOneOfFollowingValues(parameter, values),
	}
}
