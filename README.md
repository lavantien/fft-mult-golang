# Large Integer Multiplication using FFT (Fast Fourier Transform) in Go

This Go program demonstrates the multiplication of two very large integers using the Fast Fourier Transform (FFT) technique. The algorithm leverages FFT to multiply large numbers in an efficient manner, making it suitable for high-performance applications involving very large integers.

## Features

- **FFT-based Large Integer Multiplication**: Uses the Fast Fourier Transform (FFT) for efficient multiplication of large integers.
- **Command Line Input**: Accepts two large numbers as command-line arguments.
- **No External Libraries**: The program does not rely on external libraries, making it a self-contained solution.

## Prerequisites

- Golang installed on your system.
- Basic understanding of FFT and large number multiplication.

## How to Use

### Command Line Usage

1. Clone or download this repository.
2. Open a terminal and navigate to the project directory.
3. Run the program by passing two large numbers as arguments.

Example:

```bash
go run main.go 3841682376418392641823057109274012740124782374713256137851342713874613741713829704 19283701563240141570740147262317641246123469823146213946213894612642319000000000
```

This will output:

```
Product: 74081856447611460169228840027029381325989295181471445060537582314669396748922631925913987093072481750102917096092877111278600359360827827535040229643576000000000
```

### Input Format

- The two numbers should be provided as command-line arguments.
- The numbers must be positive integers and passed as strings (without commas, spaces, or other symbols).

## How It Works

1. **Convert Input Numbers to Digits**: The program converts the input numbers (strings) into slices of individual digits, with the least significant digit at index 0.
2. **FFT-based Multiplication**: The program applies the Fast Fourier Transform (FFT) to both numbers, treating them as polynomials with the digits as coefficients.
3. **Pointwise Multiplication**: It multiplies the FFT-transformed arrays pointwise.
4. **Inverse FFT**: The inverse FFT is applied to get back to the time domain, which gives the product in the form of digits.
5. **Carry Over**: It handles carries for digits greater than 10 after the inverse FFT.
6. **Display Result**: Finally, the product is printed to the console.

## Notes

- The program uses FFT to perform efficient multiplication of large integers, which helps handle larger numbers than traditional methods.
- This implementation only works for positive integers and assumes that the inputs are valid.
- The program does not handle negative numbers or non-integer inputs.

## Limitations

- The current implementation works with positive integers and assumes valid input.
