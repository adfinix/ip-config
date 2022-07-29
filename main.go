package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type publicIPAddr struct {
	PublicIP string `json:"public_ip"`
}

func applicationPort() string {
	if port, ok := os.LookupEnv("PORT"); ok {
		return port
	}
	return "80"
}

func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func userPublicIP(w http.ResponseWriter, r *http.Request)  {
	ip := readUserIP(r)

	response, _ := json.Marshal(publicIPAddr{
		PublicIP: ip,
	})
	log.Printf("Request from :%s\n", ip)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(response)
}

func main()  {
	portNumber :=  applicationPort()
	log.Printf("Application started at port %s", portNumber)
	http.HandleFunc("/", userPublicIP)
	log.Fatal(http.ListenAndServe(":" +portNumber , nil))
}

