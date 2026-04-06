package team

type TeamConfig struct {
	TeamParameterLevel TeamTeamParameterLevel
	PmatParameterLevel PmatTeamParameterLevel
	PrecParameterLevel PrecTeamParameterLevel
	ReslParameterLevel ReslTeamParameterLevel
	FlexParameterLevel FlexTeamParameterLevel
}

var defaultConfig = TeamConfig{
	TeamParameterLevel: UnifiedWhole,
	PmatParameterLevel: CMMLevel5,
	PrecParameterLevel: FullyKnownProject,
	ReslParameterLevel: Full100,
	FlexParameterLevel: GeneralGoalsOnly,
}

const basePower = 1.01

func (c *TeamConfig) Power() float64 {
	return basePower +
		(c.TeamParameterLevel.TeamParameter()+
			c.PmatParameterLevel.PmatParameter()+
			c.PrecParameterLevel.PrecParameter()+
			c.ReslParameterLevel.ReslParameter()+
			c.FlexParameterLevel.FlexParameter())/100.
}
