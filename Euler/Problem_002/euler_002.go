package main
import "fmt"

func fib(n int64) int64 {
    
    var sum,a,b,i int64
    
    sum,a,b = 0, 0, 1

    for ; i < n && a+b <= n; i++ {
        a, b = b, a+b
        
        if b%2 == 0{
            sum += b
        }
    }
    return sum
}

func main() {
    var t,c int
    var n int64
    fmt.Scanf("%d", &t)
    for c=0; c<t ;c++ {
        fmt.Scanf("%d", &n)
        fmt.Println(fib(n))
    } 
    
    
    
}
