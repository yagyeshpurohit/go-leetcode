package main

//https://leetcode.com/problems/isomorphic-strings

/*
	Approach1: Using two hash maps
	s(i) -> t(i)
	t(i) -> s(i)
*/
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// one-one mapping (S <-> T)
	charMapStoT := make(map[byte]byte)
	charMapTtoS := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		// if char of s is already in map
		if val, ok := charMapStoT[s[i]]; ok {
			// if corresponding char in t is not matching to the value stored in map, it means there exists two different mappings for the same char in s
			if val != t[i] {
				return false
			}
			continue
		} else if val, ok := charMapTtoS[t[i]]; ok {
			// if corresponding char in s is not matching to the value stored in map, it means there exists two different mappings for the same char in t
			if val != s[i] {
				return false
			}
		}
		// store the mapping s(i) -> t(i)
		charMapStoT[s[i]] = t[i]
		// store the mapping t(i) -> s(i)
		charMapTtoS[t[i]] = s[i]
	}
	return true
}
