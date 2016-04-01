package main 

import (
    "fmt"
   "encoding/hex" 
   "encoding/base64"
   "log"
)

func main() {
    strToEncode := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    bytesToEncode, err := hex.DecodeString(strToEncode)
    
    if (err != nil) {
        log.Println("Error loading string")
    }

    fmt.Println(base64.StdEncoding.EncodeToString(bytesToEncode))
}
