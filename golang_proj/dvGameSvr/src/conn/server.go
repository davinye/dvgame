package conn

type Request struct {
	Method string "method"
	Params string "params"
}

var Value int = 5

func GetRequest() *Request {
	return &Request{}
}
