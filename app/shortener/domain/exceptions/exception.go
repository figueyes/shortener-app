package exceptions

import "errors"

var (
	InvalidShortToCreateException = errorBuilder("invalid short to create")
	DataBaseException             = errorBuilder("database connection failed")
	DuplicationException          = errorBuilder("invalid short caused by duplication")
)

func errorBuilder(m string) error {
	return errors.New(m)
}
