package leetcode

func romanToInt(s string) int {
	m := map[string]int{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000}
	result := 0
	for i := 0; i < len(s); {
		if i+1 < len(s) {
			v, ok := m[s[i:i+2]]
			if ok {
				result += v
				i = i + 2
				continue
			}
		}
		v := m[s[i:i+1]]
		result += v
		i++
	}
	return result
}
