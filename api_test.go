package electme

import (
	//"google.golang.org/appengine"
	//"google.golang.org/appengine/aetest"
	"io/ioutil"
	"net/http"
	"testing"
)

//Unit Tests

//Integration Tests
func TestElectMe(t *testing.T) {
	/*
	   ctx, done, err : aetest.NewContext()

	   if err != nil{
	       t.Fatal(err)
	   }
	   defer done()
	*/

	//test that the pages load
	pageLoadTester(t, "", "index.html", "about.html")

	//test the api calls
	//we'll get to this... once we have an api
}

//PageLoadTester GETs each path and fails the test if the HTTP Status Code returned is not a 200 or 302
func pageLoadTester(t *testing.T, paths ...string) {
	for _, p := range paths {
		func() {
			res, err := http.Get("http://localhost:8080/" + p)
            if err != nil {
				t.Errorf("Error Getting page %s: %s", p, err)
				return
			}

			defer res.Body.Close()
			if res.StatusCode != 200 && res.StatusCode != 302 {
				t.Errorf("Error Getting %s: %s", p, res.Status)
				message, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Errorf("Error reading page %s body: %s", p, err)
					return
				}
				t.Errorf("Body of error response: %s", message)
			}
		}()
	}
}
