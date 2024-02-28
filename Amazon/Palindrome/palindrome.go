package main

func IsPalindrome(num int) bool {
	if num < 0 {
		return false
	}

	if num%10 == 0 && num != 0 {
		return false
	}

	if num == Reverse(num) {
		return true
	}

	return false
}

func Reverse(num int) int {
	var reverse int
	for num > 0 {
		reverse = reverse * 10 + num % 10
		num = num / 10
	}
	return reverse
}