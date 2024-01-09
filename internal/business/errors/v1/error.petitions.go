package v1

const (
	prefix = "domain: "
)

type domainError struct {
	Message    string
	StatusCode int
}

type AlreadyExistsError domainError

func (e *AlreadyExistsError) Error() string {
	return prefix + e.Message
}

type NotFoundError domainError

func (e *NotFoundError) Error() string {
	return prefix + e.Message
}

type AlreadyLikedError domainError

func (e *AlreadyLikedError) Error() string {
	return prefix + e.Message
}

type AlreadyVoicedError domainError

func (e *AlreadyVoicedError) Error() string {
	return prefix + e.Message
}

type ForbiddenError domainError

func (e *ForbiddenError) Error() string {
	return prefix + e.Message
}
