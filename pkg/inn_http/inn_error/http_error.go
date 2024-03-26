package inn_error

import "fmt"

type InnErrorResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (e *InnErrorResponse) Error() string {
	return fmt.Sprintf("INN error code: '%d'. Message: '%s'", e.Code, e.Message)
}
