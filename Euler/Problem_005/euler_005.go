package main
import "fmt"

func gcd(a int64,b int64) int64{
    var tmp int64
    for ;b!=0;{
        tmp = b
        b = a % b
        a = tmp
    }
    return a
}

func lcm (a int64,b int64) int64 {
    return (a*b)/gcd(a,b)
}

func main() {
    var t int
    var n int64
    var j int64
    var lcm_all int64
    fmt.Scanf("%d",&t)
    for i:=0;i<t;i++{
        lcm_all = 1
        fmt.Scanf("%d",&n)
        for j=1;j<=n;j++{
            lcm_all = lcm(j,lcm_all) 
        }
        fmt.Printf("%d\n",lcm_all)
    }
}
