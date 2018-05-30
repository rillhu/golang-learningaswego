package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"study_03_RPN/stack"
)

//从左往右开始，遇到符号则把符号前面的两个数字与符号运算，然后再将数字放回。
//中缀表达式：
//8-(2-1)*4+5/1
//生成的后缀表达式字符以空格分隔:
// 8 2 1 - 4 * - 5 1 / + //note the string is end with a blank, res should be 9
//REF: https://www.cnblogs.com/hu772188392-163/p/5775367.html
//REF: https://blog.csdn.net/xiaolei09bupt/article/details/44888049

/*Generate a RPN*/
func RPNGen()string {
	var reader = bufio.NewReader(os.Stdin)
	var s1 = new(stack.MyStack) //store operation code + - * / ()
	var s2 = new(stack.MyStack) //store operation number
	var token string
	s, err := reader.ReadString('\n')

	if err!= nil {
		fmt.Println("Input error RPNGen")
		return ""
	}

	s1.Push('#')
	for _, c := range s {
		switch {
		case c=='+': fallthrough
		case c=='-':
			for ch:=s1.Top();ch!='#';ch=s1.Top(){
				if ch=='('{
					break
				}else{
					s2.Push(s1.Pop())
				}
			}
			s1.Push(int(c))
		case c=='*': fallthrough
		case c=='/':
			for ch:=s1.Top();ch!='#'&&ch!='+'&&ch!='-';ch=s1.Top(){
				if ch=='('{
					break
				}else{
					s2.Push(s1.Pop())
				}
			}
			s1.Push(int(c))
		case c=='(':
			s1.Push(int(c))
		case c==')':
			for  s1.Top() != '('{
				s2.Push(s1.Pop())
			}
			s1.Pop()
		//TODO: currently only single number is supported, should support multi numbers
		case c>='0'&&c<='9': //number:
			//token = token + string(c)
			token = string(c)
			r, _ := strconv.Atoi(token)
			s2.Push(r)
			token = ""
		default:
			//
		}
	}
	//Push all rest operator into s2, s2 will have all elements
	for s1.Top()!='#'{
		s2.Push(s1.Pop())
	}

	//Convert the stack into string
	var result string
	for k:=s2.Pop();k!=0;k=s2.Pop(){
		if k>='('&&k<='/'{
			result += string(k)
		}else {
			result += strconv.Itoa(k)
		}
		result += " "
	}
	fmt.Println(result)
	//reverse a string
	bytes := []rune(result)
	//bytes2 := []rune(result)
	for from , to := 0 , len(bytes) -1 ; from < to ; from , to = from + 1, to -1{
		bytes[from] , bytes[to] = bytes[to] , bytes[from]
	}
	result = string(bytes)
	fmt.Println(result)
	//remove the 1st blank and add blank into last position
	for i:=0;i<len(bytes)-1;i++ {
		bytes[i] = bytes[i+1]
	}
	bytes[len(bytes)-1] = ' '
	result = string(bytes)
	fmt.Println(result)

	return result
}

/*Calculate the result of a RPN*/
func RPNCalc(s string) int {
	var st = new(stack.MyStack)
	var token string
	/*
	//Receive the REN from std input
	var reader = bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	*/
	for _, c := range s {
		switch {
		case c >= '0' && c <= '9':
			token = token + string(c)
		case c == ' ':
			r, _ := strconv.Atoi(token)
			st.Push(r)
			token = ""
		case c == '+':
			token = strconv.Itoa(st.Pop() + st.Pop())
			fmt.Printf("%v\n", token)
			//st.Push(token)
		case c == '*':
			token = strconv.Itoa(st.Pop() * st.Pop())
			fmt.Printf("%v\n", token)
			//st.Push(res)
		case c == '-':
			p := st.Pop()
			q := st.Pop()
			token = strconv.Itoa(q - p)
			fmt.Printf("%v\n", token)
		case c == '/':
			p := st.Pop()
			q := st.Pop()
			token = strconv.Itoa(q / p)
			fmt.Printf("%v\n", token)
		case c == 'q':
			return 0
		default:
			//error
		}
	}
	result := st.Pop()
	fmt.Println("Result is: ", result) //

	return result
}

func main() {
	str := RPNGen()
	fmt.Println(str)
	res := RPNCalc(str)
	fmt.Println(res)

}
