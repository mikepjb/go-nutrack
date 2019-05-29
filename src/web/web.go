// a web interface for nutrack
package web

import (
	"fmt"
	"github.com/mikepjb/nutrack/src/transport"
	"log"
	"net/http"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("site"))
	mux.Handle("/", fs)

	mux.HandleFunc("/test-plan.json", func(w http.ResponseWriter, r *http.Request) {
		// if r.Method == "POST" {
		// }
		transport.WriteTestPlan(w)
	})

	return mux
}

func Serve() {
	port := "8080"
	mux := routes()
	fmt.Printf("Starting serving on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
