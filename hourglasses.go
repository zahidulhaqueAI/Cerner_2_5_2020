// cerner_2^5_2020

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
    // for storing the val
    var sum int32
    sum = -10000
    // create hash to store the value
    var maxx int32
    maxx = -10000
    for row := 0; row < 4; row++{
        for col := 0; col < 4; col++{
            sum = arr[row][col] + arr[row][col+1] + arr[row][col+2] + arr[row+1][col+1] +
                  arr[row+2][col] + arr[row+2][col+1] + arr[row+2][col+2]
            // hold the max value
            if maxx <= sum{
                maxx = sum
            }
        }
    }
    return maxx
}
