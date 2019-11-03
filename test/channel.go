package main

import ( "fmt"

)


var myres=make(map[int]int,0)

func Fac(n int) {
	c := make(chan int,0)

	var res = 1
	for i := 1; i <= n; i++ {
		res *= i

	}
	myres[n] = res
	c<-0

}
func main() {
	c := make(chan int,0)
	for i := 1; i <= 20; i++ {
		go Fac(i)

	}
	for i,v:= range myres {
		fmt.Printf("myres[%d]=%d\n",i, v)
	}
	<-c


}

