package memory

import ALU "github.com/coderick14/ARMed/ALU"

type InstructionMemory struct {
	PC           uint64
	Instructions []string
}

/*
 * Method to update program counter
 */

func (instructionMemory *InstructionMemory) updatePC() {

}

/*
 * Method to check if program counter is valid (is program over or not)
 */

func (instructionMemory *InstructionMemory) isValidPC() {

}

/*
 * Function : checkInstructionType
 * Details  : checks instruction type and invokes the checkSyntax method of the particular instruction
 */

func (instructionMemory *InstructionMemory) checkInstructionType() {

}

/*
 * Function : executeInstruction
 * Details  : checks instruction type, performs syntax analysis, parses the statement and executes it
 */

func (instructionMemory *InstructionMemory) executeInstruction() {

}

/*
 * INSTRUCTION : ADDITION
 * Example : ADD X1, X2, X3
 * Meaning : X1 = X2 + X3
 */

type AddInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *AddInstruction) checkSyntax() {

}

func (instruction *AddInstruction) execute() {

}

/*
 * INSTRUCTION : SUBTRACTION
 * Example : SUB X1, X2, X3
 * Meaning : X1 = X2 - X3
 */

type SubInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *SubInstruction) checkSyntax() {

}

func (instruction *SubInstruction) execute() {

}

/*
 * INSTRUCTION : ADD IMMEDIATE
 * Example : ADDI X1, X2, 40
 * Meaning : X1 = X2 + 40
 */

type AddImmediateInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *AddImmediateInstruction) checkSyntax() {

}

func (instruction *AddImmediateInstruction) execute() {

}

/*
 * INSTRUCTION : SUB IMMEDIATE
 * Example : ADDI X1, X2, 40
 * Meaning : X1 = X2 - 40
 */

type SubImmediateInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *SubImmediateInstruction) checkSyntax() {

}

func (instruction *SubImmediateInstruction) execute() {

}

/*
 * INSTRUCTION : ADD AND SET FLAGS
 * Example : ADDS X1, X2, X3
 * Meaning : X1 = X2 + X3
 * Comments : Adds and sets condition codes
 */

type AddAndSetFlagsInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *AddAndSetFlagsInstruction) checkSyntax() {

}

func (instruction *AddAndSetFlagsInstruction) execute() {

}

/*
 * INSTRUCTION : SUB AND SET FLAGS
 * Example : ADDS X1, X2, X3
 * Meaning : X1 = X2 - X3
 * Comments : Subtracts and sets condition codes
 */

type SubAndSetFlagsInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *SubAndSetFlagsInstruction) checkSyntax() {

}

func (instruction *SubAndSetFlagsInstruction) execute() {

}

/*
 * INSTRUCTION : ADD IMMEDIATE AND SET FLAGS
 * Example : ADDIS X1, X2, 40
 * Meaning : X1 = X2 + 40
 * Comments : Adds constant and sets condition codes
 */

type AddImmediateAndSetFlagsInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *AddImmediateAndSetFlagsInstruction) checkSyntax() {

}

func (instruction *AddImmediateAndSetFlagsInstruction) execute() {

}

/*
 * INSTRUCTION : SUB IMMEDIATE AND SET FLAGS
 * Example : SUBIS X1, X2, 40
 * Meaning : X1 = X2 - 40
 * Comments : Subtracts constant and sets condition codes
 */

type SubImmediateAndSetFlagsInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *SubImmediateAndSetFlagsInstruction) checkSyntax() {

}

func (instruction *SubImmediateAndSetFlagsInstruction) execute() {

}

/*
 * INSTRUCTION : LOAD
 * Example : LDUR X1, [X2, 40]
 * Meaning : X1 = Memory[X2 + 40]
 * Comments : Doubleword from memory to register
 */

type LoadInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *LoadInstruction) checkSyntax() {

}

func (instruction *LoadInstruction) execute() {

}

/*
 * INSTRUCTION : STORE
 * Example : STUR X1, [X2, 40]
 * Meaning : Memory[X2 + 40] = X1
 * Comments : Doubleword from register to memory
 */

type StoreInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *StoreInstruction) checkSyntax() {

}

func (instruction *StoreInstruction) execute() {

}

/*
 * INSTRUCTION : LOAD WORD
 * Example : LDURSW X1, [X2, 40]
 * Meaning : X1 = Memory[X2 + 40]
 * Comments : Word from memory to register
 */

type LoadWordInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *LoadWordInstruction) checkSyntax() {

}

func (instruction *LoadWordInstruction) execute() {

}

