package question11_20

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question11_20"
	"testing"
)

func TestMaxArea(t *testing.T) {
	input:=[]int{1,8,6,2,5,4,8,3,7}
	result:=question11_20.MaxArea(input)
	fmt.Println("MaxArea","input:",input,"result:",result)
	result=question11_20.MaxArea1(input)
	fmt.Println("MaxArea1","input:",input,"result:",result)
}
