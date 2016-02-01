package main

import "fmt"
import "strconv"

func IsPalindrome(num int) bool {
    if num % 10 == 0{
        return false
    }
    s := strconv.Itoa(num)
    mid := len(s) / 2
    last := len(s) - 1
    for i := 0; i < mid; i++ {
        if s[i] != s[last-i] {
            return false
        }
    }
    return true
}

func largestPalindrome(max int) int{
    var prod int
    biggest := 0
    for i := 999; i >= 100; i--{
        for j:= 999; j>=i;j--{
            prod = i*j
            if prod > biggest && prod < max && IsPalindrome(prod){
                biggest = prod
            }
        }
    }
    return biggest


}

func main() {
    var t,n int
    fmt.Scanf("%d",&t)
    for i:=0;i<t;i++ {
        fmt.Scanf("%d",&n)

        fmt.Println(largestPalindrome(n))
    }


}
