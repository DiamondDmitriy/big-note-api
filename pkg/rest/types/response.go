package types

type (
	ResponseSuccess struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Meta   interface{} `json:"meta,omitempty"`
	}

	ResponseError struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Details interface{} `json:"details"`
	}
)
