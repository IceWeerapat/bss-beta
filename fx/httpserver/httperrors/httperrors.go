package httperrors

var (
	ErrNotFound     = &errImpl{"404", "not found"}
	ErrUnauthorized = &errImpl{"401", "unauthorized"}
	ErrInternal     = &errImpl{"500", "internal server error"}
	ErrBadRequest   = &errImpl{"400", "bad request"}
)

type errImpl struct {
	code    string
	message string
}

func (i *errImpl) Code() string  { return i.code }
func (i *errImpl) Error() string { return i.message }
func NewError(code, message string) *errImpl {
	return &errImpl{code: code, message: message}
}
