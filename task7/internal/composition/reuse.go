package composition

type ReuseParameter float64

func ReuseFromPercent(percent float64) ReuseParameter {
	if percent > 100 {
		percent = 100
	}
	if percent < 0 {
		percent = 0
	}
	return ReuseParameter(percent / 100.)
}

