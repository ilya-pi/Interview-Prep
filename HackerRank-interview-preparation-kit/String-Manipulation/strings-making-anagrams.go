package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'makeAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func makeAnagram(a string, b string) int32 {
    characters := make(map[rune]int32)
    // First populate map of characters from one string
    for _, v := range a {
        characters[v]++
    }    
    
    // Substract character from the second string in this map
    for _, v := range b {
        characters[v]--
    }
    
    absInt32 := func(v int32) int32 {
        if v < 0 {
            return -v
        }
        return v
    }
    
    var r int32
    for _, v := range characters {
        r += absInt32(v)
    }
    
    return r

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    a := readLine(reader)

    b := readLine(reader)

    res := makeAnagram(a, b)

    fmt.Fprintf(writer, "%d\n", res)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
