package chapter4_1

func zero1(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}
