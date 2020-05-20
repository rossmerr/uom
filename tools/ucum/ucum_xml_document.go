package ucum

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"golang.org/x/text/encoding/charmap"
)

type UcumXmlDocument struct {
	XMLName      xml.Name   `xml:"root"`
	Text         string     `xml:",chardata"`
	Xmlns        string     `xml:"xmlns,attr"`
	Version      string     `xml:"version,attr"`
	Revision     string     `xml:"revision,attr"`
	RevisionDate string     `xml:"revision-date,attr"`
	Prefixs       []Prefix   `xml:"prefix"`
	BaseUnits     []BaseUnit `xml:"base-unit"`
	Units         []Unit     `xml:"unit"`
}

func Import() *UcumXmlDocument {
	file, err := ioutil.ReadFile("/home/rossmerr/git/uom/internal/data/ucum-essence.xml")
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewBuffer(file)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charsetReader
	dec := xml.NewTokenDecoder(trimmer{decoder}) // trimming decoder

	ucum := &UcumXmlDocument{}

	err = dec.Decode(&ucum)
	if err != nil {
		log.Fatal(err)
	}

	return ucum
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "ascii" {
		return charmap.Windows1252.NewDecoder().Reader(input), nil
	}

	return nil, fmt.Errorf("Unknown charset: %s", charset)
}

type trimmer struct {
	*xml.Decoder
}

func (tr trimmer) Token() (xml.Token, error) {
	t, err := tr.Decoder.Token()
	if cd, ok := t.(xml.CharData); ok {
		t = xml.CharData(bytes.TrimSpace(cd))
	}
	return t, err
}