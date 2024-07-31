package cpu

import "fmt"

func (cpu *CPU) LoadRegister(register uint, value int64) error {
	if register < 16 {
		cpu.registers[register] = value
		return nil
	}
	return fmt.Errorf("Register out of bounds: 0x%016X", register)
}

func (cpu *CPU) ReadRegister(register uint) (int64, error) {
	if register < 16 {
		return cpu.registers[register], nil
	}
	return 0, fmt.Errorf("Register out of bounds: 0x%016X", register)
}
