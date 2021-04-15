package transport

import "grpc-transforms/pbmath"

// PbNumToNum convert gRPC to model
func PbNumToNum(pbNum *pbmath.Num) *Num {
	return &Num{
		Num: pbNum.GetNum(),
	}
}

// NumToPbNum convert model to gRPC
func NumToPbNum(num *Num) *pbmath.Num {
	return &pbmath.Num{
		Num: num.Num,
	}
}
