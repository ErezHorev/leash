package workarounder

import (
	"fmt"
	"os"
	"testing"
)

func TestSimpleTextParse(t *testing.T) {
	fmt.Println(findMatch("[workaround for #123]"))
	fmt.Println(findMatch("[workaound for #123]"))
	fmt.Println(findMatch("abababab[workaround for #123333333333]zxzxzxzxzxzxz"))
	fmt.Println(findMatch("            [workaround for #1]/t/t"))
}

func TestParseFiles(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Failed getting work dir")
	}
	FindWorkarounds(dir)
}
