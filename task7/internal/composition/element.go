package composition

type ObjectiveElement interface {
	Low() float64
	Normal() float64
	High() float64
}

type genericObjectiveElement struct {
	low    float64
	normal float64
	high   float64
}

func (f *genericObjectiveElement) Low() float64 {
	return f.low
}

func (f *genericObjectiveElement) Normal() float64 {
	return f.normal
}

func (f *genericObjectiveElement) High() float64 {
	return f.high
}

var (
	ViewElement   = genericObjectiveElement{low: 1, normal: 2, high: 3}
	ReportElement = genericObjectiveElement{low: 2, normal: 5, high: 8}
)

type ElementLevel string

var (
	Simple  ElementLevel = "Simple"
	Complex ElementLevel = "Complex"
	Medium  ElementLevel = "Medium"
)
