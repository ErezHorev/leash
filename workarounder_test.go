package workarounder

import (
	"os"
	"testing"
)

func TestSimpleFindMatch(t *testing.T) {
	t.Log(findMatch("[workaround for #123]"))
	t.Log(findMatch("[workaound for #123]"))
	t.Log(findMatch("abababab[workaround for #123333333333]zxzxzxzxzxzxz"))
	t.Log(findMatch("            [workaround for #1]/t/t"))
}

func TestParseFiles(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Failed getting work dir")
	}
	FindWorkarounds(dir)
}
