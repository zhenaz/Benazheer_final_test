package main

import (
	"fmt"
	"strings"
)

func main() {
	// SOAL 1
	// nums := []int{2, 3, 5, 7, 11, 13}
	// var target int
	// fmt.Println("Masukan angka target: ")
	// fmt.Scanln(&target)
	// fmt.Println(twoSum(nums, target))

	// SOAL 2
	// var input string
	// fmt.Println("Masukan kata: ")
	// fmt.Scanln(&input)
	// fmt.Println(palindrom(input))

	// SOAL 3
	// var kata string
	// var targetKata string
	// fmt.Println("Masukan kata: ")
	// fmt.Scanln(&kata)
	// fmt.Println("Masukan target: ")
	// fmt.Scanln(&target)
	// fmt.Println(hitungHuruf(kata, targetKata))

	// SOAL 5
	var inputPyramid int
	var startNum int
	fmt.Println("Masukan jumlah baris: ")
	fmt.Scanln(&inputPyramid)
	fmt.Println("Masukan angka awal: ")
	fmt.Scanln(&startNum)
	fmt.Println("Result: ")
	pyramid(inputPyramid, startNum)
}

// SOAL 1 (Kemungkinan twoSum)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return []int{}
}

// SOAL 2
func palindrom(s string) bool {
	kecil := strings.ToLower(s)
	var balik []string
	stringSplit := strings.Split(kecil, "")
	for i := len(stringSplit) - 1; i >= 0; i-- {
		balik = append(balik, stringSplit[i])
	}
	return strings.Join(balik, "") == kecil
}

// SOAL 3
func hitungHuruf(kata string, targetKata string) int {
	jumlah := 0
	i := 0

	for i <= len(kata)-len(targetKata) {
		if kata[i:i+len(targetKata)] == targetKata {
			jumlah++
			i += len(targetKata)
		} else {
			i++
		}
	}
	return jumlah
}

// SOAL 4

// SOAL 5
func pyramid(inputPyramid int, startNum int) {
	for i := 0; i < inputPyramid; i++ {
		num := startNum + i
		for j := 0; j <= i; j++ {
			fmt.Print(num+j, " ")
		}
		fmt.Println()
	}
}

