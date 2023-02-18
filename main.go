package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

type Records struct {
	Name    string
	Subject string
	Marks   int64
}

type Result struct {
	Data []Records
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("index.html")
		temp.Execute(w, nil)
		return
	}

	file, _, err := r.FormFile("csvfile")
	if err != nil {
		fmt.Fprintln(w, "error:", err)
		return
	}
	defer file.Close()

	var parsedCsv []Records
	fileReader := csv.NewReader(file)
	fileReader.FieldsPerRecord = -1
	for {
		record, err := fileReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(w, "error:", err)
			return
		}

		marks, _ := strconv.Atoi(record[2])
		student := Records{Name: record[0], Subject: record[1], Marks: int64(marks)}
		parsedCsv = append(parsedCsv, student)

	}
	output := formatedData(parsedCsv)
	w.Header().Set("content-type", "text/plain")
	w.Write([]byte(output))
}

func formatedData(unformattedSlice []Records) string {
	var output string
	for _, data := range unformattedSlice {
		output += fmt.Sprintf("%-10s%-10s%-10d\n", data.Name, data.Subject, data.Marks)

	}
	return output
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, nil)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/upload", uploadHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8888", r))

}
