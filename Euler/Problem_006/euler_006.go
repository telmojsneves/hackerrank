package main
import "fmt"

func main() {
    var t int
    var n int
    var square_sum int
    var sum_square int

    fmt.Scanf("%d",&t)

    for i:=0;i<t;i++{
        square_sum = 0
        sum_square = 0
        fmt.Scanf("%d",&n)
        for j:=1;j<=n;j++{

            sum_square += j

            square_sum += j*j
        }
        sum_square = sum_square*sum_square
        result := sum_square - square_sum
        fmt.Println(result)
    }
}
