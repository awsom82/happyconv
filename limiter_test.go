package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateLimit(t *testing.T) {

	h := rateLimit(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello, im a dummy handle func!\n")
	}))

	req := httptest.NewRequest("POST", "/", nil)
	rr := httptest.NewRecorder()

	h.ServeHTTP(rr, req)
	//Should not be limited
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	rr2 := httptest.NewRecorder()

	h.ServeHTTP(rr, req)
	h.ServeHTTP(rr, req)
	h.ServeHTTP(rr2, req)

	//Should not be limited
	if rr2.Code != http.StatusTooManyRequests {
		t.Errorf("handler returned wrong status code: got %v want %v", rr2.Code, http.StatusTooManyRequests)
	}

	// ch := make(chan int)
	// go func() {
	//   rec := httptest.NewRecorder()
	//   h.ServeHTTP(rec, req)
	//   // Should be limited
	//   if rec.Code != http.StatusTooManyRequests {
	//     t.Errorf("handler returned wrong status code: got %v want %v", rec.Code, http.StatusTooManyRequests)
	//   }
	//   // OnLimitReached should be called
	//   if counter != 1 {
	//     t.Errorf("onLimitReached was not called")
	//   }
	//   close(ch)
	// }()
	// <-ch // Block until go func is done.

}
