package memory

import (
	"errors"
	ALU "github.com/coderick14/ARMed/ALU"
	"regexp"
	"strconv"
	"strings"
)

type InstructionMemory struct {
	PC           int64
	Instructions []string
}

var InstructionMem = InstructionMemory{
	PC:           0,
	Instructions: []string{},
}

var dataMemory = DataMemory{
	Memory: make([]int32, 4096),
}

const INCREMENT int64 = 1

/*
 * Method to update program counter
 */

func (instructionMemory *InstructionMemory) updatePC(offset ...int64) {
	if len(offset) == 0 {
		instructionMemory.PC += INCREMENT
	} else {
		instructionMemory.PC += offset[0]
	}
}

/*
 * Method to check if program counter is valid (is program over or not)
 */

func (instructionMemory *InstructionMemory) isValidPC() bool {
	isValidPC := instructionMemory.PC >= 0 && instructionMemory.PC < int64(len(instructionMemory.Instructions))
	return isValidPC
}

/*
 * Function : validateAndExecuteInstruction
 * Details  : checks instruction type, performs syntax analysis, parses the statement and executes it
 */

func (instructionMemory *InstructionMemory) ValidateAndExecuteInstruction() error {

	//get next instruction to be executed from instruction memory
	currentInstruction := instructionMemory.Instructions[instructionMemory.PC]

	var err error

	if strings.HasPrefix(currentInstruction, "ADD ") {

		currentInstructionObject := AddInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "SUB ") {

		currentInstructionObject := SubInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "ADDI ") {

		currentInstructionObject := AddImmediateInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "SUBI ") {

		currentInstructionObject := SubImmediateInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "ADDS ") {

		currentInstructionObject := AddAndSetFlagsInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "SUBS ") {

		currentInstructionObject := SubAndSetFlagsInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "ADDIS ") {

		currentInstructionObject := AddImmediateAndSetFlagsInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "SUBIS ") {

		currentInstructionObject := SubImmediateAndSetFlagsInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "LDUR ") {

		currentInstructionObject := LoadInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "STUR ") {

		currentInstructionObject := StoreInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "LDURH ") {

		currentInstructionObject := LoadHalfInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "STURH ") {

		currentInstructionObject := StoreHalfInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "LDURB ") {

		currentInstructionObject := LoadByteInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "STURB ") {

		currentInstructionObject := StoreByteInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "LDXR ") {

		currentInstructionObject := LoadExclusiveInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "STXR ") {

		currentInstructionObject := StoreExclusiveInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "MOVZ ") {

		currentInstructionObject := MoveWithZeroInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "MOVK ") {

		currentInstructionObject := MoveWithKeepInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "AND ") {

		currentInstructionObject := AndInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "ORR ") {

		currentInstructionObject := OrInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "EOR ") {

		currentInstructionObject := ExclusiveOrInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "ANDI ") {

		currentInstructionObject := AndImmediateInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "ORRI ") {

		currentInstructionObject := OrImmediateInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "EORI ") {

		currentInstructionObject := ExclusiveOrImmediateInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "LSL ") {

		currentInstructionObject := LeftShiftInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "LSR ") {

		currentInstructionObject := RightShiftInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "CBZ ") {

		currentInstructionObject := BranchOnZeroInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "CBNZ ") {

		currentInstructionObject := BranchOnNonZeroInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "B.") {

		currentInstructionObject := ConditionalBranchInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "B ") {

		currentInstructionObject := BranchInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "BR ") {

		currentInstructionObject := BranchToRegisterInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "BL ") {

		currentInstructionObject := BranchWithLinkInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else {

		err = errors.New("Invalid instruction type on line number " + string(instructionMemory.PC+1))

	}

	return err
}

/*
 * All instructions implement the Instruction interface
 */

type Instruction interface {
	checkSyntax() bool
	parse()
	execute()
}

func executeInstruction(currentInstruction Instruction) error {
	isSyntaxOK := currentInstruction.checkSyntax()
	if isSyntaxOK {
		currentInstruction.parse()
		currentInstruction.execute()
	} else {
		return errors.New("Syntax error occured")
	}
	return nil
}

/*
 * INSTRUCTION : ADDITION
 * Example : ADD X1, X2, X3
 * Meaning : X1 = X2 + X3
 */

type AddInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *AddInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^ADD X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *AddInstruction) parse() {
	statement := instruction.inst
	var registers [3]int
	var i, indexX, indexComma int
	for i = 0; i < 3; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if indexComma == -1 {
			indexComma = len(statement)
		}
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])
}

