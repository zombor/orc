package main

import (
	"flag"
	"fmt"
	stdlog "log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"

	"github.com/zombor/orc"
	"github.com/zombor/orc/pb"
)

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		grpcAddr = fs.String("grpc.addr", ":8090", "Address for gRPC server")
	)

	// package log
	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC).With("caller", log.DefaultCaller)
		stdlog.SetFlags(0)                             // flags are handled by Go kit's logger
		stdlog.SetOutput(log.NewStdlibAdapter(logger)) // redirect anything using stdlib log to us
	}

	// Business domain
	var svc orc.NodeService
	{
		svc = orc.NewNodeService()
	}

	errc := make(chan error)

	go func() {
		errc <- interrupt()
	}()

	go func() {
		transportLogger := log.NewContext(logger).With("transport", "gRPC")
		ln, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errc <- err
			return
		}

		s := grpc.NewServer() // uses its own, internal context
		pb.RegisterControllerServer(s, grpcBinding{svc})
		transportLogger.Log("addr", *grpcAddr)
		errc <- s.Serve(ln)
	}()

	logger.Log("fatal", <-errc)
}

func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}
