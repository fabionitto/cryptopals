package main

import (
    "fmt"
    "bytes"
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

/* Decode base64 - output hex encoded string */
func base642hex(src []byte) []byte {
    dec := base642binary(src)
    result := make([]byte, hex.EncodedLen(len(dec)))

    hex.Encode(result, dec)

    return result
}

/* Encode binary in Base64 */
func binary2base64(src []byte) []byte {
    codetable := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
    remainder := len(src) % 3
    result := make([]byte, 0)

    /* Is src multiple of 3 ? */
    switch remainder {
        case 1:
            src = append(src, 0x00, 0x00)
        case 2:
            src = append(src, 0x00)
    }

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

    switch remainder {
        case 1:
            /* Replace last AA with == */
            result[len(result)-1] = '='
            result[len(result)-2] = '='
        case 2:
            result[len(result)-1] = '='
    }

    return result
}

/* Decode Base64 */
func base642binary(src []byte) []byte {
    codetable := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
    result := make([]byte, 0)

    for i := 0 ; i < len(src) ; i +=4 {
        byte1 := bytes.IndexByte(codetable, src[i])
        byte2 := bytes.IndexByte(codetable, src[i+1])
        byte3 := bytes.IndexByte(codetable, src[i+2])
        byte4 := bytes.IndexByte(codetable, src[i+3])

        char1 := uint8((byte1 << 2) | (byte2 >> 4))
        char2 := uint8((byte2 & 0x0f) << 4 | (byte3 >> 2))
        char3 := uint8((byte3 & 0x03) << 6 | byte4)

        result = append(result, char1, char2, char3)
    }

    if src[len(src)-1] == '=' {
        result = result[0:len(result)-1]
    }
    if src[len(src)-2] == '=' {
        result = result[0:len(result)-1]
    }

    return result
}

func main() {
    test1 := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
    test2 := []byte("foob")

    e2 := binary2base64(test2) 
    e1 := hex2base64(test1) 
    fmt.Printf("%s, len %d, cap %d\n", e1, len(e1), cap(e1))
    fmt.Printf("%s, len %d, cap %d\n", e2, len(e2), cap(e2))

    d1 := base642binary(e1)
    d2 := base642binary(e2)

    h1 := base642hex(e1)
    h2 := base642hex(e2)

    fmt.Printf("%s, len %d, cap %d\n", d1, len(d1), cap(d1))
    fmt.Printf("%s, len %d, cap %d\n", d2, len(d2), cap(d2))
    fmt.Printf("%s, len %d, cap %d\n", h1, len(h1), cap(h1))
    fmt.Printf("%s, len %d, cap %d\n", h2, len(h2), cap(h2))
}
