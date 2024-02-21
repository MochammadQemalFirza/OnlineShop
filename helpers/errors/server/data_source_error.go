package server

type DataSourceError struct {
	err   error
	cause string
}

// compile time checking
var _ ServerError = DataSourceError{}

func NewDataSourceError(err error, cause string) *DataSourceError {
	return &DataSourceError{
		err:   err,
		cause: cause,
	}
}

func (e DataSourceError) Error() string {
	return e.cause + ": " + e.err.Error()
}

func (e DataSourceError) Unwrap() error {
	return e.err
}
