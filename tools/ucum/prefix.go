package ucum

type Prefix struct {
	Text        string `xml:",chardata"`
	Code        string `xml:"Code,attr"`
	CODE        string `xml:"CODE,attr"`
	Name        string `xml:"name"`
	PrintSymbol string `xml:"printSymbol"`
	Value       struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Sup   string `xml:"sup"`
	} `xml:"value"`
}
