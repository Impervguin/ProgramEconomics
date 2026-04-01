package chars

type TimeCharacteristics struct {
	Planning           float64
	Design             float64
	Construction       float64
	CodingTesting      float64
	IntegrationTesting float64
}

var (
	TimeCharacteristicsCoefficients TimeCharacteristics = TimeCharacteristics{
		Planning:           0.36,
		Design:             0.36,
		Construction:       0.18,
		CodingTesting:      0.18,
		IntegrationTesting: 0.28,
	}
)

// time in months
func CalculateTimeCharacteristics(time float64) TimeCharacteristics {
	return TimeCharacteristics{
		Planning:           time * TimeCharacteristicsCoefficients.Planning,
		Design:             time * TimeCharacteristicsCoefficients.Design,
		Construction:       time * TimeCharacteristicsCoefficients.Construction,
		CodingTesting:      time * TimeCharacteristicsCoefficients.CodingTesting,
		IntegrationTesting: time * TimeCharacteristicsCoefficients.IntegrationTesting,
	}
}


