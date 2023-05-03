package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	switcher int
)

func RunProxy(poxyAddr string, instanceHostsAddr []string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleProxy(w, r, instanceHostsAddr)
	})
	fmt.Println("Proxy Started")
	log.Fatalln(http.ListenAndServe(poxyAddr, nil))
}

func writeResponce(w http.ResponseWriter, res *http.Response) {
	text, err := io.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(text)

}

func handleProxy(w http.ResponseWriter, r *http.Request, hostsAddr []string) {
	if switcher == len(hostsAddr) {
		switcher = 0
	}
	addr := "http://" + hostsAddr[switcher%len(hostsAddr)]
	switcher++

	switch r.Method {
	case http.MethodPost:
		response, err := http.Post(addr+r.URL.Path, "text/json", r.Body)
		if err != nil {
			return
		}
		writeResponce(w, response)

	case http.MethodGet:
		response, err := http.Get(addr + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		writeResponce(w, response)

	case http.MethodDelete:
		req, _ := http.NewRequest("DELETE", addr+r.URL.Path, nil)
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		writeResponce(w, response)

	case http.MethodPut:
		req, _ := http.NewRequest("PUT", addr+r.URL.Path, r.Body)
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		writeResponce(w, response)
	}

}
