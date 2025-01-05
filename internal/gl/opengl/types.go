package gl

func Uint16ToInt32(u16 []uint16) []int32 {
	i32 := make([]int32, len(u16))

	for i, v := range u16 {
		i32[i] = int32(v)
	}
	return i32
}
