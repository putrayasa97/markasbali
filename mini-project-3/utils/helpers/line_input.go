package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LineInput(title string, variable interface{}) {
	var err error
	fmt.Print(title)
	switch v := variable.(type) {
	case *int:
		var input string
		_, err = fmt.Scanln(&input)
		if err == nil {
			*v, err = strconv.Atoi(input)
		}
	case *string:
		var readString string
		input := bufio.NewReader(os.Stdin)
		readString, err = input.ReadString('\n')
		*v = strings.Replace(readString, "\n", "", 1)
	default:
		fmt.Println("Tipe data tidak didukung")
		return
	}

	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}
}
