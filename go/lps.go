// Longest Palindromic Substring
package main

import (
	"fmt"
)

// TODO:
// 1. 使用动态规划解决
//

func main() {
	strings := []string{"babad", "shitabcdcbalovebbb", "abbcd", "a", "aabbbbc", "aaaa", "aaaaa"}
	// fmt.Printf("%s, result: %s\n", "", longestPalindromeD(strings[5]))
	for _, v := range strings {
		fmt.Printf("%s, result: %s\n", v, longestPalindromeD(v))
	}
}

// 中心法
func longestPalindrome(s string) string {
	longest := -1
	length := len(s)
	low, high := 0, 0
	isSame := func(i, h int) bool {
		same := true
		j := h - 1
		for j >= i {
			if s[h] != s[j] {
				return false
			}
			j--
		}
		return same
	}
	for i, _ := range s {
		l, h := i-1, i+1
		count := 0
		stoped := false
		for l >= 0 && h < length {
			if s[l] != s[h] {
				if isSame(i, h) {
					h++
					continue
				}
				tLength := h - l - 1
				if tLength > longest {
					low = l + 1
					high = h
					longest = tLength
				}
				stoped = true
				break
			}
			l--
			h++
			count++
		}
		if !stoped {
			if h < length && isSame(i, h) {
				for h < length && s[h] == s[i] {
					h++
				}
			}
			tLength := h - l - 1
			if tLength > longest {
				low = l + 1
				high = h
				longest = tLength
			}
		}
		if count < 1 && (i-1) >= 0 {
			if (s[i] == s[i-1]) && longest < 2 {
				longest = 2
				low = i - 1
				high = i + 1
			}
		}
	}
	return s[low:high]
}

// 动态规划
func longestPalindromeD(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	low, high := 0, 0
	graph := make([][]bool, length)
	for i := 0; i < length; i++ {
		graph[i] = make([]bool, length)
	}

	for i := 0; i < length; i++ {
		graph[i][i] = true
	}

	for i := 1; i < length; i++ {
		if s[i-1] == s[i] {
			graph[i-1][i] = true
			// just get the first
			if low == 0 {
				low, high = i-1, i
			}
		}
	}

	for tLength := 3; tLength <= length; tLength++ {
		for i := 0; i <= length-tLength; i++ {
			j := i + tLength - 1
			if s[i] == s[j] && graph[i+1][j-1] {
				graph[i][j] = true
				// just get the first
				if tLength > (high - low + 1) {
					low, high = i, j
				}
			}
		}
	}
	// fmt.Printf("%s:\n", s)
	// printGraph(graph)

	return s[low : high+1]
}

func printGraph(graph [][]bool) {
	fmt.Printf("    ")
	for i := 0; i < len(graph); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
	for i := 0; i < len(graph); i++ {
		fmt.Printf("%.3d ", i)
		for j := 0; j < len(graph[i]); j++ {
			n := 0
			if graph[i][j] {
				n = 1
			}
			fmt.Printf("%d ", n)
		}
		fmt.Printf("\n")
	}
}

// 使用中心法则，通过中心向两端延伸逐步验证回文
// 中心的位置可能在某一元素，也可能在某两个元素中间
// 1. 中心为某一元素i，以left=right=i，逐步扩展验证(left--,right++)
// 2. 若中心为两元素之间，以这两个元素分别为left，right扩展
// 每个元素即可能为中心，也可能为中心的两边。通过二倍字符串长度取余遍历每个中心计算
func longestPalindromeX(s string) string {

	if s == "" {
		return ""
	}

	rl, rr, rw := -1, -1, -1
	var l, r, tl, tr, w int

	// 通过
	for i := 0; i < len(s)*2-1; i++ {

		l = i / 2
		r = l + i%2

		// validate palindrome

		for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
			if s[l] == s[r] {
				tl, tr = l, r
			} else {
				break
			}
		}

		w = tr - tl + 1

		if w > rw {
			rw, rl, rr = w, tl, tr
		}
	}

	return s[rl : rr+1]
}

func longestPalindromeFastest(s string) string {
	n := len(s)
	t := make([]byte, 2, 2*n+3)
	t[0], t[1] = '\n', 0
	for i := 0; i < n; i++ {
		t = append(t, s[i], 0)
	}
	t = append(t, 0)
	c, d, n := 0, 0, len(t)
	p := make([]int, n)
	for i, lim := 1, n-1; i < lim; i++ {
		mirror := 2*c - i
		if mirror < 0 {
			mirror += n
		}
		p[i] = max(0, min(p[mirror], d-i))
		for t[i+p[i]+1] == t[i-p[i]-1] {
			p[i]++
		}
		if d < i+p[i] {
			d = i + p[i]
			c = i
		}
	}
	c, d = 0, 0
	for i := 1; i < n; i++ {
		if d < p[i] {
			d = p[i]
			c = i
		}
	}
	return s[(c-d)>>1 : (c+d)>>1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
