package main

import (
	"fmt"
	// "go-vm/src/internal/assembler"
	"go-vm/src/internal/cpu"
	"go-vm/src/internal/memory"
	"go-vm/src/internal/utils"
	// "os"
)

const (
	STACK = 1024 * 1024       // 1 MB
	MEM   = 128 * 1024 * 1024 // 128 MB
)

func main() {
	// Lexer test WIP
	/*
		if len(os.Args) == 2 {
			contents, err := os.ReadFile(os.Args[1])
			cpu.Error(err)

			tokens := assembler.Lex(string(contents))
			for _, tok := range tokens {
				if tok.Value != "\n" {
					fmt.Printf("%03d:%03d  $  %s\n", tok.Row, tok.Col, tok.Value)
				}
			}
			return
		}
	*/

	// TODO: make assembler and get rid of hardcoded programs
	program := []byte{}

	// Data section
	program = append(program, []byte("test.txt")...)
	program = append(program, 0)
	program = append(program, []byte("Enter a message: ")...)
	program = append(program, 0)

	// BSS section
	bss_start := len(program)
	program = append(program, make([]byte, 128)...)

	// Text section
	text_start := len(program)
	program = append(program, append([]byte{cpu.LOAD, 0}, utils.IntToBytes(1)...)...)
	program = append(program, append([]byte{cpu.LOAD, 1}, utils.IntToBytes(1)...)...)
	program = append(program, append([]byte{cpu.LOAD, 2}, utils.IntToBytes(9)...)...)
	program = append(program, append([]byte{cpu.LOAD, 3}, utils.IntToBytes(17)...)...)
	program = append(program, cpu.SYSCALL)

	program = append(program, append([]byte{cpu.LOAD, 0}, utils.IntToBytes(0)...)...)
	program = append(program, append([]byte{cpu.LOAD, 1}, utils.IntToBytes(0)...)...)
	program = append(program, append([]byte{cpu.LOAD, 2}, utils.IntToBytes(int64(bss_start))...)...)
	program = append(program, append([]byte{cpu.LOAD, 3}, utils.IntToBytes(128)...)...)
	program = append(program, cpu.SYSCALL)

	program = append(program, append([]byte{cpu.LOAD, 0}, utils.IntToBytes(2)...)...)
	program = append(program, append([]byte{cpu.LOAD, 1}, utils.IntToBytes(0)...)...)
	program = append(program, append([]byte{cpu.LOAD, 2}, utils.IntToBytes(577)...)...)
	program = append(program, append([]byte{cpu.LOAD, 3}, utils.IntToBytes(420)...)...)
	program = append(program, cpu.SYSCALL)
	program = append(program, cpu.MOV, 15, 0)

	program = append(program, append([]byte{cpu.LOAD, 0}, utils.IntToBytes(int64(bss_start))...)...)
	program = append(program, append([]byte{cpu.LOAD, 1}, utils.IntToBytes(0)...)...)

	counter := len(program)
	program = append(program, cpu.MOV_VAL, 2, 0)
	program = append(program, append([]byte{cpu.CMP_VAL, 2}, utils.IntToBytes(0)...)...)
	program = append(program, append([]byte{cpu.JE}, utils.IntToBytes(int64(counter+35))...)...)
	program = append(program, cpu.INC, 0)
	program = append(program, cpu.INC, 1)
	program = append(program, append([]byte{cpu.JMP}, utils.IntToBytes(int64(counter))...)...)

	program = append(program, cpu.MOV, 14, 1)

	program = append(program, append([]byte{cpu.LOAD, 0}, utils.IntToBytes(1)...)...)
	program = append(program, cpu.MOV, 1, 15)
	program = append(program, append([]byte{cpu.LOAD, 2}, utils.IntToBytes(int64(bss_start))...)...)
	program = append(program, cpu.MOV, 3, 14)
	program = append(program, cpu.SYSCALL)

	program = append(program, append([]byte{cpu.LOAD, 0}, utils.IntToBytes(3)...)...)
	program = append(program, cpu.MOV, 1, 15)
	program = append(program, cpu.SYSCALL)
	program = append(program, cpu.HLT)

	mem := memory.NewMemory(uint64(len(program) + STACK + MEM))
	mem.LoadProgram(program)

	vm := cpu.NewCPU(mem, uint64(text_start), uint64(len(program)))

	vm.AddSyscall(0, (*cpu.CPU).SysRead)
	vm.AddSyscall(1, (*cpu.CPU).SysWrite)
	vm.AddSyscall(2, (*cpu.CPU).SysOpen)
	vm.AddSyscall(3, (*cpu.CPU).SysClose)
	vm.AddSyscall(4, (*cpu.CPU).SysStat)

	vm.Run()

	for i := 0; i < 16; i++ {
		reg, _ := vm.ReadRegister(uint(i))
		fmt.Printf("%%%d = %d\n", i, reg)
	}

	fmt.Printf("\nExecutable size: %d bytes\n", len(program))
}
