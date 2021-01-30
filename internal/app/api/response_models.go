package api

import (
	"errors"
	"fmt"
)

func getResponseData(responseData interface{}, err error) map[string]interface{} {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}
	return map[string]interface{}{"response": responseData, "error": errorMsg}
}

func getErrorResponseRequired(requiredField string) map[string]interface{} {
	return getResponseData(nil, errors.New(fmt.Sprintf("%s is required", requiredField)))
}

func getErrorResponseValidation(field string, validationMsg string) map[string]interface{} {
	return getResponseData(nil, errors.New(fmt.Sprintf("%s must be %s", field, validationMsg)))
}
