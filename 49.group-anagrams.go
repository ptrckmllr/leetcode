package main

import "fmt"

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (66.77%)
 * Likes:    15565
 * Dislikes: 446
 * Total Accepted:    2M
 * Total Submissions: 3M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings strs, group the anagrams together. You can return
 * the answer in any order.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 *
 *
 * Example 1:
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 * Example 2:
 * Input: strs = [""]
 * Output: [[""]]
 * Example 3:
 * Input: strs = ["a"]
 * Output: [["a"]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] consists of lowercase English letters.
 *
 *
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// generate lettercounts per string for letter count comparison
	letterCounts := []map[byte]int{}
	for _, str := range strs {
		letterMap := map[byte]int{}
		for i := 0; i < len(str); i++ {
			letterMap[str[i]]++
		}
		letterCounts = append(letterCounts, letterMap)
	}

	// find anagrams
	result := [][]string{}
	knownAnagrams := map[int]bool{}
	for strIndex1, letterCounts1 := range letterCounts {
		if knownAnagrams[strIndex1] {
			// this string is a known anagram
			continue
		}
		anagrams := []string{strs[strIndex1]}

		for strIndex2, letterCounts2 := range letterCounts {
			if strIndex1 >= strIndex2 || // skip previous indices
				len(letterCounts1) != len(letterCounts2) || // strings have different length
				knownAnagrams[strIndex2] { // this string is a known anagram
				continue
			}
			isAnagram := true
			for i := 97; i < 123; i++ { // a-z
				v1 := letterCounts1[byte(i)]
				v2 := letterCounts2[byte(i)]
				if v1 != v2 {
					isAnagram = false
					break
				}
			}
			if isAnagram {
				knownAnagrams[strIndex2] = true
				anagrams = append(anagrams, strs[strIndex2])
			}
		}
		result = append(result, anagrams)
	}
	return result
}

// @lc code=end

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	println(fmt.Sprintf("%+v", res))
}
