package servidor

import (
	"encoding/json"
	"fmt"
	"github.com/kmiloparra/detector-gen-mutante-go/service"
	"io/ioutil"
	"net/http"
)

var detector service.DetectorService = service.DetectorGenMutante{}

type Router struct {
	rules map[string]http.HandlerFunc
}

type Req struct {
	Dna []string `json:"dna"`
}


func NewRouter() *Router {
	rules := make(map[string]http.HandlerFunc)
	rules["/isMutant"] = AllArticles
	for k, v := range rules {
		http.HandleFunc(k, v)
	}
	return &Router{rules}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hola mundo!!")
}

func AllArticles(w http.ResponseWriter, r *http.Request) {

	var req Req

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode("ok")
	case http.MethodPost:
		fmt.Println(r.Body)
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal

		err = json.Unmarshal(b, &req)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(req)
		if detector.IsMutant(req.Dna) {
			json.NewEncoder(w).Encode("ok")
		} else {
			w.WriteHeader(403)
			fmt.Fprintf(w, "Forbidden")
		}
	default:
		w.WriteHeader(404)
		fmt.Fprintf(w, "Resource Not Found")

	}

}
