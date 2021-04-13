package httpserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const CONFIG_PATH = "configs"
const VERSION = "v1.0.0"

func Serve() {
	http.HandleFunc("/text", text)
	http.HandleFunc("/config", config)
	http.HandleFunc("/showip", showip)
	http.HandleFunc("/version", version)
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

func config(w http.ResponseWriter, r *http.Request) {
	files := getFiles()
	for _, file := range files {
		if strings.HasPrefix(file, CONFIG_PATH + "/.") {
			continue
		}

		fmt.Fprintf(w, "FileName: %s\n", file)
		fmt.Fprintf(w, "----------------\n")
		data, _ := ioutil.ReadFile(file)
		fmt.Fprintf(w, "%s\n", string(data))
		fmt.Fprintf(w, "----------------\n\n\n")
	}
}

func getFiles() []string {
	var files []string
	root := CONFIG_PATH
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func printEnv(w http.ResponseWriter, r *http.Request) {
	regex, _ := regexp.Compile("^[a-z]")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if regex.MatchString(pair[0]) {
			fmt.Fprintf(w, "%s:  %s\n", pair[0], pair[1])
		}
	}
}

func showip(w http.ResponseWriter, r *http.Request) {
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		// ignore err
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Fprintf(w, "IP: %s\n", ip)
		}
	}
}

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "VERSION: " + VERSION)
}
