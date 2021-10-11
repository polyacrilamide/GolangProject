/*
Алгоритм для задачи №120:triangle.
Необходимо найти путь от верха треугольника до его низа так, 
чтобы сумма пути была наименьшей; вывести эту сумму.

Алгоритм (функция minimumTotal): в каждой строке ищем наименьший элемент и 
прибавляем его к общей сумме. В условии было указано, что числа в треугольнике 
не превышают 10000. Поэтому при поиске минимума сравниваем с 10001. 

*/
package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Println("Enter the numbers of rows: ")
    fmt.Scan(&n)
    triangle := make([][]int, n)
    for i := 0; i < n; i++ {
        triangle[i] = make([]int, n)
    }
    
    fmt.Printf("Enter the lower triangular matrix:\n")
    for i, rows := range triangle {
        for j := range rows {
            fmt.Scan(&triangle[i][j])
        }
    }
    fmt.Print()
    fmt.Print("Answer: ", minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
    var sum int
    min := 10001
    e := 1
    for i := 0; i < len(triangle); i++ {
        for j := 0; j < e; j++{
            if triangle[i][j] < min {
                min = triangle[i][j]
            }
        }
        e ++
        sum += min
        min = 10000
    }
    return sum
}
