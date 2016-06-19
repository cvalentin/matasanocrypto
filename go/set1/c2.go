package main

import (
    "fmt"
    "encoding/hex"
    "os"
    "log"
)

func checkError(err error) {
    if (err != nil) {
        log.Println(err)
    }
}

//Takes in two arguments and xors them together
func main() {

    firstBuf, err := hex.DecodeString(os.Args[1])
    checkError(err)

    secondBuf, err := hex.DecodeString(os.Args[2])
    checkError(err)

    var outputBuf []byte

    for x := range firstBuf {
        outputBuf = append(outputBuf, firstBuf[x]^secondBuf[x]) 
    }

    fmt.Println(hex.EncodeToString(outputBuf))
}
