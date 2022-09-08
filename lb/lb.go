package lb

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/TakuroSugahara/load-balancer/config"
)

var mu sync.Mutex
var idx int = 0

type LB struct {
	Backends []config.Backend
}

func New(backends []config.Backend) *LB {
	return &LB{
		Backends: backends,
	}
}

func (l *LB) Handler(w http.ResponseWriter, r *http.Request) {
	maxLen := len(l.Backends)

	// Round Robin
	mu.Lock()
	targetURL, err := url.Parse(l.Backends[idx%maxLen].URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	idx++
	mu.Unlock()

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ServeHTTP(w, r)
}
