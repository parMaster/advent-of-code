package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Given a string s which consists of lowercase or uppercase letters, return
// the length of the longest palindrome that can be built with those letters.
// https://leetcode.com/problems/longest-palindrome
func longestPalindrome(s string) int {
	letters := map[rune]int{}
	for _, si := range s {
		if _, ok := letters[si]; ok {
			letters[si]++
		} else {
			letters[si] = 1
		}
	}
	sumEven := 0
	sumOdds := 0
	maxOdd := 0
	for _, lcnt := range letters {
		// all Evens building palindrome
		if lcnt%2 == 0 {
			sumEven += lcnt
		}

		// all Odds without one letter (ccc -> cc) are Evens
		// maxOdd will go in the middle whole
		if lcnt%2 != 0 {
			sumOdds += lcnt - 1
			maxOdd = max(lcnt, maxOdd)
		}
	}
	// exclude maxOdd from the sum
	if maxOdd != 0 {
		sumOdds -= (maxOdd - 1)
	}

	return sumEven + maxOdd + sumOdds
}

func Test_LongetsPalindrome(t *testing.T) {
	require.Equal(t, 7, longestPalindrome("abccccdd"))
	require.Equal(t, 3, longestPalindrome("ccc"))
	require.Equal(t, 1, longestPalindrome("a"))
	require.Equal(t, 2, longestPalindrome("bb"))
	require.Equal(t, 9, longestPalindrome("ababababa"))
	require.Equal(t, 983, longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}

// split N array into K non-empty subarrays such that the largest sum of any subarray is minimized
// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/description/
// https://leetcode.com/problems/split-array-largest-sum is the same problem
func shipWithinDays(weights []int, days int) int {

	minCap := slices.Max(weights)
	maxCap := 0
	for _, weight := range weights {
		maxCap += weight
	}
	// log.Println("Cap is within:", minCap, "-", maxCap)

	// binary search fused against inf. cycle
	for i := 0; minCap != maxCap && i < 3000; i++ {
		// new capacity value to check
		mid := (maxCap-minCap)/2 + minCap

		shipDays := 0
		runningSum := 0
		for _, w := range weights {
			runningSum += w
			if runningSum > mid {
				runningSum = w
				shipDays++
				if shipDays > days {
					break
				}
			}
		}
		shipDays++

		// log.Println("Checking min-max: ", minCap, maxCap, "mid:", mid, "ships in", shipDays)

		if shipDays <= days {
			maxCap = mid
		} else {
			minCap = mid + 1
		}
	}

	return minCap
}

func Test_ShipWithinDays(t *testing.T) {
	// Given an integer array nums and an integer k, split nums into k non-empty subarrays such that the largest sum of any subarray is minimized.
	// Return the minimized largest sum of the split.
	// A subarray is a contiguous part of the array.

	require.Equal(t, 15, shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	require.Equal(t, 6, shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
	require.Equal(t, 3, shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
	require.Equal(t, 251000, shipWithinDays([]int{500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500}, 1))
}

// 872. Leaf-Similar Trees
// https://leetcode.com/problems/leaf-similar-trees
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return slices.Equal(root1.lvs(), root2.lvs())
}

// leaf value sequence returns array of leaf values (nodes with no descendants) from left to right
func (node *TreeNode) lvs() []int {

	res := []int{}
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}
	if node.Left != nil {
		res = append(res, node.Left.lvs()...)
	}
	if node.Right != nil {
		res = append(res, node.Right.lvs()...)
	}

	return res
}

func Test_LVS(t *testing.T) {
	root1 := &TreeNode{
		3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}}, &TreeNode{1, &TreeNode{9, nil, nil}, &TreeNode{8, nil, nil}},
	}

	root2 := &TreeNode{
		3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}, &TreeNode{1, &TreeNode{4, nil, nil}, &TreeNode{2, &TreeNode{9, nil, nil}, &TreeNode{8, nil, nil}}},
	}

	require.Equal(t, []int{6, 7, 4, 9, 8}, root1.lvs())
	require.Equal(t, []int{6, 7, 4, 9, 8}, root2.lvs())

	require.True(t, leafSimilar(root1, root2))
}

