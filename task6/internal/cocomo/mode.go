package cocomo

type COCOMOMode string

const (
	BasicMode        = "basic"
	IntermediateMode = "intermediate"
	EmbeddedMode     = "embedded"
)

func defaultCOCOMOMode(kloc float64) COCOMOMode {
	if kloc < 50 {
		return BasicMode
	} else if kloc < 500 {
		return IntermediateMode
	} else {
		return EmbeddedMode
	}
}
