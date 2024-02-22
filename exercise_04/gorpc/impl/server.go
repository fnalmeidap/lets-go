package impl

type Server struct {}

type Request struct {
	message string
}

type Request struct {
	message string
}

func (s *Server) Greet(req Request, res *Response) error {
	res.message = "HTTP/1.1 200 OK"
}
