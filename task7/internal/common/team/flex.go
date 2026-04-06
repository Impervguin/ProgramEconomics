package team

type FlexTeamParameterLevel string

var (
	// Точный, строгий процесс разработки
	RigidProcess FlexTeamParameterLevel = "RigidProcess"

	// Случайные послабления в процессе
	RandomRelaxations FlexTeamParameterLevel = "RandomRelaxations"

	// Некоторые послабления в процессе
	SomeRelaxations FlexTeamParameterLevel = "SomeRelaxations"

	// Большей частью согласованный процесс
	MostlyAgreedProcess FlexTeamParameterLevel = "MostlyAgreedProcess"

	// Некоторое согласование процесса
	SomeAgreement FlexTeamParameterLevel = "SomeAgreement"

	// Заказчик определил только общие цели
	GeneralGoalsOnly FlexTeamParameterLevel = "GeneralGoalsOnly"
)

var (
	flexParameterMap = map[FlexTeamParameterLevel]float64{
		RigidProcess:        5.07,
		RandomRelaxations:   4.05,
		SomeRelaxations:     3.04,
		MostlyAgreedProcess: 2.03,
		SomeAgreement:       1.01,
		GeneralGoalsOnly:    0.,
	}
)

func (l FlexTeamParameterLevel) FlexParameter() float64 {
	return flexParameterMap[l]
}
