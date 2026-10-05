// Command friend-code 演示邀请码的生成、校验与台账查询。
//
//	go run . generate alice
//	go run . verify ABCDEF
//	go run . sample charset
package main

import (
	"fmt"
	"os"
	"strings"

	"example.com/friendcode/codec"
	"example.com/friendcode/registry"
)

func run(args []string, out *os.File, errOut *os.File) int {
	if len(args) < 2 {
		fmt.Fprintln(errOut, "用法：friend-code generate <player> | verify <code> | sample <名>")
		return 2
	}
	command := args[0]
	value := strings.Join(args[1:], " ")
	book := registry.New()
	switch command {
	case "generate":
		fmt.Fprintf(out, "player=%s code=%s\n", value, book.Assign(value))
	case "verify":
		fmt.Fprintf(out, "ok=%v\n", codec.Verify(value) == nil)
	case "sample":
		runSample(book, value, out, errOut)
	default:
		fmt.Fprintln(errOut, "未知命令")
		return 2
	}
	return 0
}

func runSample(book *registry.Registry, name string, out *os.File, errOut *os.File) {
	switch name {
	case "charset":
		confusable := 0
		for index := 0; index < 200; index++ {
			if strings.ContainsAny(codec.Encode(fmt.Sprintf("p-%d", index)), "O0I1") {
				confusable++
			}
		}
		fmt.Fprintf(out, "confusable=%d\n", confusable)
	case "checksum":
		fmt.Fprintf(out, "ok=%v\n", codec.Verify("AAAAAA") == nil)
	case "dup":
		first := book.Assign("alice")
		second := book.Assign("alice")
		fmt.Fprintf(out, "same=%v\n", first == second)
	case "case":
		code := codec.Encode("alice")
		fmt.Fprintf(out, "lowercase_ok=%v\n", codec.Verify(strings.ToLower(code)) == nil)
	case "work":
		for index := 0; index < 3000; index++ {
			book.Assign(fmt.Sprintf("p-%d", index))
		}
		for index := 0; index < 3000; index++ {
			book.Owner("AAAAAA")
		}
		fmt.Fprintf(out, "work=%d\n", book.Scanned())
	default:
		fmt.Fprintln(errOut, "需要 sample charset|checksum|dup|case|work")
	}
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