func (instruction *AddInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : SUBTRACTION
 * Example : SUB X1, X2, X3
 * Meaning : X1 = X2 - X3
 */

type SubInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *SubInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^SUB X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *SubInstruction) parse() {
	statement := instruction.inst
	var registers [3]int
	var i, indexX, indexComma int
	for i = 0; i < 3; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if indexComma == -1 {
			indexComma = len(statement)
		}
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

}

func (instruction *SubInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), -getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : ADD IMMEDIATE
 * Example : ADDI X1, X2, 40
 * Meaning : X1 = X2 + 40
 */

type AddImmediateInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *AddImmediateInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^ADDI X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *AddImmediateInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	constant, _ := strconv.Atoi(statement[indexHash+1:])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.constant = uint(constant)
}

func (instruction *AddImmediateInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : SUB IMMEDIATE
 * Example : SUBI X1, X2, 40
 * Meaning : X1 = X2 - 40
 */

type SubImmediateInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *SubImmediateInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^SUBI X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *SubImmediateInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	constant, _ := strconv.Atoi(statement[indexHash+1:])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.constant = uint(constant)

}

func (instruction *SubImmediateInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), -int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : ADD AND SET FLAGS
 * Example : ADDS X1, X2, X3
 * Meaning : X1 = X2 + X3
 * Comments : Adds and sets condition codes
 */

type AddAndSetFlagsInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *AddAndSetFlagsInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^ADDS X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *AddAndSetFlagsInstruction) parse() {
	statement := instruction.inst
	var registers [3]int
	var i, indexX, indexComma int
	for i = 0; i < 3; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if indexComma == -1 {
			indexComma = len(statement)
		}
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

}

func (instruction *AddAndSetFlagsInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)

	//set flag N
	if result < 0 {
		flagNegative = true
	} else {
		flagNegative = false
	}

	//set flag Z
	if result == 0 {
		flagZero = true
	} else {
		flagZero = false
	}

	var hasOverflowOccured bool

	//set flag V (signed addition overflow)
	hasOverflowOccured = (getRegisterValue(instruction.reg2) > 0 && getRegisterValue(instruction.reg3) > 0 && result >= int64(1<<31)) || (getRegisterValue(instruction.reg2) < 0 && getRegisterValue(instruction.reg3) < 0 && result < -int64(1<<31))
	if hasOverflowOccured {
		flagOverflow = true
	} else {
		flagOverflow = false
	}

	//set flag C (unsigned addition overflow)
	unsignedSum := ALU.UnsignedAdder(uint64(getRegisterValue(instruction.reg2)), uint64(getRegisterValue(instruction.reg3)))
	if unsignedSum >= uint64(1<<32) {
		flagCarry = true
	} else {
		flagCarry = false
	}

	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : SUB AND SET FLAGS
 * Example : SUBS X1, X2, X3
 * Meaning : X1 = X2 - X3
 * Comments : Subtracts and sets condition codes
 */

type SubAndSetFlagsInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *SubAndSetFlagsInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^SUBS X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *SubAndSetFlagsInstruction) parse() {
	statement := instruction.inst
	var registers [3]int
	var i, indexX, indexComma int
	for i = 0; i < 3; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if indexComma == -1 {
			indexComma = len(statement)
		}
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

}

