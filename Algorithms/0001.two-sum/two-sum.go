package problem0001

func twoSum(nums []int, target int) []int {
	// index 负责保存map[整数]整数的序列号
	index := make(map[int]int, len(nums))

	// 通过 for 循环，获取b的序列号
	for i, b := range nums {
		// 通过查询map，获取a = target - b的序列号
		if j, ok := index[target-b]; ok {
			// ok 为 true
			// 说明在i之前，存在 nums[j] == a
			return []int{j, i}
			// 注意，顺序是j，i
		}

		// 把i和i的值，存入map
		index[b] = i
	}

	return nil
}

// 注意slice和map的搭配使用，可以解决很多问题。(数组去重中也使用了slice 和map 的搭配问题)
