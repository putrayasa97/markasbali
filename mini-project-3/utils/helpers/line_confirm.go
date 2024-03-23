package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LineConfirm(s string) bool {
	r := bufio.NewReader(os.Stdin)

	fmt.Printf("%s [y/n]: ", s)

	res, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return false
	}

	return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
}
