package memory

import "sync"

type DataMemory struct {
	sync.RWMutex
	Memory []int32
}

var registers [32]int64

var flagNegative, flagZero, flagOverflow, flagCarry bool

/*
 * method to initiate register values
 */

 func InitRegisters() {
 	registers[XZR] = 0
 	registers[SP] = MEMORY_SIZE * 4 
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
