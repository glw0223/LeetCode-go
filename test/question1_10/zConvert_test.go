package question1_10

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question1_10"
	"testing"
)

func TestZConvert(t *testing.T){
	s:="HELLOWORLD!!!"
	result:=question1_10.ZConvert(s,3)
	fmt.Println(result)
}
