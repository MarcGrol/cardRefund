package appengine

import (
	"net/http"

	"github.com/MarcGrol/cardRefund/services/adminService"
	"github.com/MarcGrol/cardRefund/services/userService"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	{
		s := userService.CardReturnService{}
		s.HTTPHandlerWithRouter(router)
	}
	{
		s := adminService.CardReturnService{}
		s.HTTPHandlerWithRouter(router)
	}

	// register the root
	http.Handle("/", router)
}
