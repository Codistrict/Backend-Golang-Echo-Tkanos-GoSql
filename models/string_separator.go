package models

import (
	"fmt"
	"strconv"
)

func String_Separator_To_Int(str string) []int {
	str2 := str

	var by = []byte{}

	by = []byte(str)
	by2 := byte(0)
	by = append(by, by2)
	str2 = string(by)
	fmt.Println(str2)

	var new string = ""
	var i int = 0

	var data = []int{}

	for by[i] != 0 {
		var co int = 0
		new = ""
		if by[i] == 124 {
			co++
			i++
			for co < 2 {
				if by[i] == 124 {
					co++
					i++
					tempint, _ := strconv.Atoi(new)
					fmt.Println("temp:", tempint)
					data = append(data, tempint)
				} else {
					new += string(by[i])
					i++
					fmt.Println(new)
				}
			}
		} else {
			i++
		}
	}
	fmt.Println(data, i)

	return data
}

func String_Separator_To_String(str string) []string {
	str2 := str

	var by = []byte{}

	by = []byte(str)
	by2 := byte(0)
	by = append(by, by2)
	str2 = string(by)
	fmt.Println(str2)

	var new string = ""
	var i int = 0

	var data = []string{}

	for by[i] != 0 {
		var co int = 0
		new = ""
		if by[i] == 124 {
			co++
			i++
			for co < 2 {
				if by[i] == 124 {
					co++
					i++
					data = append(data, new)
				} else {
					new += string(by[i])
					i++
				}
			}
		} else {
			i++
		}
	}
	fmt.Println(data, i)

	return data
}
