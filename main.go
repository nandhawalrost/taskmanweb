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

func main() {
	// http.Handle("/", http.FileServer(http.Dir("."))) //for accessing index.html

	http.HandleFunc("/run", runExecutable)

	http.HandleFunc("/read", serveTextFile)

	log.Println("Serving on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
