package binarySearch

import (
	"fmt"
)

//非递归
func BinarySearch1(sortedArray []int, key int)(result string){
	length:=len(sortedArray)
	if length <= 0{
		result = "有序数组异常"
	}else {
		if key < sortedArray[0]{
			result = "key 小于有序数组最小值"
		}else if key > sortedArray[length-1] {
			result = "key 大于有序数组最大值"
		}else {
			index:=length/2
			max:=length
			min:=0
			for(true){
				if key == sortedArray[index]{
					result = fmt.Sprint("成功找到，是有序数组的第 ",index, " 个")
					break
				}else if key < sortedArray[index] {
					max = index
				}else if key > sortedArray[index] {
					min = index
				}

				index = (max+min)/2
				if max == min || index == min || index == max {
					result = "没有找到"
					break
				}else {
					//fmt.Println(max, "  ", index, " ", min)
				}
			}
		}
	}

	return
}

//递归法
func BinarySearch2(sortedArray []int, min, max, key int) (result string){
	if len(sortedArray) <= 0 {
		result = "数组错误"
	}else {
		if key < sortedArray[0] {
			result = "小于最小值"
		}else if key > sortedArray[len(sortedArray)-1]{
			result = "大于最大值"
		}else {
			index:=((max+min)/2)
			temp:=sortedArray[index]
			if key == temp {
				result = fmt.Sprintf("找到啦!!! 是第%d个",index)
			}else if max == min || index == min || index == max {
				result = "没有找到"
			}else {
				result = ""

				if key < temp{
					max = index
				}else if key > temp {
					min = index
				}
				result=BinarySearch2(sortedArray, min, max, key)
			}
		}
	}

	return
}

