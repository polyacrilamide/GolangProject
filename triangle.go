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
    
    fmt.Printf("������� ������ ����������� �������:\n")
    for i, rows := range triangle {
        for j := range rows {
            fmt.Scan(&triangle[i][j])
        }
    }
    fmt.Print()
    fmt.Print("�����: ", minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
    var sum int
    min := 10000
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