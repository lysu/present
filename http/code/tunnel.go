func ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if r.Method == "CONNECT" {
		handler.HandleHTTPS(w, r)
	} else {
		handler.HandleHTTP(w, r)
	}
}

func (handler ProxyHandler) HandleHTTPS(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Connection", "close")
	clientConn, _, err := w.(http.Hijacker).Hijack()
  remoteConn := getRemoteConn()//ss,pb..
	go Pipe(clientConn, remoteConn)
	log.Debugln("Sending ok reponse to client")
	clientConn.Write([]byte("HTTP/1.0 200 Connection established\r\n\r\n"))

}
