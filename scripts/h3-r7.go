package main

import (
	"encoding/json"
	"os"
	"time"

	h3 "github.com/uber/h3-go/v4"
)

type H3CellsRecord struct {
	Description string    `json:"description"`
	DateCreated time.Time `json:"dateCreated"`
	Resolution  uint16    `json:"resolution"`
	Cells       []string  `json:"cells"`
}

// get all r7 children of r6 cells
func main() {
	file, err := os.Open("./h3-cells/uk-h3-r6-trimmed.json")
	if err != nil {
		panic(err.Error())
	}
	dec := json.NewDecoder(file)
	var rec H3CellsRecord
	err = dec.Decode(&rec)
	if err != nil {
		panic(err.Error())
	}

	parents := make([]h3.Cell, len(rec.Cells))
	for i, parent := range rec.Cells {
		cell := h3.IndexFromString(parent)
		parents[i] = h3.Cell(cell)
	}

	cRecord := H3CellsRecord{
		Description: "r7 children of r6 cells.",
		DateCreated: time.Now(),
		Resolution: uint16(7),
	}
	for _, c := range parents {
		children := c.Children(7)

		for _, childCell := range children {
			cRecord.Cells = append(cRecord.Cells, childCell.String())
		}
	}

	res, err := os.Create("uk-h3-r7.json")
	if err != nil {
		panic(err.Error())
	}

	enc := json.NewEncoder(res)
	err = enc.Encode(cRecord)
	if err != nil {
		panic(err.Error())
	}
}
