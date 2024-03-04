package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "HTTP triggered function.  Please pass a name"
	name := r.URL.Query().Get("name")
	if name != "" {
		message = fmt.Sprintf("Hello %s", name)
	}
	fmt.Fprint(w, message)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/GoExample", helloHandler)
	log.Printf("GoExample listening on %s, go to http://localhost%s/api/GoExample?name=Lyena", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}