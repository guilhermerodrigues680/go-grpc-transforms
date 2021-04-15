package transport

import (
	"context"
	"grpc-transforms/pbmath"

	"github.com/sirupsen/logrus"
)

type TransportGRPC struct {
	pbmath.UnimplementedMathServiceServer
	logger *logrus.Entry
}

type Num struct {
	Num float64
}

func NewTransportGRPC(logger *logrus.Entry) *TransportGRPC {
	return &TransportGRPC{
		logger: logger,
	}
}

func (t *TransportGRPC) Add(ctx context.Context, num *pbmath.Num) (*pbmath.Num, error) {
	return NumToPbNum(add(PbNumToNum(num), PbNumToNum(num))), nil
}

func add(x, y *Num) *Num {
	return &Num{
		Num: x.Num + y.Num,
	}
}
