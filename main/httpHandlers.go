package appengine

import (
	"net/http"

	"github.com/MarcGrol/cardRefund/service"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	{
		s := service.CardReturnService{}
		s.HTTPHandlerWithRouter(router)
	}

	// register the root
	http.Handle("/", router)
}
