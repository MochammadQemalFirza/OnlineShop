package server

type ServerError interface {
	Error() string
	Unwrap() error
}
