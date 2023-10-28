package main

import (
	"fmt"
	"log"
	"os"
)

var (
	srcCodes    string
	srcInput    string
	srcCodeLen  int
	srcInputLen int
	codePos     int
	inputPos    int
)
var (
	tape    []byte
	tapePos int
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: gobrainfuck '${brainfuck_codes}' '${input_bytes}'")
	}
	srcCodes = os.Args[1]
	if len(os.Args) > 2 {
		srcInput = os.Args[2]
	}
	precheck()
	for eat() {
	}
}

func precheck() {
	ptr := 0
	rmax := 0
	brackets := 0
	for _, ch := range srcCodes {
		switch ch {
		case '<':
			ptr--
		case '>':
			ptr++
			rmax = max(rmax, ptr)
		case '[':
			brackets++
		case ']':
			brackets--
		case '.', ',', '+', '-':
		default:
			log.Fatalf("invalid char: %c\n", ch)
		}
		if ptr < 0 {
			log.Fatalln("invalid pointer, pointer cannot be neg.")
		}
	}
	if brackets != 0 {
		log.Fatalln("invalid '[' && ']', not matched.")
	}
	tape = make([]byte, rmax+1)
	srcCodeLen = len(srcCodes)
	srcInputLen = len(srcInput)
}

func eat() (hasNext bool) {
	if codePos == len(srcCodes) {
		return false
	}

	ch := srcCodes[codePos]

	switch ch {
	case '>':
		tapePos++
		codePos++
	case '<':
		tapePos--
		codePos++
	case '.':
		fmt.Fprintf(os.Stderr, "%c", tape[tapePos])
		codePos++
	case '+':
		tape[tapePos]++
		codePos++
	case '-':
		tape[tapePos]--
		codePos++
	case ',':
		b, ok := getch()
		if ok {
			tape[tapePos] = b
		}
		codePos++
	case '[':
		if tape[tapePos] == 0 {
			jumpForward()
		} else {
			codePos++
		}
	case ']':
		jumpBackward()
	}

	if tapePos >= len(tape) {
		incrTape()
	}

	return codePos < len(srcCodes)
}

func jumpForward() {
	brackets := 0
	for i := codePos; i < srcCodeLen; i++ {
		switch srcCodes[i] {
		case '[':
			brackets++
		case ']':
			brackets--
			if brackets == 0 {
				codePos = i + 1
				return
			}
		}
	}
}

func jumpBackward() {
	brackets := 0
	for i := codePos; i >= 0; i-- {
		switch srcCodes[i] {
		case ']':
			brackets++
		case '[':
			brackets--
			if brackets == 0 {
				codePos = i
				return
			}
		}
	}
}

func getch() (byte, bool) {
	if inputPos >= srcInputLen {
		return 0, false
	}
	return srcInput[inputPos], true
}

func incrTape() {
	newTape := make([]byte, len(tape)*2)
	copy(newTape, tape)
	tape = newTape
}
