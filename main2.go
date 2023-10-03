package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	data := []byte("my passs")
	hash := sha512.Sum512(data)
	hashString := hex.EncodeToString(hash[:])
	fmt.Println(time.Since(t).Nanoseconds())
	fmt.Println("SHA-512 Hash:", hashString)
	fmt.Println("SHA-512 Hash:", len(hashString))

}
