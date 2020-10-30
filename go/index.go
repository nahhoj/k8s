package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"

	pb "k8s/grpc/proto"

	"google.golang.org/grpc"
)

var (
	host           string
	portWebServer  string = ":3000"
	portgRPCServer string = ":3001"
)

type Server_gRPC struct {
	pb.UnimplementedMessageServer
}

func (s *Server_gRPC) SendMessage(ctx context.Context, in *pb.MessageReq) (*pb.MessageRes, error) {
	return &pb.MessageRes{
		Iam: host,
	}, nil
}

func main() {
	host, _ = os.Hostname()
	conn, _ := grpc.Dial("localhost:8081", grpc.WithInsecure())
	client := pb.NewMessageClient(conn)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		requ := pb.MessageReq{
			Who: host,
		}
		resMessage := ""
		resp, err := client.SendMessage(context.Background(), &requ)
		if err != nil {
			resMessage = "Connection error"
		} else {
			resMessage = resp.Iam
		}
		fmt.Println(err)
		fmt.Println(resp)
		res.Header().Set("Content-Type", "text/html; charset=UTF-8")
		res.WriteHeader(http.StatusOK)
		fmt.Fprintf(res, "<h2>Hello World <span style=\"color:red;\">GO</span> with kubernetes and microservice</h2>")
		fmt.Fprintf(res, "<h3>Platform: "+runtime.GOOS+"</h3>")
		fmt.Fprintf(res, "<h3>Hostmame: "+host+"</h3>")
		var ip net.IP
		iaddrs, _ := net.InterfaceAddrs()
		for _, addr := range iaddrs {
			a := addr.(*net.IPNet)
			if !a.IP.IsLoopback() && a.IP.To4() != nil {
				ip = a.IP
			}
		}
		fmt.Fprintf(res, "<h3>IP: "+ip.String()+"</h3>")
		fmt.Fprintf(res, "<br />")
		fmt.Fprintf(res, "<h3>Message from: "+resMessage+"</h3>")
	})
	go func() {
		fmt.Println("WebServer is starting for port " + portWebServer)
		error := http.ListenAndServe(portWebServer, nil)
		if error != nil {
			fmt.Println(error)
			os.Exit(1)
		}
	}()
	go func() {
		fmt.Println("gRPCServer is starting for port " + portgRPCServer)
		lis, error := net.Listen("tcp", portgRPCServer)
		if error != nil {
			fmt.Println(error)
			os.Exit(1)
		}
		serv := grpc.NewServer()
		pb.RegisterMessageServer(serv, &Server_gRPC{})
		if error := serv.Serve(lis); error != nil {
			fmt.Println(error)
			os.Exit(1)
		}
	}()
	select {}
}
