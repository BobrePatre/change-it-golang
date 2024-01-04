package v1

const (
	prefix = "database: "
)

type DatasourceError struct {
	Message string
}

func (e *DatasourceError) Error() string {
	return prefix + e.Message
}

type NotFoundError DatasourceError

func (e *NotFoundError) Error() string {
	return prefix + e.Message
}
