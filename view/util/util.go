package util

func FindSubStr(s string) (subStr string, length int) {
	s1 := []rune(s)
	var tmpStr []rune
	for i := 0; i < len(s1)-1; i++ {
		subStr := findStr(s1[i:])
		if len(subStr) >= len(tmpStr) {
			tmpStr = subStr
			i = i + len(tmpStr) - 1
		}
	}

	return string(tmpStr), len(tmpStr)
}

// if giving string contains two or more sub string which length is same and longtest
func FindSubStr2(s string) []rune {
	s1 := []rune(s)
	var tmpStr []rune
	var strSet []rune
	for i := 0; i < len(s1)-1; i++ {
		subStr := findStr(s1[i:])
		if len(subStr) > len(tmpStr) {
			tmpStr = subStr
			i = i + len(tmpStr) - 1
			if len(strSet) > 0 {
				strSet = strSet[:0]
				strSet = append(strSet, tmpStr...)
			}
		} else if len(subStr) == len(tmpStr) {
			strSet = append(strSet, tmpStr...)
			strSet = append(strSet, subStr...)
			i = i + len(subStr) - 1
		}
	}

	return strSet
}

func findStr(s []rune) []rune {
	length := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			break
		}
		length++
	}

	return s[:length]
}

func MergeIntSclice(i1, i2 []int) []int {
	merged := append(i1, i2...)
	selectSort(merged)

	return merged
}

func selectSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		index := i
		num := s[i]
		for j := i + 1; j < len(s); j++ {
			if num > s[j] {
				index = j
				num = s[j]
			}
		}

		if index != i {
			s[i], s[index] = s[index], s[i]
		}
	}
}

func IsAbleToLast(a []int) bool {
	if len(a) == 0 {
		return true
	}

	if a[0] == 0 && len(a) == 1 {
		return true
	}

	if a[0] == 0 && len(a) > 1 {
		return false
	}
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
		if sum >= len(a) {
			return true
		} else if a[i] == 0 {
			return false
		}
		i = sum - 1
	}

	return false
}
