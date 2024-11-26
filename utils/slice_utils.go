package utils

func IntSliceToInt64[T int | int8 | int16 | int32 | int64](arr []T) []int64 {
	result := make([]int64, len(arr), len(arr))

	for i, v := range arr {
		result[i] = int64(v)
	}

	return result
}

func UintSliceToUint64[T uint | uint8 | uint16 | uint32 | uint64](arr []T) []uint64 {
	result := make([]uint64, len(arr), len(arr))

	for i, v := range arr {
		result[i] = uint64(v)
	}
	return result
}
