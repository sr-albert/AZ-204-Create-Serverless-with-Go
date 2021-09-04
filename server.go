package main

import (
 "fmt"
 "io/ioutil"
 "log"
 "net/http"
 "os"
 "encoding/json"
)

type InvokeRequest struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}

// function sets the content type to application/json. 
// It responds with either "hello world" or the posted body, if any.
func triggerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
	  w.Write([]byte("hello world"))
	} else {
	  body, _ := ioutil.ReadAll(r.Body)
	  w.Write(body)
	}
  }

func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
	  w.Write([]byte("HELLO FROM HTTP TRIGGER"))
	} else {
	  body, _ := ioutil.ReadAll(r.Body)
	  w.Write(body)
	}
  }


func queueHandler(w http.ResponseWriter, r *http.Request) {
	// Read the body from incoming response stream and decodes it
	var invokeRequest InvokeRequest

	d := json.NewDecoder(r.Body)
	d.Decode(&invokeRequest)

	// The message itself is dug out with a call to [Unmarshal()]
	var reqData map[string]interface{}
	json.Unmarshal(invokeRequest.Data["queueItem"], &reqData)
	
	var parsedMessage string
	json.Unmarshal(invokeRequest.Data["queueItem"], &parsedMessage)
	// Print message out
	fmt.Println(parsedMessage) // your message
}

func main() {
	//  how it will read from the FUNCTIONS_CUSTOM_HANDLER_PORT environment variable
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	//  the function checks whether the port exists. If not, the function is assigned port 8080:
	if !exists {
	  customHandlerPort = "8080"
	}
	//The next code instantiates an HTTP server instance:
	mux := http.NewServeMux()
	// handle controller
	mux.HandleFunc("/api/hello", triggerHandler)
	mux.HandleFunc("/api/queueTrigger", queueHandler)
	mux.HandleFunc("/api/queueTrigger", queueHandler)

	fmt.Println("Go server Listening on: ", customHandlerPort)
	// The last line of importance is the one that starts listening to a specific port and signals that it's ready to receive requests
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
  }