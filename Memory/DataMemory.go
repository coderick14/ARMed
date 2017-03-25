package memory

import "sync"

type DataMemory struct {
	sync.RWMutex
	memory []int64
}

/*
 * Method to read data from memory
 * Guarantees mutually exclusive access
 */

func (dataMemory *DataMemory) read(address uint64) int64 {
	dataMemory.RLock()
	value := dataMemory.memory[address]
	dataMemory.RUnlock()
	return value
}

/*
 * Method to write data to memory
 * Guarantees mutually exclusive access
 */

func (dataMemory *DataMemory) write(address uint64, value int64) {
	dataMemory.Lock()
	dataMemory.memory[address] = value
	dataMemory.Unlock()
}