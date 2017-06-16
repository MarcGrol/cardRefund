// +build !appengine

package ctx

import (
	"net/http"

	"golang.org/x/net/context"
)

type regularContextFactory struct {
}

func (f regularContextFactory) CreateContext(r *http.Request) context.Context {
	return context.TODO()
}

func init() {
	New = regularContextFactory{}
}
