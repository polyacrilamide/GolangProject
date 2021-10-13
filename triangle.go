package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Println("������� ���-�� �����: ")
    fmt.Scan(&n)
    triangle := make([][]int, n)
    for i := 0; i < n; i++ {
        triangle[i] = make([]int, n)
    }
    
    fmt.Printf("������� �������� �����������:\n")
    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            fmt.Scan(&triangle[i][j])
        }
    }
    fmt.Print()
    fmt.Print("����������� ����� ����: ", minimumTotal(triangle))
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