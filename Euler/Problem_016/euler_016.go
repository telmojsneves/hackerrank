package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	var sum,t,n int
    fmt.Scanf("%d",&t)

    for ;t!=0;t--{

        fmt.Scanf("%d",&n)
        sum = 0
        i, e := big.NewInt(2), big.NewInt(int64(n))
        i.Exp(i, e, nil)
        str := i.String()

    	for _,ch := range str{
            value := string(ch)
            newvalue,_ := strconv.Atoi(value)
            sum += newvalue
        }

        fmt.Println(sum)
    }
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}
