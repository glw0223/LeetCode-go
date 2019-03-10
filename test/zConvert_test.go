package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/zConvert"
	"testing"
)

func TestZConvert(t *testing.T){
	s:="HELLOWORLD!!!"
	result:=zConvert.ZConvert(s,3)
	fmt.Println(result)
}
