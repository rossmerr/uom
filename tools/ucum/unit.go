package ucum

type Unit struct {
	Text        string      `xml:",chardata"`
	Code        string      `xml:"Code,attr"`
	CODE        string      `xml:"CODE,attr"`
	IsMetric    string      `xml:"isMetric,attr"`
	Class       string      `xml:"class,attr"`
	IsSpecial   string      `xml:"isSpecial,attr"`
	IsArbitrary string      `xml:"isArbitrary,attr"`
	Name        []string    `xml:"name"`
	PrintSymbol PrintSymbol `xml:"printSymbol"`
	Property    string      `xml:"property"`
	Value       Value       `xml:"value"`
}

type Value struct {
	Text     string `xml:",chardata"`
	Unit     string `xml:"Unit,attr"`
	UNIT     string `xml:"UNIT,attr"`
	Value    string `xml:"value,attr"`
	Function struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"name,attr"`
		Value string `xml:"value,attr"`
		Unit  string `xml:"Unit,attr"`
	} `xml:"function"`
	Sup string `xml:"sup"`
}

type PrintSymbol struct {
	Text string `xml:",chardata"`
	Sup  string `xml:"sup"`
	Sub  struct {
		Text string `xml:",chardata"`
		R    string `xml:"r"`
	} `xml:"sub"`
	I struct {
		Text string `xml:",chardata"`
		Sub  struct {
			Text string `xml:",chardata"`
			R    string `xml:"r"`
		} `xml:"sub"`
	} `xml:"i"`
}
