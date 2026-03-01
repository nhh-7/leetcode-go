package skills

func findDuplicate1(nums []int) int {
	fast, slow := 0, 0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if fast == slow {
			head := 0
			for slow != head {
				head = nums[head]
				slow = nums[slow]
			}
			return slow
		}
	}
}

// 判断有环，寻找环入口即重复的数
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			ans := nums[0]
			for ans != slow {
				ans = nums[ans]
				slow = nums[slow]
			}
			return ans
		}
	}
}
