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
