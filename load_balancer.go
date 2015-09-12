package main

import (
		"log"
		"net/http"
		"net/http/httputil"
		"math/rand"
		"strconv"
)

func main() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

				ports := []int{5555, 5556, 5557, 5558, 5559, 5560}
				for i := range ports {
						j := rand.Intn(i + 1)
						ports[i], ports[j] = ports[j], ports[i]
				}
				port := ports[0]
				director := func(req *http.Request) {
						req = r
						req.URL.Scheme = "http"
						req.URL.Host = "localhost:"+strconv.Itoa(port)
				}
				proxy := &httputil.ReverseProxy{Director: director}
				proxy.ServeHTTP(w, r)
		})
		log.Fatal(http.ListenAndServe(":8181", nil))
}
