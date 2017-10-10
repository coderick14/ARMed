package memory

import (
	"errors"
	ALU "github.com/coderick14/ARMed/ALU"
	"regexp"
	"strconv"
	"strings"
)

// Struct to represent instruction memory
type InstructionMemory struct {
	PC           int64
	Instructions []string
	Labels       map[string]int64
}

// Instance of instruction memory
var InstructionMem = InstructionMemory{
	PC:           0,
	Instructions: []string{},
	Labels:       make(map[string]int64),
}

// Instance of data memory
var dataMemory = DataMemory{
	Memory: make([]int32, MEMORY_SIZE),
}

// Method to update program counter.
func (instructionMemory *InstructionMemory) updatePC(offset ...int64) {
	if len(offset) == 0 {
		instructionMemory.PC += INCREMENT
	} else {
		instructionMemory.PC += offset[0]
	}
}

// IsValidPC is a function to check if program counter is valid.
func IsValidPC(PC int64) bool {
	isValidPC := PC >= 0 && PC < int64(len(InstructionMem.Instructions))
	return isValidPC
}

// isEmptyInstruction is a method to check for null instructions (NoOps)
func isEmptyInstruction(currentInstruction string) bool {
	return len(currentInstruction) == 0
}

// ExtractLabels is a method to extract labels from instructions.
func (instructionMemory *InstructionMemory) ExtractLabels() {

	labelRegex, _ := regexp.Compile("^([a-zA-Z][[:alnum:]]*)[[:space:]]*:")
	for counter, currentInstruction := range instructionMemory.Instructions {
		if labelRegex.MatchString(currentInstruction) {

			indexColon := strings.Index(currentInstruction, ":")
			labelName := strings.TrimSpace(currentInstruction[:indexColon])
			currentInstruction = strings.TrimSpace(currentInstruction[indexColon+1:])
			instructionMemory.Labels[labelName] = int64(counter)
			instructionMemory.Instructions[counter] = currentInstruction

		}
	}
}

// ValidateAndExecuteInstruction is a method to check instruction type, perform syntax analysis, parse the statement and execute it
func (instructionMemory *InstructionMemory) ValidateAndExecuteInstruction() error {

	//get next instruction to be executed from instruction memory
	currentInstruction := instructionMemory.Instructions[instructionMemory.PC]

	if isEmptyInstruction(currentInstruction) {
		instructionMemory.updatePC()
		return nil
	}

	var err error

	if strings.HasPrefix(currentInstruction, "ADD ") {

		currentInstructionObject := AddInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "SUB ") {

		currentInstructionObject := SubInstruction{inst: currentInstruction}
		err = executeInstruction(&currentInstructionObject)

	} else if strings.HasPrefix(currentInstruction, "MUL ") {

		currentInstructionObject := MulInstruction{inst: currentInstruction}
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

		err = errors.New("Invalid instruction type in " + currentInstruction)

	}

	return err
}

// All instructions implement the Instruction interface
type Instruction interface {
	// Checks syntax of current instruction and returns an error
	checkSyntax() error

	// Parses the current instruction and extracts register numbers, constants etc.
	parse() error

	// emulates the execution of the current instruction
	execute()
}

// Function that takes an interface as argument and executes the corresponding instruction
func executeInstruction(currentInstruction Instruction) error {
	syntaxError := currentInstruction.checkSyntax()
	if syntaxError != nil {
		return syntaxError
	} else {
		parseError := currentInstruction.parse()
		if parseError != nil {
			return parseError
		}
		currentInstruction.execute()
	}
	return nil
}

/*
INSTRUCTION : ADDITION

	Example : ADD X1, X2, X3
	Meaning : X1 = X2 + X3
*/
type AddInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *AddInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^ADD X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *AddInstruction) parse() error {
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
			registers[i] = XZR
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		if indexComma < len(statement) {
			statement = statement[indexComma+1:]
		}
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

	return nil
}

func (instruction *AddInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : SUBTRACTION

	Example : SUB X1, X2, X3
	Meaning : X1 = X2 - X3
*/
type SubInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *SubInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^SUB X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *SubInstruction) parse() error {
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
			registers[i] = XZR
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		if indexComma < len(statement) {
			statement = statement[indexComma+1:]
		}
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

	return nil
}

func (instruction *SubInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), -getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : MULTIPLICATION

	Example : MUL X1, X2, X3
	Meaning : X1 = X2 * X3
