package composition

type TeamProductivityLevel string

var (
	VeryLowProductivity  TeamProductivityLevel = "VeryLowProductivity"
	LowProductivity      TeamProductivityLevel = "LowProductivity"
	MediumProductivity   TeamProductivityLevel = "MediumProductivity"
	HighProductivity     TeamProductivityLevel = "HighProductivity"
	VeryHighProductivity TeamProductivityLevel = "VeryHighProductivity"
)

var (
	productivityMap = map[TeamProductivityLevel]float64{
		VeryLowProductivity:  4,
		LowProductivity:      7,
		MediumProductivity:   13,
		HighProductivity:     25,
		VeryHighProductivity: 50,
	}
)

func (l TeamProductivityLevel) Productivity() float64 {
	return productivityMap[l]
}
