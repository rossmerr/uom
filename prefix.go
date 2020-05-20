package uom


const (
	Yotta Prefix = 1000000000000000000000000
	Zetta Prefix = 1000000000000000000000
	Exa   Prefix = 1000000000000000000
	Peta  Prefix = 1000000000000000
	Tera  Prefix = 1000000000000
	Giga  Prefix = 1000000000
	Mega  Prefix = 1000000
	Kilo  Prefix = 1000
	Hecto Prefix = 100
	Deka  Prefix = 10
	Deci  Prefix = 0.1
	Centi Prefix = 0.01
	Milli Prefix = 0.001
	Micro Prefix = 0.000001
	Nano  Prefix = 0.0000000001
	Pico  Prefix = 0.000000000001
	Femto Prefix = 0.000000000000001
	Atto  Prefix = 0.000000000000000001
	Zepto Prefix = 0.000000000000000000001
	Yocto Prefix = 0.000000000000000000000001
)

// SI official symbol
var PrefixPrint = map[Prefix]string{
	Yotta: "Y",
	Zetta: "Z",
	Exa:   "E",
	Peta:  "P",
	Tera:  "T",
	Giga:  "G",
	Mega:  "M",
	Kilo:  "k",
	Hecto: "h",
	Deka:  "da",
	Deci:  "d",
	Centi: "c",
	Milli: "m",
	Micro: "μ",
	Nano:  "n",
	Pico:  "p",
	Femto: "f",
	Atto:  "a",
	Zepto: "z",
	Yocto: "y",
}

// ISO 2955 and ANSI X3.50, where “giga-,” “tera-,” and “peta-” have been replaced by “G,” “T,” and “PE.”
// New prefixes “yotta-,” “zetta-,” “yocto-,” and “zepto-” that were adopted by the 19th CGPM (1990) have a second letter ‘A’ and ‘O’ resp. to avoid current and future conflicts.
var PrefixSymbolCaseInsensitive = map[Prefix]string{
	Yotta: "YA",
	Zetta: "ZA",
	Exa:   "EX",
	Peta:  "PT",
	Tera:  "TR",
	Giga:  "GA",
	Mega:  "MA",
	Kilo:  "K",
	Hecto: "H",
	Deka:  "DA",
	Deci:  "D",
	Centi: "C",
	Milli: "M",
	Micro: "U",
	Nano:  "N",
	Pico:  "P",
	Femto: "F",
	Atto:  "A",
	Zepto: "ZO",
	Yocto: "YO",
}
