package client

import "fmt"

// RestError represents an error struct that occurres when underlying API is called
type RestError struct {
	Message    string
	StatusCode int
}

func (err *RestError) Error() string {
	return fmt.Sprintf("Error %d: %q", err.StatusCode, err.Message)
}
