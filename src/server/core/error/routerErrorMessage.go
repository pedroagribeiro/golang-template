package error

import "fmt"

const (
	RouterErrorMessageResourceSomethingWentWrong        string = "Something went wrong"
	RouterErrorMessageResourceNotFoundWithId            string = "No record of %s could be found with id %s"
	RouterErrorMessageParameterMandatory                string = "The '%s' field must be provided"
	RouterErrorMessageUnwritableParameter               string = "The '%s' field cannot be updated once written"
	RouterErrorInvalidAuthentication                    string = "Wrong username/email and password combination"
	RouterErrorMessageValueMustBeWithinRange            string = "The value provided for parameter '%s' must be within the [%d .. %d] range"
	RouterErrorMessageValueDoesNotMatchFormat           string = "The value '%s' for parameter '%s' must match the '%s' regex"
	RouterErrorMessageStringMustBeWithinSize            string = "The string provided for parameter '%s' must have between %d and %d characters"
	RouterErrorMessageStringMustBeAtLeastSize           string = "The string provided for parameter '%s' must be at least %d characters"
	RouterErrorMessageParameterMustOneOfFollowingValues string = "The parameter '%s' must take on of the following values: %+v"
)

func SomethingWentWrong() string {
	return RouterErrorMessageResourceSomethingWentWrong
}

func ResourceNotFoundWithId(resourceType string, resourceId any) string {
	return fmt.Sprintf(RouterErrorMessageResourceNotFoundWithId, resourceType, resourceId)
}

func ParameterMandatory(parameter string) string {
	return fmt.Sprintf(RouterErrorMessageParameterMandatory, parameter)
}

func UnwritableParameter(parameter string) string {
	return fmt.Sprintf(RouterErrorMessageUnwritableParameter, parameter)
}

func InvalidAuthentication() string {
	return RouterErrorInvalidAuthentication
}

func ValueMustBeWithinRange(parameter string, lowerBoundary int, upperBoundary int) string {
	return fmt.Sprintf(RouterErrorMessageValueMustBeWithinRange, parameter, lowerBoundary, upperBoundary)
}

func ValueDoesNotMatchFormat(value string, parameter string, regex string) string {
	return fmt.Sprintf(RouterErrorMessageValueDoesNotMatchFormat, value, parameter, regex)
}

func StringMustBeWithinSize(parameter string, lowerBoundary int, upperBoundary int) string {
	return fmt.Sprintf(RouterErrorMessageStringMustBeWithinSize, parameter, lowerBoundary, upperBoundary)
}

func StringMustBeAtLeastSize(parameter string, lowerBoundary int) string {
	return fmt.Sprintf(RouterErrorMessageStringMustBeAtLeastSize, parameter, lowerBoundary)
}

func ParameterMustOneOfFollowingValues[T string | int](parameter string, values []T) string {
	return fmt.Sprintf(RouterErrorMessageParameterMustOneOfFollowingValues, parameter, values)
}
