package main
import "fmt"

//proof : http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/runsums/triNbProof.html

func mults(p int64,n int64)  int64{
    n=n-1
    j:= n/p
    var result int64
    result =(p*j*(j+1))/2 
    return result
}

func main() {
    var t,c int
    var n,result int64
    fmt.Scanf("%d", &t)
    for c=0; c<t ;c++{
        result = 0
        fmt.Scanf("%d", &n) 
        result=mults(3,n)+mults(5,n)-mults(15,n)
        fmt.Println(result)
    } 
}
