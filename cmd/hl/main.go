package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	txt := append(loadText(), byte(0))
	_var := make([]int, 255)

	for i := 0; i < 10; i++ {
		_var[int(strconv.Itoa(i)[0])] = i
	}

	err := func(pc int) {
		fmt.Printf("syntax error : %.10s\n", string(txt[pc:]))
		os.Exit(1)
	}

	for pc := 0; txt[pc] != 0; pc++ {
		switch txt[pc] {
		case '\n', '\r', ' ', '\t', ';':
			continue
		}

		switch {
		case txt[pc+1] == '=' && txt[pc+3] == ';':
			_var[txt[pc]] = _var[txt[pc+2]]
		case txt[pc+1] == '=' && txt[pc+3] == '+' && txt[pc+5] == ';':
			_var[txt[pc]] = _var[txt[pc+2]] + _var[txt[pc+4]]
		case txt[pc+1] == '=' && txt[pc+3] == '-' && txt[pc+5] == ';':
			_var[txt[pc]] = _var[txt[pc+2]] - _var[txt[pc+4]]
		case txt[pc] == 'p' && txt[pc+1] == 'r' && txt[pc+5] == ' ' && txt[pc+7] == ';':
			fmt.Printf("%d\n", _var[txt[pc+6]])
		default:
			err(pc)
		}

		for txt[pc] != ';' {
			pc++
		}
	}
}

func loadText() []byte {
	if len(os.Args) < 2 {
		fmt.Printf("usage>%s program-file\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("os.Open error : %s\n", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("io.ReadALl error : %s\n", err.Error())
		os.Exit(1)
	}

	return b
}
