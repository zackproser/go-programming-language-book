# Usage 
`go run bitwise.go`

Output 

```bash
Bitwise Operators

a is 60 which in binary is 00111100
b is 13 which in binary is 00001101
c is 0 which in binary is 00000000

Bitwise AND copies a bit to the result if it exists in both operands:
Value of a & b is 12 or 00001100

Bitwise OR copies a bit to the result if it exists in either operand:
Value of a | b is 61 or 00111101

Bitwise XOR copies the bit to the result if it is set in one operand but not in both:
Value of a ^ b is 49 or 00110001

Binary left shift operator. The left operand's value is moved left by the number of bits specified by the right operand:
Value of a << b is 491520 or 1111000000000000000

Binary right shift operator. The left operand's value is moved right by the number of bits specified by the right operand:
Value of a >> b is 0 or 00000000

```