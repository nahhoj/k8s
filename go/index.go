package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html; charset=UTF-8")
		res.WriteHeader(http.StatusOK)
		fmt.Fprintf(res, "<h2>Hello World GO with kubernetes and microservice</h3>")
		fmt.Fprintf(res, "<h3>Platform:"+runtime.GOOS+"</h3>")
		hostName, _ := os.Hostname()
		fmt.Fprintf(res, "<h3>Hostmame:"+hostName+"</h3>")
		var ip net.IP
		iaddrs, _ := net.InterfaceAddrs()
		for _, addr := range iaddrs {
			a := addr.(*net.IPNet)
			if !a.IP.IsLoopback() && a.IP.To4() != nil {
				ip = a.IP
			}
		}
		fmt.Fprintf(res, "<h3>IP:"+ip.String()+"</h3>")
	})
	error := http.ListenAndServe(":3000", nil)
	if error != nil {
		fmt.Println(error)
	}
}
