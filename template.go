//The static pages are based on this tutorial
//http://www.alexedwards.net/blog/serving-static-sites-with-go

// Serve the application form the `app.yaml` file.
// When running locally, use port 8080 so that the server is running behind the gateway.
// Disregard if there is no gateway running/configured locally on your system.
//
// dev_appserver.py --port=8080 elect-me-1255

package electme

import (
	"errors"
	"html/template"
	"net/http"
	"os"
	"path"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//it would be good for this function to pass a token to the page in case the page has a form (a lot of them will)
func serveTemplate(w http.ResponseWriter, r *http.Request) *appError {

	// Right now the server will redirect to `index.html` if the request URL
	// is `/`. However, this is not very good for the server. What should be
	// happening is this server should be handling routes, which right now it
	// is not. That looks like the following:
	//
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//   w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	//   w.WriteHeader(http.StatusOK)
	//   w.Write([]byte("Hello World"))
	// })
	//
	// This is a callback function that points to the HTTP request and response.
	// Once the URL path request is made and there is a an adequate TCP connection,
	// the server will pass the request to the handler and the handle will use HTTP
	// to serve the request.

	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index.html", 301)
	}

	c := appengine.NewContext(r)

	log.Debugf(c, "serving a non-static request")

	lp := path.Join(basePath+"/templates", "layout.html")
	fp := path.Join(basePath+"/templates", r.URL.Path)

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			return &appError{err, "Couldn't find the specified template", http.StatusNotFound}
		}
		return &appError{err, "Unknown error finding template", http.StatusInternalServerError}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		return &appError{errors.New("Attempted to display directory " + r.URL.Path), "Can't display record", http.StatusNotFound}
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		return &appError{err, "Error parsing template", http.StatusInternalServerError}
	}

	//log.Println(len(offices))
	if err := tmpl.ExecuteTemplate(w, "layout", topTen); err != nil {
		return &appError{err, "Error executing template", http.StatusInternalServerError}
	}
	return nil
}
