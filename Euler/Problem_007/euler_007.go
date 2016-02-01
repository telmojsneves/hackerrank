package main
import "fmt"
import "math"

const max = 10000
var arr [max]int64

func IsPrime(n int64,start int64) bool{
    var i int64
    var sqrt_n int64

    sqrt_n = int64(math.Sqrt(float64(n)))
    if n%2 == 0{
        return false
    }
    i = start

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
    var countPrimes int

    arr[0] = 2
    value = 1

    for countPrimes=1;countPrimes<max;{

        value += 2

        if IsPrime(value,3){
            arr[countPrimes] = value
            countPrimes++
        }



    }


}

func main() {

    var t,n int
    fmt.Scanf("%d",&t)
    listPrimes()
    for i:=0;i<t;i++{

        fmt.Scanf("%d",&n)
        fmt.Println(arr[n-1])

    }
}
