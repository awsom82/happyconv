package webconv

import (
	"fmt"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

// NewServer creates new server and limiter
func NewServer(conf *Config) *http.Server {

	lim := &limiter.ExpirableOptions{DefaultExpirationTTL: conf.RateLimitTTL * time.Second}
	lmt := tollbooth.NewLimiter(conf.RateLimit, lim)
	lmt.SetIPLookups([]string{"X-Forwarded-For", "RemoteAddr", "X-Real-IP"})

	h := http.HandlerFunc(WebconvHandler)

	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%d", conf.Hostname, conf.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      tollbooth.LimitFuncHandler(lmt, h), // handle with third-party limiter
	}

	srv.SetKeepAlivesEnabled(conf.KeepAlive)

	return &srv
}

// WebconvHandler a http.handler function
func WebconvHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		conv := NewConv()
		conv.CopyInput(r)
		err := conv.SwapFormat()
		conv.MakeReply(w, err)

	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)

	}

}
