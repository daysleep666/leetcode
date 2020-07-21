package main

// 滑动窗口
// 双指针
func findMaxAverage(nums []int, k int) float64 {
	start, end := 0, 0
	sum, avg := 0, float64(0)
	for end <= len(nums) {
		value := end - start
		if value < k {
			sum += nums[end]
			end++
		} else if value == k {
			tmpAvg := float64(sum) / float64(k)
			// TODO: float比较
			if avg == 0 || avg < tmpAvg {
				avg = tmpAvg
			}
			sum -= nums[start]
			start++
			if end == len(nums) {
				break
			}
		}
	}
	return avg
}

// 给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。
