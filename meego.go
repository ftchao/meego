package meego

import (
	"log"
	"net/http"
)

func Run(addr string) {
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
