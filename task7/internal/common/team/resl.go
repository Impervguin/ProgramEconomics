package team

type ReslTeamParameterLevel string

var (
	// Малое (20 %)
	Little20 ReslTeamParameterLevel = "Little20"

	// Некоторое (40 %)
	Some40 ReslTeamParameterLevel = "Some40"

	// Частое (60 %)
	Frequent60 ReslTeamParameterLevel = "Frequent60"

	// В целом (75 %)
	Overall75 ReslTeamParameterLevel = "Overall75"

	// Почти полное (90 %)
	AlmostFull90 ReslTeamParameterLevel = "AlmostFull90"

	// Полное (100%)
	Full100 ReslTeamParameterLevel = "Full100"
)

var (
	reslParameterMap = map[ReslTeamParameterLevel]float64{
		Little20:     7.,
		Some40:       5.65,
		Frequent60:   4.24,
		Overall75:    2.83,
		AlmostFull90: 1.41,
		Full100:      0.,
	}
)

func (l ReslTeamParameterLevel) ReslParameter() float64 {
	return reslParameterMap[l]
}
