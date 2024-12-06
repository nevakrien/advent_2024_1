package main

import (
    "fmt"
    "os"
    "bufio"
    "sort"
)

func absInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func main() {
    // Open the file
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() // Ensure the file is closed when the function exits
    reader := bufio.NewReader(file)

    A := make([]int, 0) // Empty slice of integers
    B := make([]int, 0) // Empty slice of strings

    //read inputs
    for {
       var input_a,input_b int
        n,err := fmt.Fscanf(reader,"%d %d\n",&input_a,&input_b)
        if(n==0) {
            break
        }
        if(n!=2||err!=nil){
            fmt.Println("bad input:",err)
            os.Exit(1)
        } 
        A = append(A,input_a)
        B = append(B,input_b)
    }

    fmt.Println("got:",len(A),len(B))
    sort.Ints(A)
    sort.Ints(B)

    ans := 0

    for i:=0; i<len(A);i++ {
        ans+=absInt(A[i]-B[i])
    }
    fmt.Println("ans:",ans)
}
