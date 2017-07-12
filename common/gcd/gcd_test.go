package gcd

import "testing"
import "fmt"

func TestGCD(t *testing.T) {
	fmt.Println(Euclid(10, 6))
	fmt.Println(Stein(10, 6))
	fmt.Println(NGCD(10, 6))
}
