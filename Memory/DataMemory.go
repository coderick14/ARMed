package memory

import "sync"

type DataMemory struct {
	sync.RWMutex
	Memory []int32
}

var Registers [32]int64

var flagNegative, flagZero, flagOverflow, flagCarry bool

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
	return Registers[registerIndex]
}

/*
 * Function to write to register
 */

func setRegisterValue(registerIndex uint, value int64) {
	Registers[registerIndex] = value
}
