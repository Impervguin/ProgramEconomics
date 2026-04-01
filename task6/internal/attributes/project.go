package attributes

type ProjectAttributes struct {
	MODP float64
	TOOL float64
	SCED float64
}

var (
	modpAttributeMap = map[AttributeLevel]float64{
		VeryLow:  1.24,
		Low:      1.1,
		Nominal:  1.,
		High:     0.91,
		VeryHigh: 0.82,
	}
	toolAttributeMap = map[AttributeLevel]float64{
		VeryLow:  1.24,
		Low:      1.1,
		Nominal:  1.,
		High:     0.91,
		VeryHigh: 0.82,
	}
	scedAttributeMap = map[AttributeLevel]float64{
		VeryLow:  1.23,
		Low:      1.08,
		Nominal:  1.,
		High:     1.04,
		VeryHigh: 1.1,
	}
)

type ProjectAttributesConfig struct {
	MODP AttributeLevel
	TOOL AttributeLevel
	SCED AttributeLevel
}

func MapProjectAttributes(c *ProjectAttributesConfig) *ProjectAttributes {
	return &ProjectAttributes{
		MODP: modpAttributeMap[c.MODP],
		TOOL: toolAttributeMap[c.TOOL],
		SCED: scedAttributeMap[c.SCED],
	}
}