// returns binary tree constructed from string like "3,5,1,6,7,4,2,null,null,null,null,null,null,9,8"
// must be able to read incomplete trees like "1,null,2,2,null" and "1,null,2,2"
func ReadBST(s string) *TreeNode {
	arr := []*int{}
	for _, v := range strings.Split(s, ",") {
		if v == "null" {
			arr = append(arr, nil)
			continue
		}
		d, _ := strconv.Atoi(v)
		arr = append(arr, &d)
	}
	if len(arr) == 0 {
		return nil
	}

	var insertLevelOrder func(arr []*int, i int) (root *TreeNode)
	insertLevelOrder = func(arr []*int, i int) (root *TreeNode) {
		if i < len(arr) && arr[i] != nil {
			root = &TreeNode{*arr[i], nil, nil}
			root.Left = insertLevelOrder(arr, 2*i+1)
			root.Right = insertLevelOrder(arr, 2*i+2)
		}
		return root
	}

	return insertLevelOrder(arr, 0)
}

func Test_ReadBST(t *testing.T) {
	root1 := &TreeNode{
		3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}}, &TreeNode{1, &TreeNode{9, nil, nil}, &TreeNode{8, nil, nil}},
	}
	require.Equal(t, root1, ReadBST("3,5,1,6,2,9,8,null,null,7,4"))

	root2 := &TreeNode{
		3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}, &TreeNode{1, &TreeNode{4, nil, nil}, &TreeNode{2, &TreeNode{9, nil, nil}, &TreeNode{8, nil, nil}}},
	}
	require.Equal(t, root2, ReadBST("3,5,1,6,7,4,2,null,null,null,null,null,null,9,8"))

	root3 := &TreeNode{
		2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil},
	}
	require.Equal(t, root3, ReadBST("2,1,3"))

}

// 98. Validate Binary Search Tree
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
func (root *TreeNode) isValidBST() bool {
	var validSubtree func(root *TreeNode, minVal, maxVal int) bool
	validSubtree = func(root *TreeNode, minVal, maxVal int) bool {

		if root.Val >= maxVal || root.Val <= minVal {
			return false
		}

		var res bool = true
		if root.Left != nil {
			if root.Left.Val >= root.Val {
				return false
			}
			res = res && validSubtree(root.Left, minVal, min(maxVal, root.Val))
		}

		if root.Right != nil {
			if root.Right.Val <= root.Val {
				return false
			}
			res = res && validSubtree(root.Right, max(minVal, root.Val), maxVal)
		}

		return res
	}

	return validSubtree(root, math.MinInt, math.MaxInt)
}

func Test_validBST(t *testing.T) {
	require.True(t, ReadBST("2,1,3").isValidBST())
	require.False(t, ReadBST("2,2,2").isValidBST())
	require.False(t, ReadBST("5,1,4,null,null,3,6").isValidBST())
	require.False(t, ReadBST("5,4,6,null,null,3,7").isValidBST())
}

// 501. Find Mode in Binary Search Tree
func (root *TreeNode) findMode() []int {

	memo := map[int]int{}

	var traverse func(root *TreeNode, memo map[int]int)
	traverse = func(root *TreeNode, memo map[int]int) {
		if _, ok := memo[root.Val]; ok {
			memo[root.Val]++
		} else {
			memo[root.Val] = 1
		}

		if root.Left != nil {
			traverse(root.Left, memo)
		}

		if root.Right != nil {
			traverse(root.Right, memo)
		}
	}

	traverse(root, memo)
	maxFreq := 0
	for _, freq := range memo {
		maxFreq = max(maxFreq, freq)
	}
	res := []int{}
	for key, freq := range memo {
		if freq == maxFreq {
			res = append(res, key)
		}
	}

	return res
}

func Test_findMode(t *testing.T) {
	require.Equal(t, []int{2}, ReadBST("1,0,2,2,null").findMode())
	require.Equal(t, []int{3}, ReadBST("1,null,2,2,3,3,3").findMode())
}

