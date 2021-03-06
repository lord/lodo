package core

const ( 
    Orient_0 = 1 << iota 
    Orient_90 
    Orient_180 
    Orient_270
)


type ch struct {
	r rune
	width int
	data [7]int
}

var letters = []ch {	
{'0',3, [7]int{6,9,9,9,9,9,6}},
{'1',3, [7]int{7,2,2,2,2,6,2}},
{'2',4, [7]int{15,8,4,2,1,9,6}},
{'3',4, [7]int{14,1,1,14,1,1,14}},
{'4',4, [7]int{2,2,2,15,10,10,2}},
{'5',4, [7]int{14,1,1,14,8,8,15}},
{'6',4, [7]int{6,9,9,14,8,9,6}},
{'7',4, [7]int{8,8,4,4,2,1,15}},
{'8',4, [7]int{6,9,9,6,9,9,6}},
{'9',4, [7]int{6,9,1,7,9,9,6}},
{'A',4, [7]int{9,9,15,9,9,9,6}},
{'B',4, [7]int{14,9,9,14,9,9,14}},
{'C',4, [7]int{6,9,8,8,8,9,6}},
{'D',4, [7]int{14,9,9,9,9,9,14}},
{'E',4, [7]int{15,8,8,14,8,8,15}},
{'F',4, [7]int{8,8,8,14,8,8,15}},
{'G',4, [7]int{6,9,11,8,8,9,6}},
{'H',4, [7]int{9,9,9,15,9,9,9}},
{'I',3, [7]int{7,2,2,2,2,2,7}},
{'J',4, [7]int{4,10,2,2,2,2,7}},
{'K',4, [7]int{9,10,12,8,12,10,9}},
{'L',4, [7]int{15,8,8,8,8,8,8}},
{'M',5, [7]int{17,17,17,17,21,27,17}},
{'N',4, [7]int{9,9,9,9,11,13,9}},
{'O',4, [7]int{6,9,9,9,9,9,6}},
{'P',4, [7]int{8,8,8,14,9,9,14}},
{'Q',5, [7]int{13,18,22,18,18,18,12}},
{'R',4, [7]int{9,9,9,14,9,9,14}},
{'S',4, [7]int{6,9,1,6,8,9,6}},
{'T',4, [7]int{4,4,4,4,4,4,15}},
{'U',4, [7]int{6,9,9,9,9,9,9}},
{'V',4, [7]int{4,10,9,9,9,9,9}},
{'W',5, [7]int{10,21,17,17,17,17,17}},
{'X',4, [7]int{9,9,9,6,9,9,9}},
{'Y',5, [7]int{4,4,4,4,10,17,17}},
{'Z',4, [7]int{15,8,8,6,1,1,15}},
{'a',4, [7]int{7,9,7,1,6,0,0}},
{'b',4, [7]int{14,9,9,14,8,8,0}},
{'c',4, [7]int{6,9,8,9,6,0,0}},
{'d',4, [7]int{7,9,9,7,1,1,0}},
{'e',4, [7]int{6,8,15,9,6,0,0}},
{'f',3, [7]int{2,2,2,7,2,1,0}},
{'g',4, [7]int{6,1,7,9,9,6,0}},
{'h',4, [7]int{9,9,9,14,8,8,0}},
{'i',1, [7]int{1,1,1,1,0,1,0}},
{'j',2, [7]int{2,5,1,1,0,1,0}},
{'k',4, [7]int{9,10,12,10,9,8,8}},
{'l',2, [7]int{1,2,2,2,2,2,0}},
{'m',5, [7]int{21,21,21,21,30,0,0}},
{'n',4, [7]int{9,9,9,9,14,0,0}},
{'o',4, [7]int{6,9,9,9,6,0,0}},
{'p',4, [7]int{8,14,9,9,14,0,0}},
{'q',4, [7]int{1,7,9,9,7,0,0}},
{'r',4, [7]int{4,4,4,5,11,0,0}},
{'s',4, [7]int{14,1,6,8,7,0,0}},
{'t',3, [7]int{1,2,2,2,7,2,2}},
{'u',4, [7]int{6,9,9,9,9,0,0}},
{'v',4, [7]int{4,10,9,9,9,0,0}},
{'w',5, [7]int{10,21,17,17,17,0,0}},
{'x',4, [7]int{9,9,6,9,9,0,0}},
{'y',6, [7]int{6,1,7,9,9,0,0}},
{'z',4, [7]int{15,8,6,1,15,0,0}},
{'!',1, [7]int{1,0,1,1,1,1,0}},
{'+',3, [7]int{0,2,7,2,0,0,0}},
{' ',1, [7]int{0,0,0,0,0,0,0}},
{':',1, [7]int{0,1,0,1,0,0,0}},
{'.',1, [7]int{1,0,0,0,0,0,0}}}