package main
import "fmt"

var numberWordMapUnderTwenty = [20]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven","Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen","Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

var numberWordMapOfTens = [10]string{"Zero", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty","Seventy", "Eighty", "Ninety"}


func NumbToWords(n int64) string{

    if n < 20 {
        return numberWordMapUnderTwenty[n]
    }else if n < 100{

        module := n%10
        val := n/10

        if module == 0 {
            return numberWordMapOfTens[val]
        }
        return numberWordMapOfTens[val] + " " + NumbToWords(module)

    }else if n<1000{
        module := n%100
        val := n/100

        if module == 0 {
            return numberWordMapUnderTwenty[val] + " Hundred"
        }
        return numberWordMapUnderTwenty[val] + " Hundred " + NumbToWords(module)

    }else if n<1000000 {
        module := n%1000
        val := n/1000


        if module == 0{
            return NumbToWords(val) + " Thousand"
        }else{
            return NumbToWords(val) + " Thousand " + NumbToWords(module)
        }

    }else if n <1000000000 {
        module := n%1000000
        val := n/1000000

        if module == 0{
            return NumbToWords(val) + " Million"
        }
        return NumbToWords(val) + " Million " + NumbToWords(module)

    }else if n<1000000000000 {
        module := n%1000000000
        val := n/1000000000

        if module == 0{
            return NumbToWords(val) + " Billion"
        }
        return NumbToWords(val) + " Billion " + NumbToWords(module)

    }else {return "NOTHING"}


}


func main() {

    var t int
    var n int64

    fmt.Scanf("%d",&t)
    for ;t!=0;t--{
        fmt.Scanf("%d", &n)

        fmt.Println(NumbToWords(n))
    }

}
