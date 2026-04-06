package functional

type ParamLevel string

var (
	NoInfluence        ParamLevel = "NoInfluence"
	RandomInfluence    ParamLevel = "RandomInfluence"
	SmallInfluence     ParamLevel = "SmallInfluence"
	MediumInfluence    ParamLevel = "MediumInfluence"
	ImportantInfluence ParamLevel = "ImportantInfluence"
	MainInfluence      ParamLevel = "MainInfluence"
)

var (
	influenceMap = map[ParamLevel]float64{
		NoInfluence:        0,
		RandomInfluence:    0.01,
		SmallInfluence:     0.02,
		MediumInfluence:    0.03,
		ImportantInfluence: 0.04,
		MainInfluence:      0.05,
	}
)

type ParamConfig struct {
	DataPassing              ParamLevel
	DistributedProcessing    ParamLevel
	Efficiency               ParamLevel
	ExpluatationRequirements ParamLevel
	TransacrionFrequency     ParamLevel
	FastIO                   ParamLevel
	UserEffectiveness        ParamLevel
	UpdateEfficiency         ParamLevel
	ProcessingComplexity     ParamLevel
	Reusability              ParamLevel
	InstallationLightness    ParamLevel
	ExpluatationLightness    ParamLevel
	PlatformDiversity        ParamLevel
	Maintainability          ParamLevel
}

func (c *ParamConfig) Sum() float64 {
	arr := []ParamLevel{
		c.DataPassing,
		c.DistributedProcessing,
		c.Efficiency,
		c.ExpluatationRequirements,
		c.TransacrionFrequency,
		c.FastIO,
		c.UserEffectiveness,
		c.UpdateEfficiency,
		c.ProcessingComplexity,
		c.Reusability,
		c.InstallationLightness,
		c.ExpluatationLightness,
		c.PlatformDiversity,
		c.Maintainability,
	}
	var sum float64
	for _, level := range arr {
		sum += influenceMap[level]
	}
	return sum
}

var defaultConfig = ParamConfig{
	DataPassing:              RandomInfluence,
	DistributedProcessing:    RandomInfluence,
	Efficiency:               RandomInfluence,
	ExpluatationRequirements: RandomInfluence,
	TransacrionFrequency:     RandomInfluence,
	FastIO:                   RandomInfluence,
	UserEffectiveness:        RandomInfluence,
	UpdateEfficiency:         RandomInfluence,
	ProcessingComplexity:     RandomInfluence,
	Reusability:              RandomInfluence,
	InstallationLightness:    RandomInfluence,
	ExpluatationLightness:    RandomInfluence,
	PlatformDiversity:        RandomInfluence,
	Maintainability:          RandomInfluence,
}

type WithInfluenceOption func(config *ParamConfig)

func WithDataPassing(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.DataPassing = level
	}
}

func WithDistributedProcessing(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.DistributedProcessing = level
	}
}

func WithEfficiency(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.Efficiency = level
	}
}

func WithExpluatationRequirements(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.ExpluatationRequirements = level
	}
}

func WithTransaccrionFrequency(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.TransacrionFrequency = level
	}
}

func WithFastIO(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.FastIO = level
	}
}

func WithUserEffectiveness(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.UserEffectiveness = level
	}
}

func WithUpdateEfficiency(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.UpdateEfficiency = level
	}
}

func WithProcessingComplexity(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.ProcessingComplexity = level
	}
}

func WithReusability(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.Reusability = level
	}
}

func WithInstallationLightness(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.InstallationLightness = level
	}
}

func WithExpluatationLightness(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.ExpluatationLightness = level
	}
}

func WithPlatformDiversity(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.PlatformDiversity = level
	}
}

func WithMaintainability(level ParamLevel) WithInfluenceOption {
	return func(config *ParamConfig) {
		config.Maintainability = level
	}
}