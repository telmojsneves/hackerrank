package main
import "fmt"

//if b = (n^2 - 2*n*a)/(2*n-2*a)
//n-b-a=c and a^2 + b^2 = c^2
func find_triplets(n int) int{
    max :=0
    n_pow := n*n

    for a:=1;a<n;a++{
        if  ((n_pow - 2*n*a)%(2*n - 2*a)) == 0 {
            b := (n_pow - 2*n*a)/(2*n - 2*a)
            c := n-b-a
            result := a*b*c
            if max < result {
                max = result
            }

        }

    }
    if max == 0{
        return -1
    }
    return max


}

/*
//brute-force
func find_triplets(n int) int{

    max := -1
    var new_max int
    a_max := n/3
    b_max := n/2
    for a:=1;a<a_max;a++{
        a_sqr := a * a
        for b:=a;b<b_max;b++{

            c := n - b - a;

            if a_sqr + b*b == c*c{
                new_max = a*b*c
                if new_max >= max{
                    max = new_max
                }
            }

        }
    }
    return max

}
*/

func main() {

    var t,n int

    fmt.Scanf("%d",&t)

    for ;t>0;t--{
        fmt.Scanf("%d",&n)
        fmt.Println(find_triplets(n))

    }


}
