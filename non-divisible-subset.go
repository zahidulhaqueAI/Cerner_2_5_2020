// cerner_2^5_2020

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

func nonDivisibleSubset(k int32, s []int32) int32 {
    len := int32(len(s))
    // frequency array
    freq := make([] int32, k)
    for i :=int32(0); i < len; i++ {
        index := s[i]%k
        freq[index] += 1
    } 
    // result count
    res := int32(0)
    res = int32(math.Min(float64(freq[0]),float64(1)))  
    for i := int32(1); i <= k/2; i++ {
        if i != k-i{
            res += int32(math.Max(float64(freq[i]), float64(freq[k-i])))
        } else {
            res += int32(math.Min(float64(freq[i]),float64(1)))
        }
    }
    return res
}
