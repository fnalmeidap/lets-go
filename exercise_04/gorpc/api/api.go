package api

type Api struct {}

type Request struct {
	message string
}

type Response struct {
	message string
}

func (a *Api) Greet(req Request, res *Response) error {
	res.message = "Hello from server!"
	return nil
}
