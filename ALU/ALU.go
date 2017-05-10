package ALU

//Function to perform signed addition of two int64 numbers.
func Adder(val1, val2 int64) int64 {
	return val1 + val2
}


//Function to perform logical AND operation of two int64 numbers.
func LogicalAND(val1, val2 int64) int64 {
	return val1 & val2
}


//Function to perform logical OR operation of two int64 numbers.
func LogicalOR(val1, val2 int64) int64 {
	return val1 | val2
}


//Function to perform logical XOR operation of two int64 numbers.
func LogicalXOR(val1, val2 int64) int64 {
	return val1 ^ val2
}


//Function to perform unsigned addition of two uint64 numbers.
//Used only for setting Carry flag.
func UnsignedAdder(val1, val2 uint64) uint64 {
	return val1 + val2
}
