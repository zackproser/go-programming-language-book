package main

import "fmt"

func main() {
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0
	letters := []string{"a", "b", "c"}
	nums := []uint{a, b, c}

	fmt.Println("Bitwise Operators")
	fmt.Println()

	for i, val := range nums {
		fmt.Printf("%s is %d which in binary is %08b\n", letters[i], val, val)
	}

	fmt.Println()

	c = a & b /* 12 = 0000 1100 */
	fmt.Println("Bitwise AND copies a bit to the result if it exists in both operands:")
	fmt.Printf("Value of a & b is %d or %08b\n", c, c)
	fmt.Println()

	c = a | b /* 61 = 0011 1101 */
	fmt.Println("Bitwise OR copies a bit to the result if it exists in either operand:")
	fmt.Printf("Value of a | b is %d or %08b\n", c, c)
	fmt.Println()

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Println("Bitwise XOR copies the bit to the result if it is set in one operand but not in both:")
	fmt.Printf("Value of a ^ b is %d or %08b\n", c, c)
	fmt.Println()


	c = a << b
	fmt.Println("Binary left shift operator. The left operand's value is moved left by the number of bits specified by the right operand:")
	fmt.Printf("Value of a << b is %d or %08b\n", c, c)
	fmt.Println()


	c = a >> b
	fmt.Println("Binary right shift operator. The left operand's value is moved right by the number of bits specified by the right operand:")
	fmt.Printf("Value of a >> b is %d or %08b\n", c, c)
	fmt.Println()
}

