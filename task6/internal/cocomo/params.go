package cocomo

type COCOMOParams struct {
	WorkCoefficient float64
	WorkPower       float64
	TimeCoefficient float64
	TimePower       float64
}

var (
	paramsMap = map[COCOMOMode]*COCOMOParams{
		BasicMode:        &COCOMOParams{3.2, 1.05, 2.5, 0.38},
		IntermediateMode: &COCOMOParams{3., 1.12, 2.5, 0.35},
		EmbeddedMode:     &COCOMOParams{2.8, 1.2, 2.5, 0.32},
	}
)
