package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kmiloparra/detector-gen-mutante-go/service"
	"github.com/kmiloparra/resource-mutant/validaciones"
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
	rules["/isMutant"] = isMutant
	for k, v := range rules {
		http.HandleFunc(k, v)
	}
	return &Router{rules}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hola mundo!!")
}

func isMutant(w http.ResponseWriter, r *http.Request) {

	var req Req
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode("ok")
	case http.MethodPost:
		if evaluarGen(w, r, req) {
			return
		}
	default:
		w.WriteHeader(404)
		fmt.Fprintf(w, "Resource Not Found")

	}

}

func evaluarGen(w http.ResponseWriter, r *http.Request, req Req) bool {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return true
	}
	err = json.Unmarshal(b, &req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return true
	}

	if !validacionesDna(req.Dna) {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Bad request")
		return true
	}

	if detector.IsMutant(req.Dna) {
		json.NewEncoder(w).Encode("ok")
	} else {
		w.WriteHeader(403)
		fmt.Fprintf(w, "Forbidden")
	}
	return false
}

func validacionesDna(dna []string) bool {
	return validaciones.ValidacionFilaVacia(dna) &&
		validaciones.ValidacionNxN(dna) &&
		validaciones.ValidacionDominio(dna)
}