// 1026. Maximum Difference Between Node and Ancestor
// https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
func (root *TreeNode) maxAncestorDiff() int {
	var minmax func(root *TreeNode, minVal, maxVal int) int
	minmax = func(root *TreeNode, minVal, maxVal int) int {
		minVal = min(minVal, root.Val)
		maxVal = max(maxVal, root.Val)

		var res int = maxVal - minVal
		if root.Left != nil {
			res = max(res, minmax(root.Left, minVal, maxVal))
		}
		if root.Right != nil {
			res = max(res, minmax(root.Right, minVal, maxVal))
		}

		return res
	}

	return minmax(root, root.Val, root.Val)
}

func Test_maxAncestorDiff(t *testing.T) {
	require.Equal(t, 7, ReadBST("8,3,10,1,6,null,14,null,null,4,7,13").maxAncestorDiff())
	require.Equal(t, 3, ReadBST("1,0,2,null,null,null,0,null,null,null,null,null,null,3,null").maxAncestorDiff())
}

func (root *TreeNode) Traverse(f func(root *TreeNode)) {

	f(root)

	if root.Left != nil {
		root.Left.Traverse(f)
	}

	if root.Right != nil {
		root.Right.Traverse(f)
	}
}

// findMode with Traverse
func (root *TreeNode) findMode2() []int {
	memo := map[int]int{}

	root.Traverse(func(root *TreeNode) {
		if _, ok := memo[root.Val]; ok {
			memo[root.Val]++
		} else {
			memo[root.Val] = 1
		}
	})

	maxFreq := 0
	for _, freq := range memo {
		maxFreq = max(maxFreq, freq)
	}
	res := []int{}
	for key, freq := range memo {
		if freq == maxFreq {
			res = append(res, key)
		}
	}

	return res
}

func Test_findMode2(t *testing.T) {
	require.Equal(t, []int{2}, ReadBST("1,0,2,2,null").findMode2())
	require.Equal(t, []int{3}, ReadBST("1,null,2,2,3,3,3").findMode2())
}

func Traverse(root *TreeNode, f func(root *TreeNode)) {
	if root == nil {
		return
	}

	f(root)

	Traverse(root.Left, f)
	Traverse(root.Right, f)
}

// 1038. Binary Search Tree to Greater Sum Tree
// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
func bstToGst(root *TreeNode) *TreeNode {
	memo := map[int]int{}

	Traverse(root, func(root *TreeNode) {
		if _, ok := memo[root.Val]; ok {
			memo[root.Val]++
		} else {
			memo[root.Val] = 1
		}
	})

	Traverse(root, func(root *TreeNode) {
		rootVal := root.Val
		for key := range memo {
			if key > rootVal {
				root.Val += key * memo[key]
			}
		}
	})

	return root
}

func Test_bstToGst(t *testing.T) {
	require.Equal(t, ReadBST("30,36,21,36,35,26,15,null,null,null,33,null,null,null,8"), bstToGst(ReadBST("4,1,6,0,2,5,7,null,null,null,3,null,null,null,8")))
}

// 670. Maximum Swap
func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	maxed := []int{}
	for _, v := range s {
		maxed = append(maxed, int(v)-0x30)
	}
	in := slices.Clone(maxed)
	slices.SortFunc(maxed, func(i, j int) int {
		if i > j {
			return -1
		}
		return 1
	})

	fmt.Println("in:", in, "maxed:", maxed)

	start := 0 // start of the suboptimal part if array
	for i := 0; i < len(in); i++ {
		if in[i] != maxed[i] {
			start = i
			break
		}
	}

	// find max digit in 'in' from right to left
	maxDigit := 0
	maxDigitIdx := -1
	for j := len(in) - 1; j > start; j-- {
		if in[j] > maxDigit {
			maxDigit = in[j]
			maxDigitIdx = j
		}
	}
	fmt.Println("maxDigit:", maxDigit, "maxDigitIdx:", maxDigitIdx)

	for i := start; i < maxDigitIdx; i++ {
		if in[i] < maxDigit {
			in[i], in[maxDigitIdx] = maxDigit, in[i]
			break
		}
	}

	// output
	res := 0
	for _, v := range in {
		res = res*10 + v
	}

	return res
}

func Test_maximumSwap(t *testing.T) {
	require.Equal(t, 999, maximumSwap(999))
	require.Equal(t, 7236, maximumSwap(2736))
	require.Equal(t, 9913, maximumSwap(1993))
	require.Equal(t, 9973, maximumSwap(9973))
	require.Equal(t, 98863, maximumSwap(98368))
	require.Equal(t, 210, maximumSwap(120))
	require.Equal(t, 52341342, maximumSwap(22341345))
	require.Equal(t, 99910, maximumSwap(99901))
}

