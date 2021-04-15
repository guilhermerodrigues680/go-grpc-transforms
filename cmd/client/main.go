package main

import (
	"context"
	"grpc-transforms/pbmath"
	"os"
	"time"

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
	logger.Info("Client Starting...")

	serverAddr := "localhost:10000"
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// conn, err := grpc.Dial(serverAddr, opts...)
	conn, err := grpc.DialContext(ctx, serverAddr, opts...)
	if err != nil {
		//logger.Fatalf("fail to dial: %v", err)
		logger.WithError(err).Fatal("Fail to dial")
	}

	logger.Info("Connection created successfully!")
	defer conn.Close()
	client := pbmath.NewMathServiceClient(conn)
	add(client)
}

func add(client pbmath.MathServiceClient) {

}
