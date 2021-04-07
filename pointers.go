package uptr

// Bool takes a bool value and converts it into a pointer.
func Bool(b bool) *bool {
	return &b
}

// Float32 takes a float32 value and converts it into a pointer.
func Float32(f float32) *float32 {
	return &f
}

// Float64 takes a float64 value and converts it into a pointer.
func Float64(f float64) *float64 {
	return &f
}

// Int takes an int value and converts it into a pointer.
func Int(i int) *int {
	return &i
}

// String takes a string value and converts it into a pointer.
func String(s string) *string {
	return &s
}
