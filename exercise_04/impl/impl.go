package impl

type Api struct {}

type Request struct {
	Message string
}

type Response struct {
	Message string
}

func (a *Api) Greet(req Request, res *Response) error {
	res.Message = "HTTP/1.1 200 OK"
	return nil
}
