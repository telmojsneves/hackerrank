package main
import "fmt"

func isPrime(n int64, start int64) bool{
    i := start
    for ;i*i <= n;{
        if n%i==0{
            return false
        }
        i += 1
    }
    return true
}

func largest(n int64) int64{
    num := n
    var count int64
    count = 3
    for ;num%2 == 0;{
        num = num / 2
    }
    if num==1{
        return 2
    }
    flag := true
    for ;num != 1; {
        if flag && isPrime(num,count){
            return num
        }
        flag = false
        for ;num%count==0;{
            num/=count
            flag = true
        }
        count += 2
    }
    return count -2
}

func main() {
    var t int
    var n int64

    fmt.Scanf("%d",&t)

    for i:=0 ; i<t ; i++ {

        fmt.Scanf("%d",&n)
        fmt.Println(largest(n))
    }
}
