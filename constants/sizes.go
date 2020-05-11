package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 10995116277776
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424
	YiB // 1208925819614629174706176
)

func main() {
	fmt.Printf("KiB: %d\n", KiB)
	fmt.Printf("MiB: %d\n", MiB)
	fmt.Printf("GiB: %d\n", GiB)
	fmt.Printf("TiB: %d\n", TiB) // (exceeds 1 << 32 )
	fmt.Printf("PiB: %d\n", PiB)
	fmt.Printf("EiB: %d\n", EiB)
	fmt.Printf("ZiB: %d\n", ZiB) // (exceeds 1 << 64 )
	fmt.Printf("YiB: %d\n", YiB)
}