package main

import (
	"grpc-transforms/pbmath"
	"grpc-transforms/transport"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func getLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
	// log all
	logger.SetLevel(logrus.TraceLevel)
	// Output to stdout instead of the default stderr
	logger.SetOutput(os.Stdout)
	return logger
}

func main() {
	logger := getLogger()
	logger.Info("Server Starting...")

	serverAddr := "localhost:10000"
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	srv := transport.NewTransportGRPC(logger.WithField("transport", "gRPC"))
	pbmath.RegisterMathServiceServer(grpcServer, srv)
	logger.Infof("gRPC Listening on: %s", serverAddr)
	grpcServer.Serve(lis)
}
