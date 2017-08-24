package main

import (
//    "encoding/hex"
    "fmt"
//    "log"
)

func byteAnd(op1, op2 []byte) []byte {
    result := make([]byte, len(op1))

    for i := 0; i < len(op1) ; i++ {
        result[i] = op1[i] & op2[i]
    }

    return result
}

func main() {
    codetable := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
    test := []byte("Man")
    fmt.Printf("%032b\n", uint32(test[0])<<16)
    fmt.Printf("%032b\n", uint32(test[1])<<8)
    fmt.Printf("%032b\n", uint32(test[2]))

    b1 := uint32(test[0])
    b2 := uint32(test[1])
    b3 := uint32(test[2])
    
    c1 := b1 >> 2
    c2 := (b1 & 0x03)<<4 | (b2 >>4)
    c3 := (b2 & 0x0f)<<2 | (b3 >>6)
    c4 := b3 & 0x3f

    fmt.Printf("%08b\n", test)
    fmt.Printf("%032b\n", c1)
    fmt.Printf("%032b\n", c2)
    fmt.Printf("%032b\n", c3)
    fmt.Printf("%032b\n", c4)
    
    fmt.Printf("%c\n", codetable[c1])
    fmt.Printf("%c\n", codetable[c2])
    fmt.Printf("%c\n", codetable[c3])
    fmt.Printf("%c\n", codetable[c4])

}
