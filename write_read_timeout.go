package deadline

import (
	"net/http"
	"time"
)

func TimeoutMiddleware(next http.Handler, timeout time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// rc.SetReadDeadline rc.SetWriteDeadline : A timeout occurs with the shorter of the two.
		// If http.Server.ReadTimeout or WriteTimeout is set and not reset,
		// a timeout occurs with the http.ServerTimeout value.
		// Therefore, both are set to the same value.

		// NewResponseController Set Timeout
		rc := http.NewResponseController(w)
		if timeout > 0 {
			rc.SetReadDeadline(time.Now().Add(timeout))
			rc.SetWriteDeadline(time.Now().Add(timeout))
		}

		next.ServeHTTP(w, r)
	})
}
