package main
import "fmt"
import "math"


const MAX_SIZE int64 = 500

func main() {

    var t int64
    var lines,columns int64

    var i,j int64

    var array [MAX_SIZE+1][MAX_SIZE+1]int64

    for j=0;j<=MAX_SIZE;j++{
        array[0][j] = 1
    }
    for i=0;i<=MAX_SIZE;i++{
        array[i][0] = 1
    }


    for i=1;i<=MAX_SIZE;i++{

        for j=1;j<=MAX_SIZE;j++{

            array[i][j] = array[i][j-1] + array[i-1][j]
            array[i][j] = int64(math.Mod(float64(array[i][j]),math.Pow(10,9)+7))
        }

    }

    fmt.Scanf("%d",&t)
    for ;t!=0;t--{

        fmt.Scanf("%d %d",&lines,&columns)

        x := array[lines][columns]
        fmt.Println(int64(x))


    }

    //Enter your code here. Read input from STDIN. Print output to STDOUT
}