func (instruction *SubAndSetFlagsInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)

	//set flag N
	if result < 0 {
		flagNegative = true
	} else {
		flagNegative = false
	}

	//set flag Z
	if result == 0 {
		flagZero = true
	} else {
		flagZero = false
	}

	var hasOverflowOccured bool

	//set flag V (signed addition overflow)
	hasOverflowOccured = (getRegisterValue(instruction.reg2) > 0 && getRegisterValue(instruction.reg3) < 0 && result >= int64(1<<31)) || (getRegisterValue(instruction.reg2) < 0 && getRegisterValue(instruction.reg3) > 0 && result < -int64(1<<31))
	if hasOverflowOccured {
		flagOverflow = true
	} else {
		flagOverflow = false
	}

	//set flag C (unsigned addition overflow)
	if uint64(getRegisterValue(instruction.reg2)) < uint64(getRegisterValue(instruction.reg3)) {
		flagCarry = true
	} else {
		flagCarry = false
	}

	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : ADD IMMEDIATE AND SET FLAGS
 * Example : ADDIS X1, X2, 40
 * Meaning : X1 = X2 + 40
 * Comments : Adds constant and sets condition codes
 */

type AddImmediateAndSetFlagsInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *AddImmediateAndSetFlagsInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^ADDIS X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *AddImmediateAndSetFlagsInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	constant, _ := strconv.Atoi(statement[indexHash+1:])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.constant = uint(constant)

}

func (instruction *AddImmediateAndSetFlagsInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)

	//set flag N
	if result < 0 {
		flagNegative = true
	} else {
		flagNegative = false
	}

	//set flag Z
	if result == 0 {
		flagZero = true
	} else {
		flagZero = false
	}

	var hasOverflowOccured bool

	//set flag V (signed addition overflow)
	hasOverflowOccured = (getRegisterValue(instruction.reg2) > 0 && result >= int64(1<<31))
	if hasOverflowOccured {
		flagOverflow = true
	} else {
		flagOverflow = false
	}

	//set flag C (unsigned addition overflow)
	unsignedSum := ALU.UnsignedAdder(uint64(getRegisterValue(instruction.reg2)), uint64(instruction.constant))
	if unsignedSum >= uint64(1<<32) {
		flagCarry = true
	} else {
		flagCarry = false
	}

	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : SUB IMMEDIATE AND SET FLAGS
 * Example : SUBIS X1, X2, 40
 * Meaning : X1 = X2 - 40
 * Comments : Subtracts constant and sets condition codes
 */

type SubImmediateAndSetFlagsInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *SubImmediateAndSetFlagsInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^SUBIS X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *SubImmediateAndSetFlagsInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	constant, _ := strconv.Atoi(statement[indexHash+1:])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.constant = uint(constant)

}

func (instruction *SubImmediateAndSetFlagsInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), -int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)

	//set flag N
	if result < 0 {
		flagNegative = true
	} else {
		flagNegative = false
	}

	//set flag Z
	if result == 0 {
		flagZero = true
	} else {
		flagZero = false
	}

	var hasOverflowOccured bool

	//set flag V (signed addition overflow)
	hasOverflowOccured = (getRegisterValue(instruction.reg2) < 0 && result < -int64(1<<31))
	if hasOverflowOccured {
		flagOverflow = true
	} else {
		flagOverflow = false
	}

	//set flag C (unsigned addition overflow)
	if uint64(getRegisterValue(instruction.reg2)) < uint64(instruction.constant) {
		flagCarry = true
	} else {
		flagCarry = false
	}

	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : LOAD
 * Example : LDUR X1, [X2, #40]
 * Meaning : X1 = Memory[X2 + 40]
 * Comments : Word from memory to register
 */

type LoadInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *LoadInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^LDUR X([0-9]|1[0-9]|2[0-7]), \\[X([0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *LoadInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash, indexBracket, offset int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	indexBracket = strings.Index(statement, "]")
	offset, _ = strconv.Atoi(statement[indexHash+1 : indexBracket])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)

}

func (instruction *LoadInstruction) execute() {
	memoryIndex := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.offset)) / 4
	memoryValue := dataMemory.read(uint64(memoryIndex))
	setRegisterValue(instruction.reg1, int64(memoryValue))
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : STORE
 * Example : STUR X1, [X2, #40]
 * Meaning : Memory[X2 + 40] = X1
 * Comments : Word from register to memory
 */

type StoreInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *StoreInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^STUR X([0-9]|1[0-9]|2[0-7]), \\[X([0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *StoreInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash, indexBracket, offset int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	indexBracket = strings.Index(statement, "]")
	offset, _ = strconv.Atoi(statement[indexHash+1 : indexBracket])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)

}

func (instruction *StoreInstruction) execute() {
	memoryIndex := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.offset)) / 4
	registerValue := getRegisterValue(instruction.reg1)
	dataMemory.write(uint64(memoryIndex), int32(registerValue))
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : LOAD HALFWORD
 * Example : LDURH X1, [X2, #40]
 * Meaning : X1 = Memory[X2 + 40]
 * Comments : Halfword from memory to register
 */

type LoadHalfInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *LoadHalfInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^LDURH X([0-9]|1[0-9]|2[0-7]), \\[X([0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *LoadHalfInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash, indexBracket, offset int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	indexBracket = strings.Index(statement, "]")
	offset, _ = strconv.Atoi(statement[indexHash+1 : indexBracket])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)

}

func (instruction *LoadHalfInstruction) execute() {
	var memoryValue int16
	var shift uint = 16
	memoryIndex := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.offset))
	if memoryIndex%4 == 0 {
		// extract upper 16 bits
		memoryValue = int16(dataMemory.read(uint64(memoryIndex/4)) >> shift)
	} else {
		// extract lower 16 bits
		memoryValue = int16(dataMemory.read(uint64(memoryIndex / 4)))
	}
	setRegisterValue(instruction.reg1, int64(memoryValue))
}

/*
 * INSTRUCTION : STORE HALFWORD
 * Example : STURH X1, [X2, #40]
 * Meaning : Memory[X2 + 40] = X1
 * Comments : Halfword from register to memory
 */

type StoreHalfInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *StoreHalfInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^STURH X([0-9]|1[0-9]|2[0-7]), \\[X([0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *StoreHalfInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash, indexBracket, offset int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	indexBracket = strings.Index(statement, "]")
	offset, _ = strconv.Atoi(statement[indexHash+1 : indexBracket])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)

}

func (instruction *StoreHalfInstruction) execute() {
	var registerValue int16
	var shift uint = 16
	registerValue = int16(getRegisterValue(instruction.reg1))
	memoryIndex := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.offset))
	currentMemoryValue := dataMemory.read(uint64(memoryIndex / 4))
	if memoryIndex%4 == 0 {
		// store in upper 16 bits
		currentMemoryValue = currentMemoryValue & ((1 << shift) - 1) // clear upper 16 bits
		currentMemoryValue = currentMemoryValue | (int32(registerValue) << shift)
	} else {
		// store in lower 16 bits
		currentMemoryValue = currentMemoryValue & -(1 << shift) // clear lower 16 bits
		currentMemoryValue = currentMemoryValue | int32(registerValue)
	}
	dataMemory.write(uint64(memoryIndex/4), currentMemoryValue)
}

/*
 * INSTRUCTION : LOAD BYTE
 * Example : LDURB X1, [X2, #40]
 * Meaning : X1 = Memory[X2 + 40]
 * Comments : Byte from memory to register
 */

type LoadByteInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *LoadByteInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^LDURB X([0-9]|1[0-9]|2[0-7]), \\[X([0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *LoadByteInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash, indexBracket, offset int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	indexBracket = strings.Index(statement, "]")
	offset, _ = strconv.Atoi(statement[indexHash+1 : indexBracket])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)

}

func (instruction *LoadByteInstruction) execute() {
	var registerValue int8
	memoryIndex := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.offset))
	memoryValue := dataMemory.read(uint64(memoryIndex / 4))
	if memoryIndex%4 == 0 {
		// extract bits[31:24]
		registerValue = int8(memoryValue >> 24)
	} else if memoryIndex%4 == 1 {
		// extract bits[23:16]
		registerValue = int8(memoryValue >> 16)
	} else if memoryIndex%4 == 2 {
		// extract bits[15:8]
		registerValue = int8(memoryValue >> 8)
	} else {
		// extract bit[7:0]
		registerValue = int8(memoryValue)
	}
	setRegisterValue(instruction.reg1, int64(registerValue))
}

