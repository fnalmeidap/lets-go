package impl

type Api struct {}

type Request struct {
	message string
}

type Request struct {
	message string
}

func (a *Api) Greet(req Request, res *Response) error {
	res.message = "Hello from server!"
}
