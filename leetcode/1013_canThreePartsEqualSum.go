package leetcode

func canThreePartsEqualSum(A []int) bool {
	if len(A) < 3 {
		return false
	}
	ret := 0
	for _, v := range A {
		ret += v
	}
	if ret%3 != 0 {
		return false
	}
	s := ret / 3
	t := 0
	count := 0
	for i := 0; i < len(A); {
		t += A[i]
		if t == s {
			count++
			i = i + 1
			t = 0
			continue
		}

		j := i + 1
		for ; j < len(A); j++ {
			t += A[j]
			if t == s {
				count++
				t = 0
				break
			}
		}
		i = j + 1
	}
	if count >= 3 {
		return true
	}
	return false
}
