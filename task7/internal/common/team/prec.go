package team

type PrecTeamParameterLevel string

var (
	// Полное отсутствие прецедентов, полностью непредсказуемый проект
	NewChaoticProject PrecTeamParameterLevel = "NewChaoticProject"

	// Почти полное отсутствие прецедентов, в значительной мере непредсказуемый проект
	NewProject PrecTeamParameterLevel = "NewProject"

	// Наличие некоторого количества прецедентов
	SomeExperienceProject PrecTeamParameterLevel = "SomeExperienceProject"

	// Общее знакомство с проектом
	CommonExperienceProject PrecTeamParameterLevel = "CommonExperienceProject"

	// Значительное знакомство с проектом
	ExperiencedProject PrecTeamParameterLevel = "ExperiencedProject"

	// Полное знакомство с проектом
	FullyKnownProject PrecTeamParameterLevel = "FullyKnownProject"
)

var (
	precParameterMap = map[PrecTeamParameterLevel]float64{
		NewChaoticProject:       6.2,
		NewProject:              4.96,
		SomeExperienceProject:   3.72,
		CommonExperienceProject: 2.48,
		ExperiencedProject:      1.24,
		FullyKnownProject:       0.,
	}
)

func (l PrecTeamParameterLevel) PrecParameter() float64 {
	return precParameterMap[l]
}
