package vld

func NumberPtr(value *int) *numberVld {
	if value == nil {
		return &numberVld{
			zero: true,
		}
	}
	return &numberVld{
		value: float64(*value),
		zero:  false,
	}
}

func NumF32Ptr(value *float32) *numberVld {
	if value == nil {
		return &numberVld{
			zero: true,
		}
	}
	return &numberVld{
		value: float64(*value),
		zero:  false,
	}
}

func NumF64Ptr(value *float64) *numberVld {
	if value == nil {
		return &numberVld{
			zero: true,
		}
	}
	return &numberVld{
		value: *value,
		zero:  false,
	}
}

func NumI32Ptr(value *int32) *numberVld {
	if value == nil {
		return &numberVld{
			zero: true,
		}
	}
	return &numberVld{
		value: float64(*value),
		zero:  false,
	}
}

func NumI64Ptr(value *int64) *numberVld {
	if value == nil {
		return &numberVld{
			zero: true,
		}
	}
	return &numberVld{
		value: float64(*value),
		zero:  false,
	}
}
