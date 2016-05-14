//The api is intended to support elect-me as an app that has some SPA liek capabilities for searching and displaying information about offices.

package electme

import (
	"google.golang.org/appengine"
	"net/http"
)

func offices(w http.ResponseWriter, r *http.Request) *appError {
	c := appengine.NewContext(r)
    
    return nil;
}