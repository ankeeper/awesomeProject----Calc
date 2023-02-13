package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RomanResult(aDigit int) (rString string) {

	if aDigit == 1 {
		rString = "I"
		return rString
	} else if aDigit == 2 {
		rString := "II"
		return rString
	} else if aDigit == 3 {
		rString := "III"
		return rString
	} else if aDigit == 4 {
		rString := "IV"
		return rString
	} else if aDigit == 5 {
		rString := "V"
		return rString
	} else if aDigit == 6 {
		rString := "VI"
		return rString
	} else if aDigit == 7 {
		rString := "VII"
		return rString
	} else if aDigit == 8 {
		rString := "VIII"
		return rString
	} else if aDigit == 9 {
		rString := "IX"
		return rString
	} else if aDigit == 10 {
		rString := "X"
		return rString
	} else if aDigit == 11 {
		rString := "XI"
		return rString
	} else if aDigit == 12 {
		rString := "XII"
		return rString
	} else if aDigit == 13 {
		rString := "XIII"
		return rString
	} else if aDigit == 14 {
		rString := "XIV"
		return rString
	} else if aDigit == 15 {
		rString := "XV"
		return rString
	} else if aDigit == 16 {
		rString := "XVI"
		return rString
	} else if aDigit == 17 {
		rString := "XVII"
		return rString
	} else if aDigit == 18 {
		rString := "XVIII"
		return rString
	} else if aDigit == 19 {
		rString := "XIX"
		return rString
	} else if aDigit == 20 {
		rString := "XX"
		return rString
	} else if aDigit == 21 {
		rString := "XXI"
		return rString
	} else if aDigit == 22 {
		rString := "XXII"
		return rString
	} else if aDigit == 23 {
		rString := "XXIII"
		return rString
	} else if aDigit == 24 {
		rString := "XXIV"
		return rString
	} else if aDigit == 25 {
		rString := "XXV"
		return rString
	} else if aDigit == 26 {
		rString := "XXVI"
		return rString
	} else if aDigit == 27 {
		rString := "XXVII"
		return rString
	} else if aDigit == 28 {
		rString := "XXVIII"
		return rString
	} else if aDigit == 29 {
		rString := "XXIX"
		return rString
	} else if aDigit == 30 {
		rString := "XXX"
		return rString
	} else if aDigit == 31 {
		rString := "XXXI"
		return rString
	} else if aDigit == 32 {
		rString := "XXXII"
		return rString
	} else if aDigit == 33 {
		rString := "XXXIII"
		return rString
	} else if aDigit == 34 {
		rString := "XXXIV"
		return rString
	} else if aDigit == 35 {
		rString := "XXXV"
		return rString
	} else if aDigit == 36 {
		rString := "XXXVI"
		return rString
	} else if aDigit == 37 {
		rString := "XXXVII"
		return rString
	} else if aDigit == 38 {
		rString := "XXXVIII"
		return rString
	} else if aDigit == 39 {
		rString := "XXXIX"
		return rString
	} else if aDigit == 40 {
		rString := "XL"
		return rString
	} else if aDigit == 41 {
		rString := "XLI"
		return rString
	} else if aDigit == 42 {
		rString := "XLII"
		return rString
	} else if aDigit == 43 {
		rString := "XLIII"
		return rString
	} else if aDigit == 44 {
		rString := "XLIV"
		return rString
	} else if aDigit == 45 {
		rString := "XLV"
		return rString
	} else if aDigit == 46 {
		rString := "XLVI"
		return rString
	} else if aDigit == 47 {
		rString := "XLVII"
		return rString
	} else if aDigit == 48 {
		rString := "XLVIII"
		return rString
	} else if aDigit == 49 {
		rString := "XLIX"
		return rString
	} else if aDigit == 50 {
		rString := "L"
		return rString
	} else if aDigit == 51 {
		rString := "LI"
		return rString
	} else if aDigit == 52 {
		rString := "LII"
		return rString
	} else if aDigit == 53 {
		rString := "LIII"
		return rString
	} else if aDigit == 54 {
		rString := "LIV"
		return rString
	} else if aDigit == 55 {
		rString := "LV"
		return rString
	} else if aDigit == 56 {
		rString := "LVI"
		return rString
	} else if aDigit == 57 {
		rString := "LVII"
		return rString
	} else if aDigit == 58 {
		rString := "LVIII"
		return rString
	} else if aDigit == 59 {
		rString := "LIX"
		return rString
	} else if aDigit == 60 {
		rString := "LX"
		return rString
	} else if aDigit == 61 {
		rString := "LXI"
		return rString
	} else if aDigit == 62 {
		rString := "LXII"
		return rString
	} else if aDigit == 63 {
		rString := "LXIII"
		return rString
	} else if aDigit == 64 {
		rString := "LXIV"
		return rString
	} else if aDigit == 65 {
		rString := "LXV"
		return rString
	} else if aDigit == 66 {
		rString := "LXVI"
		return rString
	} else if aDigit == 67 {
		rString := "LXVII"
		return rString
	} else if aDigit == 68 {
		rString := "LXVIII"
		return rString
	} else if aDigit == 69 {
		rString := "LXIX"
		return rString
	} else if aDigit == 70 {
		rString := "LXX"
		return rString
	} else if aDigit == 71 {
		rString := "LXXI"
		return rString
	} else if aDigit == 72 {
		rString := "LXXII"
		return rString
	} else if aDigit == 73 {
		rString := "LXXIII"
		return rString
	} else if aDigit == 74 {
		rString := "LXXIV"
		return rString
	} else if aDigit == 75 {
		rString := "LXXV"
		return rString
	} else if aDigit == 76 {
		rString := "LXXVI"
		return rString
	} else if aDigit == 77 {
		rString := "LXXVII"
		return rString
	} else if aDigit == 78 {
		rString := "LXXVIII"
		return rString
	} else if aDigit == 79 {
		rString := "LXXIX"
		return rString
	} else if aDigit == 80 {
		rString := "LXXX"
		return rString
	} else if aDigit == 81 {
		rString := "LXXXI"
		return rString
	} else if aDigit == 82 {
		rString := "LXXXII"
		return rString
	} else if aDigit == 83 {
		rString := "LXXXIII"
		return rString
	} else if aDigit == 84 {
		rString := "LXXXIV"
		return rString
	} else if aDigit == 85 {
		rString := "LXXXV"
		return rString
	} else if aDigit == 86 {
		rString := "LXXXVI"
		return rString
	} else if aDigit == 87 {
		rString := "LXXXVII"
		return rString
	} else if aDigit == 88 {
		rString := "LXXXVIII"
		return rString
	} else if aDigit == 89 {
		rString := "LXXXVIV"
		return rString
	} else if aDigit == 90 {
		rString := "XC"
		return rString
	} else if aDigit == 91 {
		rString := "XCI"
		return rString
	} else if aDigit == 92 {
		rString := "XCII"
		return rString
	} else if aDigit == 93 {
		rString := "XCIII"
		return rString
	} else if aDigit == 94 {
		rString := "XCIV"
		return rString
	} else if aDigit == 95 {
		rString := "XCV"
		return rString
	} else if aDigit == 96 {
		rString := "XCVI"
		return rString
	} else if aDigit == 97 {
		rString := "XCVII"
		return rString
	} else if aDigit == 98 {
		rString := "XCVIII"
		return rString
	} else if aDigit == 99 {
		rString := "XCIX"
		return rString
	} else if aDigit == 100 {
		rString := "C"
		return rString
	}

	return rString
}

