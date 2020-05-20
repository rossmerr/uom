package length

import (
	"math/big"

	"github.com/RossMerr/uom"
)

// A Length represents the distances between two points in meters
type Length *big.Float

// NewLength with base unit of meters
func NewLength(v float64) Length {
	return big.NewFloat(v)
}

func Centimeter(v float64) Length{
	l := big.NewFloat(0).Mul(big.NewFloat(v), big.NewFloat(float64(uom.Yotta)))
	return l
}

// func Millimeters(v float64) Length{
// 	return NewLength(v * float64(uom.Centi))
// }
//
// func (s Length)Scale(p uom.Prefix) Length {
// 	v := int64(s) * int64(p)
// 	return Length(v)
// }
//
// func (s Length) Centimeters() uint64 {
// 	return uint64(s) + uint64(uom.Deka)
// }
//
//
// func (s Length) Millimeters() uint64 {
// 	return uint64(s) + uint64(uom.Hecto)
// }