package product

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")
	// Menampilkan header di server log
	fmt.Printf("Authorization: %s\n", authorization)
}
