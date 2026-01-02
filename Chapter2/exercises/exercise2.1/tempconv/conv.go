package tempconv

// CToF converts a Celsius temprature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temprature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temprature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k) + KelvinZeroC }