*/
type MulInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *MulInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^MUL X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *MulInstruction) parse() error {
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
			registers[i] = XZR
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		if indexComma < len(statement) {
			statement = statement[indexComma+1:]
		}
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

	return nil
}

func (instruction *MulInstruction) execute() {
	result := ALU.Multiplier(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : ADD IMMEDIATE

	Example : ADDI X1, X2, 40
	Meaning : X1 = X2 + 40
*/
type AddImmediateInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *AddImmediateInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^ADDI ((X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]))|(SP, SP)), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *AddImmediateInstruction) parse() error {
	statement := instruction.inst

	// if instruction updates stack pointer
	if strings.Index(statement, "SP") != -1 {
		indexHash := strings.Index(statement, "#")
		constant, _ := strconv.Atoi(statement[indexHash+1:])
		instruction.reg1 = SP
		instruction.reg2 = SP
		instruction.constant = uint(constant)

		address := getRegisterValue(instruction.reg2) + int64(instruction.constant)
		if address > MEMORY_SIZE*WORD_SIZE {
			return errors.New("Stack underflow error in : " + instruction.inst)
		}

		return nil
	}

	var registers [2]int
	var i, indexX, indexComma, indexHash int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = XZR
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

	return nil
}

func (instruction *AddImmediateInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : SUB IMMEDIATE

	Example : SUBI X1, X2, 40
	Meaning : X1 = X2 - 40
*/
type SubImmediateInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *SubImmediateInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^SUBI ((X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]))|(SP, SP)), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *SubImmediateInstruction) parse() error {
	statement := instruction.inst

	// if instruction updates stack pointer
	if strings.Index(statement, "SP") != -1 {
		indexHash := strings.Index(statement, "#")
		constant, _ := strconv.Atoi(statement[indexHash+1:])
		instruction.reg1 = SP
		instruction.reg2 = SP
		instruction.constant = uint(constant)

		address := getRegisterValue(instruction.reg2) + int64(instruction.constant)
		if address < (MEMORY_SIZE-STACK_SIZE)*WORD_SIZE {
			return errors.New("Stack overflow error in : " + instruction.inst)
		}

		return nil
	}

	var registers [2]int
	var i, indexX, indexComma, indexHash int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexComma = strings.Index(statement, ",")
		if statement[indexX+1:indexComma] == "ZR" {
			registers[i] = XZR
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

	return nil
}

func (instruction *SubImmediateInstruction) execute() {
	result := ALU.Adder(getRegisterValue(instruction.reg2), -int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : ADD AND SET FLAGS

	Example : ADDS X1, X2, X3
	Meaning : X1 = X2 + X3

Comments : Adds and sets condition codes
*/
type AddAndSetFlagsInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *AddAndSetFlagsInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^ADDS X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *AddAndSetFlagsInstruction) parse() error {
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
			registers[i] = XZR
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		if indexComma < len(statement) {
			statement = statement[indexComma+1:]
		}
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

	return nil
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
INSTRUCTION : SUB AND SET FLAGS

	Example : SUBS X1, X2, X3
	Meaning : X1 = X2 - X3

Comments : Subtracts and sets condition codes
*/
type SubAndSetFlagsInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *SubAndSetFlagsInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^SUBS X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *SubAndSetFlagsInstruction) parse() error {
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
			registers[i] = XZR
		} else {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		}
		if indexComma < len(statement) {
			statement = statement[indexComma+1:]
		}
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

	return nil
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
INSTRUCTION : ADD IMMEDIATE AND SET FLAGS

	Example : ADDIS X1, X2, 40
	Meaning : X1 = X2 + 40

Comments : Adds constant and sets condition codes
*/
type AddImmediateAndSetFlagsInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *AddImmediateAndSetFlagsInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^ADDIS X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *AddImmediateAndSetFlagsInstruction) parse() error {
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

	return nil
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
INSTRUCTION : SUB IMMEDIATE AND SET FLAGS

	Example : SUBIS X1, X2, 40
	Meaning : X1 = X2 - 40

Comments : Subtracts constant and sets condition codes
*/
type SubImmediateAndSetFlagsInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *SubImmediateAndSetFlagsInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^SUBIS X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *SubImmediateAndSetFlagsInstruction) parse() error {
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

	return nil
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
INSTRUCTION : LOAD

	Example : LDUR X1, [X2, #40]
	Meaning : X1 = Memory[X2 + 40]

Comments : Word from memory to register
*/
type LoadInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *LoadInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^LDUR (X([0-9]|1[0-9]|2[0-7])|LR), \\[(X([0-9]|1[0-9]|2[0-7])|SP), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *LoadInstruction) parse() error {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexLR, indexComma, indexHash, indexBracket, offset int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexLR = strings.Index(statement, "LR")
		indexComma = strings.Index(statement, ",")
		if indexX != -1 {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		} else if indexLR != -1 {
			registers[i] = LR
		} else {
			registers[i] = SP
		}
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	indexBracket = strings.Index(statement, "]")
	offset, _ = strconv.Atoi(statement[indexHash+1 : indexBracket])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)

	//check for alignment restriction
	if (getRegisterValue(instruction.reg2)+int64(instruction.offset))%4 != 0 {
		return errors.New("Alignment restriction violation in : " + instruction.inst)
	}

	return nil
}

func (instruction *LoadInstruction) execute() {
	memoryIndex := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.offset)) / 4
	memoryValue := dataMemory.read(uint64(memoryIndex))
	setRegisterValue(instruction.reg1, int64(memoryValue))
	InstructionMem.updatePC()
}

