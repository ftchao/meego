package meego

import (
	"net/http"
)

func Router(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, f)
}
