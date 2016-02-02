package main
import "fmt"
import "math"

const max = 30000
var arr [max]int

func IsPrime(n int) bool{
    var i int
    var sqrt_n int

    sqrt_n = int(math.Sqrt(float64(n)))
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

    var value int
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


func check(number int)int {
    nod := 1
    remain := number

    for i:=0;i<len(arr);i++ {

        // In case there is a remainder this is a prime factor as well
        // The exponent of that factor is 1
        if arr[i] * arr[i] > number {
            return nod * 2
        }

        exponent := 1

        for ;remain % arr[i] == 0; {
            exponent++
            remain = remain/arr[i]
        }
        nod *= exponent

        //If there is no remainder, return the count
        if (remain == 1) {
            return nod;
        }
    }
    return nod;
}


/*bf
func check (triangle_number int) int{
    count := 1
    triangle_sqrt := int(math.Sqrt(float64(triangle_number)))
    for i:=2;i<=triangle_sqrt;i++{
        if triangle_number % i == 0{
            count+=2
        }
    }
    if triangle_sqrt * triangle_sqrt == triangle_number{
        count -=1
    }
    return count
}
*/

func main() {

    listPrimes();

    var t,n int
    fmt.Scanf("%d",&t)

    for ;t!=0;t--{

        fmt.Scanf("%d",&n)
        triangle_number := 1
        i:= 2
        for check(triangle_number) <= n{
            triangle_number+=i
            i++
        }
        if n==1{
            triangle_number = 3
        }
        fmt.Println(triangle_number)


    }
}
