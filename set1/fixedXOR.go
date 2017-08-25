package main

import (
    "fmt"
    "log"
    "encoding/hex"
)

func hexfixedxor(h1, h2 []byte) []byte {
    r1 := make([]byte, hex.DecodedLen(len(h1)))
    r2 := make([]byte, hex.DecodedLen(len(h2)))

    _, err := hex.Decode(r1, h1)
    if err != nil {
        log.Fatal(err)
    }

    _, err = hex.Decode(r2, h2)
    if err != nil {
        log.Fatal(err)
    }

    xor := fixedxor(r1, r2)

    result := make([]byte, hex.EncodedLen(len(xor)))
    hex.Encode(result, xor)

    return result
}

func fixedxor(op1, op2 []byte) []byte {
    /* implement error handling */

    result := make([]byte, len(op1))

    for i:=0; i < len(op1); i++ {
        result[i] = op1[i] ^ op2[i]
    }

    return result
}

func main () {
    op1 := []byte("1c0111001f010100061a024b53535009181c")
    op2 := []byte("686974207468652062756c6c277320657965")

    xor := hexfixedxor(op1, op2)
    fmt.Printf("%s\n", xor)

}
