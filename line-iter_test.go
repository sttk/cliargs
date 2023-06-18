package cliargs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineIter_Next_emptyText(t *testing.T) {
	text := ""
	iter := newLineIter(text, 20)

	line, status := iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, text)

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_oneCharText(t *testing.T) {
	text := "a"
	iter := newLineIter(text, 20)

	line, status := iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, text)

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_lessThanLineWidth(t *testing.T) {
	text := "1234567890123456789"
	iter := newLineIter(text, 20)

	line, status := iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, text)

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_equalToLineWidth(t *testing.T) {
	text := "12345678901234567890"
	iter := newLineIter(text, 20)

	line, status := iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, text)

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_breakAtLineBreakOppotunity(t *testing.T) {
	text := "1234567890 abcdefghij"
	iter := newLineIter(text, 20)

	line, status := iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, text[0:11])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, text[11:21])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_removeHeadingSpaceOfEachLine(t *testing.T) {
	text := "12345678901234567890   abcdefghij"
	iter := newLineIter(text, 20)

	line, status := iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, text[0:20])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, text[23:])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_thereIsNoLineBreakOppotunity(t *testing.T) {
	text := "12345678901234567890abcdefghij"
	iter := newLineIter(text, 20)

	line, status := iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, text[0:20])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, text[20:])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "")
}

func TestLineIter_setIndent(t *testing.T) {
	text := "12345678901234567890abcdefghij"
	iter := newLineIter(text, 10)

	line, status := iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, text[0:10])

	iter.setIndent(3)

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "   "+text[10:17])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "   "+text[17:24])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "   "+text[24:])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "")
}

func TestLineIter_resetText(t *testing.T) {
	text := "12345678901234567890"
	iter := newLineIter(text, 12)

	line, status := iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, text[0:12])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, text[12:])

	text = "abcdefghijklmnopqrstuvwxyz"
	iter.resetText(text)

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, text[0:12])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, text[12:24])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, text[24:])

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "")
}

// This text is quoted from https://go.dev/doc/
const longText string = `The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.`

func TestLineIter_Next_tryLongText(t *testing.T) {
	iter := newLineIter(longText, 20)

	line, status := iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "The Go programming ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "language is an open ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "source project to ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "make programmers ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "more productive.")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "Go is expressive, ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "concise, clean, and ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "efficient. Its ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "concurrency ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "mechanisms make it ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "easy to write ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "programs that get ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "the most out of ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "multicore and ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "networked machines, ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "while its novel type")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "system enables ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "flexible and modular")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "program ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "construction. Go ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "compiles quickly to ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "machine code yet has")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "the convenience of ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "garbage collection ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "and the power of ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "run-time reflection.")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "It's a fast, ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "statically typed, ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "compiled language ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "that feels like a ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "dynamically typed, ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, line, "interpreted ")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "language.")

	line, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_printLongText(t *testing.T) {
	iter := newLineIter(longText, 20)

	for {
		line, status := iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
}

func TestLineIter_setIndentToLongText(t *testing.T) {
	iter := newLineIter(longText, 40)

	line, status := iter.Next()
	fmt.Println(line)

	iter.setIndent(8)

	for {
		if status == ITER_NO_MORE {
			break
		}
		line, status = iter.Next()
		fmt.Println(line)
	}
}

func TestLineIter_textContainsNonPrintChar(t *testing.T) {
	text := "abcdefg\u0002hijklmn"
	iter := newLineIter(text, 10)

	line, status := iter.Next()
	assert.Equal(t, line, "abcdefghij")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "klmn")
	assert.Equal(t, status, ITER_NO_MORE)
}
