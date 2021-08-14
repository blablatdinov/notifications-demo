package notifications

import "net/http"

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(hadler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":8000",
		Handler: hadler,
	}
	return s.httpServer.ListenAndServe()
}
