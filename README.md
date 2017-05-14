# ARMed
---
#### A basic ARM emulator written in Golang 
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/coderick14/ARMed/issues)
[![Open Source Love](https://badges.frapsoft.com/os/v2/open-source.svg?v=103)](https://github.com/coderick14/ARMed/) 
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)
[![GoDoc](https://godoc.org/github.com/coderick14/ARMed?status.svg)](https://godoc.org/github.com/coderick14/ARMed)

![ARMed logo](https://github.com/coderick14/ARMed/blob/dev/images/logo.png "ARMed - an ARM emulator written in Golang" )

##### Build instructions
Requires Go to be [installed](https://golang.org/doc/install) and [configured](https://golang.org/doc/install#testing)
```
go get github.com/coderick14/ARMed
```
##### Usage
`ARMed --help` will show all the usage details

---

##### Contributions
Found a bug? Or maybe add support for some more instructions? Feel free to open up a pull request or raise an issue!!

---

![ARMed demo GIF](https://github.com/coderick14/ARMed/blob/dev/images/demo.gif "Simple demo of how ARMed works" )

#### Instructions supported in v1.0

```
INSTRUCTION : ADDITION
Example : ADD X1, X2, X3
Meaning : X1 = X2 + X3
```

```
INSTRUCTION : SUBTRACTION
Example : SUB X1, X2, X3
Meaning : X1 = X2 - X3
```

```
INSTRUCTION : ADD IMMEDIATE
Example : ADDI X1, X2, #40
Meaning : X1 = X2 + 40
```

```
INSTRUCTION : SUB IMMEDIATE
Example : SUBI X1, X2, #40
Meaning : X1 = X2 - 40
```

```
INSTRUCTION : ADD AND SET FLAGS
Example : ADDS X1, X2, X3
Meaning : X1 = X2 + X3
Comments : Adds and sets condition codes
```

```
INSTRUCTION : SUB AND SET FLAGS
Example : SUBS X1, X2, X3
Meaning : X1 = X2 - X3
Comments : Subtracts and sets condition codes
```

```
INSTRUCTION : ADD IMMEDIATE AND SET FLAGS
Example : ADDIS X1, X2, #40
Meaning : X1 = X2 + 40
Comments : Adds constant and sets condition codes
```

```
INSTRUCTION : SUB IMMEDIATE AND SET FLAGS
Example : SUBIS X1, X2, #40
Meaning : X1 = X2 - 40
Comments : Subtracts constant and sets condition codes
```

```
INSTRUCTION : LOAD
Example : LDUR X1, [X2, #40]
Meaning : X1 = Memory[X2 + 40]
Comments : Word from memory to register
```

```
INSTRUCTION : STORE
Example : STUR X1, [X2, #40]
Meaning : Memory[X2 + 40] = X1
Comments : Word from register to memory
```

```
INSTRUCTION : LOAD HALFWORD
Example : LDURH X1, [X2, #40]
Meaning : X1 = Memory[X2 + 40]
Comments : Halfword from memory to register
```

```
INSTRUCTION : STORE HALFWORD
Example : STURH X1, [X2, #40]
Meaning : Memory[X2 + 40] = X1
Comments : Halfword from register to memory
```

```
INSTRUCTION : LOAD BYTE
Example : LDURB X1, [X2, #40]
Meaning : X1 = Memory[X2 + 40]
Comments : Byte from memory to register
```

```
INSTRUCTION : STORE BYTE
Example : STURB X1, [X2, #40]
Meaning : Memory[X2 + 40] = X1
Comments : Byte from register to memory
```

```
INSTRUCTION : MOVE WITH ZERO
Example : MOVZ X1, 20, LSL 0
Meaning : X1 = 20 or 20*(2^16) or 20*(2^32) or 20*(2^48)
Comments : Loads 16-bit constant, rest zeroes
```

```
INSTRUCTION : MOVE WITH KEEP
Example : MOVK X1, 20, LSL 0
Meaning : X1 = 20 or 20*(2^16) or 20*(2^32) or 20*(2^48)
Comments : Loads 16-bit constant, rest unchanged
```

```
INSTRUCTION : LOGICAL AND
Example : AND X1, X2, X3
Meaning : X1 = X2 & X3
Comments : Bitwise-And of X2 and X3, stores result in X1
```

```
INSTRUCTION : LOGICAL OR
Example : ORR X1, X2, X3
Meaning : X1 = X2 | X3
Comments : Bitwise-Or of X2 and X3, stores result in X1
```

```
INSTRUCTION : LOGICAL EXCLUSIVE-OR
Example : EOR X1, X2, X3
Meaning : X1 = X2 ^ X3
Comments : Bitwise-Xor of X2 and X3, stores result in X1
```

```
INSTRUCTION : LOGICAL AND IMMEDIATE
Example : ANDI X1, X2, #20
Meaning : X1 = X2 & 20
Comments : Bitwise-And of X2 with a constant, stores result in X1
```

```
INSTRUCTION : LOGICAL OR IMMEDIATE
Example : ORRI X1, X2, #20
Meaning : X1 = X2 | 20
Comments : Bitwise-Or of X2 with a constant, stores result in X1
```

```
INSTRUCTION : LOGICAL EXCLUSIVE-OR IMMEDIATE
Example : EORI X1, X2, #20
Meaning : X1 = X2 ^ 20
Comments : Bitwise-Xor of X2 with a constant, stores result in X1
```

```
INSTRUCTION : LOGICAL LEFT SHIFT
Example : LSL X1, X2, #10
Meaning : X1 = X2 << 10
Comments : Left shifts X2 by a constant, stores result in X1
```

```
INSTRUCTION : LOGICAL RIGHT SHIFT
Example : LSR X1, X2, #10
Meaning : X1 = X2 >> 10
Comments : Right shifts X2 by a constant, stores result in X1
```

```
INSTRUCTION : COMPARE AND BRANCH ON EQUAL 0
Example : CBZ X1, label
Meaning : if (X1 == 0) go to label
Comments : Equal 0 test; PC-relative branch
```

```
INSTRUCTION : COMPARE AND BRANCH ON NOT EQUAL 0
Example : CBNZ X1, label
Meaning : if (X1 != 0) go to label
Comments : NotEqual 0 test; PC-relative branch
```

```
INSTRUCTION : CONDITIONAL BRANCH
Example : B.cond label
Meaning : if (condition true) go to label
Comments : Test condition codes; if true, then branch
```

```
INSTRUCTION : UNCONDITIONAL BRANCH
Example : B label
Meaning : go to label
Comments : Branch to PC-relative target address
```

```
INSTRUCTION : UNCONDITIONAL BRANCH TO REGISTER
Example : BR LR
Meaning : go to address stored in LR
Comments : Branch to address stored in register. Used for switch, procedure return
```

```
INSTRUCTION : UNCONDITIONAL BRANCH WITH LINK
Example : BL label
Meaning : X30 = PC + 4; go to label
Comments : For procedure call (PC-relative)
```