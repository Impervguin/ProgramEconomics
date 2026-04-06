package functional

type FunctionalElement interface {
	Low() float64
	Normal() float64
	High() float64
}

type genericFunctionalElement struct {
	low    float64
	normal float64
	high   float64
}

func (f *genericFunctionalElement) Low() float64 {
	return f.low
}

func (f *genericFunctionalElement) Normal() float64 {
	return f.normal
}

func (f *genericFunctionalElement) High() float64 {
	return f.high
}

var (
	externalInputElement         = genericFunctionalElement{low: 3, normal: 4, high: 6}
	externalOutputElement        = genericFunctionalElement{low: 4, normal: 5, high: 6}
	externalQueryElement         = genericFunctionalElement{low: 3, normal: 4, high: 6}
	internalLogicalFileElement   = genericFunctionalElement{low: 7, normal: 10, high: 15}
	externalInterfaceFileElement = genericFunctionalElement{low: 5, normal: 7, high: 10}
)

type ElementLevel string

var (
	Low    ElementLevel = "Low"
	Normal ElementLevel = "Normal"
	High   ElementLevel = "High"
)
