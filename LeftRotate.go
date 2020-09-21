// cerner_2^5_2020

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func rotateLeft(d int32, arr []int32) []int32 {
    // Write your code here
    x := int32(len(arr))
   // var Arr []int32
    Arr := make([]int32, x)
    for i := 0; i < len(arr);  i++ {
        j := (int32(i) + x - d)% x
        Arr[j] = arr[i]
    }
    return Arr
}
