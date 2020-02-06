package main

// n, m the length of s, sub string
// return first matched string start index
// return -1 if sub not in s
// now just for ascii strings
func bf(s string, n int, sub string, m int) int {
	if m > n {
		return -1
	}
	if m == n {
		if s == sub {
			return 0
		}
		return -1
	}

	bc := map[byte]int{}
	generateBC(bc, sub, m)
	prefix, suffix := generateGSIndex(sub, m)
	for i := 0; i <= n-m; {
		j := m - 1
		for j >= 0 && sub[j] == s[i+j] {
			j--
		}
		// find the matched index
		if j == -1 {
			return i
		}

		// bad caractor rule
		x := j - bc[s[i+j]]
		// good suffix rule
		y := 0
		if j < m-1 {
			y = moveByGS(sub, m, j, prefix, suffix)
		}
		if y > x {
			x = y
		}
		i += x
	}

	return -1
}

// key is char ascii code, value is char highest index
func generateBC(bc map[byte]int, s string, n int) {
	for i, c := range s {
		bc[byte(c)] = i
	}
}

// generate good suffix index: prefix, suffix
// prefix, suffix: index is the sub string length
// each length stands for a uniq suffix sub string
// suffix value: matched string start index, -1 if no matched string
// prefix value: true if matched prefix sub string
func generateGSIndex(s string, m int) ([]bool, []int) {
	prefix := make([]bool, m-1)
	suffix := make([]int, m-1)
	for i, _ := range suffix {
		suffix[i] = -1
	}

	for i := 0; i < m-1; i++ {
		j := i
		for j >= 0 && s[j] == s[m-1-j] {
			suffix[i-j+1] = j
			j--
		}
		if j == -1 {
			prefix[i-j] = true
		}
	}
	return prefix, suffix
}

// good suffix rule
// j: bad caractor index
func moveByGS(sub string, m, j int, prefix []bool, suffix []int) int {
	k := m - 1 - j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}

	for i := k - 1; i > 0; i-- {
		if prefix[i] {
			return j + k - i
		}
	}
	return m
}
