package _001_two_sum

/*
   by wangwei
   2018/12/5 16:43
*/

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for k, v := range nums {
		if j, ok := index[target-v]; ok {
			return []int{j, k}
		}
		index[v] = k
	}
	return nil
}
