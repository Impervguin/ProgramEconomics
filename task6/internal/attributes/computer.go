package attributes

// COCOMO computer attributes
type ComputerAttributes struct {
	Time float64
	Stor float64
	Virt float64
	Turn float64
}

var (
	timeAttributeMap = map[AttributeLevel]float64{
		VeryLow:  1.,
		Low:      1.,
		Nominal:  1.,
		High:     1.11,
		VeryHigh: 1.5,
	}
	storageAttributeMap = map[AttributeLevel]float64{
		VeryLow:  1.,
		Low:      1.,
		Nominal:  1.,
		High:     1.06,
		VeryHigh: 1.21,
	}
	virtualizationAttributeMap = map[AttributeLevel]float64{
		VeryLow:  0.87,
		Low:      0.87,
		Nominal:  1.,
		High:     1.15,
		VeryHigh: 1.3,
	}
	turnaroundAttributeMap = map[AttributeLevel]float64{
		VeryLow:  0.87,
		Low:      0.87,
		Nominal:  1.,
		High:     1.07,
		VeryHigh: 1.15,
	}
)

type ComputerAttributesConfig struct {
	Time AttributeLevel
	Stor AttributeLevel
	Virt AttributeLevel
	Turn AttributeLevel
}

func MapComputerAttributes(c *ComputerAttributesConfig) *ComputerAttributes {
	return &ComputerAttributes{
		Time: timeAttributeMap[c.Time],
		Stor: storageAttributeMap[c.Stor],
		Virt: virtualizationAttributeMap[c.Virt],
		Turn: turnaroundAttributeMap[c.Turn],
	}
}