/*
INSTRUCTION : STORE

	Example : STUR X1, [X2, #40]
	Meaning : Memory[X2 + 40] = X1

Comments : Word from register to memory
*/
type StoreInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *StoreInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^STUR (X([0-9]|1[0-9]|2[0-7])|LR), \\[(X([0-9]|1[0-9]|2[0-7])|SP), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *StoreInstruction) parse() error {
	statement := instruction.inst
	var registers [2]int
	var i, indexX, indexLR, indexComma, indexHash, indexBracket, offset int
	for i = 0; i < 2; i++ {
		indexX = strings.Index(statement, "X")
		indexLR = strings.Index(statement, "LR")
		indexComma = strings.Index(statement, ",")
		if indexX != -1 {
			registers[i], _ = strconv.Atoi(statement[indexX+1 : indexComma])
		} else if indexLR != -1 {
			registers[i] = LR
		} else {
			registers[i] = SP
		}
		statement = statement[indexComma+1:]
	}
	indexHash = strings.Index(statement, "#")
	indexBracket = strings.Index(statement, "]")
	offset, _ = strconv.Atoi(statement[indexHash+1 : indexBracket])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)

	//check for alignment restriction
	if (getRegisterValue(instruction.reg2)+int64(instruction.offset))%4 != 0 {
		return errors.New("Alignment restriction violation in : " + instruction.inst)
	}

	return nil
}

func (instruction *StoreInstruction) execute() {
	memoryIndex := ALU.Adder(getRegisterValue(instruction.reg2), int64(instruction.offset)) / 4
	registerValue := getRegisterValue(instruction.reg1)
	dataMemory.write(uint64(memoryIndex), int32(registerValue))
	InstructionMem.updatePC()
}

