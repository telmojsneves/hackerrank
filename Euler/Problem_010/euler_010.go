
package main
import "fmt"
import "math"

const max = 1000000
var arr [max]int64

func IsPrime(n int64) bool{
    var i int64
    var sqrt_n int64

    sqrt_n = int64(math.Sqrt(float64(n)))
    if n%2 == 0{
        return false
    }
    i = 3 //start

    for ;i<= sqrt_n;{
        if n%i==0 {
            return false
        }
        i+=2
    }
    return true
}

func listPrimes() {

    var value int64
    index:=1

    arr[0] = 2

    for value=3;value<max;{

        if IsPrime(value){
            arr[index] = value
            index++
        }
        value += 2

    }


}

func main() {

    var t int
    var n,sum int64

    fmt.Scanf("%d",&t)
    listPrimes()
    for ;t!=0;t--{

        sum = 0
        fmt.Scanf("%d",&n)
        for i:=0;arr[i]<=n;i++{
            sum = sum + arr[i]
        }
        fmt.Println(sum)

    }
}

/*
func IsPrime(n int64) bool{
    var start int64
    var sqrt_n int64

    start = 3

    sqrt_n = int64(math.Sqrt(float64(n)))

    for ;start<= sqrt_n;{
        if n%start==0 {
            return false
        }
        start+=2
    }
    return true
}

func sum(n int64) int64 {

    var sum_value int64
    var value int64

    if n==1{
        return 0
    }

    sum_value = 2

    for value=3;value<=n;value+=2{
        if IsPrime(value){
            sum_value = sum_value + value
        }
    }
    return sum_value

}

func main() {

    var t,n int64
    fmt.Scanf("%d",&t)
    for ;t!=0;t--{

        fmt.Scanf("%d",&n)
        fmt.Println(sum(n))

    }
}*/
