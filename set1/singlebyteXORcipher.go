	package main

import (
    "fmt"
    "log"
    "encoding/hex"
    "strings"
    "math"
)

func hexFixedXor(h1, h2 []byte) []byte {
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

    xor := fixedXor(r1, r2)

    result := make([]byte, hex.EncodedLen(len(xor)))
    hex.Encode(result, xor)

    return result
}

func fixedXor(op1, op2 []byte) []byte {
    /* implement error handling */

    result := make([]byte, len(op1))

    for i:=0; i < len(op1); i++ {
        result[i] = op1[i] ^ op2[i]
    }

    return result
}


func englishScore(text string) float64 {
    score := 0.0
    engFreqTable := map[byte]float64{
        'a': 8.167,
        'b': 1.492,
        'c': 2.782,
        'd': 4.253,
        'e': 12.702,
        'f': 2.228,
        'g': 2.015,
        'h': 6.094,
        'i': 6.966,
        'j': 0.153,
        'k': 0.772,
        'l': 4.025,
        'm': 2.406,
        'n': 6.749,
        'o': 7.507,
        'p': 1.929,
        'q': 0.095,
        'r': 5.987,
        's': 6.327,
        't': 9.056,
        'u': 2.758,
        'v': 0.978,
        'w': 2.360,
        'x': 0.150,
        'y': 1.974,
        'z': 0.074,
    }

    textFreqTable := make(map[byte]float64)

    size := len(text)
    lower := strings.ToLower(text)

    for i := 0 ; i < size ; i ++ {
        textFreqTable[lower[i]] = ((float64(size)*textFreqTable[lower[i]])+1.0)/float64(size)
    }

    for key, value := range engFreqTable {
        score += math.Abs(value - textFreqTable[key])
    }

    for key, _ := range textFreqTable {
        if (key < 20) || (key > 126) {
            /* Penalty */
            score += 1.0
        }
    }

    return score
}

func chiChi(text string) float64 {
    score := 0.0
    engFreqTable := map[byte]float64{
        'a': 8.167,
        'b': 1.492,
        'c': 2.782,
        'd': 4.253,
        'e': 12.702,
        'f': 2.228,
        'g': 2.015,
        'h': 6.094,
        'i': 6.966,
        'j': 0.153,
        'k': 0.772,
        'l': 4.025,
        'm': 2.406,
        'n': 6.749,
        'o': 7.507,
        'p': 1.929,
        'q': 0.095,
        'r': 5.987,
        's': 6.327,
        't': 9.056,
        'u': 2.758,
        'v': 0.978,
        'w': 2.360,
        'x': 0.150,
        'y': 1.974,
        'z': 0.074,
    }
    textFreqTable := make(map[byte]int)

    size := len(text)

    lower := strings.ToLower(text)

    for i:=0; i < size; i++ {
        textFreqTable[lower[i]] = textFreqTable[lower[i]]+1
    }

    for k, v := range textFreqTable {
        expected := float64(size) * engFreqTable[k]
        if engFreqTable[k] > 0 {
            score += math.Pow((float64(v)- expected),2)/expected
        } else {
            score += math.Pow((float64(v)- expected),2)/0.00000001 
        }
        //fet.Printf("%c, %d, %f, %f\n", k,textFreqTable[k], expected, score)
    }

    return score
}

func expandByte(char byte, length int) string {
    result := make([]byte, 0)
    for i := 0; i < length; i++ {
        result = append(result, char)
    }

    return string(result)
}

func analyzeSingleByteXor(text string) byte {
/* Analyzes text, assuming its a singlebyteXored encrypted text, and returns the most probable char as the key */
    keyScore := make(map[int]float64)

    for i := 65 ; i <= 127 ; i++ {
        key := expandByte(byte(i), len(text))

        decryptedText := fixedXor([]byte(key), []byte(text))
        keyScore[i] = chiChi(string(decryptedText))
        fmt.Printf("Key: %c, ChiChi: %f\n", i, keyScore[i])
    }

    result := 0
    score := keyScore['a']
    for k,v := range keyScore {
        if v < score {
            score = v
            result = k
        }
    }

    return byte(result)
}

func main () {
    text := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    r := make([]byte, hex.DecodedLen(len(text)))

    _, err := hex.Decode(r, text)
    if err != nil {
        log.Fatal(err)
    }

    k := analyzeSingleByteXor(string(r))
    key := expandByte(k, len(r))

    decryptedText := fixedXor([]byte(key), []byte(r))

    fmt.Printf("The char key is: %c and the decrypted text is probably: %s.\n", k, decryptedText)
}
