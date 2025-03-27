package httpx

import "github.com/boostgo/errorx"

const (
	statusSuccess = "Success"
	statusFailure = "Failure"
)

type FailureResponse struct {
	Status  string         `json:"status"`
	Type    string         `json:"type,omitempty"`
	Message string         `json:"message"`
	Inner   string         `json:"inner,omitempty"`
	Context map[string]any `json:"context,omitempty"`
}

func NewFailure(err error) FailureResponse {
	const defaultErrorType = "ERROR"

	response := FailureResponse{
		Status: statusFailure,
	}

	// build/collect error output
	custom, ok := errorx.TryGet(err)
	if ok {
		response.Message = custom.Message()
		response.Type = custom.Type()
		response.Context = custom.Context()
		if custom.InnerError() != nil {
			response.Inner = custom.InnerError().Error()
		}
	} else {
		response.Message = err.Error()
		response.Type = defaultErrorType
	}

	// clear from trace
	if response.Context != nil {
		if _, traceExist := response.Context["trace"]; traceExist {
			delete(response.Context, "trace")
		}
	}

	return response
}

type CreatedResponse struct {
	ID any `json:"id"`
}

type SuccessResponse struct {
	Status string `json:"status"`
	Body   any    `json:"body"`
}

func NewSuccess(body any) SuccessResponse {
	return SuccessResponse{
		Status: statusSuccess,
		Body:   body,
	}
}
