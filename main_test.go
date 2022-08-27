package main

import "testing"

type TestCase struct {
	InputData int // то что будет подаваться на вход
	Answer    int // то что вернет тестируемая функция (в нашем случае)
	Expected  int // то что ожидаем получить

}

var cases []TestCase = []TestCase{
	{
		InputData: 0,
		Expected:  1,
	},
	{
		InputData: 1,
		Expected:  1,
	},
	{
		InputData: 2,
		Expected:  2,
	},
	{
		InputData: 3,
		Expected:  6,
	},
}

func TestFactorial(t *testing.T) {
	for id, test := range cases {
		if test.Answer = factorial(test.InputData); test.Answer != test.Expected {
			t.Errorf("Test error in id %d give value %d end got answer %v expected %d", id, test.InputData, test.Answer, test.Expected)
		}
	}
}
