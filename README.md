# FRACTIONS CALCULATOR

###### INTEGER
	[1-9]\d*\_
	For example:
	  - 123_
	  - 23423_
	  - 9_
	  - NOT 0

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
	  - *
	  - /
	  - +
	  - -
	
###### ELEMENT
	(INTEGER|FRACTION|INTEGERFRACTION)
	(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*))
	For example:
	  - 12_
	  - 2/5
	  - 3_65/695
	  - NOT 0_1/3
	  - NOT 33_0/98
	  - NOT 56_4566/0 
	  - NOT 0_0/0
	
###### ECUATION
	ELEMENT(OPERATORELEMENT)*
	^(INTEGER|FRACTION|INTEGERFRACTION)(OPERATOR(INTEGER|FRACTION|INTEGERFRACTION))*$
	^(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*))(\s(\*|\/|\+|\-)\s(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*)))*$
	For example:
	  - 1/2 + 3/4 * 5_34/345 / 3_ * 55/123 + 334_234/1234 + 3432 / 33_4/13 * 23/66 - 33_ * 1/3 + 33_
	  - 5_ * 4/5 * 6/97 * 89_34/701 * 5/604 * 3756_ * 40/287 + 5_6/32 * 40_  * 8/75 * 6_10/27 * 8/64 * 8/5 * 2_1/698 * 60_1/249
	  - 90/120 / 93_ / 87_ / 40/983 / 4_ / 70/597 / 8_4/6 / 3_ / 578/3624 / 9_5/934 / 5_9/21
	  - 5 - 9_3/48 - 6/5 - 98_7/9 - 63_4/59 - 3/240 - 5/98 - 3_4/8 - 4_3/8 - 5/6
	  - 30_ + 1_7/48 + 6/38 + 47_2/594 + 3_8/7 + 10_2/96 + 5/42 + 33_ + 4_ + 3/60 + 9_ + 5/67
	  - 40_ / 2/67 * 896/5 / 40_1/2 - 9_5/6 + 23/4250 * 1_ - 2_8/7 / 6/5 - 40/1234 * 7/8 - 5/6 + 1_3/24  78/7 + 6_ - 5_
	  - 50_
	  - 5/8

###### PROCESS
[CASE01]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/Case01.png "CASE 01"
![CASE 01][CASE01]
---
[CASE02]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/Case02.png "CASE 02"
![CASE 02][CASE02]
---
[CASE03]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/Case03.png "CASE 03"
![CASE 03][CASE03]
---
[CASE04]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/Case04.png "CASE 04"
![CASE 04][CASE04]
---
[CASE05]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/Case05.png "CASE 05"
![CASE 05][CASE05]
---
[CASE06]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/Case06.png "CASE 06"
![CASE 06][CASE06]

