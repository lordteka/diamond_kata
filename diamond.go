package main

import (
	"fmt"
	"os"
	"errors"
	"strings"
)

var Diamond [26]string = [26]string {
	"                         A",
	"                        B B",
	"                       C   C",
	"                      D     D",
	"                     E       E",
	"                    F         F",
	"                   G           G",
	"                  H             H",
	"                 I               I",
	"                J                 J",
	"               K                   K",
	"              L                     L",
	"             M                       M",
	"            N                         N",
	"           O                           O",
	"          P                             P",
	"         Q                               Q",
	"        R                                 R",
	"       S                                   S",
	"      T                                     T",
	"     U                                       U",
	"    V                                         V",
	"   W                                           W",
	"  X                                             X",
	" Y                                               Y",
	"Z                                                 Z",
}

var a byte = 'a'
var z byte = 'z'

func add_reverse(diamond []string) []string {
	for i := len(diamond) - 2; i >= 0; i-- {
		diamond = append(diamond, diamond[i])
	}
	return diamond
}

func make_diamond(letter byte) (string) {
	var diamond []string
	stop := int(letter) - int(a)
	start := int(z) - int(letter)

	for i := 0; i <= stop ; i++ {
		diamond = append(diamond, Diamond[i][start:])
	}
	return strings.Join(add_reverse(diamond), "\n")
}

func get_arg() (byte, error) {
	if len(os.Args[1:]) != 1 {
		return 0, errors.New("Bad number of argument : one argument is expected")
	}
	arg := os.Args[1]
	if len(arg) > 1 {
		return 0, errors.New("Bad argument : expected a letter")
	}
	letter := arg[0]
	if (letter < 'A' || letter > 'Z') && (letter < 'a' || letter > 'z') {
		return 0, errors.New("Bad argument : expected a letter")
	}
	return letter, nil
}

func main() {
	letter, err := get_arg()
	if err != nil {
		fmt.Println(err)
		return
	}
	if letter >= 'A' && letter <= 'Z' {
		a = 'A'
		z = 'Z'
	}
	println(make_diamond(letter))
}
