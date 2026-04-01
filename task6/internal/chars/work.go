package chars

type WorkCharacteristics struct {
	Planning           float64
	Design             float64
	Construction       float64
	CodingTesting      float64
	IntegrationTesting float64
}

var (
	WorkCharacteristicsCoefficients WorkCharacteristics = WorkCharacteristics{
		Planning:           0.08,
		Design:             0.18,
		Construction:       0.25,
		CodingTesting:      0.26,
		IntegrationTesting: 0.31,
	}
)

// work in human-months
func CalculateWorkCharacteristics(work float64) WorkCharacteristics {
	return WorkCharacteristics{
		Planning:           work * WorkCharacteristicsCoefficients.Planning,
		Design:             work * WorkCharacteristicsCoefficients.Design,
		Construction:       work * WorkCharacteristicsCoefficients.Construction,
		CodingTesting:      work * WorkCharacteristicsCoefficients.CodingTesting,
		IntegrationTesting: work * WorkCharacteristicsCoefficients.IntegrationTesting,
	}
}
