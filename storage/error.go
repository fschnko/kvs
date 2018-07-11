package storage

// Error represents a constant error.
type Error string

func (e Error) Error() string {
	return string(e)
}
