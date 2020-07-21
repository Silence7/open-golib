package public

func MergeSilce(src1 []string, src2 []string) []string {
	var result []string

	if nil == src1 {
		copy(result, src2)
		return result
	}

	if nil == src2 {
		copy(result, src1)
		return result
	}

	// src1 = append(src1, src2...)
	copy(result, src1)
	copy(result[len(src1):], src2)

	return src1
}

func RemoveDuplicateByLoop(src []string) []string {
	var result []string
	for i := range src {
		flag := true
		for j := range result {
			if src[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, src[i])
		}
	}
	return result
}

func RemoveDuplicateByMap(src []string) []string {
	var result []string
	tempMap := map[string]byte{}
	for _, e := range src{
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l{
			result = append(result, e)
		}
	}

	return result
}

func RemoveDuplicate(src []string) []string {
	if len(src) < 1024 {
		return RemoveDuplicateByLoop(src)
	} else {
		return RemoveDuplicateByMap(src)
	}
}
