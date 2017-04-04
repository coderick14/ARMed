package ALU

/*
 * Function to perform signed addition
 */
func Adder(val1, val2 int64) int64 {
	return val1 + val2
}

/*
 * Function to perform logical AND operation
 */

func LogicalAND(val1, val2 int64) int64 {
	return val1 & val2
}

/*
 * Function to perform logical OR operation
 */

func LogicalOR(val1, val2 int64) int64 {
	return val1 | val2
}

/*
 * Function to perform logical XOR operation
 */

func LogicalXOR(val1, val2 int64) int64 {
	return val1 ^ val2
}

/*
 * Function to perform unsigned addition
 * Used only for setting Carry flag
 */
func UnsignedAdder(val1, val2 uint64) uint64 {
	return val1 + val2
}
