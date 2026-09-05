/*
===========================================================
GO SLICES — 50 DAY PREDICT THE OUTPUT CHALLENGE
===========================================================

GOAL
----
Improve:
- Go syntax understanding
- Slice syntax
- len() and cap()
- append()
- copy()
- make()
- nil slices
- array vs slice
- backing arrays
- slice aliasing
- mutation
- functions and slices
- range
- variadic functions
- full slice expressions
- dynamic length
- subtle slice behavior
- DSA-style reasoning

===========================================================
RULES
===========================================================

1. DO NOT run the code before predicting the output.

2. For each day:
   - Read the code.
   - Write your predicted output.
   - Explain WHY you think that is the output.
   - Run the code.
   - Compare your prediction with the actual output.
   - If wrong, write what concept you misunderstood.

3. Only uncomment ONE challenge at a time.

4. Do NOT modify the challenge before making your prediction.

5. You MAY add comments explaining your thinking.

6. After solving a challenge, commit it to GitHub.

Example:

    git add main.go
    git commit -m "day 01: slice basics"
    git push

7. Your commit history should look like:

    Day 01
    Day 02
    Day 03
    ...
    Day 50

===========================================================
IMPORTANT
===========================================================

Some challenges intentionally contain confusing behavior.

Pay special attention to:

    []int{}
    var s []int
    make([]int, len)
    make([]int, len, cap)

    len()
    cap()

    append()

    copy()

    s[a:b]
    s[a:b:c]

    array[:]
    slice[:]

    range

    ...  // variadic / slice expansion

The difficult questions are designed to test whether you
actually understand what is happening in memory.

===========================================================
*/

// package main

// import "fmt"

// ==========================================================
// DAY 01 — BASIC SLICE
// ==========================================================

// func main() {
// 	s := []int{10, 20, 30}
//
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }
//answer:[10,20,30]
3 length 
3 capacity 

// ==========================================================
// DAY 02 — NIL SLICE
// ==========================================================

// func main() {
// 	var s []int
//
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// 	fmt.Println(s == nil)
// }
//answer:[]
0
0
true 
// ==========================================================
// DAY 03 — MAKE
// ==========================================================

// func main() {
// 	s := make([]int, 3)
//
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }
//answer: [0,0,0]
3
3
// ==========================================================
// DAY 04 — MAKE WITH CAPACITY
// ==========================================================

// func main() {
// 	s := make([]int, 3, 8)
//
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }
//answer:[0,0,0]
3
8
// ==========================================================
// DAY 05 — APPEND
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3}
//
// 	s = append(s, 4)
// 	s = append(s, 5, 6)
//
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// }
[1,2,3,4,4,5,6]
leght =7
cap=7
// ==========================================================
// DAY 06 — LENGTH VS CAPACITY
// ==========================================================

// func main() {
// 	s := make([]int, 2, 5)
// len=2,cap=5
// 	s[0] = 10
// 	s[1] = 20
//s=[0,0,10,20]
// 	s = append(s, 30)
//s=[0,0,10,20,30]

// 	fmt.Println(s) s=[0,0,10,20,30]
// 	fmt.Println(len(s)) len=5
// 	fmt.Println(cap(s)) cap=5
// }

// ==========================================================
// DAY 07 — BASIC SLICING
// ==========================================================

// func main() {
// 	s := []int{10, 20, 30, 40, 50}
//
// 	a := s[1:4]
//
// 	fmt.Println(a)
// 	fmt.Println(len(a))
// 	fmt.Println(cap(a))
// }

// ==========================================================
// DAY 08 — SLICE EXPRESSIONS
// ==========================================================

// func main() {
// 	s := []int{10, 20, 30, 40, 50}
//
// 	fmt.Println(s[:3])
// 	fmt.Println(s[2:])
// 	fmt.Println(s[:])
// }

// ==========================================================
// DAY 09 — SHARED BACKING ARRAY
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
//
// 	a := s[:2]
// 	b := s[2:4]
//
// 	a[1] = 99
// 	b[0] = 88
//
// 	fmt.Println(s)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// ==========================================================
// DAY 10 — ARRAY TO SLICE
// ==========================================================

// func main() {
// 	a := [5]int{10, 20, 30, 40, 50}
//
// 	s := a[1:3]
//
// 	s[0] = 99
//
// 	fmt.Println(a)
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 11 — CAPACITY OF A SLICE
// ==========================================================

// func main() {
// 	a := [6]int{1, 2, 3, 4, 5, 6}
//
// 	s := a[2:4]
//
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }

// ==========================================================
// DAY 12 — APPEND INTO BACKING ARRAY
// ==========================================================

// func main() {
// 	a := [6]int{1, 2, 3, 4, 5, 6}
//
// 	s := a[2:4]
//
// 	s = append(s, 99)
//
// 	fmt.Println(a)
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 13 — APPEND TWICE
// ==========================================================

// func main() {
// 	a := [6]int{1, 2, 3, 4, 5, 6}
//
// 	s := a[2:4]
//
// 	s = append(s, 99)
// 	s = append(s, 100)
//
// 	fmt.Println(a)
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 14 — APPEND AND REALLOCATION
// ==========================================================

