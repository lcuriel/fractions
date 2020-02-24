# FRACTIONS CALCULATOR
   Is a command line that will take operations on fractions as an input and produce a fractional result. 
   
   Assumptions 
   - Valid operators shall be *, /, +, - (multiply, division, addition and subtract)
   - Operands and operators shall be separated by one or more spaces. The command has the ability to eliminate unnecessary spaces
   - Operand is defined by, later it will be explained more thoroughly by regular expressions: 
      + Single whole. 3_ 345_
      + Fraction 2/5 54/535
      + Mixed numbers 5_55/67
  - When is passed a single operand, the command calculate a reduction. By example if the input is 44_56/5 the command will return 55_1/5 

### DEFINITIONS   
###### INTEGER
    An integer is a number that can be written without a fractional component and is greater than.
	[1-9]\d*\_
	For example:
	  - 123_
	  - 23423_
	  - 9_
	  - NOT 0_

###### FRACTION
	([1-9]\d*\/[1-9]\d*)
	For example:
	  - 1/2
	  - 23/345
	  - 12/65
	  - 23423423/342352435
	  - NOT 0/12
	  - NOT 445/0
	  
###### OPERATOR
	(\s(\*|\/|\+|\-)\s)
	For example
	  - * (Multiplication)
	  - / (Division)
	  - + (Addition)
	  - - (Subtraction)
	
###### ELEMENT
    An element represent an Integer, a fraction or the combination of both
	(INTEGER|FRACTION|INTEGERFRACTION)
	(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*))
	For example:
	  - 12_
	  - 2/5
	  - 3_65/695
	  - NOT 0_1/3
	  - NOT 3_0/8
	  - NOT 56_4/0 
	  - NOT 0_0/0
	
###### ECUATION
	ELEMENT(OPERATORELEMENT)*
	^(INTEGER|FRACTION|INTEGERFRACTION)(OPERATOR(INTEGER|FRACTION|INTEGERFRACTION))*$
	^(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*))(\s(\*|\/|\+|\-)\s(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*)))*$
	For example:
	  - 5_ * 4/5 * 6/97 * 9_3/7 * 5/4 * 6_ * 4/87 + 5_6/12 * 40_  * 8/75 * 6_10/27 * 8/64 * 8/5 * 2_1/8 * 60_1/9
	  - 1/2 + 3/4 * 5_34/345 / 3_ * 55/123 + 33_4/34 + 343_ / 33_4/13 * 23/66 - 33_ * 1/3 + 33_
	  - 90/12 / 3_ / 7_ / 40/83 / 4_ / 70/97 / 8_4/6 / 3_ / 5/24 / 9_5/34 / 5_9/21
	  - 5_ - 9_3/48 - 6/5 - 98_7/9 - 63_4/59 - 3/240 - 5/98 - 3_4/8 - 4_3/8 - 5/6
	  - 30_ + 1_7/48 + 6/38 + 47_2/594 + 3_8/7 + 10_2/96 + 5/42 + 33_ + 4_ + 3/60 + 9_ + 5/67
	  - 40_ / 2/67 * 6/5 / 40_1/2 - 9_5/6 + 2/50 * 1_ - 2_8/7 / 6/5 - 40/12 * 7/8 - 5/6 + 1_3/24 + 78/7 + 6_ - 5_
	  - 50_
	  - 5/8
	  - 55/8
	  - 230/12


### Examples of use
INPUT | RESULT
:--- | ---:
./fractions "11_"                                                                                                       | 11_
./fractions "20/5"                                                                                                      | 4_
./fractions "341156172/28576800"                                                                                        | 11_319183/340200
./fractions "5_ * 4/5"                                                                                                  | 4_
./fractions "5_1/3 * 4/5"                                                                                               | 4_4/15
./fractions "5_6/6 * 4_3/4"                                                                                             | 28_1/2
./fractions "1/4 &nbsp;&nbsp;&nbsp; * &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1/8 &nbsp;&nbsp;&nbsp;&nbsp; / 1/2 &nbsp;&nbsp;&nbsp;&nbsp; + &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1/16 &nbsp;&nbsp;&nbsp;&nbsp; - &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1/32 * 1/64"                                                                      | 255/2048
./fractions "90/12 / 3_ / 7_ / 40/83 / 4_ / 70/97 / 8_4/6 / 3_ / 5/24 / 9_5/34 / 5_9/21"                                | 410601/430175200
./fractions "5_ - 9_3/48 - 6/5 - 98_7/9 - 63_4/59 - 3/240 - 5/98 - 3_4/8 - 4_3/8 - 5/6"                                 | -202437322260480/1150997299200
./fractions "40_ / 2/67 * 6/5 / 40_1/2 - 9_5/6 + 2/50 * 1_ - 2_8/7 / 6/5 - 40/12 * 7/8 - 5/6 + 1_3/24"                  | 10_69617/340200
./fractions "1/2 + 3/4 * 5_34/345 / 3_ * 55/123 + 33_4/34 + 343_ / 33_4/13 * 23/66 - 33_ * 1/3 + 33_"                   | 22_151883/8097100
./fractions "30_ + 1_7/48 + 6/38 + 47_2/594 + 3_8/7 + 10_2/96 + 5/42 + 33_ + 4_ + 3/60 + 9_ + 5/67"                     | 138_37817327/52931340
./fractions "5_ * 4/5 * 6/97 * 9_3/7 * 5/4 * 6_ * 4/87 + 5_6/12 * 40_  * 8/75 * 6_10/27 * 8/64 * 8/5 * 2_1/8 * 60_1/9"  | 3819_320603111/358868475

### COMPILING TO DIFFERENTS OS
    # OSX
    GOOS=darwin GOARCH=386 go build fractions.go
    
    # LINUX
    GOOS=linux GOARCH=arm go build fractions.go
