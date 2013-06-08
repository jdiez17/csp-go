package csp

func contains(slc []Value, v Value) bool {
	for _, i := range slc {
		if i == v {
			return true
		}
	}

	return false
}