/*
INSTRUCTION : LOAD HALFWORD

	Example : LDURH X1, [X2, #40]
	Meaning : X1 = Memory[X2 + 40]

Comments : Halfword from memory to register
*/
type LoadHalfInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *LoadHalfInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^LDURH X([0-9]|1[0-9]|2[0-7]), \\[X([0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *LoadHalfInstruction) parse() error {
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

	//check for alignment restriction
	if (getRegisterValue(instruction.reg2)+int64(instruction.offset))%4 != 0 && (getRegisterValue(instruction.reg2)+int64(instruction.offset))%4 != 2 {
		return errors.New("Alignment restriction violation in : " + instruction.inst)
	}

	return nil
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

	InstructionMem.updatePC()
}

/*
INSTRUCTION : STORE HALFWORD

	Example : STURH X1, [X2, #40]
	Meaning : Memory[X2 + 40] = X1

Comments : Halfword from register to memory
*/
type StoreHalfInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *StoreHalfInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^STURH X([0-9]|1[0-9]|2[0-7]), \\[X([0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *StoreHalfInstruction) parse() error {
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

	//check for alignment restriction
	if (getRegisterValue(instruction.reg2)+int64(instruction.offset))%4 != 0 && (getRegisterValue(instruction.reg2)+int64(instruction.offset))%4 != 2 {
		return errors.New("Alignment restriction violation in : " + instruction.inst)
	}

	return nil
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

	InstructionMem.updatePC()
}

/*
INSTRUCTION : LOAD BYTE

	Example : LDURB X1, [X2, #40]
	Meaning : X1 = Memory[X2 + 40]

Comments : Byte from memory to register
*/
type LoadByteInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *LoadByteInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^LDURB X([0-9]|1[0-9]|2[0-7]), \\[X([0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *LoadByteInstruction) parse() error {
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

	return nil
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

	InstructionMem.updatePC()
}

/*
INSTRUCTION : STORE BYTE

	Example : STURB X1, [X2, #40]
	Meaning : Memory[X2 + 40] = X1

Comments : Byte from register to memory
*/
type StoreByteInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *StoreByteInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^STURB X([0-9]|1[0-9]|2[0-7]), \\[X([0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *StoreByteInstruction) parse() error {
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

	return nil
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
		currentMemoryValue = currentMemoryValue | (int32(registerValue) << 16)

	} else if memoryIndex%4 == 2 {

		// store in bits[15:8]
		currentMemoryValue = currentMemoryValue & (((1 << 8) - 1) | -(1 << 16)) // clear bits[15:8]
		currentMemoryValue = currentMemoryValue | (int32(registerValue) << 8)

	} else {

		// store in bits[7:0]
		currentMemoryValue = currentMemoryValue & -(1 << 8) // clear bits[7:0]
		currentMemoryValue = currentMemoryValue | int32(registerValue)

	}
	dataMemory.write(uint64(memoryIndex/4), currentMemoryValue)

	InstructionMem.updatePC()
}

/*
INSTRUCTION : LOAD EXCLUSIVE REGISTER

	Example : LDXR X1, [X2, #0]
	Meaning : X1 = Memory[X2]

Comments : Load; first half of atomic swap
*/
// type LoadExclusiveInstruction struct {
// 	inst string
// 	reg1 uint
// 	reg2 uint
// }

// func (instruction *LoadExclusiveInstruction) checkSyntax() error {
// 	r, _ := regexp.Compile("^LDXR X([0-9]|1[0-9]|2[0-7]), \\[X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
// 	if r.MatchString(instruction.inst) == false {
// 		return errors.New("Syntax error occurred in " + instruction.inst)
// 	}
// 	return nil
// }

// func (instruction *LoadExclusiveInstruction) parse() error {

// 	return nil
// }

// func (instruction *LoadExclusiveInstruction) execute() {

// }

/*
INSTRUCTION : STORE EXCLUSIVE REGISTER

	Example : STXR X1, X3, [X2, #0]
	Meaning : Memory[X2] = X1; X3 = 0 or 1

Comments : Store; second half of atomic swap
*/
// type StoreExclusiveInstruction struct {
// 	inst string
// 	reg1 uint
// 	reg2 uint
// 	reg3 uint
// }

// func (instruction *StoreExclusiveInstruction) checkSyntax() error {
// 	r, _ := regexp.Compile("^STXR X([0-9]|1[0-9]|2[0-7]), X([0-9]|1[0-9]|2[0-7]), \\[X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)\\]$")
// 	if r.MatchString(instruction.inst) == false {
// 		return errors.New("Syntax error occurred in " + instruction.inst)
// 	}
// 	return nil
// }

// func (instruction *StoreExclusiveInstruction) parse() error {

// 	return nil
// }

// func (instruction *StoreExclusiveInstruction) execute() {

// }

/*
INSTRUCTION : MOVE WITH ZERO

	Example : MOVZ X1, 20, LSL 0
	Meaning : X1 = 20 or 20*(2^16) or 20*(2^32) or 20*(2^48)

Comments : Loads 16-bit constant, rest zeroes
*/
type MoveWithZeroInstruction struct {
	inst     string
	reg1     uint
	constant uint16
	offset   uint
}

func (instruction *MoveWithZeroInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^MOVZ X([0-9]|1[0-9]|2[0-7]), (0|[1-9][0-9]*), LSL (0|1|2|3)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *MoveWithZeroInstruction) parse() error {
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

	return nil
}

func (instruction *MoveWithZeroInstruction) execute() {
	value := int64(instruction.constant)
	offset := uint(16 * instruction.offset)
	value = value << offset
	setRegisterValue(instruction.reg1, value)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : MOVE WITH KEEP

	Example : MOVK X1, 20, LSL 0
	Meaning : X1 = 20 or 20*(2^16) or 20*(2^32) or 20*(2^48)

Comments : Loads 16-bit constant, rest unchanged
*/
type MoveWithKeepInstruction struct {
	inst     string
	reg1     uint
	constant uint16
	offset   uint
}

func (instruction *MoveWithKeepInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^MOVK X([0-9]|1[0-9]|2[0-7]), (0|[1-9][0-9]*), LSL (0|1|2|3)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *MoveWithKeepInstruction) parse() error {
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

	return nil
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
INSTRUCTION : LOGICAL AND

	Example : AND X1, X2, X3
	Meaning : X1 = X2 & X3

Comments : Bitwise-And of X2 and X3, stores result in X1
*/
type AndInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *AndInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^AND X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *AndInstruction) parse() error {
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
		if indexComma < len(statement) {
			statement = statement[indexComma+1:]
		}
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

	return nil
}

func (instruction *AndInstruction) execute() {
	result := ALU.LogicalAND(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : LOGICAL OR

	Example : ORR X1, X2, X3
	Meaning : X1 = X2 | X3

Comments : Bitwise-Or of X2 and X3, stores result in X1
*/
type OrInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *OrInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^ORR X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *OrInstruction) parse() error {
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
		if indexComma < len(statement) {
			statement = statement[indexComma+1:]
		}
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

	return nil
}

func (instruction *OrInstruction) execute() {
	result := ALU.LogicalOR(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : LOGICAL EXCLUSIVE-OR

	Example : EOR X1, X2, X3
	Meaning : X1 = X2 ^ X3

Comments : Bitwise-Xor of X2 and X3, stores result in X1
*/
type ExclusiveOrInstruction struct {
	inst string
	reg1 uint
	reg2 uint
	reg3 uint
}

func (instruction *ExclusiveOrInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^EOR X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7])$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *ExclusiveOrInstruction) parse() error {
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
		if indexComma < len(statement) {
			statement = statement[indexComma+1:]
		}
	}
	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.reg3 = uint(registers[2])

	return nil
}

func (instruction *ExclusiveOrInstruction) execute() {
	result := ALU.LogicalXOR(getRegisterValue(instruction.reg2), getRegisterValue(instruction.reg3))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : LOGICAL AND IMMEDIATE

	Example : ANDI X1, X2, #20
	Meaning : X1 = X2 & 20

Comments : Bitwise-And of X2 with a constant, stores result in X1
*/
type AndImmediateInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *AndImmediateInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^ANDI X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *AndImmediateInstruction) parse() error {
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

	return nil
}

func (instruction *AndImmediateInstruction) execute() {
	result := ALU.LogicalAND(getRegisterValue(instruction.reg2), int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : LOGICAL OR IMMEDIATE

	Example : ORRI X1, X2, #20
	Meaning : X1 = X2 | 20

Comments : Bitwise-Or of X2 with a constant, stores result in X1
*/
type OrImmediateInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *OrImmediateInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^ORRI X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *OrImmediateInstruction) parse() error {
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

	return nil
}

func (instruction *OrImmediateInstruction) execute() {
	result := ALU.LogicalOR(getRegisterValue(instruction.reg2), int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : LOGICAL EXCLUSIVE-OR IMMEDIATE

	Example : EORI X1, X2, #20
	Meaning : X1 = X2 ^ 20

Comments : Bitwise-Xor of X2 with a constant, stores result in X1
*/
type ExclusiveOrImmediateInstruction struct {
	inst     string
	reg1     uint
	reg2     uint
	constant uint
}

func (instruction *ExclusiveOrImmediateInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^EORI X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), #(0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *ExclusiveOrImmediateInstruction) parse() error {
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

	return nil
}

func (instruction *ExclusiveOrImmediateInstruction) execute() {
	result := ALU.LogicalXOR(getRegisterValue(instruction.reg2), int64(instruction.constant))
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : LOGICAL LEFT SHIFT

	Example : LSL X1, X2, #10
	Meaning : X1 = X2 << 10

Comments : Left shifts X2 by a constant, stores result in X1
*/
type LeftShiftInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *LeftShiftInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^LSL X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), (0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *LeftShiftInstruction) parse() error {
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
	offset, _ := strconv.Atoi(statement[indexHash+1:])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)

	return nil
}

func (instruction *LeftShiftInstruction) execute() {
	result := getRegisterValue(instruction.reg2) << instruction.offset
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : LOGICAL RIGHT SHIFT

	Example : LSR X1, X2, #10
	Meaning : X1 = X2 >> 10

Comments : Right shifts X2 by a constant, stores result in X1
*/
type RightShiftInstruction struct {
	inst   string
	reg1   uint
	reg2   uint
	offset uint
}

func (instruction *RightShiftInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^LSR X([0-9]|1[0-9]|2[0-7]), X(ZR|[0-9]|1[0-9]|2[0-7]), (0|[1-9][0-9]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *RightShiftInstruction) parse() error {
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
	offset, _ := strconv.Atoi(statement[indexHash+1:])

	instruction.reg1 = uint(registers[0])
	instruction.reg2 = uint(registers[1])
	instruction.offset = uint(offset)

	return nil
}

func (instruction *RightShiftInstruction) execute() {
	result := getRegisterValue(instruction.reg2) >> instruction.offset
	setRegisterValue(instruction.reg1, result)
	InstructionMem.updatePC()
}

/*
INSTRUCTION : COMPARE AND BRANCH ON EQUAL 0

	Example : CBZ X1, label
	Meaning : if (X1 == 0) go to label

Comments : Equal 0 test; PC-relative branch
*/
type BranchOnZeroInstruction struct {
	inst   string
	reg1   uint
	offset int64
}

func (instruction *BranchOnZeroInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^CBZ X([0-9]|1[0-9]|2[0-7]), ([a-zA-Z][[:alnum:]]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *BranchOnZeroInstruction) parse() error {
	statement := instruction.inst
	var indexX, indexComma int

	indexX = strings.Index(statement, "X")
	indexComma = strings.Index(statement, ",")

	register, _ := strconv.Atoi(statement[indexX+1 : indexComma])
	labelName := strings.TrimSpace(statement[indexComma+1:])
	labelPC, isValidLabel := InstructionMem.Labels[labelName]

	if !isValidLabel {
		return errors.New("Invalid label name " + labelName + " in " + instruction.inst)
	}

	instruction.reg1 = uint(register)
	instruction.offset = labelPC - InstructionMem.PC

	return nil
}

func (instruction *BranchOnZeroInstruction) execute() {
	if getRegisterValue(instruction.reg1) == 0 {
		InstructionMem.updatePC(instruction.offset)
	} else {
		InstructionMem.updatePC()
	}
}

/*
INSTRUCTION : COMPARE AND BRANCH ON NOT EQUAL 0

	Example : CBNZ X1, label
	Meaning : if (X1 != 0) go to label

Comments : NotEqual 0 test; PC-relative branch
*/
type BranchOnNonZeroInstruction struct {
	inst   string
	reg1   uint
	offset int64
}

func (instruction *BranchOnNonZeroInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^CBNZ X([0-9]|1[0-9]|2[0-7]), ([a-zA-Z][[:alnum:]]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *BranchOnNonZeroInstruction) parse() error {
	statement := instruction.inst
	var indexX, indexComma int

	indexX = strings.Index(statement, "X")
	indexComma = strings.Index(statement, ",")

	register, _ := strconv.Atoi(statement[indexX+1 : indexComma])
	labelName := strings.TrimSpace(statement[indexComma+1:])
	labelPC, isValidLabel := InstructionMem.Labels[labelName]

	if !isValidLabel {
		return errors.New("Invalid label name " + labelName + " in " + instruction.inst)
	}

	instruction.reg1 = uint(register)
	instruction.offset = labelPC - InstructionMem.PC

	return nil
}

func (instruction *BranchOnNonZeroInstruction) execute() {
	if getRegisterValue(instruction.reg1) != 0 {
		InstructionMem.updatePC(instruction.offset)
	} else {
		InstructionMem.updatePC()
	}
}

/*
INSTRUCTION : CONDITIONAL BRANCH

	Example : B.cond label
	Meaning : if (condition true) go to label

Comments : Test condition codes; if true, then branch
*/
type ConditionalBranchInstruction struct {
	inst      string
	offset    int64
	condition string
}

func (instruction *ConditionalBranchInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^B\\.(EQ|NE|LT|LE|GT|GE|LO|LS|HI|HS) ([a-zA-Z][[:alnum:]]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *ConditionalBranchInstruction) parse() error {
	statement := instruction.inst

	conditionCode := statement[2:4]
	labelName := strings.TrimSpace(statement[5:])
	labelPC, isValidLabel := InstructionMem.Labels[labelName]

	if !isValidLabel {
		return errors.New("Invalid label name " + labelName + " in " + instruction.inst)
	}

	instruction.condition = conditionCode
	instruction.offset = labelPC - InstructionMem.PC

	return nil
}

func (instruction *ConditionalBranchInstruction) execute() {
	is_branching := false

	switch instruction.condition {

	case "EQ":
		is_branching = flagZero
		break
	case "NE":
		is_branching = !flagZero
		break
	case "LT":
		is_branching = (flagNegative != flagOverflow)
		break
	case "LE":
		is_branching = !(flagZero == false && flagNegative == flagOverflow)
		break
	case "GT":
		is_branching = (flagZero == false && flagNegative == flagOverflow)
		break
	case "GE":
		is_branching = (flagNegative == flagOverflow)
		break
	case "LO":
		is_branching = !flagCarry
		break
	case "LS":
		is_branching = !(flagZero == false && flagCarry == true)
		break
	case "HI":
		is_branching = (flagZero == false && flagCarry == true)
		break
	case "HS":
		is_branching = flagCarry
		break

	}

	if is_branching {
		InstructionMem.updatePC(instruction.offset)
	} else {
		InstructionMem.updatePC()
	}
}

/*
INSTRUCTION : UNCONDITIONAL BRANCH

	Example : B label
	Meaning : go to label

Comments : Branch to PC-relative target address
*/
type BranchInstruction struct {
	inst   string
	offset int64
}

func (instruction *BranchInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^B ([a-zA-Z][[:alnum:]]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *BranchInstruction) parse() error {
	statement := instruction.inst

	labelName := strings.TrimSpace(statement[2:])
	labelPC, isValidLabel := InstructionMem.Labels[labelName]

	if !isValidLabel {
		return errors.New("Invalid label name " + labelName + " in " + instruction.inst)
	}

	instruction.offset = labelPC - InstructionMem.PC

	return nil
}

func (instruction *BranchInstruction) execute() {
	InstructionMem.updatePC(instruction.offset)
}

/*
INSTRUCTION : UNCONDITIONAL BRANCH TO REGISTER

	Example : BR LR
	Meaning : go to address stored in LR

Comments : Branch to address stored in register. Used for switch, procedure return
*/
type BranchToRegisterInstruction struct {
	inst   string
	reg1   uint
	offset int64
}

func (instruction *BranchToRegisterInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^BR (X([0-9]|1[0-9]|2[0-7])|LR)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *BranchToRegisterInstruction) parse() error {
	statement := instruction.inst
	registerName := strings.TrimSpace(statement[3:])
	var register uint

	if strings.Compare(registerName, "LR") == 0 {
		register = 30
	} else {
		registerValue, _ := strconv.Atoi(statement[4:])
		register = uint(registerValue)
	}

	if !IsValidPC(getRegisterValue(register)) {
		return errors.New("Invalid address in register " + registerName + " in " + instruction.inst)
	}

	instruction.reg1 = register
	instruction.offset = getRegisterValue(register) - InstructionMem.PC

	return nil
}

func (instruction *BranchToRegisterInstruction) execute() {
	InstructionMem.updatePC(instruction.offset)
}

/*
INSTRUCTION : UNCONDITIONAL BRANCH WITH LINK

	Example : BL label
	Meaning : X30 = PC + 4; go to label

Comments : For procedure call (PC-relative)
*/
type BranchWithLinkInstruction struct {
	inst   string
	offset int64
}

func (instruction *BranchWithLinkInstruction) checkSyntax() error {
	r, _ := regexp.Compile("^BL ([a-zA-Z][[:alnum:]]*)$")
	if r.MatchString(instruction.inst) == false {
		return errors.New("Syntax error occurred in " + instruction.inst)
	}
	return nil
}

func (instruction *BranchWithLinkInstruction) parse() error {
	statement := instruction.inst

	labelName := strings.TrimSpace(statement[3:])
	labelPC, isValidLabel := InstructionMem.Labels[labelName]

	if !isValidLabel {
		return errors.New("Invalid label name " + labelName + " in " + instruction.inst)
	}

	instruction.offset = labelPC - InstructionMem.PC

	return nil
}

func (instruction *BranchWithLinkInstruction) execute() {
	setRegisterValue(30, InstructionMem.PC+INCREMENT)
	InstructionMem.updatePC(instruction.offset)
}
