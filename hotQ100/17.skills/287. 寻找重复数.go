package skills

func findDuplicate(nums []int) int {
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
