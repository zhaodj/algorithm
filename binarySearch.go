package main

func search(arr []int,des int){
	middle := len(arr)/2
	println(middle)
}

func main(){
	sl := make([]int,9)
	for i :=0;i<len(sl);i++ {
		sl[i]=i;
	}
	search(sl,4)
}
