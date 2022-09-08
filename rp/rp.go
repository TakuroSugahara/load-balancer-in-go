package rp

import (
	"fmt"
	"net/http"
)

// Reverse Proxy
type RP struct {
	Server http.Server
}

func New(port string, httpHander http.Handler) *RP {
	return &RP{
		Server: http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: httpHander,
		},
	}
}

func (r *RP) ListenAndServe() error {
	if err := r.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