func roman2arabic(line string) (rDigit int) {
	//fmt.Println("функция:", line)
	ClearLine := strings.TrimSpace(line)
	if ClearLine == "I" {
		rDigit := 1
		return rDigit
	} else if ClearLine == "II" {
		rDigit := 2
		return rDigit
	} else if ClearLine == "III" {
		rDigit := 3
		return rDigit
	} else if ClearLine == "IV" {
		rDigit := 4
		return rDigit
	} else if ClearLine == "V" {
		rDigit := 5
		return rDigit
	} else if ClearLine == "VI" {
		rDigit := 6
		return rDigit
	} else if ClearLine == "VII" {
		rDigit := 7
		return rDigit
	} else if ClearLine == "VIII" {
		rDigit := 8
		return rDigit
	} else if ClearLine == "IX" {
		rDigit := 9
		return rDigit
	} else if ClearLine == "X" {
		rDigit := 10
		return rDigit
	}
	return rDigit
}

func main() {

	//var testdigit0 int
	//	var testdigit1 int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	//	fmt.Println("captured:", line)

	//	words := strings.Split(line, "+")
	//	fmt.Printf("%q\n", words[0], words[1])

	if strings.Contains(line, "+") == true { /* поиск символа +,*/
		//	fmt.Println("found +")

		words := strings.Split(line, "+") /* разделение фразы по разделителю ,*/
		//	fmt.Printf("%q\n", words[0], words[1])

		ClearWords0 := strings.TrimSpace(words[0]) /* очистка от пробелов*/
		ClearWords1 := strings.TrimSpace(words[1])

		testdigit0, err1 := strconv.Atoi(ClearWords0) /* конвертация в число*/
		testdigit1, err2 := strconv.Atoi(ClearWords1)
		testsumm := testdigit0 + testdigit1

		if err1 == nil && err2 == nil {
			fmt.Println("результат: ", testsumm) /* вывод результата,*/
		} else {
			//fmt.Println("Roman?")

			rDigit0 := roman2arabic(words[0])
			rDigit1 := roman2arabic(words[1])
			testsummRoman := rDigit0 + rDigit1
			aDigit := RomanResult(testsummRoman)
			fmt.Println("сумма римская: ", aDigit)

		}

	} else if strings.Contains(line, "-") == true {

		//	fmt.Println("found -")

		words := strings.Split(line, "-") /* разделение фразы по разделителю ,*/
		//	fmt.Printf("%q\n", words[0], words[1])

		ClearWords0 := strings.TrimSpace(words[0]) /* очистка от пробелов*/
		ClearWords1 := strings.TrimSpace(words[1])

		testdigit0, err1 := strconv.Atoi(ClearWords0) /* конвертация в число*/
		testdigit1, err2 := strconv.Atoi(ClearWords1)
		testsumm := testdigit0 - testdigit1

		if err1 == nil && err2 == nil {
			fmt.Println("результат: ", testsumm) /* вывод результата,*/
		} else {
			//fmt.Println("Roman?")

			rDigit0 := roman2arabic(words[0])
			rDigit1 := roman2arabic(words[1])
			testsummRoman := rDigit0 - rDigit1

			if testsummRoman < 1 {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			} else {
				aDigit := RomanResult(testsummRoman)
				fmt.Println("результат: ", aDigit)
			}
		}

	} else if strings.Contains(line, "/") == true {
		// fmt.Println("found /")

		words := strings.Split(line, "/") /* разделение фразы по разделителю ,*/
		//	fmt.Printf("%q\n", words[0], words[1])

		ClearWords0 := strings.TrimSpace(words[0]) /* очистка от пробелов*/
		ClearWords1 := strings.TrimSpace(words[1])

		testdigit0, err1 := strconv.Atoi(ClearWords0) /* конвертация в число*/
		testdigit1, err2 := strconv.Atoi(ClearWords1)

		if err1 == nil && err2 == nil {
			testsumm := testdigit0 / testdigit1
			fmt.Println("результат: ", testsumm) /* вывод результата,*/
		} else {
			//fmt.Println("Roman?")

			rDigit0 := roman2arabic(words[0])
			rDigit1 := roman2arabic(words[1])
			testsummRoman := rDigit0 / rDigit1
			aDigit := RomanResult(testsummRoman)
			fmt.Println("результат: ", aDigit)

		}

	} else if strings.Contains(line, "*") == true {
		//fmt.Println("found *")
		words := strings.Split(line, "*") /* разделение фразы по разделителю ,*/
		//	fmt.Printf("%q\n", words[0], words[1])

		ClearWords0 := strings.TrimSpace(words[0]) /* очистка от пробелов*/
		ClearWords1 := strings.TrimSpace(words[1])

		testdigit0, err1 := strconv.Atoi(ClearWords0) /* конвертация в число*/
		testdigit1, err2 := strconv.Atoi(ClearWords1)

		if err1 == nil && err2 == nil {
			testsumm := testdigit0 * testdigit1
			fmt.Println("результат: ", testsumm) /* вывод результата,*/
		} else {
			//fmt.Println("Roman?")

			rDigit0 := roman2arabic(words[0])
			rDigit1 := roman2arabic(words[1])
			testsummRoman := rDigit0 * rDigit1
			aDigit := RomanResult(testsummRoman)
			fmt.Println("результат: ", aDigit)

		}

	} else {

		fmt.Println("no luck")

	}

}