/*
 * INSTRUCTION : STORE WORD
 * Example : STURW X1, [X2, 40]
 * Meaning : Memory[X2 + 40] = X1
 * Comments : Word from register to memory
 */

type StoreWordInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *StoreWordInstruction) checkSyntax() {

}

func (instruction *StoreWordInstruction) execute() {

}

/*
 * INSTRUCTION : LOAD HALFWORD
 * Example : LDURH X1, [X2, 40]
 * Meaning : X1 = Memory[X2 + 40]
 * Comments : Halfword from memory to register
 */

type LoadHalfInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *LoadHalfInstruction) checkSyntax() {

}

func (instruction *LoadHalfInstruction) execute() {

}

/*
 * INSTRUCTION : STORE HALFWORD
 * Example : STURH X1, [X2, 40]
 * Meaning : Memory[X2 + 40] = X1
 * Comments : Halfword from register to memory
 */

type StoreHalfInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *StoreHalfInstruction) checkSyntax() {

}

func (instruction *StoreHalfInstruction) execute() {

}

/*
 * INSTRUCTION : LOAD BYTE
 * Example : LDURH X1, [X2, 40]
 * Meaning : X1 = Memory[X2 + 40]
 * Comments : Byte from memory to register
 */

type LoadByteInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *LoadByteInstruction) checkSyntax() {

}

func (instruction *LoadByteInstruction) execute() {

}

/*
 * INSTRUCTION : STORE BYTE
 * Example : STURH X1, [X2, 40]
 * Meaning : Memory[X2 + 40] = X1
 * Comments : Byte from register to memory
 */

type StoreByteInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *StoreByteInstruction) checkSyntax() {

}

func (instruction *StoreByteInstruction) execute() {

}

/*
 * INSTRUCTION : LOAD EXCLUSIVE REGISTER
 * Example : LDXR X1, [X2, 0]
 * Meaning : X1 = Memory[X2]
 * Comments : Load; first half of atomic swap
 */

type LoadExclusiveInstruction struct {
	reg1 int64
	reg2 int64
}

func (instruction *LoadExclusiveInstruction) checkSyntax() {

}

func (instruction *LoadExclusiveInstruction) execute() {

}

/*
 * INSTRUCTION : STORE EXCLUSIVE REGISTER
 * Example : STXR X1, X3, [X2]
 * Meaning : Memory[X2] = X1; X3 = 0 or 1
 * Comments : Store; second half of atomic swap
 */

type StoreExclusiveInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *StoreExclusiveInstruction) checkSyntax() {

}

func (instruction *StoreExclusiveInstruction) execute() {

}

/*
 * INSTRUCTION : MOVE WITH ZERO
 * Example : MOVZ X1, 20, LSL 0
 * Meaning : X1 = 20 or 20*(2^16) or 20*(2^32) or 20*(2^48)
 * Comments : Loads 16-bit constant, rest zeroes
 */

type MoveWithZeroInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *MoveWithZeroInstruction) checkSyntax() {

}

func (instruction *MoveWithZeroInstruction) execute() {

}

/*
 * INSTRUCTION : MOVE WITH KEEP
 * Example : MOVK X1, 20, LSL 0
 * Meaning : X1 = 20 or 20*(2^16) or 20*(2^32) or 20*(2^48)
 * Comments : Loads 16-bit constant, rest unchanged
 */

type MoveWithKeepInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *MoveWithKeepInstruction) checkSyntax() {

}

func (instruction *MoveWithKeepInstruction) execute() {

}

/*
 * INSTRUCTION : LOGICAL AND
 * Example : AND X1, X2, X3
 * Meaning : X1 = X2 & X3
 * Comments : Bitwise-And of X2 and X3, stores result in X1
 */

type AndInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *AndInstruction) checkSyntax() {

}

func (instruction *AndInstruction) execute() {

}

/*
 * INSTRUCTION : LOGICAL OR
 * Example : ORR X1, X2, X3
 * Meaning : X1 = X2 | X3
 * Comments : Bitwise-Or of X2 and X3, stores result in X1
 */

type OrInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *OrInstruction) checkSyntax() {

}

func (instruction *OrInstruction) execute() {

}

/*
 * INSTRUCTION : LOGICAL EXCLUSIVE-OR
 * Example : EOR X1, X2, X3
 * Meaning : X1 = X2 ^ X3
 * Comments : Bitwise-Xor of X2 and X3, stores result in X1
 */

type ExclusiveOrInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *ExclusiveOrInstruction) checkSyntax() {

}

func (instruction *ExclusiveOrInstruction) execute() {

}

