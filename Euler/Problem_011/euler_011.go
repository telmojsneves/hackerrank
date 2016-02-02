package main
import "fmt"

const max_size = 20

func main() {

    greatest_product := 0
    value := 0
    values := [max_size][max_size]int{}
    input := 0
    for i:=0 ; i<max_size;i++{
        for j:=0 ; j<max_size;j++{
            fmt.Scanf("%d",&input)
            values[i][j] = input
        }
    }
    for i:=0 ; i<max_size;i++{
        for j:=0 ; j<max_size;j++{
            if values[i][j] == 0{
                continue
            }
            if j+3<max_size{
                value = values[i][j] * values[i][j+1] * values[i][j+2] * values[i][j+3]
                if value > greatest_product{
                    greatest_product = value
                }
            }

            if i+3<max_size{

                value = values[i][j] * values[i+1][j] * values[i+2][j] * values[i+3][j]
                if value > greatest_product{
                    greatest_product = value
                }

                if j+3<max_size{

                    value = values[i][j] * values[i+1][j+1] * values[i+2][j+2] * values[i+3][j+3]
                    if value > greatest_product{
                        greatest_product = value
                    }
                }
                if j-3>0{
                    value = values[i][j] * values[i+1][j-1] * values[i+2][j-2] * values[i+3][j-3]
                    if value > greatest_product{
                        greatest_product = value
                    }
                }

            }
        }
    }
    fmt.Println(greatest_product)

}
