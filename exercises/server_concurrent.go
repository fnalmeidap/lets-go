package main



func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))

	w.Write([]byte("Message received!"))
}

func serverImpl(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePostRequest(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func main() {
	http.HandleFunc("/", serverImpl)
	fmt.Println("Starting server on http://localhost" + SERVER_PORT)

	err := http.ListenAndServe(SERVER_PORT, nil)
	if err != nil {
		panic(err)
	}
}