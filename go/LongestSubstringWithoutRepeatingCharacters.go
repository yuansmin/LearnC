package main

import "fmt"

func lengthOfLongestSubstring_slow(s string) int {
	max, cur := 0, 0
	heap := make([]byte, 0)
	chars := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := chars[s[i]]; ok {
			if cur > max {
				max = cur
			}
			for {
				key := heap[0]
				delete(chars, key)
				heap = heap[1:]
				cur -= 1
				if key == s[i] {
					break
				}
			}
		}
		cur += 1
		chars[s[i]] = 0
		heap = append(heap, s[i])
	}
	if cur > max {
		max = cur
	}
	return max
}

func lengthOfLongestSubstring(s string) int {
	max, counter := 0, 0
	sub := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		oldi, ok := sub[s[i]]
		if ok && oldi >= i-counter {
			counter = i - oldi - 1
		}
		counter += 1
		sub[s[i]] = i
		if counter > max {
			max = counter
		}
	}
	return max
}

func testLengthOfLongestSubstring() {
	cases := []string{"abcabcbb", "bbbbb", "pwwkew", "dvdf"}
	for i := 0; i < len(cases); i++ {
		fmt.Printf("%s: %d\n", cases[i], lengthOfLongestSubstring(cases[i]))
	}
}
