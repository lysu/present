package main
import ...

func main() {
	var srv http.Server
	srv.Addr = "0.0.0.0:2333"
	http.Handle("/", NewReverseProxy())
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	select {}
}

func NewReverseProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = req.Host
		req.URL.Path = req.URL.Path
		if req.URL.RawQuery != "" {
			req.URL.RawPath = req.URL.RawQuery
		}
	}
	return &httputil.ReverseProxy{Director: director}
}
