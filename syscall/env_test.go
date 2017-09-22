package syscalltest

import "testing"
import "fmt"
import "syscall"

func TestEnv(t *testing.T) {
	fmt.Println(syscall.Environ())
}
