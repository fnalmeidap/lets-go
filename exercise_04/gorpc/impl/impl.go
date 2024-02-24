package impl

type Api struct {}

type Request struct {
	Message string
}

type Response struct {
	Message string
}

func (a *Api) Greet(req Request, res *Response) error {
	res.Message = "Hello from server!"
	return nil
}
