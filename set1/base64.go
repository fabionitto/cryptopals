package main

import (
    "fmt"
    "encoding/hex"
    "log"
)

/* Encode hex to base64 */
func hex2base64(src []byte) []byte {
    /* local variable declaration */

    rawinput := make([]byte, hex.DecodedLen(len(src)))

    _, err := hex.Decode(rawinput, src)
    if err != nil {
        log.Fatal(err)
    }

    result := binary2base64(rawinput)
    return result
}

func binary2base64(src []byte) []byte {
    codetable := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

    /* Is src multiple of 3 ? */
    switch len(src) % 3 {
        case 1:
            src = append(src, 0x00, 0x00)
        case 2:
            src = append(src, 0x00)
    }

    result := make([]byte, len(src)/3*4)

    for i := 0 ; i < len(src) ; i +=3 {
        byte1 := uint32(src[i])
        byte2 := uint32(src[i+1])
        byte3 := uint32(src[i+2])

        char1 := byte1 >> 2
        char2 := (byte1 & 0x03)<<4 | (byte2 >> 4)
        char3 := (byte2 & 0x0f)<<2 | (byte3 >> 6)
        char4 := byte3 & 0x3f

        result = append(result, codetable[char1], codetable[char2], codetable[char3], codetable[char4])
    }

    return result
}

func main() {
    test1 := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
    test2 := []byte("fo")

    b64 := binary2base64(test2) 
    r := hex2base64(test1) 
    fmt.Printf("%s\n", b64)
    fmt.Printf("%s\n", r)
}
