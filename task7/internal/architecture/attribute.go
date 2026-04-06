package architecture

type AttributeLevel string

var (
	ExtraLow  AttributeLevel = "ExtraLow"
	VeryLow   AttributeLevel = "VeryLow"
	Low       AttributeLevel = "Low"
	Nominal   AttributeLevel = "Nominal"
	High      AttributeLevel = "High"
	VeryHigh  AttributeLevel = "VeryHigh"
	ExtraHigh AttributeLevel = "ExtraHigh"
)

type architectureAttributes struct {
	PERS float64
	RCPX float64
	RUSE float64
	PDIF float64
	PREX float64
	FCIL float64
	SCED float64
}

func (a *architectureAttributes) Multiplier() float64 {
	return a.PERS * a.RCPX * a.RUSE * a.PDIF * a.PREX * a.FCIL * a.SCED
}

type ArchitectureAttributesConfig struct {
	PERS AttributeLevel
	RCPX AttributeLevel
	RUSE AttributeLevel
	PDIF AttributeLevel
	PREX AttributeLevel
	FCIL AttributeLevel
	SCED AttributeLevel
}

var defaultConfig = ArchitectureAttributesConfig{
	PERS: Nominal,
	RCPX: Nominal,
	RUSE: Nominal,
	PDIF: Nominal,
	PREX: Nominal,
	FCIL: Nominal,
	SCED: Nominal,
}

func DefaultArchitectureAttributesConfig() *ArchitectureAttributesConfig {
	var c = defaultConfig
	return &c
}

func mapArchitectureAttributes(c *ArchitectureAttributesConfig) *architectureAttributes {
	return &architectureAttributes{
		PERS: persAttributeMap[c.PERS],
		RCPX: rcpxAttributeMap[c.RCPX],
		RUSE: ruseAttributeMap[c.RUSE],
		PDIF: pdifAttributeMap[c.PDIF],
		PREX: prexAttributeMap[c.PREX],
		FCIL: fcilAttributeMap[c.FCIL],
		SCED: scedAttributeMap[c.SCED],
	}
}

var (
	persAttributeMap = map[AttributeLevel]float64{
		ExtraLow:  2.12,
		VeryLow:   1.62,
		Low:       1.26,
		Nominal:   1.0,
		High:      0.83,
		VeryHigh:  0.63,
		ExtraHigh: 0.5,
	}
	rcpxAttributeMap = map[AttributeLevel]float64{
		ExtraLow:  0.49,
		VeryLow:   0.60,
		Low:       0.83,
		Nominal:   1.0,
		High:      1.33,
		VeryHigh:  1.91,
		ExtraHigh: 2.72,
	}
	ruseAttributeMap = map[AttributeLevel]float64{
		ExtraLow:  0.95,
		VeryLow:   0.95,
		Low:       0.95,
		Nominal:   1.0,
		High:      1.07,
		VeryHigh:  1.15,
		ExtraHigh: 1.24,
	}
	pdifAttributeMap = map[AttributeLevel]float64{
		ExtraLow:  0.87,
		VeryLow:   0.87,
		Low:       0.87,
		Nominal:   1.0,
		High:      1.29,
		VeryHigh:  1.81,
		ExtraHigh: 2.61,
	}
	prexAttributeMap = map[AttributeLevel]float64{
		ExtraLow:  1.59,
		VeryLow:   1.33,
		Low:       1.22,
		Nominal:   1.0,
		High:      0.87,
		VeryHigh:  0.73,
		ExtraHigh: 0.62,
	}
	fcilAttributeMap = map[AttributeLevel]float64{
		ExtraLow:  1.43,
		VeryLow:   1.30,
		Low:       1.10,
		Nominal:   1.0,
		High:      0.87,
		VeryHigh:  0.73,
		ExtraHigh: 0.62,
	}
	scedAttributeMap = map[AttributeLevel]float64{
		ExtraLow:  1.43,
		VeryLow:   1.43,
		Low:       1.14,
		Nominal:   1.0,
		High:      1.0,
		VeryHigh:  1.0,
		ExtraHigh: 1.0,
	}
)
