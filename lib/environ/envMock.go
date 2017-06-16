// +build !appengine

package environ

import "golang.org/x/net/context"

// IsDevMode determines if the development server is running
func IsDevMode() bool {
	return true
}

// AppID determines the id of the application
func AppID(_ context.Context) string {
	return "mockApp"
}

func GetDomainName() string {
	return "http://localhost:8080"
}
