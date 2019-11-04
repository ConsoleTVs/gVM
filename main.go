package main

import "fmt"

// Memory size.
const memorySize = 32

// The kind of the opcodes.
type opcodeKind = int8

const (
	loadi opcodeKind = iota
	addi
	compare
	jump
	branch
	exit
)

// Represents an opcode with 3 operands.
type opcode struct {
	kind opcodeKind
	op1  int64
	op2  int64
	op3  int64
}

// Entry point.
func main() {
	// Program memory (registers)
	var memory [memorySize]int64
	// Program code to execute.
	code := []opcode{
		opcode{kind: loadi, op1: 0, op2: 1000000000},  // r0 = 1000000000;
		opcode{kind: loadi, op1: 1, op2: 0},           // r1 = 0;
		opcode{kind: compare, op1: 2, op2: 0, op3: 1}, // r2 = r0 == r1;
		opcode{kind: branch, op1: 2, op2: 2},          // if (r2 == 0) goto +2;
		opcode{kind: addi, op1: 1, op2: 1, op3: 1},    // r0 = r0 + 1;
		opcode{kind: jump, op1: -4},                   // goto -4;
		opcode{kind: exit},
	}
	// Program Counter.
	var pc int64 = 0
	// The VM itself.
loop:
	for {
		var op opcode = code[pc]
		switch op.kind {
		case loadi:
			memory[op.op1] = op.op2
		case addi:
			memory[op.op1] = memory[op.op2] + op.op3
		case compare:
			if memory[op.op2] == memory[op.op3] {
				memory[op.op1] = 1
			} else {
				memory[op.op1] = 0
			}
		case jump:
			pc += op.op1
		case branch:
			if memory[op.op1] != 0 {
				pc += op.op1
			}
		case exit:
			break loop
		}
		pc++
	}
	fmt.Printf("Result: %d\n", memory[1])
}
