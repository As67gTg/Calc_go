package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := strings.Split(input.Text(), " ")
	if len(str) > 3 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор")

	} else if len(str) < 3 {
		panic("Проверьте корректность вводимых данных.")

	}
	num1 := str[0]
	num2 := str[1]
	num3 := str[2]

	if rometest(num1, num3) == false {
		op1, err := strconv.Atoi(num1)
		if err != nil {
			panic("Ошибка: введены некорректные данные")

		}
		op2, err2 := strconv.Atoi(num3)
		if err2 != nil {
			panic("Ошибка: введены некорректные данные")

		}
		if arabTest(op1, op2) == true {
			fmt.Println(arabOper(op1, num2, op2))
		} else {
			panic("На входе числа от 1 до 10")
		}
	} else {
		resRome := romeOper(num1, num2, num3)

		if resRome < 1 {
			panic("В римской системе только положительные числа")
		} else {
			dict_decades := map[int]string{1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC"}
			dict_units := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX"}
			unit := resRome % 10
			decade := resRome / 10
			if decade == 10 {
				fmt.Println("C")
			} else if decade < 1 {
				fmt.Println(dict_units[unit])
			} else {
				fmt.Print(dict_decades[decade], dict_units[unit])
			}

		}

	}

}

func rometest(k1 string, k2 string) bool {
	test1 := false
	test2 := false
	test := false
	romeNums := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, i := range romeNums {
		if k1 == i {
			test1 = true
		}
		if k2 == i {
			test2 = true
		}
		if (test1 == true) && (test2 == true) {
			test = true
		}
	}
	return test

}
func romeOper(k1 string, k2 string, k3 string) int {

	dict := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	ch1 := dict[k1]
	ch2 := dict[k3]
	if k2 == "+" {
		res := ch1 + ch2
		return res
	} else if k2 == "-" {
		res := ch1 - ch2
		return res
	} else if k2 == "*" {
		res := ch1 * ch2
		return res
	} else if k2 == "/" {
		res := ch1 / ch2
		return res
	} else {
		panic("Проблема с операндом_R")
	}

}
func arabTest(op1 int, op2 int) bool {
	test := true
	if (op1 > 10) || (op1 < 1) || (op2 > 10) || (op2 < 1) {
		test = false
	}
	return test
}
func arabOper(op1 int, num2 string, op2 int) int {
	if num2 == "+" {
		res := op1 + op2
		return res
	} else if num2 == "-" {
		res := op1 - op2
		return res
	} else if num2 == "*" {
		res := op1 * op2
		return res
	} else if num2 == "/" {
		res := op1 / op2
		return res
	} else {
		panic("Проблема с операндом")

	}

}
