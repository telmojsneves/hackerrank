package main
import "fmt"

//BOTTOM UP solution (best approach)

const MAX int = 15

var matrix [MAX][MAX]int

func getMax(x int,y int) int{
    if x > y{
        return x
    }
    return y
}

func bottomUp(line int) int {

    for i:= line-1;i>=0;i--{
        for j:=0;j<=i;j++{
            matrix[i][j] += getMax(matrix[i+1][j],matrix[i+1][j+1])

        }

    }
    return matrix[0][0]


}

func main() {

    var t,n int
    var result int
    var num int

    fmt.Scanf("%d",&t)

    for ;t!=0;t--{
        fmt.Scanf("%d",&n)
        for i:=0;i<n;i++{
            for j:=0;j<=i;j++{

                fmt.Scanf("%d",&num)
                matrix[i][j] = num

            }
        }

        result = bottomUp(n-1)

        fmt.Println(result)

    }
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}
