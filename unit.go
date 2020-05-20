package uom


type Unit struct {
	isBase bool
	name string
	csCode string
	ciCode string
	property string
	// The magnitude of the unit, e.g., 3600/3937 for a yard,
	// where a yard - 3600/3973 * m(eter).  The Dimension
	// property specifies the meter - which is the unit on which
	// a yard is based, and this magnitude specifies how to figure
	// this unit based on the base unit.
	magnitude float64
	dim *Dimension
	printSymbol string
	class string
	isMetric bool
	variable string
	cnv Conversion
	isSpecial bool
	isArbitrary bool
	moleExp int
}

func (s *Unit) ConvertFrom() {

}

type Conversion func()

type Dimension struct {

}