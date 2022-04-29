package leetcode

func TwoSum(nums []int, target int) []int {
	l := len(nums)
	resp := []int{}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				resp = append(resp, i, j)
			}
		}
	}
	return resp
}