/*
 * INSTRUCTION : STORE BYTE
 * Example : STURB X1, [X2, #40]
 * Meaning : Memory[X2 + 40] = X1
 * Comments : Byte from register to memory
 */

type StoreByteInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *StoreByteInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^STURB X([0-9]|1[0-9]|2[0-7]), \\[X([0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *StoreByteInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash, indexBracket, offset int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	indexBracket = strings.Index(statement, "]")
	offset, _ = strconv.Atoi(statement[indexHash+1 : indexBracket])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)

}

func (instruction *StoreByteInstruction) execute() {
	var registerValue int8
	registerValue = int8(getRegisterValue(instruction.reg1))
	memoryIndex := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.offset))
	currentMemoryValue := dataMemory.read(uint64(memoryIndex / 4))
	if memoryIndex%4 == 0 {

		// store in bits[31:24]
		currentMemoryValue = currentMemoryValue & ((1 << 24) - 1) // clear bits[31:24]
		currentMemoryValue = currentMemoryValue | (int32(registerValue) << 24)

	} else if memoryIndex%4 == 1 {

		// store in bits[23:16]
		currentMemoryValue = currentMemoryValue & (((1 << 16) - 1) | -(1 << 24)) // clear bits[23:16]
		currentMemoryValue = currentMemoryValue | int32(registerValue)

	} else if memoryIndex%4 == 2 {

		// store in bits[15:8]
		currentMemoryValue = currentMemoryValue & (((1 << 8) - 1) | -(1 << 16)) // clear bits[15:8]
		currentMemoryValue = currentMemoryValue | int32(registerValue)

	} else {

		// store in bits[7:0]
		currentMemoryValue = currentMemoryValue & -(1 << 8) // clear bits[7:0]
		currentMemoryValue = currentMemoryValue | int32(registerValue)

	}
	dataMemory.write(uint64(memoryIndex/4), currentMemoryValue)
}

/*
 * INSTRUCTION : LOAD EXCLUSIVE REGISTER
 * Example : LDXR X1, [X2, #0]
 * Meaning : X1 = Memory[X2]
 * Comments : Load; first half of atomic swap
 */

type LoadExclusiveInstruction struct {
	inst string
	reg1 uint
	reg2 uint
}

