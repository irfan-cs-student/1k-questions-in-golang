// ============================================================
// GO ARRAYS — 30 LOGIC, OUTPUT & SYNTAX CHALLENGES
// ============================================================
//
// Mission:
// 1. Predict the output
// 2. Find syntax/compile errors
// 3. Find runtime errors
// 4. Find logic errors
// 5. Understand deep Go array concepts
//
// IMPORTANT:
// Try to solve each question WITHOUT running the code first.
// ============================================================

// ============================================================
// QUESTION 1 — Basic Array Output
// ============================================================

package main

import "fmt"

func main() {

	// Q1
	nums := [5]int{10, 20, 30, 40, 50}

	fmt.Println(nums[2])
}

// ============================================================
// QUESTION 2 — Default Values
// ============================================================

func question2() {

	var nums [5]int

	nums[2] = 10

	fmt.Println(nums)
}

// ============================================================
// QUESTION 3 — Array Length
// ============================================================

func question3() {

	nums := [5]int{10, 20, 30}

	fmt.Println(len(nums))
}

// ============================================================
// QUESTION 4 — Index Out of Range
// ============================================================

func question4() {

	nums := [5]int{10, 20, 30, 40, 50}

	fmt.Println(nums[5])
}

// ============================================================
// QUESTION 5 — Reverse Loop
// ============================================================

func question5() {

	nums := [5]int{10, 20, 30, 40, 50}

	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Print(nums[i], " ")
	}
}

// ============================================================
// QUESTION 6 — Modify Through Loop
// ============================================================

func question6() {

	nums := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * 2
	}

	fmt.Println(nums)
}

// ============================================================
// QUESTION 7 — Range Copy Concept
// ============================================================

func question7() {

	nums := [4]int{10, 20, 30, 40}

	for _, value := range nums {
		value = value * 2
	}

	fmt.Println(nums)
}

// ============================================================
// QUESTION 8 — Range With Index
// ============================================================

func question8() {

	nums := [4]int{10, 20, 30, 40}

	for i := range nums {
		nums[i] += 5
	}

	fmt.Println(nums)
}

// ============================================================
// QUESTION 9 — Array Comparison
// ============================================================

func question9() {

	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}

	fmt.Println(a == b)
}

// ============================================================
// QUESTION 10 — Different Array Sizes
// ============================================================

func question10() {

	a := [3]int{1, 2, 3}
	b := [4]int{1, 2, 3, 4}

	fmt.Println(a == b)
}

// ============================================================
// QUESTION 11 — Array Assignment
// ============================================================

func question11() {

	a := [3]int{10, 20, 30}

	b := a

	b[0] = 100

	fmt.Println(a)
	fmt.Println(b)
}

// ============================================================
// QUESTION 12 — Function Receives Array
// ============================================================

func change(nums [3]int) {

	nums[0] = 100
}

func question12() {

	a := [3]int{10, 20, 30}

	change(a)

	fmt.Println(a)
}

// ============================================================
// QUESTION 13 — Function Returning Array
// ============================================================

func create() [3]int {

	return [3]int{5, 10, 15}
}

func question13() {

	nums := create()

	fmt.Println(nums)
}

// ============================================================
// QUESTION 14 — Return Array + Value
// ============================================================

func getData() ([3]int, int) {

	nums := [3]int{10, 20, 30}

	return nums, nums[1]
}

func question14() {

	a, b := getData()

	fmt.Println(a)
	fmt.Println(b)
}

// ============================================================
// QUESTION 15 — Multiple Return Trap
// ============================================================

func getData2() ([3]int, int) {

	return [3]int{10, 20, 30}, 50
}

func question15() {

	a := getData2()

	fmt.Println(a)
}

// ============================================================
// QUESTION 16 — Count Specific Number
// ============================================================

func question16() {

	nums := [7]int{2, 5, 2, 8, 2, 9, 5}

	count := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == 2 {
			count++
		}
	}

	fmt.Println(count)
}

// ============================================================
// QUESTION 17 — Find Duplicate
// ============================================================

func question17() {

	nums := [6]int{10, 20, 10, 30, 20, 40}

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i] == nums[j] {
				fmt.Println(nums[i])
			}
		}
	}
}

// ============================================================
// QUESTION 18 — Duplicate + Break
// ============================================================