/*
 * INSTRUCTION : LOGICAL AND IMMEDIATE
 * Example : ANDI X1, X2, 20
 * Meaning : X1 = X2 & 20
 * Comments : Bitwise-And of X2 with a constant, stores result in X1
 */

type AndImmediateInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *AndImmediateInstruction) checkSyntax() {

}

func (instruction *AndImmediateInstruction) execute() {

}

/*
 * INSTRUCTION : LOGICAL OR IMMEDIATE
 * Example : ORRI X1, X2, 20
 * Meaning : X1 = X2 | 20
 * Comments : Bitwise-Or of X2 with a constant, stores result in X1
 */

type OrImmediateInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *OrImmediateInstruction) checkSyntax() {

}

func (instruction *OrImmediateInstruction) execute() {

}

/*
 * INSTRUCTION : LOGICAL EXCLUSIVE-OR IMMEDIATE
 * Example : ERRI X1, X2, 20
 * Meaning : X1 = X2 ^ 20
 * Comments : Bitwise-Xor of X2 with a constant, stores result in X1
 */

type ExclusiveOrImmediateInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *ExclusiveOrImmediateInstruction) checkSyntax() {

}

func (instruction *ExclusiveOrImmediateInstruction) execute() {

}

/*
 * INSTRUCTION : LOGICAL LEFT SHIFT
 * Example : LSL X1, X2, 10
 * Meaning : X1 = X2 << 10
 * Comments : Left shifts X2 by a constant, stores result in X1
 */

type LeftShiftInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *LeftShiftInstruction) checkSyntax() {

}

func (instruction *LeftShiftInstruction) execute() {

}

/*
 * INSTRUCTION : LOGICAL RIGHT SHIFT
 * Example : LSR X1, X2, 10
 * Meaning : X1 = X2 >> 10
 * Comments : Right shifts X2 by a constant, stores result in X1
 */

type RightShiftInstruction struct {
	reg1 int64
	reg2 int64
	reg3 int64
}

func (instruction *RightShiftInstruction) checkSyntax() {

}

func (instruction *RightShiftInstruction) execute() {

}

/*
 * INSTRUCTION : COMPARE AND BRANCH ON EQUAL 0
 * Example : CBZ X1, 25
 * Meaning : if (X1 == 0) go to PC+100
 * Comments : Equal 0 test; PC-relative branch
 */

type BranchOnZeroInstruction struct {
	reg1 int64
	reg2 int64
}

func (instruction *BranchOnZeroInstruction) checkSyntax() {

}

func (instruction *BranchOnZeroInstruction) execute() {

}

/*
 * INSTRUCTION : COMPARE AND BRANCH ON NOT EQUAL 0
 * Example : CBNZ X1, 25
 * Meaning : if (X1 != 0) go to PC+100
 * Comments : NotEqual 0 test; PC-relative branch
 */

type BranchOnNonZeroInstruction struct {
	reg1 int64
	reg2 int64
}

func (instruction *BranchOnNonZeroInstruction) checkSyntax() {

}

func (instruction *BranchOnNonZeroInstruction) execute() {

}

/*
 * INSTRUCTION : CONDITIONAL BRANCH
 * Example : B.cond 25
 * Meaning : if (condition true) go to PC+100
 * Comments : Test condition codes; if true, then branch
 */

type ConditionalBranchInstruction struct {
	reg1 int64
	reg2 int64
}

func (instruction *ConditionalBranchInstruction) checkSyntax() {

}

func (instruction *ConditionalBranchInstruction) execute() {

}

/*
 * INSTRUCTION : UNCONDITIONAL BRANCH
 * Example : B 25
 * Meaning : go to PC+100
 * Comments : Branch to PC-relative target address
 */

type BranchInstruction struct {
	reg1 int64
}

func (instruction *BranchInstruction) checkSyntax() {

}

func (instruction *BranchInstruction) execute() {

}

/*
 * INSTRUCTION : UNCONDITIONAL BRANCH TO REGISTER
 * Example : BR X30
 * Meaning : go to X30
 * Comments : Branch to address stored in register. Used for switch, procedure return
 */

type BranchToRegisterInstruction struct {
	reg1 int64
}

func (instruction *BranchToRegisterInstruction) checkSyntax() {

}

func (instruction *BranchToRegisterInstruction) execute() {

}

/*
 * INSTRUCTION : UNCONDITIONAL BRANCH WITH LINK
 * Example : BL 2500
 * Meaning : X30 = PC + 4; go to PC + 10000
 * Comments : For procedure call (PC-relative)
 */

type BranchWithLinkInstruction struct {
	reg1 int64
}

func (instruction *BranchWithLinkInstruction) checkSyntax() {

}

func (instruction *BranchWithLinkInstruction) execute() {

}
