/*
ARMed version 1.0

Author : https://github.com/coderick14

	ARMed is a very basic emulator of the ARM instruction set written in Golang
	USAGE : ARMed [OPTIONS]... SOURCE_FILE

	Example SOURCE_FILE :

		ADDI X0, XZR, #3;
		BL fact;
		B Exit;
		fact: SUBI SP, SP, #8;
		STUR LR, [SP, #4];
		STUR X0, [SP, #0];
		SUBIS X9, X0, #1;
		B.GE L1;
		ADDI X1, XZR, #1;
		ADDI SP, SP, #8;
		BR LR;
		L1: SUBI X0, X0, #1;
		BL fact;
		LDUR X0, [SP, #0];
		LDUR LR, [SP, #4];
		ADDI SP, SP, #8;
		ADD X1, X0, X1;
		BR LR;
		Exit:;


	--all 		show all register values after an instruction, with updated ones in color
	--end 		show updated registers only once, at the end of the program. Overrides --all
	--no-log 	suppress logs of statements being executed
	--help 		display help

Found a bug? Feel free to raise an issue on https://github.com/coderick14/ARMed

Contributions welcome :)
*/
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	Memory "github.com/coderick14/ARMed/Memory"
	"io"
	"os"
	"strings"
)

const helpString = `ARMed version 1.0
Author : https://github.com/coderick14

ARMed is a very basic emulator of the ARM instruction set written in Golang
USAGE : ARMed [OPTIONS]... SOURCE_FILE

--all 		show all register values after an instruction, with updated ones in color
--end 		show updated registers only once, at the end of the program. Overrides --all
--no-log 	suppress logs of statements being executed
--help 		display help

Found a bug? Feel free to raise an issue on https://github.com/coderick14/ARMed
Contributions welcome :)`

func main() {
	var err error
	helpPtr := flag.Bool("help", false, "Display help")
	allPtr := flag.Bool("all", false, "Display all registers after each instruction")
	endPtr := flag.Bool("end", false, "Display registers only at end")
	logPtr := flag.Bool("no-log", false, "Suppress log messages")

	flag.Parse()

	if *helpPtr == true {
		fmt.Println(helpString)
		return
	}

	if len(flag.Args()) == 0 {
		err = errors.New("Error : Missing filename.\n Type ARMed --help for further help")
		fmt.Println(err)
		return
	}

	fileName := flag.Args()[0]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file : ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString(';')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error while reading file : ", err)
			return
		}
		line = strings.TrimSpace(strings.TrimRight(line, ";"))
		Memory.InstructionMem.Instructions = append(Memory.InstructionMem.Instructions, line)
	}

	Memory.InitRegisters()
	Memory.InstructionMem.ExtractLabels()

	if *endPtr == true {
		Memory.SaveRegisters()
		for Memory.IsValidPC(Memory.InstructionMem.PC) {
			if *logPtr == false {
				fmt.Println("Executing :", Memory.InstructionMem.Instructions[Memory.InstructionMem.PC])
			}
			err = Memory.InstructionMem.ValidateAndExecuteInstruction()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		Memory.ShowRegisters(false)

	} else {

		for Memory.IsValidPC(Memory.InstructionMem.PC) {
			Memory.SaveRegisters()
			if *logPtr == false {
				fmt.Println("Executing :", Memory.InstructionMem.Instructions[Memory.InstructionMem.PC])
			}
			err = Memory.InstructionMem.ValidateAndExecuteInstruction()
			if err != nil {
				fmt.Println(err)
				return
			}
			Memory.ShowRegisters(*allPtr)
		}

	}

}
