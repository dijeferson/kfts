package provider

import "fmt"

// ErrFailedToParsePosition is thrown when it fail to process the position data.
type ErrFailedToParsePosition struct {
	err  error
	data string
}

func (e *ErrFailedToParsePosition) Error() string {
	return fmt.Sprintf("There was an error trying to parse the position data.\nData: %v.\nError: %+v", e.data, e.err)
}

// ErrFailedToProcessBookInfo is thrown when there's an error processing the book info.
type ErrFailedToProcessBookInfo struct {
	err  error
	data string
}

func (e *ErrFailedToProcessBookInfo) Error() string {
	return fmt.Sprintf("Failed to process the book info.\nData: %v.\nError: %+v", e.data, e.err)
}

// ErrFailedToProcessPosition is thrown when there's an error processing the clip position.
type ErrFailedToProcessPosition struct {
	err  error
	data string
}

func (e *ErrFailedToProcessPosition) Error() string {
	return fmt.Sprintf("Failed to process the clip position.\nData: %v.\nError: %+v", e.data, e.err)
}

// ErrFailedToProcessCreationDate is thrown when there's an error processing the book info.
type ErrFailedToProcessCreationDate struct {
	err  error
	data string
}

func (e *ErrFailedToProcessCreationDate) Error() string {
	return fmt.Sprintf("Failed to process the clip creation date.\nData: %v.\nError: %+v", e.data, e.err)
}

// ErrInvalidOperation is thrown when there's an error determining the clipping operation
type ErrInvalidOperation struct {
	data string
}

func (e *ErrInvalidOperation) Error() string {
	return fmt.Sprintf("Failed to determine the clipping operation.\nData: %v.", e.data)
}
