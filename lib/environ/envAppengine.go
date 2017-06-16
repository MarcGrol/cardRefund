// +build appengine

package environ

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

// IsDevMode determines if the development server is running
func IsDevMode() bool {
	return appengine.IsDevAppServer()
}

// AppID determines the id of the application
func AppID(c context.Context) string {
	return appengine.AppID(c)
}

func GetDomainName(c context.Context) string {
	return fmt.Sprintf("https://%s.appspot.com", appengine.AppID(c))
}
