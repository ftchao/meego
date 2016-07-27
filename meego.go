package meego

import (
	"fmt"
	"log"
	"net/http"
)

func Run(params ...string) {
	fmt.Print("meego.run\n")

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
