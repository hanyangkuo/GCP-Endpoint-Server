package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)
type EnrollmentMessage struct {
	Node_key string
	Node_invalid bool
}

type ConfigurationMessage struct {
	Schedule string
	System_info QueryNeme
	Node_invalid bool
}

type QueryNeme struct {
	Query string
	Interval int64
}
func EnrollmentServer(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	log.Println("EnrollmentServer")
	log.Println("\n"+string(body))
	w.Header().Set("Content-Type", "application/json")
	//raw:= json.RawMessage(`{"node_key":"","node_invalid":false}`)
	json.NewEncoder(w).Encode(json.RawMessage(`{"node_key":"","node_invalid":false}`))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func ConfigurationServer(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	log.Println("ConfigurationServer")
	log.Println("\n"+string(body))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(json.RawMessage("{\"schedule\":{\"system_info\":{\"query\":\"SELECT * FROM system_info;\",\"interval\":10}},\"node_invalid\":false}"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func LogServer(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	if strings.Contains(string(body),`"log_type":"result"`){
		log.Println("LogServer")
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, body, "", "\t")
		log.Println(string(prettyJSON.Bytes()))
		//log.Println("\n"+string(body))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(json.RawMessage("{\"node_invalid\":false}"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	http.HandleFunc("/api/v1/osquery/enroll", EnrollmentServer)
	http.HandleFunc("/api/v1/osquery/config", ConfigurationServer)
	http.HandleFunc("/api/v1/osquery/log", LogServer)
	err := http.ListenAndServeTLS(":9527", "cert/cert.pem", "cert/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}