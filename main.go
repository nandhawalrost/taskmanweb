package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func runExecutable(w http.ResponseWriter, r *http.Request) {
	// Path to your executable
	exePath := "./app2.exe"

	// Execute the external application
	cmd := exec.Command(exePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error running executable but it works anyway XD : %v", err), http.StatusInternalServerError)
		return
	}

	// Return the output of the executable
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func serveTextFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./file.txt")
}

func serveHtmlFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow  all origins for simplicity. You can customize this.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/run", runExecutable)
	mux.HandleFunc("/text", serveTextFile)
	mux.HandleFunc("/read", serveHtmlFile)

	// Apply CORS middleware to all handlers
	handler := corsMiddleware(mux)

	server := http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	log.Println("Serving on :8081")
	log.Fatal(server.ListenAndServe())
}
