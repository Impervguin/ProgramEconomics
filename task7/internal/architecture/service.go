package architecture

import (
	"math"
	"task7/internal/common/team"
)

type ArchitectureService struct {
}

type ArchitectureModelDto struct {
	ArchitectureAttributesConfig *ArchitectureAttributesConfig
	TeamConfig                   team.TeamConfig
	// kilo lines of code
	Kloc       float64
	AveragePay float64
}

type ArchitectureModelResult struct {
	Work     float64 // person-months
	Time     float64 // months
	Cost     float64 // money
	TeamSize uint
}

func (s *ArchitectureService) CalculateArchitectureModel(dto ArchitectureModelDto) ArchitectureModelResult {
	return s.calculateArchitectureModel(dto)
}

const workMultiplier = 2.94
const timeMultiplier = 3.67

func (s *ArchitectureService) calculateArchitectureModel(dto ArchitectureModelDto) ArchitectureModelResult {
	arch := mapArchitectureAttributes(dto.ArchitectureAttributesConfig)
	p := dto.TeamConfig.Power()
	work := workMultiplier *arch.Multiplier() * math.Pow(dto.Kloc, p)
	time := 3.67 * math.Pow(work, 0.28+0.2*(p-1.01))
	cost := dto.AveragePay * work
	teamSize := uint(math.Ceil(work / time))
	return ArchitectureModelResult{
		Work:     work,
		Time:     time,
		Cost:     cost,
		TeamSize: teamSize,
	}
}
