package attributes

// COCOMO staff attributes
type StaffAttributes struct {
	ACAP float64
	AEXP float64
	PCAP float64
	VEXP float64
	LEXP float64
}

var (
	acapAttributeMap = map[AttributeLevel]float64{
		VeryLow:  1.46,
		Low:      1.19,
		Nominal:  1.,
		High:     0.86,
		VeryHigh: 0.71,
	}
	aexpAttributeMap = map[AttributeLevel]float64{
		VeryLow:  1.29,
		Low:      1.15,
		Nominal:  1.,
		High:     0.91,
		VeryHigh: 0.82,
	}
	pcapAttributeMap = map[AttributeLevel]float64{
		VeryLow:  1.42,
		Low:      1.17,
		Nominal:  1.,
		High:     0.86,
		VeryHigh: 0.7,
	}
	vexpAttributeMap = map[AttributeLevel]float64{
		VeryLow:  1.21,
		Low:      1.1,
		Nominal:  1.,
		High:     0.9,
		VeryHigh: 0.9,
	}
	lexpAttributeMap = map[AttributeLevel]float64{
		VeryLow:  1.14,
		Low:      1.07,
		Nominal:  1.,
		High:     0.95,
		VeryHigh: 0.95,
	}
)

type StaffAttributesConfig struct {
	ACAP AttributeLevel
	AEXP AttributeLevel
	PCAP AttributeLevel
	VEXP AttributeLevel
	LEXP AttributeLevel
}

func MapStaffAttributes(c *StaffAttributesConfig) *StaffAttributes {
	return &StaffAttributes{
		ACAP: acapAttributeMap[c.ACAP],
		AEXP: aexpAttributeMap[c.AEXP],
		PCAP: pcapAttributeMap[c.PCAP],
		VEXP: vexpAttributeMap[c.VEXP],
		LEXP: lexpAttributeMap[c.LEXP],
	}
}
