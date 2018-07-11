package server

// Method represents methods of the server.
type Method string

// Server supported methods.
const (
	Get    Method = "GET"
	Set    Method = "SET"
	Delete Method = "DELETE"
	Exists Method = "EXISTS"
)

const (
	SuccessResult = "success"
	ExistResult   = "exists"
	NotExist      = "not exists"
)

type Request struct {
	Method Method `json:"method"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type Response struct {
	Method Method `json:"method"`
	Key    string `json:"key"`
	Value  string `json:"value,omitempty"`
	Result string `json:"result,omitempty"`
	Error  error  `json:"error,omitempty"`
}
