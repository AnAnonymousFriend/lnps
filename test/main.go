package main

func main()  {

	var balance = []int{2,7,11,15}
	//twoSum(balance,18)
	twoSum1(balance,18)
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}



func twoSum(nums []int, target int) []int {




	var sum []int
	var lennum = len(nums) - 1
	for i := 0; i <= lennum; i++ {

		if i+1 >= lennum -1 {
			return nil
		}
		println(i+1)
		if nums[i] + nums[i+1] == target {
			sum[0] = i
			sum[1] = i+1
		}

	}
	return sum
}
