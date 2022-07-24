package main

import(
	
	"fmt"
	"net"
	"net/http"
	"log"
	"os"
	
	
)


//index page handler
func index(w http.ResponseWriter, r *http.Request) {
	//set os version
	//set version in response header
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)
    //exact header from request
	//set it in response header
	for k, v := range r.Header {
		//fmt.Printf("key: %s ,value:%s \n",k,v)
		for _, vv := range v {
			//fmt.Printf("Header value: %s \n", vv)
			w.Header().Set(k,vv)
		}
	}

	//log output
	clientIP := getCurrentIP(r)

	log.Printf("Succeed! Response Status Code: %d",200)
	log.Printf("Succeed! Client IP: %s",clientIP)

}

func getCurrentIP(r *http.Request) string{
	remoteAddr := r.RemoteAddr
	
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr

}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Succeed! Status Code: 200"))


}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",index)
	mux.HandleFunc("/healthz",healthz)
	if err := http.ListenAndServe(":8080",mux); err != nil {
		log.Fatalf("start http server failed, error: %s \n",err.Error())
	}
}
