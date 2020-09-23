// cerner_2^5_2020

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math/big"
)
// Complete the extraLongFactorials function below.
func extraLongFactorials(n int32) {
    var fact = new(big.Int)
    fact.MulRange(1,int64(n))
    fmt.Println(fact)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    extraLongFactorials(n)
}
