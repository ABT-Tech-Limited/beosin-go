package beosin

import "fmt"

// Common error codes from Beosin API
const (
	ErrCodeParameterError        = 40001
	ErrCodePlatformNotSupported  = 40021
	ErrCodeAddressError          = 40022
	ErrCodeTxHashError           = 40023
	ErrCodeTxHashNotExist        = 41023
	ErrCodeNonERC20NotSupported  = 41024
	ErrCodeContractNotSupported  = 41026
	ErrCodeTokenNotInBasket      = 41035
	ErrCodeTaskExecuting         = 41038
)

// APIError represents an error returned by the Beosin API
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("beosin api error: code=%d, message=%s", e.Code, e.Message)
}

// IsParameterError checks if the error is a parameter error
func (e *APIError) IsParameterError() bool {
	return e.Code == ErrCodeParameterError
}

// IsPlatformNotSupported checks if the platform is not supported
func (e *APIError) IsPlatformNotSupported() bool {
	return e.Code == ErrCodePlatformNotSupported
}

// IsAddressError checks if the address is invalid
func (e *APIError) IsAddressError() bool {
	return e.Code == ErrCodeAddressError
}

// IsTxHashError checks if the transaction hash is invalid
func (e *APIError) IsTxHashError() bool {
	return e.Code == ErrCodeTxHashError
}

// IsTxHashNotExist checks if the transaction hash does not exist
func (e *APIError) IsTxHashNotExist() bool {
	return e.Code == ErrCodeTxHashNotExist
}

// IsTaskExecuting checks if the task is still executing
func (e *APIError) IsTaskExecuting() bool {
	return e.Code == ErrCodeTaskExecuting
}

// IsTokenNotInBasket checks if the token is not in the basket
func (e *APIError) IsTokenNotInBasket() bool {
	return e.Code == ErrCodeTokenNotInBasket
}

// NewAPIError creates a new APIError from code and message
func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}
