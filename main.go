package main

func main() {
	// return "12323"
}

func factorial(number int) int {
	if number <= 1 {
		return 1
	}

	answerNumber := 1

	for i := 1; i <= number; i++ {
		answerNumber *= i
	}

	return answerNumber

}