// func main() {
// 	a := [4]int{1, 2, 3, 4}
//
// 	s := a[:2]
//
// 	s = append(s, 99)
// 	s = append(s, 100)
//
// 	fmt.Println(a)
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }

// ==========================================================
// DAY 15 — SLICE ASSIGNMENT
// ==========================================================

// func main() {
// 	s := make([]int, 2, 4)
//
// 	s[0] = 10
// 	s[1] = 20
//
// 	t := s
//
// 	t = append(t, 30)
//
// 	t[0] = 99
//
// 	fmt.Println(s)
// 	fmt.Println(t)
// }

// ==========================================================
// DAY 16 — APPEND WITH NO EXTRA CAPACITY
// ==========================================================

// func main() {
// 	s := make([]int, 2, 2)
//
// 	s[0] = 10
// 	s[1] = 20
//
// 	t := s
//
// 	t = append(t, 30)
//
// 	t[0] = 99
//
// 	fmt.Println(s)
// 	fmt.Println(t)
// }

// ==========================================================
// DAY 17 — THREE SLICE VARIABLES
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3}
//
// 	a := s
// 	b := s
//
// 	a[0] = 100
// 	b[1] = 200
//
// 	fmt.Println(s)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// ==========================================================
// DAY 18 — TWO APPENDS FROM SAME SLICE
// ==========================================================

// func main() {
// 	s := make([]int, 2, 4)
//
// 	s[0] = 10
// 	s[1] = 20
//
// 	a := append(s, 30)
// 	b := append(s, 40)
//
// 	fmt.Println(s)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// ==========================================================
// DAY 19 — PREALLOCATED SLICE
// ==========================================================

// func main() {
// 	s := make([]int, 0, 5)
//
// 	s = append(s, 1)
// 	s = append(s, 2)
// 	s = append(s, 3)
//
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }

// ==========================================================
// DAY 20 — CAPACITY EXCEEDED
// ==========================================================

// func main() {
// 	s := make([]int, 0, 2)
//
// 	s = append(s, 1, 2)
//
// 	fmt.Println(len(s), cap(s))
//
// 	s = append(s, 3)
//
// 	fmt.Println(len(s), cap(s))
// }

// ==========================================================
// DAY 21 — COPY
// ==========================================================

// func main() {
// 	a := []int{1, 2, 3}
// 	b := make([]int, 3)
//
// 	n := copy(b, a)
//
// 	fmt.Println(n)
// 	fmt.Println(b)
// }

// ==========================================================
// DAY 22 — COPY INTO SMALLER SLICE
// ==========================================================

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	b := make([]int, 2)
//
// 	n := copy(b, a)
//
// 	fmt.Println(n)
// 	fmt.Println(b)
// }

// ==========================================================
// DAY 23 — OVERLAPPING COPY
// ==========================================================

// func main() {
// 	a := []int{1, 2, 3, 4}
//
// 	copy(a[1:], a[:3])
//
// 	fmt.Println(a)
// }

// ==========================================================
// DAY 24 — FUNCTION MUTATING SLICE
// ==========================================================

// func change(s []int) {
// 	s[0] = 100
// }
//
// func main() {
// 	s := []int{1, 2, 3}
//
// 	change(s)
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 25 — FUNCTION APPEND
// ==========================================================

// func change(s []int) {
// 	s = append(s, 100)
// }
//
// func main() {
// 	s := []int{1, 2, 3}
//
// 	change(s)
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 26 — FUNCTION APPEND WITH RETURN
// ==========================================================

// func change(s []int) []int {
// 	s = append(s, 100)
// 	return s
// }
//
// func main() {
// 	s := []int{1, 2, 3}
//
// 	change(s)
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 27 — ASSIGNING RETURNED SLICE
// ==========================================================

// func change(s []int) []int {
// 	s = append(s, 100)
// 	return s
// }
//
// func main() {
// 	s := []int{1, 2, 3}
//
// 	s = change(s)
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 28 — RANGE VALUE
// ==========================================================

// func main() {
// 	s := []int{10, 20, 30}
//
// 	for _, v := range s {
// 		v++
// 	}
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 29 — RANGE INDEX
// ==========================================================

// func main() {
// 	s := []int{10, 20, 30}
//
// 	for i := range s {
// 		s[i]++
// 	}
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 30 — APPEND WHILE RANGING
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3}
//
// 	for _, v := range s {
// 		s = append(s, v+10)
// 	}
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 31 — EMPTY VS NIL
// ==========================================================

// func main() {
// 	var a []int
// 	b := []int{}
//
// 	fmt.Println(a == nil)
// 	fmt.Println(b == nil)
// }

// ==========================================================
// DAY 32 — NIL SLICE + APPEND
// ==========================================================

// func main() {
// 	s := []int(nil)
//
// 	fmt.Println(s == nil)
//
// 	s = append(s, 10)
//
// 	fmt.Println(s)
// 	fmt.Println(s == nil)
// }

// ==========================================================
// DAY 33 — SLICE TO ZERO LENGTH
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
//
// 	s = s[:0]
//
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }

