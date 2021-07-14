package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protos "github.com/sssaang/go-cafe/currency/protos/currency"
	"github.com/sssaang/go-cafe/currency/server"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)
	
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}