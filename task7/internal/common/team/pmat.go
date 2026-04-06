package team

type PmatTeamParameterLevel string

var (
	// Уровень 1 СММ
	CMMLevel1 PmatTeamParameterLevel = "CMMLevel1"

	// Уровень 1+ СММ
	CMMLevel1Plus PmatTeamParameterLevel = "CMMLevel1Plus"

	// Уровень 2 СММ
	CMMLevel2 PmatTeamParameterLevel = "CMMLevel2"

	// Уровень 3 СММ
	CMMLevel3 PmatTeamParameterLevel = "CMMLevel3"

	// Уровень 4 СММ
	CMMLevel4 PmatTeamParameterLevel = "CMMLevel4"

	// Уровень 5 СММ
	CMMLevel5 PmatTeamParameterLevel = "CMMLevel5"
)

var (
	pmatParameterMap = map[PmatTeamParameterLevel]float64{
		CMMLevel1:     7.8,
		CMMLevel1Plus: 6.24,
		CMMLevel2:     4.68,
		CMMLevel3:     3.12,
		CMMLevel4:     1.56,
		CMMLevel5:     0.,
	}
)

func (l PmatTeamParameterLevel) PmatParameter() float64 {
	return pmatParameterMap[l]
}
