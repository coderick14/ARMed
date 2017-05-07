package memory

import (
	tablewriter "github.com/olekukonko/tablewriter"
	color "github.com/fatih/color"
	"os"
	"fmt"
	"sync"
	"strconv"
)

type DataMemory struct {
	sync.RWMutex
	Memory []int32
}

var registers, buffer [32]int64

var flagNegative, flagZero, flagOverflow, flagCarry bool

/*
 * method to initiate register values
 */

func InitRegisters() {
	registers[XZR] = 0
	registers[SP] = MEMORY_SIZE * 4
}

/*
 * function to store register values in a buffer
 */

func SaveRegisters() {
	var i int
	for i = 0; i < 32; i++ {
		buffer[i] = registers[i]
	}
}

/*
 * function to show register values
 */

func ShowRegisters(showAll bool) {
	var i int
	var registerNum, prevRegisterVal, newRegisterVal string
	table := tablewriter.NewWriter(os.Stdout)
	if showAll == true {
		table.SetHeader([]string{"Register", "Value"})

		for i = 0; i < 32; i++ {
			registerNum = strconv.Itoa(i)
			newRegisterVal = strconv.FormatInt(getRegisterValue(uint(i)), 10)
			if getRegisterValue(uint(i)) != buffer[i] {	
				table.Append([]string{color.CyanString("R"+registerNum), color.CyanString(newRegisterVal)})
			} else {
				table.Append([]string{"R"+registerNum, newRegisterVal})
			}
		}
	} else {
		table.SetHeader([]string{"Register", "Previous Value", "New Value"})

		for i = 0; i < 32; i++ {
			if getRegisterValue(uint(i)) != buffer[i] {
				registerNum = strconv.Itoa(i)
				prevRegisterVal = strconv.FormatInt(buffer[i], 10)
				newRegisterVal = strconv.FormatInt(getRegisterValue(uint(i)), 10)
				table.Append([]string{color.CyanString("R"+registerNum), color.RedString(prevRegisterVal), color.GreenString(newRegisterVal)})
			}
		}
	} 
	

	table.Render()
	fmt.Printf("\n")
}

/*
 * Method to read data from memory
 * Guarantees mutually exclusive access
 */

func (dataMemory *DataMemory) read(address uint64) int32 {
	dataMemory.RLock()
	value := dataMemory.Memory[address]
	dataMemory.RUnlock()
	return value
}

/*
 * Method to write data to memory
 * Guarantees mutually exclusive access
 */

func (dataMemory *DataMemory) write(address uint64, value int32) {
	dataMemory.Lock()
	dataMemory.Memory[address] = value
	dataMemory.Unlock()
}

/*
 * Function to read from register
 */

func getRegisterValue(registerIndex uint) int64 {
	return registers[registerIndex]
}

/*
 * Function to write to register
 */

func setRegisterValue(registerIndex uint, value int64) {
	registers[registerIndex] = value
}
