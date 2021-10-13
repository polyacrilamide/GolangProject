/*
Алгоритм для задачи №120: triangle. 
Необходимо найти путь от верхнего числа треугольника до любого числа
в его основании так, чтобы сумма чисел, лежащих на пути, была минимальной. 

Алгоритм: начнем с предпоследней строки чисел. Прибавим к каждому числу
наименьшего соседа из последней строки (соседом является число, расположенное левее
или правее по диагонали от текущего числа).

После этого перейдем на строку выше и снова прибавим к каждому числу
наименьшего соседа из нижней строки. 

Продолжим выполнять это до первой строки включительно. Самое верхнее число
окажется равным наименьшей сумме пути. 

*/
package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Println("Введите количество строк: ")
    fmt.Scan(&n)
    triangle := make([][]int, n)
    for i := 0; i < n; i++ {
        triangle[i] = make([]int, n)
    }
    
    fmt.Printf("Введите числовой треугольник:\n")
    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            fmt.Scan(&triangle[i][j])
        }
    }
    fmt.Print()
    fmt.Print("Минимальная сумма пути: ", minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
    for i := len(triangle) - 2; i >= 0; i--{ 
        for j := 0; j <= i; j++{
            triangle[i][j] += minElement(triangle[i + 1][j], triangle[i + 1][j + 1])
        }
    }
    return triangle[0][0]
}

func minElement(a int, b int) int{
    if a < b{
        return a
    }
    return b
}

/* 
func printMatrix(randMatrix [][]int) {
    for _, val := range randMatrix {
        fmt.Println(val)
    }
}
*/
