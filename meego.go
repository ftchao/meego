package meego

import (
	"fmt"
	"log"
	"net/http"
)

func Run(addr string) {
	fmt.Print("meego.run\n")

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
