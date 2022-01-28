# FRACTIONS CALCULATOR
   Is a command line that will take operations on fractions as an input and produce a fractional result. 
   
   Assumptions 
   - Valid operators shall be *, /, +, - (multiply, division, addition and subtract)
   - Operands and operators shall be separated by one or more spaces. The command has the ability to eliminate unnecessary spaces
   - Operand is defined by, later it will be explained more thoroughly by regular expressions: 
      + Single whole. 3 345
      + Fraction 2/5 54/535
      + Mixed numbers 5_55/67
  - When is passed a single operand, the command calculate a reduction. By example if the input is 44_56/5 the command will return 55_1/5
  - [TEST](#examples-of-use)
  
###### UPDATE
In this version it was added 
   - Use negative numbers like -5, -5_3/5, -3/5, 3/-5
   - According to the rules; if the fraction have the numerator and denominator is negative, then both numbers will be positive
   - Logger to see how the equations is reduced according to the PEMDAS
    
    ./fractions "1/4 - 1/8 + 1/2 / 2/5 * 1/4" logs
    0  =>  Current operator: '*'  Element: 1/4 - 1/8 + 1/2 / 2/5 * 1/4
    1  =>  Current operator: '*'  Element: 1/4 - 1/8 + 1/2 / 1/10
    2  =>  Current operator: '/'  Element: 1/4 - 1/8 + 1/2 / 1/10
    3  =>  Current operator: '/'  Element: 1/4 - 1/8 + 5
    4  =>  Current operator: '+'  Element: 1/4 - 1/8 + 5
    5  =>  Current operator: '+'  Element: 1/4 - 5_1/8
    6  =>  Current operator: '-'  Element: 1/4 - 5_1/8
    7  =>  Current operator: '-'  Element: -4_-28/32
    The result is:  -4_-28/32
    
        
### DEFINITIONS   
###### INTEGER
    An integer is a number that can be written without a fractional component and is greater than.
	\-?[1-9]\d*
	For example:
	  - 123
	  - 23423
	  - 9
	  - NOT 0

###### FRACTION
	\-?[1-9]\d*\/\-?[1-9]\d*
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
	((\-?[1-9]\d*)|(\-?[1-9]\d*\/\-?[1-9]\d*)|(\-?[1-9]\d*\_\-?[1-9]\d*\/\-?[1-9]\d*))
	For example:
	  - 12
	  - 2/5
	  - 3_65/695
	  - NOT 0_1/3
	  - NOT 3_0/8
	  - NOT 56_4/0 
	  - NOT 0_0/0
	
###### ECUATION
	ELEMENT(OPERATORELEMENT)*
	^(INTEGER|FRACTION|INTEGERFRACTION)(OPERATOR(INTEGER|FRACTION|INTEGERFRACTION))*$
	^((\-?[1-9]\d*)|(\-?[1-9]\d*\/\-?[1-9]\d*)|(\-?[1-9]\d*\_\-?[1-9]\d*\/\-?[1-9]\d*))(\s(\*|\/|\+|\-)\s((\-?[1-9]\d*)|(\-?[1-9]\d*\/\-?[1-9]\d*)|(\-?[1-9]\d*\_\-?[1-9]\d*\/\-?[1-9]\d*)))*$
	For example:
	  - 5 * 4/5 * 6/97 * 9_3/7 * 5/4 * 6 * 4/87 + 5_6/12 * 40  * 8/75 * 6_10/27 * 8/64 * 8/5 * 2_1/8 * 60_1/9
	  - 1/2 + 3/4 * 5_34/345 / 3 * 55/123 + 33_4/34 + 343 / 33_4/13 * 23/66 - 33 * 1/3 + 33
	  - 90/12 / 3 / 7 / 40/83 / 4 / 70/97 / 8_4/6 / 3 / 5/24 / 9_5/34 / 5_9/21
	  - 5 - 9_3/48 - 6/5 - 98_7/9 - 63_4/59 - 3/240 - 5/98 - 3_4/8 - 4_3/8 - 5/6
	  - 30 + 1_7/48 + 6/38 + 47_2/594 + 3_8/7 + 10_2/96 + 5/42 + 33 + 4 + 3/60 + 9 + 5/67
	  - 40 / 2/67 * 6/5 / 40_1/2 - 9_5/6 + 2/50 * 1 - 2_8/7 / 6/5 - 40/12 * 7/8 - 5/6 + 1_3/24 + 78/7 + 6 - 5
	  - 50
	  - 5/8
	  - 55/8
	  - 230/12


### Examples of use
INPUT | RESULT
:--- | ---:
./fractions "11" logs                                                                                                      | 11
./fractions "20/5"                                                                                                      | 4
./fractions "341156172/28576800"                                                                                        | 11_319183/340200
./fractions "5 * 4/5"                                                                                                  | 4
./fractions "5_1/3 * 4/5"                                                                                               | 4_4/15
./fractions "5_6/6 * 4_3/4"                                                                                             | 28_1/2
./fractions "1/4 &nbsp;&nbsp;&nbsp; * &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1/8 &nbsp;&nbsp;&nbsp;&nbsp; / 1/2 &nbsp;&nbsp;&nbsp;&nbsp; + &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1/16 &nbsp;&nbsp;&nbsp;&nbsp; - &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1/32 * 1/64"                                                                      | 255/2048
./fractions "90/12 / 3 / 7 / 40/83 / 4 / 70/97 / 8_4/6 / 3 / 5/24 / 9_5/34 / 5_9/21"                                | 410601/430175200
./fractions "5 - 9_3/48 - 6/5 - 98_7/9 - 63_4/59 - 3/240 - 5/98 - 3_4/8 - 4_3/8 - 5/6"                                 | -202437322260480/1150997299200
./fractions "40 / 2/67 * 6/5 / 40_1/2 - 9_5/6 + 2/50 * 1_ - 2_8/7 / 6/5 - 40/12 * 7/8 - 5/6 + 1_3/24"                  | 10_69617/340200
./fractions "1/2 + 3/4 * 5_34/345 / 3 * 55/123 + 33_4/34 + 343 / 33_4/13 * 23/66 - 33 * 1/3 + 33"                   | 22_151883/8097100
./fractions "30 + 1_7/48 + 6/38 + 47_2/594 + 3_8/7 + 10_2/96 + 5/42 + 33 + 4 + 3/60 + 9 + 5/67"                     | 138_37817327/52931340
./fractions "5 * 4/5 * 6/97 * 9_3/7 * 5/4 * 6 * 4/87 + 5_6/12 * 40  * 8/75 * 6_10/27 * 8/64 * 8/5 * 2_1/8 * 60_1/9"  | 3819_320603111/358868475

### COMPILING TO DIFFERENTS OS
    # OSX
    GOOS=darwin GOARCH=386 go build fractions.go
    
    # LINUX
    GOOS=linux GOARCH=arm go build fractions.go
