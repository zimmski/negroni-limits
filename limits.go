package limits

import (
	"net/http"
)

type Limits struct {
	options *Options
}

type Options struct {
	MaxRequestSize int64
}

func New(options *Options) *Limits {
	if options.MaxRequestSize < 0 {
		options.MaxRequestSize = 0
	}

	return &Limits{
		options: options,
	}
}

func (l *Limits) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if l.options.MaxRequestSize != 0 && r.ContentLength > l.options.MaxRequestSize {
		http.Error(rw, http.StatusText(http.StatusRequestEntityTooLarge), http.StatusRequestEntityTooLarge)

		return
	}

	next(rw, r)
}
