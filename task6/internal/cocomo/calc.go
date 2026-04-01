package cocomo

import (
	"math"
	"task6/internal/attributes"
)

type COCOMOCalculator struct {
	mode  COCOMOMode
	Kloc  float64
	attrs *attributes.COCOMOAttributes
}

func NewCOCOMOCalculator(kloc float64, attrs *attributes.COCOMOAttributes) *COCOMOCalculator {
	return &COCOMOCalculator{
		mode:  defaultCOCOMOMode(kloc),
		Kloc:  kloc,
		attrs: attrs,
	}
}

func (c *COCOMOCalculator) SetMode(mode COCOMOMode) {
	c.mode = mode
}

// Returns work in human-months and time in months
func (c *COCOMOCalculator) Calculate() (float64, float64) {
	params := paramsMap[c.mode]
	work := params.WorkCoefficient * c.attrs.EAF() * math.Pow(c.Kloc, params.WorkPower)
	time := params.TimeCoefficient * math.Pow(work, params.TimePower)
	return work, time
}
