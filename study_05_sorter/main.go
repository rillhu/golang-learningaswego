package main

import "fmt"

/*Define interface*/
type Sorter interface {
	Len() int
	Less(i,j int) bool
	Swap(i,j int)
}
/*Define the slice type for sorter*/
type Xi []int
type Xs []string

/*type int implementation*/
func (p Xi) Len()int{
	return len(p)
}

func (p Xi) Less(i, j int) bool{

	return p[i] < p[j]
}

func (p Xi) Swap(i,j int){
	p[i], p[j] = p[j],p[i]
}

/*type string implementation*/
func (p Xs) Len()int{
	return len(p)
}

func (p Xs) Less(i, j int) bool{

	return p[i] < p[j]
}

func (p Xs) Swap(i,j int){
	p[i], p[j] = p[j],p[i]
}

/*function bubble sorter*/
func Sort(x Sorter){
	for i:=0;i<x.Len();i++{
		for j:=0;j<i;j++{
			if x.Less(i,j){
				x.Swap(i,j)
			}
		}
	}
}


/*A interface example
Implement the sort for type int and type string
*/
func main() {
	intNum := Xi{1,2,3,9,5,6,8,7}
	fmt.Println(intNum)
	Sort(intNum)
	fmt.Println(intNum)

	strArr := Xs{"Ago","Beego","Car", "Zoo","Human","Good"}
	fmt.Println(strArr)
	Sort(strArr)
	fmt.Println(strArr)
}
