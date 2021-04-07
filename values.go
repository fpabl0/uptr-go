package uptr

// BoolValueWithFallback tries to get the bool value from a bool pointer, however
// if it is nil, then it will return the fallback bool.
func BoolValueWithFallback(b *bool, fallback bool) bool {
	if b == nil {
		return fallback
	}
	return *b
}

// IntValueWithFallback tries to get the integer value from an int pointer, however
// if it is nil, then it will return the fallback int.
func IntValueWithFallback(i *int, fallback int) int {
	if i == nil {
		return fallback
	}
	return *i
}

// Float32ValueWithFallback tries to get the float32 value from a float32 pointer, however
// if it is nil, then it will return the fallback float32.
func Float32ValueWithFallback(f *float32, fallback float32) float32 {
	if f == nil {
		return fallback
	}
	return *f
}

// Float64ValueWithFallback tries to get the float64 value from a float64 pointer, however
// if it is nil, then it will return the fallback float64.
func Float64ValueWithFallback(f *float64, fallback float64) float64 {
	if f == nil {
		return fallback
	}
	return *f
}

// StringValueWithFallback tries to get the string value from a string pointer, however
// if it is nil, then it will return the fallback string.
func StringValueWithFallback(s *string, fallback string) string {
	if s == nil {
		return fallback
	}
	return *s
}

// ===============================================================
// Zero Fallback
// ===============================================================

// BoolValue tries to get the bool value from a bool pointer, however
// if it is nil, then it will return the bool zero value.
func BoolValue(b *bool) bool {
	var zero bool
	return BoolValueWithFallback(b, zero)
}

// IntValue tries to get the integer value from an int pointer, however
// if it is nil, then it will return the int zero value.
func IntValue(i *int) int {
	var zero int
	return IntValueWithFallback(i, zero)
}

// Float32Value tries to get the float32 value from a float32 pointer, however
// if it is nil, then it will return the float32 zero value.
func Float32Value(f *float32) float32 {
	var zero float32
	return Float32ValueWithFallback(f, zero)
}

// Float64Value tries to get the float64 value from a float64 pointer, however
// if it is nil, then it will return the float64 zero value.
func Float64Value(f *float64) float64 {
	var zero float64
	return Float64ValueWithFallback(f, zero)
}

// StringValue tries to get the string value from a string pointer, however
// if it is nil, then it will return the string zero value.
func StringValue(s *string) string {
	var zero string
	return StringValueWithFallback(s, zero)
}
