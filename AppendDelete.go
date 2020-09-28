// cerner_2^5_2020

import (
    "bufio"
    "fmt"
    "math"
)

func appendAndDelete(s string, t string, k int32) string {
    lens := len(s)
    lent := len(t)
    var count int
    count = 0
    // loop to find the common letters
    for i := float64(0); i<math.Min(float64(lens),float64(lent));i++{
        if s[int(i)] == t[int(i)]{
            count++
        } else {
            break
        }
    }
    diff_s := lens - count
    diff_t := lent - count

    if int32(diff_s + diff_t) > k {
        return "No"
    }
    return "Yes"
}
