package main

func main() {
	func(x []int) {
		sum := 0
		for _, v := range x {
			sum += v
		}
		print(sum)
	}([]int{1, 2, 3, 4})
}
