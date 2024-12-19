package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"strconv"
)

// FFT function that applies the Fast Fourier Transform
func FFT(a []complex128) []complex128 {
	n := len(a)
	if n <= 1 {
		return a
	}

	// Split into even and odd parts
	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = a[2*i]
		odd[i] = a[2*i+1]
	}

	// Recursively apply FFT
	even = FFT(even)
	odd = FFT(odd)

	// Combine results
	result := make([]complex128, n)
	for i := 0; i < n/2; i++ {
		// Calculate the complex twiddle factor
		twiddle := cmplx.Exp(complex(0, -2*math.Pi*float64(i)/float64(n)))
		result[i] = even[i] + twiddle*odd[i]
		result[i+n/2] = even[i] - twiddle*odd[i]
	}
	return result
}

// Inverse FFT
func IFFT(a []complex128) []complex128 {
	n := len(a)
	// Take complex conjugate, apply FFT and then take complex conjugate again
	for i := range a {
		a[i] = cmplx.Conj(a[i])
	}
	a = FFT(a)
	for i := range a {
		a[i] = cmplx.Conj(a[i]) / complex(float64(n), 0)
	}
	return a
}

// Multiply two large numbers represented by slices of digits
func multiplyLargeNumbers(a, b []int) []int {
	// Size of the result
	n := 1
	for n < len(a)+len(b)-1 {
		n *= 2
	}

	// Convert digits to complex numbers
	A := make([]complex128, n)
	B := make([]complex128, n)
	for i := 0; i < len(a); i++ {
		A[i] = complex(float64(a[i]), 0)
	}
	for i := 0; i < len(b); i++ {
		B[i] = complex(float64(b[i]), 0)
	}

	// Apply FFT
	A = FFT(A)
	B = FFT(B)

	// Multiply pointwise
	C := make([]complex128, n)
	for i := 0; i < n; i++ {
		C[i] = A[i] * B[i]
	}

	// Inverse FFT to get the result
	C = IFFT(C)

	// Extract the real part and convert back to integers
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = int(real(C[i]) + 0.5)
	}

	// Carry over (handle digits greater than 10)
	for i := 0; i < n-1; i++ {
		if result[i] >= 10 {
			result[i+1] += result[i] / 10
			result[i] %= 10
		}
	}

	// Find the actual size of the result
	size := n
	for size > 1 && result[size-1] == 0 {
		size--
	}

	return result[:size]
}

// Convert a string of digits to an array of integers
func stringToDigits(s string) ([]int, error) {
	var digits []int
	for i := len(s) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return nil, fmt.Errorf("invalid character in number: %v", s[i])
		}
		digits = append(digits, digit)
	}
	return digits, nil
}

func main() {
	// Check if we have the correct number of arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <number1> <number2>")
		return
	}

	// Get the two numbers from the command line arguments
	num1 := os.Args[1]
	num2 := os.Args[2]

	// Convert the string numbers to slices of digits
	digits1, err := stringToDigits(num1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	digits2, err := stringToDigits(num2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Multiply the two large numbers
	result := multiplyLargeNumbers(digits1, digits2)

	// Print the result
	fmt.Print("Product: ")
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
	fmt.Println()
}