func (instruction *LoadExclusiveInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^LDXR X([0-9]|1[0-9]|2[0-7]), \\[X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *LoadExclusiveInstruction) parse() {

}

func (instruction *LoadExclusiveInstruction) execute() {

}

/*
 * INSTRUCTION : STORE EXCLUSIVE REGISTER
 * Example : STXR X1, X3, [X2, #0]
 * Meaning : Memory[X2] = X1; X3 = 0 or 1
 * Comments : Store; second half of atomic swap
 */

type StoreExclusiveInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *StoreExclusiveInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^STXR X([0-9]|1[0-9]|2[0-7]), X([0-9]|1[0-9]|2[0-7]), \\[X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *StoreExclusiveInstruction) parse() {

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
	inst     string
	reg1     uint
	constant uint16
	offset   uint
}

func (instruction *MoveWithZeroInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^MOVZ X([0-9]|1[0-9]|2[0-7]), (0|[1-9][0-9]*), LSL (0|1|2|3)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *MoveWithZeroInstruction) parse() {
	statement := instruction.inst
	var indexX, indexComma int

	indexX = strings.Index(statement, "X")
	indexComma = strings.Index(statement, ",")
	register, _ := strconv.Atoi(statement[indexX+1 : indexComma])

	statement = strings.TrimSpace(statement[indexComma+1:])
	indexComma = strings.Index(statement, ",")
	constant, _ := strconv.Atoi(statement[:indexComma])

	offset := uint(statement[len(statement)-1] - '0')

	instruction.reg1 = uint(register)
	instruction.constant = uint16(constant)
	instruction.offset = offset
}

func (instruction *MoveWithZeroInstruction) execute() {
	value := int64(instruction.constant)
	offset := uint(16 * instruction.offset)
	value = value << offset
	setRegisterValue(instruction.reg1, value)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : MOVE WITH KEEP
 * Example : MOVK X1, 20, LSL 0
 * Meaning : X1 = 20 or 20*(2^16) or 20*(2^32) or 20*(2^48)
 * Comments : Loads 16-bit constant, rest unchanged
 */

type MoveWithKeepInstruction struct {
	inst     string
	reg1     uint
	constant uint16
	offset   uint
}

func (instruction *MoveWithKeepInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^MOVK X([0-9]|1[0-9]|2[0-7]), (0|[1-9][0-9]*), LSL (0|1|2|3)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *MoveWithKeepInstruction) parse() {
	statement := instruction.inst
	var indexX, indexComma int

	indexX = strings.Index(statement, "X")
	indexComma = strings.Index(statement, ",")
	register, _ := strconv.Atoi(statement[indexX+1 : indexComma])

	statement = strings.TrimSpace(statement[indexComma+1:])
	indexComma = strings.Index(statement, ",")
	constant, _ := strconv.Atoi(statement[:indexComma])

	offset := uint(statement[len(statement)-1] - '0')

	instruction.reg1 = uint(register)
	instruction.constant = uint16(constant)
	instruction.offset = offset

}

func (instruction *MoveWithKeepInstruction) execute() {
	value := int64(instruction.constant)
	offset := uint(16 * instruction.offset)
	value = value << offset
	registerValue := getRegisterValue(instruction.reg1)

	var lastBitIndex uint
	lastBitIndex = offset + 15

	for i := offset; i <= lastBitIndex; i++ {
		if value&(1<<i) != 0 {
			registerValue = registerValue | (1 << i)
		} else {
			registerValue = registerValue &^ (1 << i)
		}
	}

	setRegisterValue(instruction.reg1, registerValue)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : LOGICAL AND
 * Example : AND X1, X2, X3
 * Meaning : X1 = X2 & X3
 * Comments : Bitwise-And of X2 and X3, stores result in X1
 */

type AndInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *AndInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^AND X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *AndInstruction) parse() {
	statement := instruction.inst
	var registers [3]int
	var i, indexX, indexComma int
	for i = 0; i < 3; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if indexComma == -1 {
			indexComma = len(statement)
		}
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

}

func (instruction *AndInstruction) execute() {
	result := ALU.LogicalAND(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : LOGICAL OR
 * Example : ORR X1, X2, X3
 * Meaning : X1 = X2 | X3
 * Comments : Bitwise-Or of X2 and X3, stores result in X1
 */

type OrInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *OrInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^ORR X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *OrInstruction) parse() {
	statement := instruction.inst
	var registers [3]int
	var i, indexX, indexComma int
	for i = 0; i < 3; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if indexComma == -1 {
			indexComma = len(statement)
		}
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

}

func (instruction *OrInstruction) execute() {
	result := ALU.LogicalOR(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : LOGICAL EXCLUSIVE-OR
 * Example : EOR X1, X2, X3
 * Meaning : X1 = X2 ^ X3
 * Comments : Bitwise-Xor of X2 and X3, stores result in X1
 */

type ExclusiveOrInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *ExclusiveOrInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^EOR X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *ExclusiveOrInstruction) parse() {
	statement := instruction.inst
	var registers [3]int
	var i, indexX, indexComma int
	for i = 0; i < 3; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if indexComma == -1 {
			indexComma = len(statement)
		}
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

}

func (instruction *ExclusiveOrInstruction) execute() {
	result := ALU.LogicalXOR(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : LOGICAL AND IMMEDIATE
 * Example : ANDI X1, X2, #20
 * Meaning : X1 = X2 & 20
 * Comments : Bitwise-And of X2 with a constant, stores result in X1
 */

type AndImmediateInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *AndImmediateInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^ANDI X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *AndImmediateInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	constant, _ := strconv.Atoi(statement[indexHash+1:])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.constant = uint(constant)

}

func (instruction *AndImmediateInstruction) execute() {
	result := ALU.LogicalAND(getRegisterValue(instruction.reg2), int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : LOGICAL OR IMMEDIATE
 * Example : ORRI X1, X2, #20
 * Meaning : X1 = X2 | 20
 * Comments : Bitwise-Or of X2 with a constant, stores result in X1
 */

type OrImmediateInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *OrImmediateInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^ORRI X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *OrImmediateInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	constant, _ := strconv.Atoi(statement[indexHash+1:])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.constant = uint(constant)

}

func (instruction *OrImmediateInstruction) execute() {
	result := ALU.LogicalOR(getRegisterValue(instruction.reg2), int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : LOGICAL EXCLUSIVE-OR IMMEDIATE
 * Example : EORI X1, X2, #20
 * Meaning : X1 = X2 ^ 20
 * Comments : Bitwise-Xor of X2 with a constant, stores result in X1
 */

type ExclusiveOrImmediateInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *ExclusiveOrImmediateInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^EORI X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *ExclusiveOrImmediateInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma, indexHash int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	constant, _ := strconv.Atoi(statement[indexHash+1:])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.constant = uint(constant)

}

func (instruction *ExclusiveOrImmediateInstruction) execute() {
	result := ALU.LogicalXOR(getRegisterValue(instruction.reg2), int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : LOGICAL LEFT SHIFT
 * Example : LSL X1, X2, 10
 * Meaning : X1 = X2 << 10
 * Comments : Left shifts X2 by a constant, stores result in X1
 */

type LeftShiftInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *LeftShiftInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^LSL X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), (0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *LeftShiftInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma int

	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	statement = strings.TrimSpace(statement[indexComma+1:])
	offset, _ := strconv.Atoi(statement)

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)
}

func (instruction *LeftShiftInstruction) execute() {
	result := getRegisterValue(instruction.reg2) << instruction.offset
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : LOGICAL RIGHT SHIFT
 * Example : LSR X1, X2, 10
 * Meaning : X1 = X2 >> 10
 * Comments : Right shifts X2 by a constant, stores result in X1
 */

type RightShiftInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *RightShiftInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^LSR X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), (0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *RightShiftInstruction) parse() {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexComma int

	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = 31
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		statement = statement[indexComma+1:]
	}
	statement = strings.TrimSpace(statement[indexComma+1:])
	offset, _ := strconv.Atoi(statement)

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)
}

func (instruction *RightShiftInstruction) execute() {
	result := getRegisterValue(instruction.reg2) >> instruction.offset
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
 * INSTRUCTION : COMPARE AND BRANCH ON EQUAL 0
 * Example : CBZ X1, 25
 * Meaning : if (X1 == 0) go to PC+100
 * Comments : Equal 0 test; PC-relative branch
 */

type BranchOnZeroInstruction struct {
	inst   string
	reg1   uint
	offset int64
}

func (instruction *BranchOnZeroInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^CBZ X([0-9]|1[0-9]|2[0-7]), ([1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *BranchOnZeroInstruction) parse() {

}

func (instruction *BranchOnZeroInstruction) execute() {
	if getRegisterValue(instruction.reg1) == 0 {
		InstructionMem.updatePC(instruction.offset)
	} else {
		InstructionMem.updatePC()
	}
}

/*
 * INSTRUCTION : COMPARE AND BRANCH ON NOT EQUAL 0
 * Example : CBNZ X1, 25
 * Meaning : if (X1 != 0) go to PC+100
 * Comments : NotEqual 0 test; PC-relative branch
 */

type BranchOnNonZeroInstruction struct {
	inst   string
	reg1   uint
	offset int64
}

func (instruction *BranchOnNonZeroInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^CBNZ X([0-9]|1[0-9]|2[0-7]), ([1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *BranchOnNonZeroInstruction) parse() {

}

func (instruction *BranchOnNonZeroInstruction) execute() {
	if getRegisterValue(instruction.reg1) != 0 {
		InstructionMem.updatePC(instruction.offset)
	} else {
		InstructionMem.updatePC()
	}
}

/*
 * INSTRUCTION : CONDITIONAL BRANCH
 * Example : B.cond 25
 * Meaning : if (condition true) go to PC+100
 * Comments : Test condition codes; if true, then branch
 */

type ConditionalBranchInstruction struct {
	inst      string
	offset    int64
	condition string
}

func (instruction *ConditionalBranchInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^B\\.(EQ|NE|LT|LE|GT|GE|LO|LS|HI|HS) ([1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *ConditionalBranchInstruction) parse() {

}

func (instruction *ConditionalBranchInstruction) execute() {
	is_branching := false
	switch instruction.condition {

	case "EQ":
		is_branching = flagZero
	case "NE":
		is_branching = !flagZero
	case "LT":
		is_branching = (flagNegative != flagOverflow)
	case "LE":
		is_branching = !(flagZero == false && flagNegative == flagOverflow)
	case "GT":
		is_branching = (flagZero == false && flagNegative == flagOverflow)
	case "GE":
		is_branching = (flagNegative == flagOverflow)
	case "LO":
		is_branching = !flagCarry
	case "LS":
		is_branching = !(flagZero == false && flagCarry == true)
	case "HI":
		is_branching = (flagZero == false && flagCarry == true)
	case "HS":
		is_branching = flagCarry

	}

	if is_branching {
		InstructionMem.updatePC(instruction.offset)
	} else {
		InstructionMem.updatePC()
	}
}

/*
 * INSTRUCTION : UNCONDITIONAL BRANCH
 * Example : B 25
 * Meaning : go to PC+100
 * Comments : Branch to PC-relative target address
 */

type BranchInstruction struct {
	inst   string
	offset int64
}

func (instruction *BranchInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^B ([1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *BranchInstruction) parse() {

}

func (instruction *BranchInstruction) execute() {
	InstructionMem.updatePC(instruction.offset)
}

/*
 * INSTRUCTION : UNCONDITIONAL BRANCH TO REGISTER
 * Example : BR X30
 * Meaning : go to X30
 * Comments : Branch to address stored in register. Used for switch, procedure return
 */

type BranchToRegisterInstruction struct {
	inst string
	reg1 uint
}

func (instruction *BranchToRegisterInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^BR X([0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *BranchToRegisterInstruction) parse() {

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
	inst   string
	offset int64
}

func (instruction *BranchWithLinkInstruction) checkSyntax() bool {
	r, _ := regexp.Compile("^BL ([1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return false
	}
	return true
}

func (instruction *BranchWithLinkInstruction) parse() {

}

func (instruction *BranchWithLinkInstruction) execute() {
	setRegisterValue(30, InstructionMem.PC+INCREMENT)
	InstructionMem.updatePC(instruction.offset)
}
