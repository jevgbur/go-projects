package conversions

type Temperature struct {
	Value float64
	Unit  string
}

func CelsiusToFahrenheit(t *Temperature) *Temperature {
	t.Value = (t.Value * 1.8) + 32
	t.Unit = "F"
	return &Temperature{Value: t.Value, Unit: t.Unit}
}

func FahrenheitToCelsius(t *Temperature) *Temperature {
	t.Value = (t.Value - 32) / 1.8
	t.Unit = "C"
	return &Temperature{Value: t.Value, Unit: t.Unit}
}
