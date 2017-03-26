package memory

import "sync"

type DataMemory struct {
	sync.RWMutex
	Memory []int64
}

var Registers [32]int64

/*
 * Method to read data from memory
 * Guarantees mutually exclusive access
 */

func (dataMemory *DataMemory) read(address uint64) int64 {
	dataMemory.RLock()
	value := dataMemory.Memory[address]
	dataMemory.RUnlock()
	return value
}

/*
 * Method to write data to memory
 * Guarantees mutually exclusive access
 */

func (dataMemory *DataMemory) write(address uint64, value int64) {
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
