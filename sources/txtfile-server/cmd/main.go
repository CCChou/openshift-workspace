package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/text", text)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}

func text(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.FormValue("text")
	fileName := getFileName()
	err := ioutil.WriteFile("files/"+fileName+".txt", []byte(text), 0666)
	if err != nil {
		log.Print(err)
	}
	fmt.Fprintf(w, "{\"message\":\"success\"}")
}

func getFileName() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
