package team

type TeamTeamParameterLevel string

var (
	// Сильно затрудненное взаимодействие
	SeverelyHindered TeamTeamParameterLevel = "SeverelyHindered"

	// Несколько затрудненное взаимодействие
	SomewhatHindered TeamTeamParameterLevel = "SomewhatHindered"

	// Некоторая согласованность
	SomeCohesion TeamTeamParameterLevel = "SomeCohesion"

	// Повышенная согласованность
	IncreasedCohesion TeamTeamParameterLevel = "IncreasedCohesion"

	// Высокая согласованность
	HighCohesion TeamTeamParameterLevel = "HighCohesion"

	// Взаимодействие как в едином целом
	UnifiedWhole TeamTeamParameterLevel = "UnifiedWhole"
)

var (
	teamParameterMap = map[TeamTeamParameterLevel]float64{
		SeverelyHindered:  5.48,
		SomewhatHindered:  4.38,
		SomeCohesion:      3.29,
		IncreasedCohesion: 2.19,
		HighCohesion:      1.1,
		UnifiedWhole:      0.,
	}
)

func (l TeamTeamParameterLevel) TeamParameter() float64 {
	return teamParameterMap[l]
}
