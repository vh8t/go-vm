package cpu

const (
	NOP          byte = iota // Not OpCode
	LOAD                     // `mov reg, value`     | Load number/address into register
	LOAD_VAL                 // `mov reg, [addr]`    | Load value from address into register
	MOV                      // `mov reg1, reg2`     | Move value from register 2 to register 1
	MOV_VAL                  // `mov reg1, [reg2]`   | Move value from register 2 address into register 1
	PUSH                     // `push value`         | Push value onto the stack
	PUSH_VAL                 // `push [addr]`        | Push value from address onto the stack
	PUSH_REG                 // `push reg`           | Push value from register onto the stack
	PUSH_REG_VAL             // `push [reg]`         | Push value from register address onto the stack
	STORE                    // `store addr, reg`    | Store value from register in memory
	STORE_VAL                // `store addr, byte`   | Store value in memory
	INC                      // `inc reg`            | Increment register value by 1
	ADD                      // `add reg1, reg2`     | Add two registers together and store answer in register 1
	SUB                      // `sub reg1, reg2`     | Subtract two registers from eachother and store answer in register 1
	MUL                      // `mul reg1, reg2`     | Multiply two registers with eachother and store answer in register 1
	DIV                      // `div reg1, reg2`     | Divide two registers with eachother and store answe in register 1 and remainder in register 2
	JMP                      // `jmp label`          | Jump to specified byte in bytecode
	CMP                      // `cmp reg1, reg2`     | Compare two registers with each other
	CMP_VAL                  // `cmp reg, value`     | Compare value to a register
	JE                       // `je label`           | Jump to specified byte if compare resulted in 0
	JNE                      // `jne label`          | Jump to specified byte if compare didn't result in 0
	JG                       // `jg label`           | Jump to specified byte if compare resulted in 1
	JL                       // `jl label`           | Jump to specified byte if compare resulted in -1
	JGE                      // `jge label`          | Jump to specified byte if compare didn't result in -1
	JLE                      // `jle label`          | Jump to specified byte if compare didn't result in 1
	NEG                      // `neg reg`            | Negate value in register
	AND                      // `and reg1, reg2`     | Perform bitwise AND on two registers and store it in register 1
	OR                       // `or reg1, reg2`      | Perform bitwise OR on two registers and store it in register 1
	XOR                      // `xor reg1, reg2`     | Perform XOR on two registers and store it in register 1
	NOT                      // `not reg`            | Invert bits of register
	SYSCALL                  // `syscall`            | Perform a syscall
	HLT                      // `halt`               | Stop the vm execution immidiatelly
)
