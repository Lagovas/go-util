package uslice

//#begin-gt -gen.gt N:F64 T:float64

//	Appends v to sl only if sl does not already contain v.
func F64AppendUnique(ref *[]float64, v float64) {
	for _, sv := range *ref {
		if sv == v {
			return
		}
	}
	*ref = append(*ref, v)
}

//	Returns the position of val in slice.
func F64At(slice []float64, val float64) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

func F64EnsureCap(ref *[]float64, capacity int) {
	if cap(*ref) < capacity {
		F64SetCap(ref, capacity)
	}
}

func F64EnsureLen(ref *[]float64, length int) {
	if len(*ref) < length {
		F64SetLen(ref, length)
	}
}

//	Returns true if one and two only contain identical values, regardless of ordering.
func F64Equivalent(one, two []float64) bool {
	if len(one) != len(two) {
		return false
	}
	for _, v := range one {
		if F64At(two, v) < 0 {
			return false
		}
	}
	return true
}

//	Returns true if val is in slice.
func F64Has(slice []float64, val float64) bool {
	return F64At(slice, val) >= 0
}

//	Returns whether one of the specified vals is contained in slice.
func F64HasAny(slice []float64, vals ...float64) bool {
	for _, v1 := range vals {
		for _, v2 := range slice {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func F64Remove(ref *[]float64, v float64, all bool) {
	for i := 0; i < len(*ref); i++ {
		if (*ref)[i] == v {
			before, after := (*ref)[:i], (*ref)[i+1:]
			*ref = append(before, after...)
			if !all {
				break
			}
		}
	}
}

func F64SetCap(ref *[]float64, capacity int) {
	nu := make([]float64, len(*ref), capacity)
	copy(nu, *ref)
	*ref = nu
}

func F64SetLen(ref *[]float64, length int) {
	nu := make([]float64, length)
	copy(nu, *ref)
	*ref = nu
}

//	Removes all withoutVals from slice.
func F64Without(slice []float64, keepOrder bool, withoutVals ...float64) []float64 {
	if len(withoutVals) > 0 {
		for _, w := range withoutVals {
			for pos := F64At(slice, w); pos >= 0; pos = F64At(slice, w) {
				if keepOrder {
					slice = append(slice[:pos], slice[pos+1:]...)
				} else {
					slice[pos] = slice[len(slice)-1]
					slice = slice[:len(slice)-1]
				}
			}
		}
	}
	return slice
}

//#end-gt