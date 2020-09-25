// cerner_2^5_2020

import (
    "bufio"
    "io"
    "os"
    "strconv"
    "strings"
)
func gradingStudents(grades []int32) []int32 {
    for i := range grades {
        // get the modulus of 5 for each to know about the nearest value of 5
        diff := 5 - (grades[i] % 5)
        if (grades[i] >= 38) && (diff < 3) {
            grades[i] += diff
        }
    }
    return grades
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    defer stdout.Close()
    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
    gradesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    var grades []int32
    for i := 0; i < int(gradesCount); i++ {
        gradesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        gradesItem := int32(gradesItemTemp)
        grades = append(grades, gradesItem)
    }
