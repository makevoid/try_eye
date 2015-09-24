package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"

	"log"
	"net/http"
	"net/http/httputil"
	"math/rand"
	"strconv"
)

type Configuration struct {
	PortStart     int `json:"port_start"`
	ProcessesNum  int `json:"processes_num"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./config/thin.json")
	check(err)
	configuration := Configuration{}

	json.Unmarshal([]byte(dat), &configuration)
	fmt.Println("Config loaded:")
	fmt.Println(configuration.PortStart)
	fmt.Println(configuration.ProcessesNum)

	var ports []int
	ports = make([]int, configuration.ProcessesNum, configuration.ProcessesNum)

	for i := 0; i <= configuration.ProcessesNum-1; i++ {
		ports[i] = configuration.PortStart + i
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
		go proxy.ServeHTTP(w, r)
	})
	fmt.Println("ReverseProxy started!")
	log.Fatal(http.ListenAndServe(":8181", nil))
}
