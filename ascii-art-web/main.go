package main

import (
	asciiart "ascii-art-web/ascii-art"
	"fmt"
	"net/http"
	"os"
)

//---------------------------------------------------------------------------------------|

var errFile string

func main() {
	content, err := os.ReadFile("templates/error.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	errFile = string(content)

	// ip := "10.25.0.59"
	// port := 45674

	// addr := fmt.Sprintf("%s:%d", ip, port)
	fmt.Printf("server runing http://localhost%s\n", ":45674")
	err = http.ListenAndServe(":45674", &mainHandler{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

//---------------------------------------------------------------------------------------|

// Главный обработчик для всех запросов
type mainHandler struct{}

func (h *mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	fmt.Printf("DEBUG: Request path: %s\n", path)

	switch path {
	case "/":
		fmt.Printf("DEBUG: Handling home page\n")
		handlerHome(w, r)
	case "/sabik":
		fmt.Printf("DEBUG: Handling sabik page\n")
		handlerSabik(w, r)
	case "/ascii-art":
		fmt.Printf("DEBUG: Handling ascii-art page\n")
		handlerAsciiart(w, r)
	default:
		fmt.Printf("DEBUG: Path not found, returning 404\n")
		// Если путь не соответствует ни одному из известных маршрутов
		handlerError(w, http.StatusNotFound)
	}
}

//---------------------------------------------------------------------------------------|

func handlerHome(w http.ResponseWriter, r *http.Request) {

	if r.URL.RawQuery != "" {
		handlerError(w, http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodGet {
		handlerError(w, http.StatusMethodNotAllowed)
		return
	}

	indexPage, err := os.ReadFile("templates/index.html")
	if err != nil {
		handlerError(w, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(indexPage))
}

//---------------------------------------------------------------------------------------|

func handlerSabik(w http.ResponseWriter, r *http.Request) {
	sabikPage, err := os.ReadFile("templates/sabik.html")
	if err != nil {
		handlerError(w, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(sabikPage))
}

//---------------------------------------------------------------------------------------|

func handlerAsciiart(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// проверить что text и банер не пустые
	// проверить что банер существует
	fmt.Printf("method %s\n", r.Method)

	if r.Method != http.MethodPost {
		handlerError(w, http.StatusMethodNotAllowed)
		return
	}

	if banner != "shadow" && banner != "thinkertoy" && banner != "standard" {
		handlerError(w, http.StatusBadRequest)
		return
	}

	for _, r := range text {
		if !((r >= 32 && r <= 126) || r == 13 || r == 10) {
			handlerError(w, http.StatusBadRequest)
			return
		}
	}

	asciiPage, err := os.ReadFile("templates/asciiart.html")
	if err != nil {
		handlerError(w, http.StatusInternalServerError)
		return
	}
	result, err := asciiGenerator(text, banner)
	if err != nil {
		handlerError(w, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(asciiPage), result)
}

//---------------------------------------------------------------------------------------|

func handlerError(w http.ResponseWriter, errCode int) {
	w.WriteHeader(errCode)
	fmt.Fprintf(w, errFile, errCode)
}

//---------------------------------------------------------------------------------------|

func asciiGenerator(text, banner string) (string, error) {

	asciiArt, err := asciiart.Generate(text, banner)
	if err != nil {
		return "", err
	}

	return asciiArt, nil
}
