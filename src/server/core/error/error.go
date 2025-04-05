package error

import "fmt"

type GenericError struct {
	Error     error
	RestError RouterResponseError
}

func GenericErrorSomethingWentWrong(message string) GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s | Cause: %s", SomethingWentWrong(), message),
		RestError: RestErrorSomethingWentWrong(message),
	}
}

func GenericErrorResourceNotFoundWithId(parameter string, id any) GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s", ResourceNotFoundWithId(parameter, id)),
		RestError: RestErrorResourceNotFoundWithId(parameter, id),
	}
}

func GenericErrorParameterMandatory(parameter string) GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s", ParameterMandatory(parameter)),
		RestError: RestErrorParameterMandatory(parameter),
	}
}

func GenericErrorUnwritableParameter(parameter string) GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s", UnwritableParameter(parameter)),
		RestError: RestErrorUnwritableParameter(parameter),
	}
}

func GenericErrorBadRequestValidationError(message string) GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s", message),
		RestError: RestErrorBadRequestValidationError(message),
	}
}

func GenericErrorInvalidAuthentication() GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s", InvalidAuthentication()),
		RestError: RestErrorInvalidAuthentication(),
	}
}

func GenericErrorValueMustBeWithinRange(parameter string, lowerBoundary int, upperBoundary int) GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s", ValueMustBeWithinRange(parameter, lowerBoundary, upperBoundary)),
		RestError: RestErrorValueMustBeWithinRange(parameter, lowerBoundary, upperBoundary),
	}
}

func GenericErrorValueDoesNotMatchFormat(parameter string, value string, regex string) GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s", ValueDoesNotMatchFormat(parameter, value, regex)),
		RestError: RestErrorValueDoesNotMatchFormat(parameter, value, regex),
	}
}

func GenericErrorStringMustBeWithinSize(parameter string, lowerBoundary int, upperBoundary int) GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s", StringMustBeWithinSize(parameter, lowerBoundary, upperBoundary)),
		RestError: RestErrorStringMustBeWithinSize(parameter, lowerBoundary, upperBoundary),
	}
}

func GenericErrorStringMustBeAtLeastSize(parameter string, lowerBoundary int) GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s", StringMustBeAtLeastSize(parameter, lowerBoundary)),
		RestError: RestErrorStringMustBeAtLeastSize(parameter, lowerBoundary),
	}
}

func GenericErrorParameterMustOneOfFollowingValues[T string | int](parameter string, values []T) GenericError {
	return GenericError{
		Error:     fmt.Errorf("%s", ParameterMustOneOfFollowingValues(parameter, values)),
		RestError: RestErrorParameterMustOneOfFollowingValues(parameter, values),
	}
}
