package lb

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/TakuroSugahara/load-balancer/backend"
)

var mu sync.Mutex
var idx int = 0

type LB struct {
	Backends []backend.Backend
}

func New(backends []backend.Backend) *LB {
	return &LB{
		Backends: backends,
	}
}

func (l *LB) Handler(w http.ResponseWriter, r *http.Request) {
	maxLen := len(l.Backends)

	// Round Robin
	mu.Lock()
	currentBackend := l.Backends[idx%maxLen]
	if currentBackend.IsDead {
		idx++
	}

	targetURL, err := url.Parse(l.Backends[idx%maxLen].URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	idx++
	mu.Unlock()

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		// NOTE: It is better to implement retry.
		log.Printf("%v is dead.", targetURL)
		currentBackend.SetDead(true)
		l.Handler(w, r)
	}
	reverseProxy.ServeHTTP(w, r)
}