// ==========================================================
// DAY 34 — ZERO LENGTH THEN APPEND
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
//
// 	s = s[:0]
//
// 	s = append(s, 99)
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 35 — EMPTY MIDDLE SLICE
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
//
// 	s = s[1:1]
//
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }

// ==========================================================
// DAY 36 — APPEND TO MIDDLE SLICE
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
//
// 	s = s[1:3]
//
// 	s = append(s, 99)
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 37 — FULL SLICE EXPRESSION
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
//
// 	s = s[1:3:3]
//
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
//
// 	s = append(s, 99)
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 38 — NORMAL VS FULL SLICE
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
//
// 	a := s[1:3]
// 	b := s[1:3:3]
//
// 	a = append(a, 99)
// 	b = append(b, 88)
//
// 	fmt.Println(s)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// ==========================================================
// DAY 39 — APPEND AND ORIGINAL SLICE
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3}
//
// 	x := append(s, 4)
//
// 	s[0] = 99
//
// 	fmt.Println(s)
// 	fmt.Println(x)
// }

// ==========================================================
// DAY 40 — APPEND MAY CREATE NEW ARRAY
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3}
//
// 	x := append(s, 4)
//
// 	x[0] = 99
//
// 	fmt.Println(s)
// 	fmt.Println(x)
// }

// ==========================================================
// DAY 41 — ARRAY BACKING + APPEND
// ==========================================================

// func main() {
// 	s := make([]int, 3, 5)
//
// 	s[0] = 10
// 	s[1] = 20
// 	s[2] = 30
//
// 	x := s[:2]
//
// 	x = append(x, 99)
//
// 	fmt.Println(s)
// 	fmt.Println(x)
// }

// ==========================================================
// DAY 42 — SHARED BACKING ARRAY
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3, 4}
//
// 	x := s[1:3]
//
// 	y := append(x, 99)
//
// 	fmt.Println(s)
// 	fmt.Println(x)
// 	fmt.Println(y)
// }

// ==========================================================
// DAY 43 — TWO APPENDS
// ==========================================================

// func main() {
// 	s := make([]int, 2, 4)
//
// 	s[0] = 10
// 	s[1] = 20
//
// 	x := append(s, 30)
// 	y := append(s, 40)
//
// 	x[0] = 99
//
// 	fmt.Println(s)
// 	fmt.Println(x)
// 	fmt.Println(y)
// }

// ==========================================================
// DAY 44 — VARIADIC FUNCTION
// ==========================================================

// func add(s []int, values ...int) []int {
// 	return append(s, values...)
// }
//
// func main() {
// 	s := []int{1, 2}
//
// 	s = add(s, 3, 4, 5)
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 45 — SLICE EXPANSION
// ==========================================================

// func main() {
// 	a := []int{1, 2, 3}
// 	b := []int{4, 5}
//
// 	a = append(a, b...)
//
// 	fmt.Println(a)
// }

// ==========================================================
// DAY 46 — ALIASING + APPEND
// ==========================================================

// func main() {
// 	a := []int{1, 2, 3}
//
// 	b := a[:2]
//
// 	b = append(b, 99)
//
// 	c := append(a, 100)
//
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// ==========================================================
// DAY 47 — DELETE WHILE ITERATING
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
//
// 	for i := 0; i < len(s); i++ {
// 		if s[i]%2 == 0 {
// 			s = append(s[:i], s[i+1:]...)
// 			i--
// 		}
// 	}
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 48 — DELETE FROM SLICE
// ==========================================================

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
//
// 	s = append(s[:1], s[3:]...)
//
// 	fmt.Println(s)
// }

// ==========================================================
// DAY 49 — MULTIPLE REFERENCES
// ==========================================================

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
//
// 	x := a[1:3]
// 	y := a[2:4]
//
// 	x[0] = 99
//
// 	fmt.Println("A:", a)
// 	fmt.Println("X:", x)
// 	fmt.Println("Y:", y)
//
// 	x = append(x, 100)
//
// 	fmt.Println("A:", a)
// 	fmt.Println("X:", x)
// 	fmt.Println("Y:", y)
//
// 	y = append(y, 200)
//
// 	fmt.Println("A:", a)
// 	fmt.Println("X:", x)
// 	fmt.Println("Y:", y)
// }

// ==========================================================
// DAY 50 — FINAL BOSS
// ==========================================================

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
//
// 	x := a[1:3]
// 	y := a[2:4]
//
// 	x[0] = 99
//
// 	fmt.Println("A:", a)
// 	fmt.Println("X:", x)
// 	fmt.Println("Y:", y)
//
// 	x = append(x, 100)
//
// 	fmt.Println("A:", a)
// 	fmt.Println("X:", x)
// 	fmt.Println("Y:", y)
//
// 	y = append(y, 200)
//
// 	fmt.Println("A:", a)
// 	fmt.Println("X:", x)
// 	fmt.Println("Y:", y)
//
// 	z := append([]int{}, a...)
//
// 	z[0] = 500
//
// 	fmt.Println("A:", a)
// 	fmt.Println("Z:", z)
// }