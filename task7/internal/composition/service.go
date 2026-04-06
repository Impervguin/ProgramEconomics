package composition

import (
	"math"
	"task7/internal/common/team"
)

type CompositionService struct {
}

type ElementOption struct {
	Element ObjectiveElement
	Count   uint
	Level   ElementLevel
}

type CompositionModelDto struct {
	Elements          []ElementOption
	ModuleCount       uint
	ReuseParam        ReuseParameter
	ProductivityLevel TeamProductivityLevel
	AveragePay        float64
	TeamConfig        team.TeamConfig
}

type CompositionModelResult struct {
	Work     float64 // person-months
	Time     float64 // months
	Cost     float64 // money
	TeamSize uint
}

func (s *CompositionService) CalculateCompositionModel(dto CompositionModelDto) CompositionModelResult {
	return s.calculateCompositionModel(dto)
}

func (s *CompositionService) calculateCompositionModel(dto CompositionModelDto) CompositionModelResult {
	op := s.objectivePoints(dto.Elements, dto.ModuleCount)
	nop := op * (1. - float64(dto.ReuseParam))
	work := nop / dto.ProductivityLevel.Productivity()
	time := 3.67 * math.Pow(work, 0.28+0.2*(dto.TeamConfig.Power()-1.01))
	cost := dto.AveragePay * work
	teamSize := uint(math.Ceil(work / time))
	return CompositionModelResult{
		Work:     work,
		Time:     time,
		Cost:     cost,
		TeamSize: teamSize,
	}
}

func (s *CompositionService) objectivePoints(options []ElementOption, moduleCount uint) float64 {
	var points float64
	for _, option := range options {
		switch option.Level {
		case Simple:
			points += float64(option.Count) * option.Element.Low()
		case Complex:
			points += float64(option.Count) * option.Element.High()
		case Medium:
			points += float64(option.Count) * option.Element.Normal()
		}
	}
	points += 10. * float64(moduleCount)
	return points
}
