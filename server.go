package main

import (
	"fmt"
	"time"
	"html/template"
	"nasawebgo/nasa"
	"net/http"
)

//ApodDay NASA
func ApodDay(w http.ResponseWriter, r *http.Request) {
	apod, err := nasa.GetNasaAPODToday()
	if err == nil {
		generateHTML(w, apod, "header", "apod", "apod_day", "footer")
	} else {
		generateHTML(w, nil, "header", "error", "warning", "footer")
	}
}

//Home NASA
func Home(w http.ResponseWriter, r *http.Request) {
	apods, err := nasa.GetNasaAPODS()
	if err == nil {
		generateHTML(w, apods, "header", "home", "apods", "footer")
	} else {
		generateHTML(w, nil, "header", "error", "warning", "footer")
	}
}

func main() {
	server := http.Server{
		Addr: ":80",
    		ReadTimeout:  5 * time.Second,
    		WriteTimeout: 10 * time.Second,
    		IdleTimeout:  120 * time.Second,
	}
	files := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", files))

	http.HandleFunc("/ApodDay", ApodDay)
	http.HandleFunc("/", Home)
	server.ListenAndServe()
}

//utils
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
