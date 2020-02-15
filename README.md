# FRACTIONS CALCULATOR

###### INTEGER
	[1-9]\d*\_
	
###### FRACTION
	([1-9]\d*\/[1-9]\d*)
	
###### OPERATOR
	(\s(\*|\/|\+|\-)\s)
###### ELEMENT
	(INTEGER|FRACTION|INTEGERFRACTION)
	(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*))
###### ECUATION
	ELEMENT(ELEMENT)*
	(INTEGER|FRACTION|INTEGERFRACTION)(OPERATOR(INTEGER|FRACTION|INTEGERFRACTION))*
	(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*))(\s(\*|\/|\+|\-)\s(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*)))*

###### PROCESS
[CASE01]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/CASE01.png "CASE 01"
[CASE02]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/CASE01.png "CASE 02"
[CASE03]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/CASE01.png "CASE 03"
[CASE04]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/CASE01.png "CASE 04"
[CASE05]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/CASE01.png "CASE 05"
[CASE06]: https://raw.githubusercontent.com/lcuriel/fractions/dev/readme/CASE01.png "CASE 06"

![CASE 01][CASE01]
---
![CASE 02][CASE02]
---
![CASE 03][CASE03]
---
![CASE 04][CASE04]
---
![CASE 05][CASE05]
---
![CASE 06][CASE06]

