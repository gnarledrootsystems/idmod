package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("You must pass a mode string ( Example: -rw-r--r-- )")
		return
	}

	arg := args[0]

	mode := strings.Split(arg, "")

	if len(mode) < 9 {
		fmt.Println("Please pass a valid 9 or 10 character long mode value ( Example: -rw-r--r-- )")
		return
	}

	if len(mode) == 10 {
		mode = mode[1:]
	}

	// Owner
	oRead := get_mode_number(mode[0])
	oWrite := get_mode_number(mode[1])
	oExectute := get_mode_number(mode[2])
	oTotal := oRead + oWrite + oExectute

	// Group
	gRead := get_mode_number(mode[3])
	gWrite := get_mode_number(mode[4])
	gExectute := get_mode_number(mode[5])
	gTotal := gRead + gWrite + gExectute

	// Public
	pRead := get_mode_number(mode[6])
	pWrite := get_mode_number(mode[7])
	pExectute := get_mode_number(mode[8])
	pTotal := pRead + pWrite + pExectute

	modeValue := fmt.Sprintf("%d%d%d", oTotal, gTotal, pTotal)

	fmt.Println(modeValue)

	return
}

/*
  - Regular file.
    b     Block special file.
    c     Character special file.
    d     Directory.
    l     Symbolic link.
    p     FIFO.
    s     Socket.
    w     Whiteout.

    4	Read (r)
    2	Write (w)
    1	Execute (x)
*/
func get_mode_number(value string) int {
	switch strings.ToLower(value) {
	case "-":
		return 0
	case "r":
		return 4
	case "w":
		return 2
	case "x":
		return 1
	default:
		return 0
	}
}
