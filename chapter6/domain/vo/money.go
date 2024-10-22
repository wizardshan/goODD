package vo

type Money struct {
}

func (o Money) ConvertFen(v float64) int64 {
	return int64(v * 100)
}

func (o Money) ConvertYuan(v int64) float64 {
	return float64(v) / 100
}
