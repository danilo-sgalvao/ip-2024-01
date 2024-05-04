package main

import (
    "fmt"
)

func main() {
    var num, countPar, countImpar int
    var sumPar, sumImpar float64

    for {
        fmt.Scan(&num)
        if num == 0 {
            break
        }

        if num%2 == 0 {
            sumPar += float64(num)
            countPar++
        } else {
            sumImpar += float64(num)
            countImpar++
        }
    }

    if countPar > 0 {
        mediaPar := sumPar / float64(countPar)
        fmt.Printf("MEDIA PAR = %.2f\n", mediaPar)
    } else {
        fmt.Println("MEDIA PAR = 0.00")
    }

    if countImpar > 0 {
        mediaImpar := sumImpar / float64(countImpar)
        fmt.Printf("MEDIA IMPAR = %.2f\n", mediaImpar)
    } else {
        fmt.Println("MEDIA IMPAR = 0.00")
    }
}
