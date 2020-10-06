package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	port      string = "8888"
	directory string = "."
)

func main() {

	port = *(flag.String("p", port, "port to serve on"))
	directory = *(flag.String("d", directory, "the directory of static file to host"))
	flag.Parse()

	log.Printf("Serving %s\n", directory)
	log.Printf("Visit http://localhost:%s\n", port)

	err := http.ListenAndServe(":"+port, handleFunc())
	if err != nil {
		log.Fatal(err)
	}
}

func handleFunc() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {

		userInputPath := req.URL.Query().Get("path")
		if IsValidURL(userInputPath) {
			resp, err := http.Get(userInputPath)
			if err != nil {
				log.Fatal(err.Error())
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.Write(body)
		} else {
			http.FileServer(http.Dir(directory)).ServeHTTP(w, req)
		}
	}
	return http.HandlerFunc(fn)

}

// IsValidURL tests a string to determine if it is a well-structured url or not.
func IsValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
