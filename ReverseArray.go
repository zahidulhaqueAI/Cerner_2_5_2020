// cerner_2^5_2020
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the reverseArray function below.
func reverseArray(a []int32) []int32 {
    // hold the result
    var res []int32

    // get the length of an array
    Len := len(a)
    i := Len -1
     for ; i>=0  ; i--  {
        res = append(res, a[i])
    }
    return res
}