// 1704. Determine if String Halves Are Alike
func halvesAreAlike(s string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	cnt := 0
	for i, v := range s {
		if slices.Contains(vowels, v) {
			if i < len(s)/2 {
				cnt++
			} else {
				cnt--
			}
		}
	}

	return cnt == 0
}

func Test_halvesAreAlike(t *testing.T) {
	require.True(t, halvesAreAlike("book"))
	require.False(t, halvesAreAlike("textbook"))
	require.False(t, halvesAreAlike("MerryChristmas"))
	require.True(t, halvesAreAlike("AbCdEfGh"))
}

// 2115. Find All the possible recipes for the given supplies
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {

	memo := map[int]bool{}

	var cook func(recipe int, cooking []int) bool
	cook = func(recipe int, cooking []int) bool {

		if res, ok := memo[recipe]; ok {
			return res
		}

		if slices.Contains(cooking, recipe) {
			// already cooking, cycle detected
			memo[recipe] = false
			return false
		}

		cookable := true
		for _, product := range ingredients[recipe] {
			if slices.Contains(supplies, product) {
				// basic supply
				continue
			}
			// check if we have recipe for this product
			if r := slices.Index(recipes, product); r != -1 {
				cookable = cookable && cook(r, append(cooking, recipe))
			} else {
				// no recipe for this product
				cookable = false
				break
			}
		}

		memo[recipe] = cookable
		return cookable
	}

	res := []string{}
	for ir, recipe := range recipes {
		if cook(ir, []int{}) {
			res = append(res, recipe)
		}
	}

	return res
}

func TestFindRecipes(t *testing.T) {
	// Example 1:
	// Input: recipes = ["bread"], ingredients = [["yeast","flour"]], supplies = ["yeast","flour","corn"]
	// Output: ["bread"]
	// Explanation:
	// We can create "bread" since we have the ingredients "yeast" and "flour".
	require.Equal(t, []string{"bread"}, findAllRecipes([]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"}))

	// 	Example 2:
	// Input: recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]], supplies = ["yeast","flour","meat"]
	// Output: ["bread","sandwich"]
	// Explanation:
	// We can create "bread" since we have the ingredients "yeast" and "flour".
	// We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
	require.Equal(t, []string{"bread", "sandwich"}, findAllRecipes([]string{"bread", "sandwich"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}}, []string{"yeast", "flour", "meat"}))
	require.Equal(t, []string{"bread"}, findAllRecipes([]string{"bread", "sandwich"}, [][]string{{"yeast", "flour"}, {"bread", "meat", "marzipan"}}, []string{"yeast", "flour", "meat"}))

	// 	Example 3:
	// Input: recipes = ["bread","sandwich","burger"], ingredients = [["yeast","flour"],["bread","meat"],["sandwich","meat","bread"]], supplies = ["yeast","flour","meat"]
	// Output: ["bread","sandwich","burger"]
	// Explanation:
	// We can create "bread" since we have the ingredients "yeast" and "flour".
	// We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
	// We can create "burger" since we have the ingredient "meat" and can create the ingredients "bread" and "sandwich".
	require.Equal(t, []string{"bread", "sandwich", "burger"}, findAllRecipes([]string{"bread", "sandwich", "burger"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}, []string{"yeast", "flour", "meat"}))

	// Example 4:
	require.Equal(t, []string{"ju", "fzjnm", "q"}, findAllRecipes([]string{"ju", "fzjnm", "x", "e", "zpmcz", "h", "q"}, [][]string{{"d"}, {"hveml", "f", "cpivl"}, {"cpivl", "zpmcz", "h", "e", "fzjnm", "ju"}, {"cpivl", "hveml", "zpmcz", "ju", "h"}, {"h", "fzjnm", "e", "q", "x"}, {"d", "hveml", "cpivl", "q", "zpmcz", "ju", "e", "x"}, {"f", "hveml", "cpivl"}}, []string{"f", "hveml", "cpivl", "d"}))

}
