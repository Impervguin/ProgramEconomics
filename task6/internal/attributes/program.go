package attributes

// COCOMO program attributes
type ProgramAttributes struct {
	Rely float64
	Data float64
	CPLX float64
}

var (
	relyAttributeMap = map[AttributeLevel]float64{
		VeryLow:  0.75,
		Low:      0.86,
		Nominal:  1.,
		High:     1.15,
		VeryHigh: 1.3,
	}
	dataAttributeMap = map[AttributeLevel]float64{
		VeryLow:  0.94,
		Low:      0.94,
		Nominal:  1.,
		High:     1.08,
		VeryHigh: 1.16,
	}
	complexityAttributeMap = map[AttributeLevel]float64{
		VeryLow:  0.7,
		Low:      0.85,
		Nominal:  1.,
		High:     1.15,
		VeryHigh: 1.3,
	}
)

type ProgramAttributesConfig struct {
	Rely AttributeLevel
	Data AttributeLevel
	CPLX AttributeLevel
}

func MapProgramAttributes(c *ProgramAttributesConfig) *ProgramAttributes {
	return &ProgramAttributes{
		Rely: relyAttributeMap[c.Rely],
		Data: dataAttributeMap[c.Data],
		CPLX: complexityAttributeMap[c.CPLX],
	}
}
