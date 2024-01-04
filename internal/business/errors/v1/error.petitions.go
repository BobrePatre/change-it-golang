package v1

const (
	prefix = "domain: "
)

type DomainError struct {
	Message    string
	StatusCode int
}

type AlreadyExistsError DomainError

func (e *AlreadyExistsError) Error() string {
	return prefix + e.Message
}

type NotFoundError DomainError

func (e *NotFoundError) Error() string {
	return prefix + e.Message
}
