package ctx

import (
	"net/http"

	"golang.org/x/net/context"
)

type contextFactory interface {
	CreateContext(r *http.Request) context.Context
}

var (
	// New exposes a suitable context depending on environemt
	New contextFactory
)
