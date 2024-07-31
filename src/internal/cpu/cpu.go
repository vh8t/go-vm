package cpu

import "go-vm/src/internal/memory"

type Syscall func(*CPU)

type CPU struct {
	registers  [16]int64
	pc         uint64
	sp         uint64
	memory     *memory.Memory
	halted     bool
	cmpState   int
	syscallMap map[int64]Syscall
}

func NewCPU(mem *memory.Memory, start, stack uint64) *CPU {
	return &CPU{
		registers:  [16]int64{0},
		pc:         start,
		sp:         stack,
		memory:     mem,
		halted:     false,
		cmpState:   0,
		syscallMap: map[int64]Syscall{},
	}
}

func (c *CPU) AddSyscall(code int64, syscall Syscall) {
	c.syscallMap[code] = syscall
}

func (cpu *CPU) Run() {
	for !cpu.halted {
		cpu.executeInstruction()
	}
}