func question18() {

	nums := [5]int{1, 2, 2, 2, 3}

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i] == nums[j] {

				fmt.Println("Duplicate:", nums[i])

				break
			}
		}
	}
}

// ============================================================
// QUESTION 19 — Find Largest
// ============================================================

func question19() {

	nums := [5]int{12, 45, 7, 89, 23}

	largest := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] > largest {
			largest = nums[i]
		}
	}

	fmt.Println(largest)
}

// ============================================================
// QUESTION 20 — Largest Initialization Trap
// ============================================================

func question20() {

	nums := [5]int{-10, -20, -5, -30, -15}

	largest := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] > largest {
			largest = nums[i]
		}
	}

	fmt.Println(largest)
}

// ============================================================
// QUESTION 21 — Second Largest
// ============================================================

func question21() {

	nums := [5]int{10, 50, 30, 40, 20}

	largest := nums[0]
	second := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] > largest {

			second = largest
			largest = nums[i]
		}
	}

	fmt.Println(largest)
	fmt.Println(second)
}

// ============================================================
// QUESTION 22 — Swap Array Elements
// ============================================================

func question22() {

	nums := [5]int{10, 20, 30, 40, 50}

	nums[0], nums[4] = nums[4], nums[0]

	fmt.Println(nums)
}

// ============================================================
// QUESTION 23 — Reverse In Place
// ============================================================

func question23() {

	nums := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums)/2; i++ {

		nums[i], nums[len(nums)-1-i] =
			nums[len(nums)-1-i], nums[i]
	}

	fmt.Println(nums)
}

// ============================================================
// QUESTION 24 — Even Numbers
// ============================================================

func question24() {

	nums := [6]int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(nums); i++ {

		if nums[i]%2 == 0 {
			fmt.Print(nums[i], " ")
		}
	}
}

// ============================================================
// QUESTION 25 — Wrong Index
// ============================================================

func question25() {

	nums := [5]int{10, 20, 30, 40, 50}

	sum := 0

	for i := 1; i <= len(nums); i++ {

		sum += nums[i]
	}

	fmt.Println(sum)
}

// ============================================================
// QUESTION 26 — Nested Loop Logic
// ============================================================

func question26() {

	nums := [4]int{1, 2, 3, 4}

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums); j++ {

			if nums[i] < nums[j] {
				fmt.Println(nums[i], nums[j])
			}
		}
	}
}

// ============================================================
// QUESTION 27 — Array Passed To Function
// ============================================================

func count(nums [5]int, target int) int {

	count := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == target {
			count++
		}
	}

	return count
}

func question27() {

	nums := [5]int{2, 4, 2, 6, 2}

	fmt.Println(count(nums, 2))
}

// ============================================================
// QUESTION 28 — Array Type Mismatch
// ============================================================

func count2(nums [5]int) int {

	return len(nums)
}

func question28() {

	nums := [7]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(count2(nums))
}

// ============================================================
// QUESTION 29 — len() Instead Of Fixed Number
// ============================================================

func question29() {

	nums := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(nums); i++ {

		fmt.Println(nums[i])
	}
}

// ============================================================
// QUESTION 30 — 🔥 DEEP ARRAY CHALLENGE
// ============================================================

func modify(nums [5]int) {

	for i := 0; i < len(nums); i++ {

		nums[i] += 10
	}
}

func question30() {

	nums := [5]int{1, 2, 3, 4, 5}

	modify(nums)

	for i := range nums {

		fmt.Print(nums[i], " ")
	}
}

// ============================================================
// MISSION
// ============================================================
//
// For every question determine:
//
// 1. What is the output?
// 2. Is there a syntax error?
// 3. Is there a compile-time error?
// 4. Is there a runtime error?
// 5. Is there a logic error?
// 6. WHY?
//
// Deep questions:
//
// Q7  -> Why doesn't range modify the original array?
// Q10 -> Why can't different array sizes be compared?
// Q11 -> Does array assignment copy the array?
// Q12 -> Does passing an array to a function copy it?
// Q20 -> Why is largest := 0 dangerous?
// Q21 -> Is the second-largest logic actually correct?
// Q23 -> Why len(nums)/2?
// Q25 -> What happens at nums[len(nums)]?
// Q28 -> Why can't [7]int be passed to [5]int?
// Q30 -> How can we modify the original array?
//
// ============================================================
