// Generated automatically by golangAnnotations: do not edit manually

package service

import (
	"github.com/MarcGrol/golangAnnotations/generator/rest/testcase"
	"github.com/gorilla/mux"
)

// HTTPTestHandlerWithRouter registers endpoint in existing router
func HTTPTestHandlerWithRouter(router *mux.Router, results testcase.TestSuiteDescriptor) *mux.Router {
	subRouter := router.PathPrefix("/").Subrouter()

	subRouter.HandleFunc("/logs.md", testcase.WriteTestLogsAsMarkdown(results)).Methods("GET")

	return router
}
