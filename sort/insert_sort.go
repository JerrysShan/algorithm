package main
import "fmt"

func main() {
	arr := []int{1, 23, 4, 5, 2, 9, 10}
    insertSort(arr)
	for _, item := range arr {
		fmt.Println(item)
	}
}

func insertSort(arr []int){
	length:=len(arr)
	for i:=1;i<length;i++{
       for j:=i;j>0;j--{
		   if arr[j]<arr[j-1]{
            arr[j],arr[j-1]=arr[j-1],arr[j]
		   }
	   }
	}
}