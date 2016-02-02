package main
import "fmt"

const MAXVALUE int = 5000001

var m map[int]int
var array [MAXVALUE]int

func calcBiggest(n int) int{
//    if value,ok := m[n]; ok{
//        return value

    if n > MAXVALUE{
        if n%2 == 0{
            return calcBiggest(n/2) + 1
        }else{
            return calcBiggest(3*n+1) + 1
        }
    }else{
        if m[n] == 0{
            if n%2 == 0{
                m[n] = calcBiggest(n/2) + 1
            }else{
                m[n] = calcBiggest(3*n+1) + 1
            }
        }
        return m[n]
    }
}

func main() {

    var t,i int
    var n int
    m = make(map[int]int)

    m[1] = 1
    array[1] = 1

    for i=2;i<=MAXVALUE-1;i++{
        if calcBiggest(i) >= m[array[i-1]]{
            array[i] = i
        }else{
            array[i] = array[i-1]

        }
    }

    fmt.Scanf("%d",&t)

    for ;t!=0;t--{

        fmt.Scanf("%d",&n)

        fmt.Println(array[n])



    }

}
