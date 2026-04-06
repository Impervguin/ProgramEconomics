package functional

type FunctionalPointsService struct {
}

type ElementOption struct {
	Element FunctionalElement
	Count   uint
	Level   ElementLevel
}

func (s *FunctionalPointsService) CalculatePoints(options []ElementOption) float64 {
	return s.points(options)
}

func (s *FunctionalPointsService) points(options []ElementOption) float64 {
	var points float64
	for _, option := range options {
		switch option.Level {
		case Low:
			points += float64(option.Count) * option.Element.Low()
		case Normal:
			points += float64(option.Count) * option.Element.Normal()
		case High:
			points += float64(option.Count) * option.Element.High()
		}
	}
	return points
}

func (s *FunctionalPointsService) CorrectedPoints(options []ElementOption, withParams ...WithInfluenceOption) float64 {
	points := s.points(options)
	config := defaultConfig
	for _, with := range withParams {
		with(&config)
	}
	mult := 0.65 + config.Sum()
	return points * mult
}

func (s *FunctionalPointsService) FPtoLoC(fp float64, language Language) float64 {
	return fp * float64(locPerFP[language])
}
