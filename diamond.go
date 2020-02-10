package main

import (
//	"fmt"
	"os"
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

func add_reverse(diamond []string) []string {
	for i := len(diamond) - 2; i >= 0; i-- {
		diamond = append(diamond, diamond[i])
	}
	return diamond
}

func make_diamond(letter byte) (string) {
	var c byte = 'A'
	var diamond []string
	stop := int(letter) - int(c)
	start := int('Z') - int(letter)

	for i := 0; i <= stop ; i++ {
		diamond = append(diamond, Diamond[i][start:])
	}
	return strings.Join(add_reverse(diamond), "\n")
}

func get_arg() byte {
	//arg = os.Arg[1:]
	return byte(os.Args[1][0])
}

func main() {
	println(make_diamond(get_arg()))
}
