package rp

import (
	"fmt"
	"net/http"
)

// Reverse Proxy
type RP struct {
	Server http.Server
}

func New(port string, handler http.HandlerFunc) *RP {
	return &RP{
		Server: http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: http.HandlerFunc(handler),
		},
	}
}

func (r *RP) ListenAndServe() error {
	if err := r.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
