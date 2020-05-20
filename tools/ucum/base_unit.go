package ucum

type BaseUnit struct {
	Text        string `xml:",chardata"`
	Code        string `xml:"Code,attr"`
	CODE        string `xml:"CODE,attr"`
	Dim         string `xml:"dim,attr"`
	Name        string `xml:"name"`
	PrintSymbol string `xml:"printSymbol"`
	Property    string `xml:"property"`
}
