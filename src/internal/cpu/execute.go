package cpu

import (
	"fmt"
	"go-vm/src/internal/utils"
	"os"
)

func Error(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func (cpu *CPU) executeInstruction() {
	opcode, err := cpu.memory.Read(cpu.pc)
	Error(err)

	cpu.pc++
	switch opcode {
	case LOAD:
		reg, err := cpu.memory.Read(cpu.pc)
		Error(err)
		num, err := cpu.memory.ReadBytes(cpu.pc+1, 8)
		Error(err)
		err = cpu.LoadRegister(uint(reg), utils.BytesToInt(num))
		Error(err)
		cpu.pc += 9
	case LOAD_VAL:
		reg, err := cpu.memory.Read(cpu.pc)
		Error(err)
		addr, err := cpu.memory.ReadBytes(cpu.pc+1, 8)
		Error(err)
		value, err := cpu.memory.Read(uint64(utils.BytesToInt(addr)))
		Error(err)
		err = cpu.LoadRegister(uint(reg), int64(value))
		Error(err)
		cpu.pc += 9
	case MOV:
		reg1, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg2, err := cpu.memory.Read(cpu.pc + 1)
		Error(err)
		reg2_val, err := cpu.ReadRegister(uint(reg2))
		Error(err)
		err = cpu.LoadRegister(uint(reg1), reg2_val)
		Error(err)
		cpu.pc += 2
	case MOV_VAL:
		reg1, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg2, err := cpu.memory.Read(cpu.pc + 1)
		Error(err)
		reg2_val, err := cpu.ReadRegister(uint(reg2))
		Error(err)
		value, err := cpu.memory.Read(uint64(reg2_val))
		Error(err)
		err = cpu.LoadRegister(uint(reg1), int64(value))
		Error(err)
		cpu.pc += 2
	case STORE:
		bytes, err := cpu.memory.ReadBytes(cpu.pc, 8)
		Error(err)
		addr := utils.BytesToInt(bytes)
		reg, err := cpu.memory.Read(cpu.pc + 8)
		Error(err)
		regVal, err := cpu.ReadRegister(uint(reg))
		Error(err)
		if regVal > 255 {
			Error(fmt.Errorf("%d could not be interpreted as byte", regVal))
		}
		err = cpu.memory.Write(uint64(addr), byte(regVal))
		Error(err)
		cpu.pc += 9
	case STORE_VAL:
		bytes, err := cpu.memory.ReadBytes(cpu.pc, 8)
		Error(err)
		addr := utils.BytesToInt(bytes)
		value, err := cpu.memory.Read(cpu.pc + 8)
		Error(err)
		err = cpu.memory.Write(uint64(addr), value)
		Error(err)
		cpu.pc += 9
	case INC:
		reg, err := cpu.memory.Read(cpu.pc)
		Error(err)
		regVal, err := cpu.ReadRegister(uint(reg))
		Error(err)
		err = cpu.LoadRegister(uint(reg), regVal+1)
		Error(err)
		cpu.pc++
	case ADD:
		reg1, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg1_val, err := cpu.ReadRegister(uint(reg1))
		Error(err)
		reg2, err := cpu.memory.Read(cpu.pc + 1)
		Error(err)
		reg2_val, err := cpu.ReadRegister(uint(reg2))
		Error(err)
		err = cpu.LoadRegister(uint(reg1), reg1_val+reg2_val)
		Error(err)
		cpu.pc += 2
	case SUB:
		reg1, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg1_val, err := cpu.ReadRegister(uint(reg1))
		Error(err)
		reg2, err := cpu.memory.Read(cpu.pc + 1)
		Error(err)
		reg2_val, err := cpu.ReadRegister(uint(reg2))
		Error(err)
		err = cpu.LoadRegister(uint(reg1), reg1_val-reg2_val)
		Error(err)
		cpu.pc += 2
	case MUL:
		reg1, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg1_val, err := cpu.ReadRegister(uint(reg1))
		Error(err)
		reg2, err := cpu.memory.Read(cpu.pc + 1)
		Error(err)
		reg2_val, err := cpu.ReadRegister(uint(reg2))
		Error(err)
		err = cpu.LoadRegister(uint(reg1), reg1_val*reg2_val)
		Error(err)
		cpu.pc += 2
	case DIV:
		reg1, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg1_val, err := cpu.ReadRegister(uint(reg1))
		Error(err)
		reg2, err := cpu.memory.Read(cpu.pc + 1)
		Error(err)
		reg2_val, err := cpu.ReadRegister(uint(reg2))
		Error(err)
		err = cpu.LoadRegister(uint(reg1), reg1_val/reg2_val)
		Error(err)
		err = cpu.LoadRegister(uint(reg2), reg1_val%reg2_val)
		Error(err)
		cpu.pc += 2
	case NEG:
		reg, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg_val, err := cpu.ReadRegister(uint(reg))
		Error(err)
		err = cpu.LoadRegister(uint(reg), ^reg_val+1)
		Error(err)
		cpu.pc++
	case AND:
		reg1, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg1_val, err := cpu.ReadRegister(uint(reg1))
		Error(err)
		reg2, err := cpu.memory.Read(cpu.pc + 1)
		Error(err)
		reg2_val, err := cpu.ReadRegister(uint(reg2))
		Error(err)
		err = cpu.LoadRegister(uint(reg1), reg1_val&reg2_val)
		Error(err)
		cpu.pc += 2
	case OR:
		reg1, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg1_val, err := cpu.ReadRegister(uint(reg1))
		Error(err)
		reg2, err := cpu.memory.Read(cpu.pc + 1)
		Error(err)
		reg2_val, err := cpu.ReadRegister(uint(reg2))
		Error(err)
		err = cpu.LoadRegister(uint(reg1), reg1_val|reg2_val)
		Error(err)
		cpu.pc += 2
	case XOR:
		reg1, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg1_val, err := cpu.ReadRegister(uint(reg1))
		Error(err)
		reg2, err := cpu.memory.Read(cpu.pc + 1)
		Error(err)
		reg2_val, err := cpu.ReadRegister(uint(reg2))
		Error(err)
		err = cpu.LoadRegister(uint(reg1), reg1_val^reg2_val)
		Error(err)
		cpu.pc += 2
	case NOT:
		reg, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg_val, err := cpu.ReadRegister(uint(reg))
		Error(err)
		err = cpu.LoadRegister(uint(reg), ^reg_val)
		Error(err)
		cpu.pc++
	case JMP:
		addr, err := cpu.memory.ReadBytes(cpu.pc, 8)
		Error(err)
		cpu.pc = uint64(utils.BytesToInt(addr))
	case CMP:
		reg1, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg1_val, err := cpu.ReadRegister(uint(reg1))
		Error(err)
		reg2, err := cpu.memory.Read(cpu.pc + 1)
		Error(err)
		reg2_val, err := cpu.ReadRegister(uint(reg2))
		Error(err)
		if reg1_val > reg2_val {
			cpu.cmpState = 1
		} else if reg1_val < reg2_val {
			cpu.cmpState = -1
		} else {
			cpu.cmpState = 0
		}
		cpu.pc += 2
	case CMP_VAL:
		reg, err := cpu.memory.Read(cpu.pc)
		Error(err)
		reg1_val, err := cpu.ReadRegister(uint(reg))
		Error(err)
		bytes, err := cpu.memory.ReadBytes(cpu.pc+1, 8)
		Error(err)
		value := utils.BytesToInt(bytes)
		if reg1_val > value {
			cpu.cmpState = 1
		} else if reg1_val < value {
			cpu.cmpState = -1
		} else {
			cpu.cmpState = 0
		}
		cpu.pc += 9
	case JE:
		addr, err := cpu.memory.ReadBytes(cpu.pc, 8)
		Error(err)
		if cpu.cmpState == 0 {
			cpu.pc = uint64(utils.BytesToInt(addr))
		} else {
			cpu.pc += 8
		}
	case JNE:
		addr, err := cpu.memory.ReadBytes(cpu.pc, 8)
		Error(err)
		if cpu.cmpState != 0 {
			cpu.pc = uint64(utils.BytesToInt(addr))
		} else {
			cpu.pc += 8
		}
	case JG:
		addr, err := cpu.memory.ReadBytes(cpu.pc, 8)
		Error(err)
		if cpu.cmpState == 1 {
			cpu.pc = uint64(utils.BytesToInt(addr))
		} else {
			cpu.pc += 8
		}
	case JL:
		addr, err := cpu.memory.ReadBytes(cpu.pc, 8)
		Error(err)
		if cpu.cmpState == -1 {
			cpu.pc = uint64(utils.BytesToInt(addr))
		} else {
			cpu.pc += 8
		}
	case JGE:
		addr, err := cpu.memory.ReadBytes(cpu.pc, 8)
		Error(err)
		if cpu.cmpState != -1 {
			cpu.pc = uint64(utils.BytesToInt(addr))
		} else {
			cpu.pc += 8
		}
	case JLE:
		addr, err := cpu.memory.ReadBytes(cpu.pc, 8)
		Error(err)
		if cpu.cmpState != 1 {
			cpu.pc = uint64(utils.BytesToInt(addr))
		} else {
			cpu.pc += 8
		}
	case SYSCALL:
		sysnum, _ := cpu.ReadRegister(0)
		if fn, exists := cpu.syscallMap[sysnum]; exists {
			fn(cpu)
		} else {
			Error(fmt.Errorf("unknown syscall %d @ %d", sysnum, cpu.pc))
		}
	case HLT:
		cpu.halted = true
	default:
		fmt.Printf("Unknown Opcode @ %d: 0x%02X\n", cpu.pc, opcode)
		cpu.halted = true
	}
}
