package distribution

import (
	"math"
	"task6/internal/chars"
)

type HumanDistributionCalculator struct {
	work      float64
	time      float64
	workChars *chars.WorkCharacteristics
	timeChars *chars.TimeCharacteristics
}

// len is time of project in months (rounded up)
type HumanDistribution []int

func NewHumanDistributionCalculator(workChars *chars.WorkCharacteristics, timeChars *chars.TimeCharacteristics) *HumanDistributionCalculator {
	return &HumanDistributionCalculator{
		work:      workChars.Planning + workChars.Design + workChars.Construction + workChars.CodingTesting + workChars.IntegrationTesting,
		time:      timeChars.Planning + timeChars.Design + timeChars.Construction + timeChars.CodingTesting + timeChars.IntegrationTesting,
		workChars: workChars,
		timeChars: timeChars,
	}
}

func (c *HumanDistributionCalculator) Calculate() HumanDistribution {
	res := make([]int, int(math.Ceil(c.time)))

	periods := []struct {
		work float64
		time float64
	}{
		{c.workChars.Planning, c.timeChars.Planning},
		{c.workChars.Design, c.timeChars.Design},
		{c.workChars.Construction, c.timeChars.Construction},
		{c.workChars.CodingTesting, c.timeChars.CodingTesting},
		{c.workChars.IntegrationTesting, c.timeChars.IntegrationTesting},
	}

	startMonth := 0
	for _, p := range periods {
		if p.time <= 0 {
			continue
		}
		humans := int(math.Ceil(p.work / p.time))
		months := int(math.Ceil(p.time))
		for j := 0; j < months && startMonth+j < len(res); j++ {
			res[startMonth+j] = humans
		}
		startMonth += months
		if startMonth >= len(res) {
			break
		}
	}

	return res
}
