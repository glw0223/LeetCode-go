package zConvert

import "fmt"

func ZConvert(s string, numRows int) (result string) {
	fmt.Println("input is ",s,"row is",numRows)
	if numRows == 1 {
		return s
	}
	n := len(s)
	cycleLen := 2 * numRows - 2
	for i := 0; i < numRows; i++ {
		for j:= 0; j + i < n; j += cycleLen {
			result += s[j+i:j+i+1]
			if i != 0 && i != numRows - 1 && j + cycleLen - i < n{
				result+=s[j+cycleLen-i:j+cycleLen-i+1]
			}
		}
		result+="\n"
	}
	return
}
