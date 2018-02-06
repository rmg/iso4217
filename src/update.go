package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"text/template"
)

// <ISO_4217 Pblshd="2014-08-15">
//   <CcyTbl>
//     <CcyNtry>
//       <CtryNm>AFGHANISTAN</CtryNm>
//       <CcyNm>Afghani</CcyNm>
//       <Ccy>AFN</Ccy>
//       <CcyNbr>971</CcyNbr>
//       <CcyMnrUnts>2</CcyMnrUnts>
//     </CcyNtry>
//     ...

type entry struct {
	Ccy        string     `xml:"Ccy"`
	CcyNbr     int        `xml:"CcyNbr"`
	CcyMnrUnts minorUnits `xml:"CcyMnrUnts"`
}

type minorUnits uint

func (m *minorUnits) UnmarshalText(b []byte) error {
	newInt, err := strconv.ParseUint(string(b), 10, 0)
	if err == nil {
		*m = minorUnits(newInt)
	} else {
		*m = 0
	}
	return nil
}

func main() {
	dump := struct {
		Entries []entry `xml:"CcyTbl>CcyNtry,name"`
	}{}
	decoder := xml.NewDecoder(os.Stdin)
	err := decoder.Decode(&dump)
	strings := make(map[int]entry)
	for _, entry := range dump.Entries {
		strings[entry.CcyNbr] = entry
	}
	const tpl = `package iso4217

var names = map[int]string {
	{{range $num, $entry := .}}{{$num}}: "{{$entry.Ccy}}",
	{{end}}
}

var codes = map[string]int {
	{{range $num, $entry := .}}"{{$entry.Ccy}}": {{$num}},
	{{end}}
}

var minorUnits = map[int]int {
	{{range $num, $entry := .}}{{$num}}: {{$entry.CcyMnrUnts}},
	{{end}}
}

func ByCode(n int) (string, int) {
	return names[n], minorUnits[n]
}

func ByName(s string) (int, int) {
	code := codes[s]
	return code, minorUnits[code]
}
`
	t := template.Must(template.New("constants.go.in").Parse(tpl))
	if err == nil {
		t.Execute(os.Stdout, strings)
	} else {
		fmt.Printf("%+v\n", err)
	}
}
