package leetcode

func addBinary(a, b string) string {
	ret := ""
	i := len(a) - 1
	j := len(b) - 1
	sign := 0
	for i > -1 || j > -1 {
		if i > -1 && a[i] == '1' {
			sign++
		}
		i--
		if j > -1 && b[j] == '1' {
			sign++

		}
		j--
		if sign%2 == 1 {
			ret += "1"
		} else {
			ret += "0"
		}
		sign /= 2
	}
	if sign == 1 {
		ret += "1"
	}
	temp := ""
	for i := len(ret) - 1; i > -1; i-- {
		temp += string(ret[i])
	}
	return temp
}
