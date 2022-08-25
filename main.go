package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	arg := args[0]

	fmt.Println(arg)

	/**
		 	   -     Regular file.
	           b     Block special file.
	           c     Character special file.
	           d     Directory.
	           l     Symbolic link.
	           p     FIFO.
	           s     Socket.
	           w     Whiteout.
	*/

	/**
	4	Read (r)
	2	Write (w)
	1	Execute (x)
	*/


	/**
	-rwxr-xr-x
	=
	-
	4
	2
	1
	4
	-
	1
	4
	-
	1
	=
	755
}
