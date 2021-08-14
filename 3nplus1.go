package main

import ("fmt"; "strconv")

func calc_sequence(num int) int{
	if num % 2 == 0{
		num = num / 2
	}else {
		num = (num * 3) + 1
	}
	return num
}

func take_input()int{
	var in_num string
	var out_num int
	fmt.Scanln(&in_num)
	out_num, error := strconv.Atoi(in_num)
	_ = error
	return out_num
}

func main(){
	var num int
	fmt.Println("Please enter a number: ")
	num = take_input()
	for{
		num = calc_sequence(num)
		fmt.Println(num)
		if num == 1{
			break
		}
	}
}